/*
log_monitor.go provides a real-time Terminal User Interface (TUI) for the Kafka demo.
It continuously tails 'tracker.log' and 'tracker.events' to provide
live visualization of system metrics and business events.

Features:
- **Live Dashboard**: Visualizes throughput (msg/s) and success rates.
- **Log Streaming**: Displays recent system logs and audit events.
- **Health Indicators**: Provides color-coded status for quick system assessment.
- **Interactive TUI**: Built with termui for a responsive terminal experience.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// MonitorLogEntry est un alias pour LogEntry (défini dans models.go).
// Utilisé pour la surveillance des logs système.
type MonitorLogEntry = LogEntry

// MonitorEventEntry est un alias pour EventEntry (défini dans models.go).
// Utilisé pour la surveillance des événements d'audit.
type MonitorEventEntry = EventEntry

// HealthStatus définit les niveaux de santé pour les indicateurs du tableau de bord.
// Il est utilisé pour déterminer la couleur et le texte à afficher pour chaque métrique.
type HealthStatus int

const (
	HealthGood     HealthStatus = iota // Indique une condition saine, typiquement affichée en vert.
	HealthWarning                      // Indique un avertissement, typiquement affiché en jaune.
	HealthCritical                     // Indique un état critique, typiquement affiché en rouge.
)

// Note: Les constantes sont définies dans constants.go pour éviter les duplications.
// Aliases locaux pour la lisibilité (référencent les constantes de constants.go)
const (
	MaxRecentLogs      = MonitorMaxRecentLogs
	MaxRecentEvents    = MonitorMaxRecentEvents
	MaxHistorySize     = MonitorMaxHistorySize
	LogChannelBuffer   = MonitorLogChannelBuffer
	EventChannelBuffer = MonitorEventChannelBuffer

	SuccessRateExcellent = MonitorSuccessRateExcellent
	SuccessRateGood      = MonitorSuccessRateGood

	ThroughputNormal = MonitorThroughputNormal
	ThroughputLow    = MonitorThroughputLow

	ErrorTimeoutCritical = MonitorErrorTimeoutCritical
	ErrorTimeoutWarning  = MonitorErrorTimeoutWarning

	QualityThroughputHigh   = MonitorQualityThroughputHigh
	QualityThroughputMedium = MonitorQualityThroughputMedium
	QualityThroughputLow    = MonitorQualityThroughputLow

	QualityScoreExcellent = MonitorQualityScoreExcellent
	QualityScoreGood      = MonitorQualityScoreGood
	QualityScoreMedium    = MonitorQualityScoreMedium

	FileCheckInterval = MonitorFileCheckInterval
	FilePollInterval  = MonitorFilePollInterval
	UIUpdateInterval  = MonitorUIUpdateInterval

	MaxLogRowLength   = MonitorMaxLogRowLength
	MaxEventRowLength = MonitorMaxEventRowLength
	TruncateSuffix    = MonitorTruncateSuffix
)

// Metrics agrège et gère l'état de toutes les métriques collectées par le moniteur.
// L'accès à cette structure est protégé par un RWMutex pour garantir la sécurité
// lors des lectures et écritures concurrentes.
type Metrics struct {
	mu                    sync.RWMutex        // Mutex pour un accès concurrent sécurisé.
	StartTime             time.Time           // Heure de démarrage du moniteur.
	MessagesReceived      int64               // Nombre total de messages reçus.
	MessagesProcessed     int64               // Nombre de messages traités avec succès.
	MessagesFailed        int64               // Nombre de messages qui ont échoué au traitement.
	MessagesPerSecond     []float64           // Historique des débits de messages par seconde pour le graphique.
	SuccessRateHistory    []float64           // Historique des taux de succès pour le graphique.
	RecentLogs            []MonitorLogEntry   // Slice des dernières entrées de log de `tracker.log`.
	RecentEvents          []MonitorEventEntry // Slice des derniers événements de `tracker.events`.
	LastUpdateTime        time.Time           // Heure de la dernière mise à jour des métriques.
	Uptime                time.Duration       // Durée de fonctionnement du moniteur.
	CurrentMessagesPerSec float64             // Valeur actuelle du débit de messages.
	CurrentSuccessRate    float64             // Valeur actuelle du taux de succès.
	ErrorCount            int64               // Nombre total d'erreurs détectées.
	LastErrorTime         time.Time           // Heure de la dernière erreur enregistrée.
}

var monitorMetrics = &Metrics{
	StartTime:          time.Now(),
	RecentLogs:         make([]MonitorLogEntry, 0, MaxRecentLogs),
	RecentEvents:       make([]MonitorEventEntry, 0, MaxRecentEvents),
	MessagesPerSecond:  make([]float64, 0, MaxHistorySize),
	SuccessRateHistory: make([]float64, 0, MaxHistorySize),
	LastErrorTime:      time.Time{},
}

// waitForFile attend que le fichier spécifié existe et retourne un handle ouvert.
func waitForFile(filename string) *os.File {
	for {
		file, err := os.Open(filename)
		if err == nil {
			return file
		}
		time.Sleep(FileCheckInterval)
	}
}

// waitForFileRecreation attend que le fichier supprimé soit recréé.
func waitForFileRecreation(filename string) *os.File {
	for {
		time.Sleep(FileCheckInterval)
		file, err := os.Open(filename)
		if err == nil {
			return file
		}
	}
}

// parseAndSendLogEntry parse une ligne JSON et l'envoie sur le canal approprié.
func parseAndSendLogEntry(line string, logChan chan<- MonitorLogEntry) {
	var entry MonitorLogEntry
	if err := json.Unmarshal([]byte(line), &entry); err == nil {
		select {
		case logChan <- entry:
		default:
			// Canal plein, ignorer
		}
	}
}

// parseAndSendEventEntry parse une ligne JSON et l'envoie sur le canal approprié.
func parseAndSendEventEntry(line string, eventChan chan<- MonitorEventEntry) {
	var entry MonitorEventEntry
	if err := json.Unmarshal([]byte(line), &entry); err == nil {
		select {
		case eventChan <- entry:
		default:
			// Canal plein, ignorer
		}
	}
}

// readNewLines lit les nouvelles lignes du fichier et les envoie sur les canaux.
// Retourne la nouvelle position dans le fichier, ou -1 en cas d'erreur.
func readNewLines(file *os.File, filename string, currentPos int64, logChan chan<- MonitorLogEntry, eventChan chan<- MonitorEventEntry) int64 {
	_, err := file.Seek(currentPos, 0)
	if err != nil {
		// Erreur de seek - retourner la position actuelle pour réessayer
		return currentPos
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		if filename == TrackerLogFile {
			parseAndSendLogEntry(line, logChan)
		} else if filename == TrackerEventsFile {
			parseAndSendEventEntry(line, eventChan)
		}
	}

	// Vérifier les erreurs du scanner
	if err := scanner.Err(); err != nil {
		// Erreur de lecture - retourner la position actuelle
		return currentPos
	}

	// Obtenir la nouvelle position
	newPos, err := file.Seek(0, os.SEEK_CUR)
	if err != nil {
		return currentPos
	}
	return newPos
}

// monitorFile surveille un fichier en continu, similaire à la commande `tail -f`.
// Il lit les nouvelles lignes ajoutées au fichier et les envoie sur des canaux
// appropriés pour un traitement asynchrone. La fonction gère aussi la recréation
// et la troncature du fichier.
//
// Paramètres:
//
//	filename (string): Le chemin du fichier à surveiller.
//	logChan (chan<- MonitorLogEntry): Canal pour envoyer les entrées de `tracker.log`.
//	eventChan (chan<- MonitorEventEntry): Canal pour envoyer les entrées de `tracker.events`.
func monitorFile(filename string, logChan chan<- MonitorLogEntry, eventChan chan<- MonitorEventEntry) {
	file := waitForFile(filename)
	var currentPos int64

	for {
		// Vérifier si le fichier existe encore
		stat, err := os.Stat(filename)
		if err != nil {
			// Fichier supprimé, attendre qu'il soit recréé
			file.Close()
			file = waitForFileRecreation(filename)
			currentPos = 0
			continue
		}

		// Si le fichier a été tronqué, repartir du début
		if stat.Size() < currentPos {
			file.Close()
			file = waitForFile(filename)
			currentPos = 0
		}

		// Lire les nouvelles lignes
		if currentPos < stat.Size() {
			newPos := readNewLines(file, filename, currentPos, logChan, eventChan)
			file.Close()
			file = waitForFile(filename)
			currentPos = newPos
		} else {
			time.Sleep(FilePollInterval)
		}
	}
}

// processLog traite une entrée de log provenant de `tracker.log`.
// Elle met à jour l'état global des métriques de manière concurrente-sûre.
//
// Paramètres:
//
//	entry (MonitorLogEntry): L'entrée de log à traiter.
func processLog(entry MonitorLogEntry) {
	monitorMetrics.mu.Lock()
	defer monitorMetrics.mu.Unlock()

	// Ajouter aux logs récents
	monitorMetrics.RecentLogs = append(monitorMetrics.RecentLogs, entry)
	if len(monitorMetrics.RecentLogs) > MaxRecentLogs {
		monitorMetrics.RecentLogs = monitorMetrics.RecentLogs[1:]
	}

	// Compter les erreurs
	if entry.Level == LogLevelERROR {
		monitorMetrics.ErrorCount++
		monitorMetrics.LastErrorTime = time.Now()
	}

	// Extraire les métriques périodiques
	if entry.Message == "Métriques système périodiques" && entry.Metadata != nil {
		if msgsReceived, ok := entry.Metadata["messages_received"].(float64); ok {
			monitorMetrics.MessagesReceived = int64(msgsReceived)
		}
		if msgsProcessed, ok := entry.Metadata["messages_processed"].(float64); ok {
			monitorMetrics.MessagesProcessed = int64(msgsProcessed)
		}
		if msgsFailed, ok := entry.Metadata["messages_failed"].(float64); ok {
			monitorMetrics.MessagesFailed = int64(msgsFailed)
		}
		if mpsStr, ok := entry.Metadata["messages_per_second"].(string); ok {
			if mps, err := strconv.ParseFloat(mpsStr, 64); err == nil {
				monitorMetrics.MessagesPerSecond = append(monitorMetrics.MessagesPerSecond, mps)
				if len(monitorMetrics.MessagesPerSecond) > MaxHistorySize {
					monitorMetrics.MessagesPerSecond = monitorMetrics.MessagesPerSecond[1:]
				}
				monitorMetrics.CurrentMessagesPerSec = mps
			}
		}
		if srStr, ok := entry.Metadata["success_rate_percent"].(string); ok {
			if sr, err := strconv.ParseFloat(srStr, 64); err == nil {
				monitorMetrics.SuccessRateHistory = append(monitorMetrics.SuccessRateHistory, sr)
				if len(monitorMetrics.SuccessRateHistory) > MaxHistorySize {
					monitorMetrics.SuccessRateHistory = monitorMetrics.SuccessRateHistory[1:]
				}
				monitorMetrics.CurrentSuccessRate = sr
			}
		}
	}

	monitorMetrics.LastUpdateTime = time.Now()
}

// processEvent traite une entrée d'événement provenant de `tracker.events`.
// Elle met à jour l'état global des métriques de manière concurrente-sûre.
//
// Paramètres:
//
//	entry (MonitorEventEntry): L'événement à traiter.
func processEvent(entry MonitorEventEntry) {
	monitorMetrics.mu.Lock()
	defer monitorMetrics.mu.Unlock()

	// Ajouter aux événements récents
	monitorMetrics.RecentEvents = append(monitorMetrics.RecentEvents, entry)
	if len(monitorMetrics.RecentEvents) > MaxRecentEvents {
		monitorMetrics.RecentEvents = monitorMetrics.RecentEvents[1:]
	}

	// Mettre à jour les compteurs
	if entry.Deserialized {
		monitorMetrics.MessagesProcessed++
	} else {
		monitorMetrics.MessagesFailed++
		monitorMetrics.ErrorCount++
		monitorMetrics.LastErrorTime = time.Now()
	}
	monitorMetrics.MessagesReceived++

	// Recalculer les métriques en temps réel
	uptime := time.Since(monitorMetrics.StartTime)
	if uptime.Seconds() > 0 {
		monitorMetrics.CurrentMessagesPerSec = float64(monitorMetrics.MessagesReceived) / uptime.Seconds()
	}
	if monitorMetrics.MessagesReceived > 0 {
		monitorMetrics.CurrentSuccessRate = float64(monitorMetrics.MessagesProcessed) / float64(monitorMetrics.MessagesReceived) * 100
	}

	monitorMetrics.LastUpdateTime = time.Now()
}

// createMetricsTable initialise et configure le widget de tableau pour les métriques principales.
//
// Retourne:
//
//	(*widgets.Table): Un pointeur vers le widget de tableau configuré.
func createMetricsTable() *widgets.Table {
	table := widgets.NewTable()
	table.Rows = [][]string{
		{"Métrique", "Valeur"},
		{"Messages reçus", "0"},
		{"Messages traités", "0"},
		{"Messages échoués", "0"},
		{"Débit (msg/s)", "0.00"},
		{"Taux de succès", "0.00%"},
		{"Dernière mise à jour", "-"},
	}
	table.TextStyle = ui.NewStyle(ui.ColorWhite)
	table.RowStyles[0] = ui.NewStyle(ui.ColorYellow, ui.ColorClear, ui.ModifierBold)
	table.SetRect(0, 0, 50, 9)
	table.ColumnWidths = []int{30, 20}
	return table
}

// createHealthDashboard initialise le widget de tableau pour le tableau de bord de santé.
//
// Retourne:
//
//	(*widgets.Table): Un pointeur vers le widget de tableau configuré.
func createHealthDashboard() *widgets.Table {
	table := widgets.NewTable()
	table.Rows = [][]string{
		{"Indicateur", "Statut"},
		{"Santé globale", "●"},
		{"Taux de succès", "●"},
		{"Débit", "●"},
		{"Erreurs", "●"},
		{"Uptime", "-"},
		{"Qualité", "-"},
	}
	table.TextStyle = ui.NewStyle(ui.ColorWhite)
	table.RowStyles[0] = ui.NewStyle(ui.ColorYellow, ui.ColorClear, ui.ModifierBold)
	table.SetRect(50, 0, 110, 9)
	table.ColumnWidths = []int{25, 35}
	return table
}

// StatusThreshold définit un seuil pour l'évaluation du statut.
type StatusThreshold struct {
	MinValue float64      // Valeur minimale pour ce seuil
	Status   HealthStatus // Statut associé
	Text     string       // Texte à afficher
	Color    ui.Color     // Couleur pour l'affichage
}

// evaluateStatus évalue une valeur par rapport à des seuils ordonnés (du plus élevé au plus bas).
// Les seuils doivent être ordonnés par MinValue décroissante.
func evaluateStatus(value float64, thresholds []StatusThreshold) (HealthStatus, string, ui.Color) {
	for _, t := range thresholds {
		if value >= t.MinValue {
			return t.Status, t.Text, t.Color
		}
	}
	// Retourner le dernier seuil si aucun ne correspond
	if len(thresholds) > 0 {
		last := thresholds[len(thresholds)-1]
		return last.Status, last.Text, last.Color
	}
	return HealthCritical, "● INCONNU", ui.ColorRed
}

// Seuils prédéfinis pour les différentes métriques
var (
	healthThresholds = []StatusThreshold{
		{SuccessRateExcellent, HealthGood, "● EXCELLENT", ui.ColorGreen},
		{SuccessRateGood, HealthWarning, "● BON", ui.ColorYellow},
		{0, HealthCritical, "● CRITIQUE", ui.ColorRed},
	}

	throughputThresholds = []StatusThreshold{
		{ThroughputNormal, HealthGood, "● NORMAL", ui.ColorGreen},
		{ThroughputLow, HealthWarning, "● FAIBLE", ui.ColorYellow},
		{0, HealthCritical, "● ARRÊTÉ", ui.ColorRed},
	}
)

// getHealthStatus évalue le taux de succès et retourne un statut de santé.
func getHealthStatus(successRate float64) (HealthStatus, string, ui.Color) {
	return evaluateStatus(successRate, healthThresholds)
}

// getThroughputStatus évalue le débit de messages et retourne un statut de santé.
func getThroughputStatus(mps float64) (HealthStatus, string, ui.Color) {
	return evaluateStatus(mps, throughputThresholds)
}

// getErrorStatus évalue le nombre d'erreurs et le temps écoulé depuis la dernière erreur.
// Cette fonction a une logique spécifique qui ne peut pas être généralisée avec evaluateStatus.
func getErrorStatus(errorCount int64, lastErrorTime time.Time) (HealthStatus, string, ui.Color) {
	if errorCount == 0 {
		return HealthGood, "● AUCUNE", ui.ColorGreen
	}

	timeSinceError := time.Since(lastErrorTime)
	if timeSinceError > ErrorTimeoutWarning {
		return HealthGood, "● AUCUNE", ui.ColorGreen
	} else if timeSinceError > ErrorTimeoutCritical {
		return HealthWarning, "● RÉCENTES", ui.ColorYellow
	}
	return HealthCritical, "● ACTIVES", ui.ColorRed
}

// calculateQualityScore calcule un score de qualité global (0-100) basé sur plusieurs métriques.
//
// Paramètres:
//
//	successRate (float64): Le taux de succès.
//	mps (float64): Le débit de messages par seconde.
//	errorCount (int64): Le nombre d'erreurs.
//	uptime (time.Duration): La durée de fonctionnement.
//
// Retourne:
//
//	(float64): Le score de qualité calculé.
func calculateQualityScore(successRate, mps float64, errorCount int64, uptime time.Duration) float64 {
	// Score basé sur le taux de succès (0-50 points)
	successScore := (successRate / 100.0) * 50.0

	// Score basé sur le débit (0-30 points)
	throughputScore := 0.0
	if mps >= QualityThroughputHigh {
		throughputScore = 30.0
	} else if mps >= QualityThroughputMedium {
		throughputScore = 25.0
	} else if mps >= QualityThroughputLow {
		throughputScore = 15.0
	} else if mps > 0 {
		throughputScore = 10.0
	}

	// Score basé sur les erreurs (0-20 points)
	errorScore := 20.0
	if errorCount > 0 {
		errorPenalty := float64(errorCount) * 2.0
		if errorPenalty > 20.0 {
			errorPenalty = 20.0
		}
		errorScore = 20.0 - errorPenalty
		if errorScore < 0 {
			errorScore = 0
		}
	}

	return successScore + throughputScore + errorScore
}

// createLogList initialise le widget de liste pour afficher les logs récents de `tracker.log`.
//
// Retourne:
//
//	(*widgets.List): Un pointeur vers le widget de liste configuré.
func createLogList() *widgets.List {
	list := widgets.NewList()
	list.Title = "Logs Récents (tracker.log)"
	list.Rows = []string{"En attente de logs..."}
	list.TextStyle = ui.NewStyle(ui.ColorWhite)
	list.SelectedRowStyle = ui.NewStyle(ui.ColorBlack, ui.ColorWhite)
	list.WrapText = true
	list.SetRect(0, 9, 80, 19)
	return list
}

// createEventList initialise le widget de liste pour afficher les événements récents de `tracker.events`.
//
// Retourne:
//
//	(*widgets.List): Un pointeur vers le widget de liste configuré.
func createEventList() *widgets.List {
	list := widgets.NewList()
	list.Title = "Événements Récents (tracker.events)"
	list.Rows = []string{"En attente d'événements..."}
	list.TextStyle = ui.NewStyle(ui.ColorWhite)
	list.SelectedRowStyle = ui.NewStyle(ui.ColorBlack, ui.ColorWhite)
	list.WrapText = true
	list.SetRect(80, 9, 160, 19)
	return list
}

// createMessagesPerSecondChart initialise le widget de graphique pour le débit de messages.
//
// Retourne:
//
//	(*widgets.Plot): Un pointeur vers le widget de graphique configuré.
func createMessagesPerSecondChart() *widgets.Plot {
	plot := widgets.NewPlot()
	plot.Title = "Débit de Messages (msg/s)"
	plot.Data = [][]float64{{}}
	plot.SetRect(0, 19, 80, 29)
	plot.AxesColor = ui.ColorWhite
	plot.LineColors[0] = ui.ColorGreen
	plot.Marker = widgets.MarkerDot
	return plot
}

// createSuccessRateChart initialise le widget de graphique pour le taux de succès.
//
// Retourne:
//
//	(*widgets.Plot): Un pointeur vers le widget de graphique configuré.
func createSuccessRateChart() *widgets.Plot {
	plot := widgets.NewPlot()
	plot.Title = "Taux de Succès (%)"
	plot.Data = [][]float64{{}}
	plot.SetRect(80, 19, 160, 29)
	plot.AxesColor = ui.ColorWhite
	plot.LineColors[0] = ui.ColorBlue
	plot.Marker = widgets.MarkerDot
	return plot
}

// updateMetricsTable met à jour le tableau des métriques principales.
func updateMetricsTable(table *widgets.Table, m *Metrics) {
	table.Rows = [][]string{
		{"Métrique", "Valeur"},
		{"Messages reçus", fmt.Sprintf("%d", m.MessagesReceived)},
		{"Messages traités", fmt.Sprintf("%d", m.MessagesProcessed)},
		{"Messages échoués", fmt.Sprintf("%d", m.MessagesFailed)},
		{"Débit (msg/s)", fmt.Sprintf("%.2f", m.CurrentMessagesPerSec)},
		{"Taux de succès", fmt.Sprintf("%.2f%%", m.CurrentSuccessRate)},
		{"Dernière mise à jour", m.LastUpdateTime.Format("15:04:05")},
	}
}

// getGlobalHealthStatus détermine la santé globale à partir des statuts individuels.
func getGlobalHealthStatus(successStatus, throughputStatus, errorStatus HealthStatus) (HealthStatus, string, ui.Color) {
	globalStatus := successStatus
	if throughputStatus > globalStatus {
		globalStatus = throughputStatus
	}
	if errorStatus > globalStatus {
		globalStatus = errorStatus
	}

	switch globalStatus {
	case HealthWarning:
		return globalStatus, "● ATTENTION", ui.ColorYellow
	case HealthCritical:
		return globalStatus, "● CRITIQUE", ui.ColorRed
	default:
		return HealthGood, "● EXCELLENT", ui.ColorGreen
	}
}

// getQualityText retourne le texte et la couleur pour un score de qualité.
func getQualityText(qualityScore float64) (string, ui.Color) {
	if qualityScore >= QualityScoreExcellent {
		return fmt.Sprintf("EXCELLENT (%.0f)", qualityScore), ui.ColorGreen
	} else if qualityScore >= QualityScoreGood {
		return fmt.Sprintf("BON (%.0f)", qualityScore), ui.ColorYellow
	} else if qualityScore >= QualityScoreMedium {
		return fmt.Sprintf("MOYEN (%.0f)", qualityScore), ui.ColorYellow
	}
	return fmt.Sprintf("FAIBLE (%.0f)", qualityScore), ui.ColorRed
}

// formatUptime formate la durée de fonctionnement en chaîne lisible.
func formatUptime(uptime time.Duration) string {
	if uptime.Hours() >= 1 {
		return fmt.Sprintf("%.1fh", uptime.Hours())
	} else if uptime.Minutes() >= 1 {
		return fmt.Sprintf("%.0fm", uptime.Minutes())
	}
	return fmt.Sprintf("%.0fs", uptime.Seconds())
}

// updateHealthDashboard met à jour le tableau de bord de santé.
func updateHealthDashboard(dashboard *widgets.Table, m *Metrics) {
	successStatus, successText, successColor := getHealthStatus(m.CurrentSuccessRate)
	throughputStatus, throughputText, throughputColor := getThroughputStatus(m.CurrentMessagesPerSec)
	errorStatus, errorText, errorColor := getErrorStatus(m.ErrorCount, m.LastErrorTime)

	_, globalText, globalColor := getGlobalHealthStatus(successStatus, throughputStatus, errorStatus)

	qualityScore := calculateQualityScore(m.CurrentSuccessRate, m.CurrentMessagesPerSec, m.ErrorCount, m.Uptime)
	qualityText, qualityColor := getQualityText(qualityScore)
	uptimeStr := formatUptime(m.Uptime)

	dashboard.Rows = [][]string{
		{"Indicateur", "Statut"},
		{"Santé globale", globalText},
		{"Taux de succès", successText},
		{"Débit", throughputText},
		{"Erreurs", errorText},
		{"Uptime", uptimeStr},
		{"Qualité", qualityText},
	}

	dashboard.RowStyles = make(map[int]ui.Style)
	dashboard.RowStyles[0] = ui.NewStyle(ui.ColorYellow, ui.ColorClear, ui.ModifierBold)
	dashboard.RowStyles[1] = ui.NewStyle(globalColor, ui.ColorClear, ui.ModifierBold)
	dashboard.RowStyles[2] = ui.NewStyle(successColor, ui.ColorClear)
	dashboard.RowStyles[3] = ui.NewStyle(throughputColor, ui.ColorClear)
	dashboard.RowStyles[4] = ui.NewStyle(errorColor, ui.ColorClear)
	dashboard.RowStyles[5] = ui.NewStyle(ui.ColorCyan, ui.ColorClear)
	dashboard.RowStyles[6] = ui.NewStyle(qualityColor, ui.ColorClear, ui.ModifierBold)
}

// formatLogRow formate une entrée de log pour l'affichage.
func formatLogRow(log MonitorLogEntry) string {
	levelIcon := "🟢"
	if log.Level == LogLevelERROR {
		levelIcon = "🔴"
	}

	timeStr := log.Timestamp
	if len(timeStr) > 19 {
		timeStr = timeStr[11:19]
	}

	row := fmt.Sprintf("%s [%s] %s", levelIcon, timeStr, log.Message)
	if len(row) > MaxLogRowLength {
		row = row[:MaxLogRowLength-len(TruncateSuffix)] + TruncateSuffix
	}
	return row
}

// updateLogList met à jour la liste des logs récents.
func updateLogList(list *widgets.List, logs []MonitorLogEntry) {
	rows := make([]string, 0, len(logs))
	for i := len(logs) - 1; i >= 0; i-- {
		rows = append(rows, formatLogRow(logs[i]))
	}
	if len(rows) == 0 {
		rows = []string{"En attente de logs..."}
	}
	list.Rows = rows
}

// formatEventRow formate une entrée d'événement pour l'affichage.
func formatEventRow(event MonitorEventEntry) string {
	status := "❌"
	if event.Deserialized {
		status = "✅"
	}

	timeStr := event.Timestamp
	if len(timeStr) > 19 {
		timeStr = timeStr[11:19]
	}

	row := fmt.Sprintf("%s [%s] Offset: %d | %s", status, timeStr, event.KafkaOffset, event.EventType)
	if len(row) > MaxEventRowLength {
		row = row[:MaxEventRowLength-len(TruncateSuffix)] + TruncateSuffix
	}
	return row
}

// updateEventList met à jour la liste des événements récents.
func updateEventList(list *widgets.List, events []MonitorEventEntry) {
	rows := make([]string, 0, len(events))
	for i := len(events) - 1; i >= 0; i-- {
		rows = append(rows, formatEventRow(events[i]))
	}
	if len(rows) == 0 {
		rows = []string{"En attente d'événements..."}
	}
	list.Rows = rows
}

// updateCharts met à jour les graphiques de débit et de taux de succès.
func updateCharts(mpsChart, srChart *widgets.Plot, mps, sr []float64) {
	if len(mps) > 0 {
		mpsChart.Data = [][]float64{mps}
	} else {
		mpsChart.Data = [][]float64{{0}}
	}

	if len(sr) > 0 {
		srChart.Data = [][]float64{sr}
	} else {
		srChart.Data = [][]float64{{0}}
	}
}

// updateUI rafraîchit tous les widgets de l'interface utilisateur avec les dernières métriques.
// Cette fonction est appelée périodiquement pour mettre à jour l'affichage.
func updateUI(table *widgets.Table, healthDashboard *widgets.Table, logList *widgets.List, eventList *widgets.List, mpsChart *widgets.Plot, srChart *widgets.Plot) {
	monitorMetrics.mu.RLock()
	defer monitorMetrics.mu.RUnlock()

	updateMetricsTable(table, monitorMetrics)
	updateHealthDashboard(healthDashboard, monitorMetrics)
	updateLogList(logList, monitorMetrics.RecentLogs)
	updateEventList(eventList, monitorMetrics.RecentEvents)
	updateCharts(mpsChart, srChart, monitorMetrics.MessagesPerSecond, monitorMetrics.SuccessRateHistory)
}

// Note: La fonction main() est définie dans cmd_monitor.go avec le build tag "monitor"
// Pour compiler: go build -tags monitor -o log_monitor.exe
