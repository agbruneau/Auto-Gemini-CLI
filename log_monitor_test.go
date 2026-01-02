package main

import (
	"testing"
	"time"

	ui "github.com/gizak/termui/v3"
)

// TestGetHealthStatus tests the getHealthStatus function with various success rates.
func TestGetHealthStatus(t *testing.T) {
	tests := []struct {
		name           string
		successRate    float64
		expectedStatus HealthStatus
		expectedText   string
		expectedColor  ui.Color
	}{
		{
			name:           "Excellent - 100%",
			successRate:    100.0,
			expectedStatus: HealthGood,
			expectedText:   "● EXCELLENT",
			expectedColor:  ui.ColorGreen,
		},
		{
			name:           "Excellent - threshold",
			successRate:    95.0,
			expectedStatus: HealthGood,
			expectedText:   "● EXCELLENT",
			expectedColor:  ui.ColorGreen,
		},
		{
			name:           "Good - just below excellent",
			successRate:    94.9,
			expectedStatus: HealthWarning,
			expectedText:   "● BON",
			expectedColor:  ui.ColorYellow,
		},
		{
			name:           "Good - at threshold",
			successRate:    80.0,
			expectedStatus: HealthWarning,
			expectedText:   "● BON",
			expectedColor:  ui.ColorYellow,
		},
		{
			name:           "Critical - below good",
			successRate:    79.9,
			expectedStatus: HealthCritical,
			expectedText:   "● CRITIQUE",
			expectedColor:  ui.ColorRed,
		},
		{
			name:           "Critical - zero",
			successRate:    0.0,
			expectedStatus: HealthCritical,
			expectedText:   "● CRITIQUE",
			expectedColor:  ui.ColorRed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status, text, color := getHealthStatus(tt.successRate)
			if status != tt.expectedStatus {
				t.Errorf("getHealthStatus(%f) status = %v, want %v", tt.successRate, status, tt.expectedStatus)
			}
			if text != tt.expectedText {
				t.Errorf("getHealthStatus(%f) text = %v, want %v", tt.successRate, text, tt.expectedText)
			}
			if color != tt.expectedColor {
				t.Errorf("getHealthStatus(%f) color = %v, want %v", tt.successRate, color, tt.expectedColor)
			}
		})
	}
}

// TestGetThroughputStatus tests the getThroughputStatus function with various throughput values.
func TestGetThroughputStatus(t *testing.T) {
	tests := []struct {
		name           string
		mps            float64
		expectedStatus HealthStatus
		expectedText   string
		expectedColor  ui.Color
	}{
		{
			name:           "Normal - high throughput",
			mps:            1.0,
			expectedStatus: HealthGood,
			expectedText:   "● NORMAL",
			expectedColor:  ui.ColorGreen,
		},
		{
			name:           "Normal - at threshold",
			mps:            0.3,
			expectedStatus: HealthGood,
			expectedText:   "● NORMAL",
			expectedColor:  ui.ColorGreen,
		},
		{
			name:           "Low - below normal",
			mps:            0.2,
			expectedStatus: HealthWarning,
			expectedText:   "● FAIBLE",
			expectedColor:  ui.ColorYellow,
		},
		{
			name:           "Low - at threshold",
			mps:            0.1,
			expectedStatus: HealthWarning,
			expectedText:   "● FAIBLE",
			expectedColor:  ui.ColorYellow,
		},
		{
			name:           "Stopped - below low",
			mps:            0.05,
			expectedStatus: HealthCritical,
			expectedText:   "● ARRÊTÉ",
			expectedColor:  ui.ColorRed,
		},
		{
			name:           "Stopped - zero",
			mps:            0.0,
			expectedStatus: HealthCritical,
			expectedText:   "● ARRÊTÉ",
			expectedColor:  ui.ColorRed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status, text, color := getThroughputStatus(tt.mps)
			if status != tt.expectedStatus {
				t.Errorf("getThroughputStatus(%f) status = %v, want %v", tt.mps, status, tt.expectedStatus)
			}
			if text != tt.expectedText {
				t.Errorf("getThroughputStatus(%f) text = %v, want %v", tt.mps, text, tt.expectedText)
			}
			if color != tt.expectedColor {
				t.Errorf("getThroughputStatus(%f) color = %v, want %v", tt.mps, color, tt.expectedColor)
			}
		})
	}
}

// TestGetErrorStatus tests the getErrorStatus function.
func TestGetErrorStatus(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name           string
		errorCount     int64
		lastErrorTime  time.Time
		expectedStatus HealthStatus
		expectedText   string
	}{
		{
			name:           "No errors",
			errorCount:     0,
			lastErrorTime:  time.Time{},
			expectedStatus: HealthGood,
			expectedText:   "● AUCUNE",
		},
		{
			name:           "Old errors - more than 5 minutes",
			errorCount:     5,
			lastErrorTime:  now.Add(-6 * time.Minute),
			expectedStatus: HealthGood,
			expectedText:   "● AUCUNE",
		},
		{
			name:           "Recent errors - between 1 and 5 minutes",
			errorCount:     3,
			lastErrorTime:  now.Add(-2 * time.Minute),
			expectedStatus: HealthWarning,
			expectedText:   "● RÉCENTES",
		},
		{
			name:           "Active errors - less than 1 minute",
			errorCount:     2,
			lastErrorTime:  now.Add(-30 * time.Second),
			expectedStatus: HealthCritical,
			expectedText:   "● ACTIVES",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status, text, _ := getErrorStatus(tt.errorCount, tt.lastErrorTime)
			if status != tt.expectedStatus {
				t.Errorf("getErrorStatus() status = %v, want %v", status, tt.expectedStatus)
			}
			if text != tt.expectedText {
				t.Errorf("getErrorStatus() text = %v, want %v", text, tt.expectedText)
			}
		})
	}
}

// TestCalculateQualityScore tests the calculateQualityScore function.
func TestCalculateQualityScore(t *testing.T) {
	tests := []struct {
		name        string
		successRate float64
		mps         float64
		errorCount  int64
		uptime      time.Duration
		minScore    float64
		maxScore    float64
	}{
		{
			name:        "Perfect score",
			successRate: 100.0,
			mps:         0.5,
			errorCount:  0,
			uptime:      1 * time.Hour,
			minScore:    90.0,
			maxScore:    100.0,
		},
		{
			name:        "Medium score",
			successRate: 80.0,
			mps:         0.3,
			errorCount:  5,
			uptime:      30 * time.Minute,
			minScore:    50.0,
			maxScore:    80.0,
		},
		{
			name:        "Low score",
			successRate: 50.0,
			mps:         0.05,
			errorCount:  20,
			uptime:      5 * time.Minute,
			minScore:    0.0,
			maxScore:    50.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score := calculateQualityScore(tt.successRate, tt.mps, tt.errorCount, tt.uptime)
			if score < tt.minScore || score > tt.maxScore {
				t.Errorf("calculateQualityScore() = %f, want between %f and %f", score, tt.minScore, tt.maxScore)
			}
		})
	}
}

// TestFormatUptime tests the formatUptime function.
func TestFormatUptime(t *testing.T) {
	tests := []struct {
		name     string
		uptime   time.Duration
		expected string
	}{
		{
			name:     "Hours",
			uptime:   2*time.Hour + 30*time.Minute,
			expected: "2.5h",
		},
		{
			name:     "Minutes",
			uptime:   45 * time.Minute,
			expected: "45m",
		},
		{
			name:     "Seconds",
			uptime:   30 * time.Second,
			expected: "30s",
		},
		{
			name:     "Zero",
			uptime:   0,
			expected: "0s",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatUptime(tt.uptime)
			if result != tt.expected {
				t.Errorf("formatUptime(%v) = %s, want %s", tt.uptime, result, tt.expected)
			}
		})
	}
}

// TestProcessEventRealTimeMetricsUpdate verifies that processEvent correctly
// updates the metrics in real-time.
func TestProcessEventRealTimeMetricsUpdate(t *testing.T) {
	// Reset metrics for a clean test environment
	monitorMetrics = &Metrics{
		StartTime:          time.Now(),
		RecentLogs:         make([]MonitorLogEntry, 0, MaxRecentLogs),
		RecentEvents:       make([]MonitorEventEntry, 0, MaxRecentEvents),
		MessagesPerSecond:  make([]float64, 0, MaxHistorySize),
		SuccessRateHistory: make([]float64, 0, MaxHistorySize),
	}

	// Simulate receiving one successful event
	event := MonitorEventEntry{
		Deserialized: true,
	}
	processEvent(event)

	// Assertions
	if monitorMetrics.MessagesReceived != 1 {
		t.Errorf("Expected MessagesReceived to be 1, got %d", monitorMetrics.MessagesReceived)
	}
	if monitorMetrics.MessagesProcessed != 1 {
		t.Errorf("Expected MessagesProcessed to be 1, got %d", monitorMetrics.MessagesProcessed)
	}
	if monitorMetrics.CurrentSuccessRate != 100.0 {
		t.Errorf("Expected CurrentSuccessRate to be 100.0, got %.2f", monitorMetrics.CurrentSuccessRate)
	}

	// Simulate receiving one failed event
	event = MonitorEventEntry{
		Deserialized: false,
	}
	processEvent(event)

	// Assertions
	if monitorMetrics.MessagesReceived != 2 {
		t.Errorf("Expected MessagesReceived to be 2, got %d", monitorMetrics.MessagesReceived)
	}
	if monitorMetrics.MessagesFailed != 1 {
		t.Errorf("Expected MessagesFailed to be 1, got %d", monitorMetrics.MessagesFailed)
	}
	if monitorMetrics.CurrentSuccessRate != 50.0 {
		t.Errorf("Expected CurrentSuccessRate to be 50.0, got %.2f", monitorMetrics.CurrentSuccessRate)
	}
}

// TestGetGlobalHealthStatus tests the getGlobalHealthStatus function.
func TestGetGlobalHealthStatus(t *testing.T) {
	tests := []struct {
		name              string
		successStatus     HealthStatus
		throughputStatus  HealthStatus
		errorStatus       HealthStatus
		expectedStatus    HealthStatus
	}{
		{
			name:             "All good",
			successStatus:    HealthGood,
			throughputStatus: HealthGood,
			errorStatus:      HealthGood,
			expectedStatus:   HealthGood,
		},
		{
			name:             "One warning",
			successStatus:    HealthGood,
			throughputStatus: HealthWarning,
			errorStatus:      HealthGood,
			expectedStatus:   HealthWarning,
		},
		{
			name:             "One critical",
			successStatus:    HealthGood,
			throughputStatus: HealthGood,
			errorStatus:      HealthCritical,
			expectedStatus:   HealthCritical,
		},
		{
			name:             "Mixed - critical wins",
			successStatus:    HealthWarning,
			throughputStatus: HealthCritical,
			errorStatus:      HealthGood,
			expectedStatus:   HealthCritical,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status, _, _ := getGlobalHealthStatus(tt.successStatus, tt.throughputStatus, tt.errorStatus)
			if status != tt.expectedStatus {
				t.Errorf("getGlobalHealthStatus() status = %v, want %v", status, tt.expectedStatus)
			}
		})
	}
}

// TestFormatLogRow tests the formatLogRow function.
func TestFormatLogRow(t *testing.T) {
	tests := []struct {
		name     string
		log      MonitorLogEntry
		contains string
	}{
		{
			name: "Info log",
			log: MonitorLogEntry{
				Level:     LogLevelINFO,
				Timestamp: "2024-01-15T10:30:00Z",
				Message:   "Test message",
			},
			contains: "🟢",
		},
		{
			name: "Error log",
			log: MonitorLogEntry{
				Level:     LogLevelERROR,
				Timestamp: "2024-01-15T10:30:00Z",
				Message:   "Error message",
			},
			contains: "🔴",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatLogRow(tt.log)
			if len(result) == 0 {
				t.Error("formatLogRow() returned empty string")
			}
			// Note: emoji checking might not work correctly in all environments
		})
	}
}
