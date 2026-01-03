# Fibonacci Benchmark Suite - Generate Report
# PowerShell Script for Windows

param(
    [string]$ReportDir = ""
)

$ErrorActionPreference = "Stop"

Write-Host "[REPORT] Fibonacci Report Generator" -ForegroundColor Cyan
Write-Host "====================================" -ForegroundColor Cyan
Write-Host ""

# Determine report directory
if (-not $ReportDir) {
    $timestamp = Get-Date -Format "yyyy-MM-dd_HH-mm-ss"
    $ReportDir = "reports\$timestamp"
}
$dataDir = "$ReportDir\data"

Write-Host "[DIR] Report directory: $ReportDir" -ForegroundColor Yellow
New-Item -ItemType Directory -Force -Path $dataDir | Out-Null

# Build release first
Write-Host "[BUILD] Building in release mode..." -ForegroundColor Yellow
cargo build --release -p fib-viz
if ($LASTEXITCODE -ne 0) {
    Write-Host "[ERROR] Build failed!" -ForegroundColor Red
    exit 1
}

# Run fib-viz to generate CSV data
Write-Host "[VIZ] Generating visualization data..." -ForegroundColor Yellow
cargo run --release -p fib-viz
if ($LASTEXITCODE -ne 0) {
    Write-Host "[ERROR] Visualization generation failed!" -ForegroundColor Red
    exit 1
}

# Move generated CSVs to report directory
Write-Host "[COPY] Moving data files..." -ForegroundColor Yellow
if (Test-Path "results") {
    Copy-Item -Path "results\*" -Destination $dataDir -Recurse -Force
    Write-Host "   [OK] Data files copied to $dataDir" -ForegroundColor Green
} else {
    Write-Host "   [WARN] No results directory found" -ForegroundColor Yellow
}

Write-Host ""
Write-Host "[DONE] Report generation complete!" -ForegroundColor Green
Write-Host "   Data saved to: $dataDir" -ForegroundColor Cyan
Write-Host ""

# Copy HTML templates
Write-Host "[HTML] Copying HTML report templates..." -ForegroundColor Yellow
if (Test-Path "templates") {
    Copy-Item -Path "templates\*" -Destination $ReportDir -Recurse -Force
    Write-Host "   [OK] HTML templates copied to $ReportDir" -ForegroundColor Green
} else {
    Write-Host "   [WARN] No templates directory found" -ForegroundColor Yellow
}

# Embed CSV data into HTML for offline viewing
Write-Host "[EMBED] Embedding data into HTML for offline viewing..." -ForegroundColor Yellow
$htmlPath = "$ReportDir\report.html"
if ((Test-Path $htmlPath) -and (Test-Path $dataDir)) {
    $htmlContent = Get-Content $htmlPath -Raw
    
    # Read CSV files
    $complexityData = ""
    $binetData = ""
    $goldenData = ""
    
    if (Test-Path "$dataDir\complexity_comparison.csv") {
        $complexityData = (Get-Content "$dataDir\complexity_comparison.csv" -Raw).Trim()
    }
    if (Test-Path "$dataDir\binet_accuracy.csv") {
        $binetData = (Get-Content "$dataDir\binet_accuracy.csv" -Raw).Trim()
    }
    if (Test-Path "$dataDir\golden_ratio_convergence.csv") {
        $goldenData = (Get-Content "$dataDir\golden_ratio_convergence.csv" -Raw).Trim()
    }
    
    # Replace placeholder with timestamp
    $timestamp = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
    $htmlContent = $htmlContent -replace "{{TIMESTAMP}}", $timestamp
    
    # Build the data script with proper string concatenation
    $bt = '`'  # backtick character for JS template literals
    $dataScript = "<script>`n// Embedded data for offline viewing`nwindow.EMBEDDED_DATA = {`n"
    $dataScript += "    complexity: $bt$complexityData$bt,`n"
    $dataScript += "    binet: $bt$binetData$bt,`n"
    $dataScript += "    golden: $bt$goldenData$bt`n"
    $dataScript += "};`n</script>"
    
    # Insert before closing </head> tag
    $htmlContent = $htmlContent -replace "</head>", "$dataScript`n</head>"
    
    Set-Content -Path $htmlPath -Value $htmlContent -Encoding UTF8
    Write-Host "   [OK] Data embedded into HTML" -ForegroundColor Green
}

Write-Host ""
Write-Host "Generated files:" -ForegroundColor Yellow
Get-ChildItem $dataDir -File | ForEach-Object { Write-Host "   - data\$($_.Name)" -ForegroundColor White }
Get-ChildItem $ReportDir -File -ErrorAction SilentlyContinue | ForEach-Object { Write-Host "   - $($_.Name)" -ForegroundColor White }

Write-Host ""
Write-Host "To view the report, open:" -ForegroundColor Cyan
Write-Host "   $ReportDir\report.html" -ForegroundColor White
