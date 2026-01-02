//go:build monitor
// +build monitor

/*
Package main provides the entry point for the Log Monitor TUI service.

The Monitor provides real-time visualization of system metrics and logs.
To compile: go build -tags monitor -o log_monitor.exe
*/

package main

import (
	"fmt"
	"os"
	"time"

	ui "github.com/gizak/termui/v3"
)

func main() {
	if err := ui.Init(); err != nil {
		fmt.Printf("Erreur lors de l'initialisation de l'interface: %v\n", err)
		os.Exit(1)
	}
	defer ui.Close()

	// Canaux pour les logs et événements
	logChan := make(chan MonitorLogEntry, LogChannelBuffer)
	eventChan := make(chan MonitorEventEntry, EventChannelBuffer)

	// Démarrer la surveillance des fichiers
	go monitorFile(TrackerLogFile, logChan, nil)
	go monitorFile(TrackerEventsFile, nil, eventChan)

	// Traiter les logs et événements
	go func() {
		for {
			select {
			case log := <-logChan:
				processLog(log)
			case event := <-eventChan:
				processEvent(event)
			}
		}
	}()

	// Créer les widgets
	metricsTable := createMetricsTable()
	healthDashboard := createHealthDashboard()
	logList := createLogList()
	eventList := createEventList()
	mpsChart := createMessagesPerSecondChart()
	srChart := createSuccessRateChart()

	// Gérer le redimensionnement
	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(UIUpdateInterval)
	defer ticker.Stop()

	monitorMetrics.StartTime = time.Now()

	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<Resize>":
				metricsTable.SetRect(0, 0, 50, 9)
				healthDashboard.SetRect(50, 0, 110, 9)
				logList.SetRect(0, 9, 80, 19)
				eventList.SetRect(80, 9, 160, 19)
				mpsChart.SetRect(0, 19, 80, 29)
				srChart.SetRect(80, 19, 160, 29)
				ui.Clear()
			}
		case <-ticker.C:
			monitorMetrics.mu.Lock()
			monitorMetrics.Uptime = time.Since(monitorMetrics.StartTime)
			monitorMetrics.mu.Unlock()
			updateUI(metricsTable, healthDashboard, logList, eventList, mpsChart, srChart)
			ui.Render(metricsTable, healthDashboard, logList, eventList, mpsChart, srChart)
		}
	}
}
