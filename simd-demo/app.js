// SIMD Fibonacci Demo - JavaScript

document.addEventListener('DOMContentLoaded', () => {
    initCpuFeatures();
    initEventListeners();
});

// Fibonacci calculation (scalar - for demo purposes)
function fibScalar(n) {
    if (n === 0n) return 0n;
    if (n === 1n) return 1n;
    
    let a = 0n, b = 1n;
    for (let i = 2n; i <= n; i++) {
        const temp = a + b;
        a = b;
        b = temp;
    }
    return b;
}

// Simulated SIMD batch calculation (JavaScript doesn't have true SIMD for bigints)
function fibSimdBatch(indices) {
    return indices.map(n => fibScalar(BigInt(n)));
}

// Batch calculation
function fibBatch(indices) {
    return indices.map(n => fibScalar(BigInt(n)));
}

// Initialize CPU features display (simulated for browser)
function initCpuFeatures() {
    // In a real scenario, we can't detect CPU SIMD features from JavaScript
    // This is just for demonstration
    const features = {
        'sse2': true,    // Assume SSE2 is available on modern CPUs
        'avx': true,     // Most modern CPUs have AVX
        'avx2': true,    // Common on CPUs from 2013+
        'avx512': false  // Less common
    };
    
    Object.entries(features).forEach(([feature, supported]) => {
        const el = document.getElementById(`feat-${feature}`);
        if (el) {
            el.querySelector('.icon').textContent = supported ? '✅' : '❌';
            el.classList.add(supported ? 'supported' : 'unsupported');
        }
    });
}

// Initialize event listeners
function initEventListeners() {
    document.getElementById('btn-calculate').addEventListener('click', calculate);
    document.getElementById('btn-benchmark').addEventListener('click', benchmark);
    document.getElementById('btn-clear').addEventListener('click', clearResults);
    
    // Enter key in input
    document.getElementById('batch-input').addEventListener('keypress', (e) => {
        if (e.key === 'Enter') calculate();
    });
}

// Parse input string to array of numbers
function parseInput() {
    const input = document.getElementById('batch-input').value;
    const indices = input
        .split(',')
        .map(s => s.trim())
        .filter(s => s && !isNaN(parseInt(s)))
        .map(s => parseInt(s, 10));
    return indices;
}

// Calculate Fibonacci numbers
function calculate() {
    const indices = parseInput();
    
    if (indices.length === 0) {
        alert('Please enter valid Fibonacci indices');
        return;
    }
    
    const results = fibBatch(indices);
    displayResults(indices, results);
}

// Display results in table
function displayResults(indices, results) {
    const section = document.getElementById('results-section');
    const table = document.getElementById('results-table');
    
    let html = `
        <table>
            <thead>
                <tr>
                    <th>Index (n)</th>
                    <th>F(n)</th>
                </tr>
            </thead>
            <tbody>
    `;
    
    indices.forEach((n, i) => {
        const result = results[i].toString();
        const displayResult = result.length > 30 
            ? result.substring(0, 15) + '...' + result.substring(result.length - 15)
            : result;
        
        html += `
            <tr>
                <td>${n}</td>
                <td title="${result}">${displayResult}</td>
            </tr>
        `;
    });
    
    html += '</tbody></table>';
    table.innerHTML = html;
    section.style.display = 'block';
}

// Run benchmark
function benchmark() {
    const indices = parseInput();
    
    if (indices.length === 0) {
        alert('Please enter valid Fibonacci indices');
        return;
    }
    
    const iterations = 100;
    
    // Time SIMD (simulated)
    const simdStart = performance.now();
    for (let i = 0; i < iterations; i++) {
        fibSimdBatch(indices);
    }
    const simdTime = (performance.now() - simdStart) / iterations;
    
    // Time scalar
    const scalarStart = performance.now();
    for (let i = 0; i < iterations; i++) {
        fibBatch(indices);
    }
    const scalarTime = (performance.now() - scalarStart) / iterations;
    
    displayBenchmark(simdTime, scalarTime);
}

// Display benchmark results
function displayBenchmark(simdTime, scalarTime) {
    const section = document.getElementById('benchmark-section');
    const speedup = scalarTime / simdTime;
    
    document.getElementById('simd-time').textContent = simdTime.toFixed(3) + ' ms';
    document.getElementById('scalar-time').textContent = scalarTime.toFixed(3) + ' ms';
    document.getElementById('speedup').textContent = speedup.toFixed(2) + 'x';
    
    section.style.display = 'block';
    
    // Draw simple bar chart
    drawChart(simdTime, scalarTime);
}

// Draw benchmark chart
function drawChart(simdTime, scalarTime) {
    const canvas = document.getElementById('benchmark-chart');
    const ctx = canvas.getContext('2d');
    
    // Set canvas size
    canvas.width = canvas.parentElement.clientWidth;
    canvas.height = 200;
    
    const maxTime = Math.max(simdTime, scalarTime);
    const barWidth = 80;
    const gap = 60;
    const startX = (canvas.width - (barWidth * 2 + gap)) / 2;
    const maxHeight = 150;
    const baseY = 180;
    
    // Clear canvas
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    
    // Draw SIMD bar
    const simdHeight = (simdTime / maxTime) * maxHeight;
    const simdGradient = ctx.createLinearGradient(0, baseY - simdHeight, 0, baseY);
    simdGradient.addColorStop(0, '#667eea');
    simdGradient.addColorStop(1, '#764ba2');
    ctx.fillStyle = simdGradient;
    ctx.fillRect(startX, baseY - simdHeight, barWidth, simdHeight);
    
    // Draw Scalar bar
    const scalarHeight = (scalarTime / maxTime) * maxHeight;
    const scalarGradient = ctx.createLinearGradient(0, baseY - scalarHeight, 0, baseY);
    scalarGradient.addColorStop(0, '#10b981');
    scalarGradient.addColorStop(1, '#059669');
    ctx.fillStyle = scalarGradient;
    ctx.fillRect(startX + barWidth + gap, baseY - scalarHeight, barWidth, scalarHeight);
    
    // Labels
    ctx.fillStyle = '#f1f5f9';
    ctx.font = '14px Inter, sans-serif';
    ctx.textAlign = 'center';
    ctx.fillText('SIMD', startX + barWidth / 2, baseY + 20);
    ctx.fillText('Scalar', startX + barWidth + gap + barWidth / 2, baseY + 20);
    
    // Time labels on bars
    ctx.fillText(simdTime.toFixed(2) + 'ms', startX + barWidth / 2, baseY - simdHeight - 10);
    ctx.fillText(scalarTime.toFixed(2) + 'ms', startX + barWidth + gap + barWidth / 2, baseY - scalarHeight - 10);
}

// Clear results
function clearResults() {
    document.getElementById('results-section').style.display = 'none';
    document.getElementById('benchmark-section').style.display = 'none';
    document.getElementById('results-table').innerHTML = '';
}
