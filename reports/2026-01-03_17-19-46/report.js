// Fibonacci Benchmark Report - JavaScript

document.addEventListener('DOMContentLoaded', function() {
    // Tab navigation
    const tabButtons = document.querySelectorAll('.tab-btn');
    const tabContents = document.querySelectorAll('.tab-content');

    tabButtons.forEach(button => {
        button.addEventListener('click', () => {
            const tabId = button.dataset.tab;
            
            tabButtons.forEach(btn => btn.classList.remove('active'));
            tabContents.forEach(content => content.classList.remove('active'));
            
            button.classList.add('active');
            document.getElementById(tabId).classList.add('active');
        });
    });

    // Load and visualize data
    loadComplexityData();
    loadBinetData();
    loadGoldenData();
});

async function loadComplexityData() {
    try {
        const response = await fetch('data/complexity_comparison.csv');
        const text = await response.text();
        const data = parseCSV(text);
        
        const ctx = document.getElementById('complexityChart').getContext('2d');
        new Chart(ctx, {
            type: 'line',
            data: {
                labels: data.map(row => row.n),
                datasets: [
                    {
                        label: 'Iterative (ns)',
                        data: data.map(row => row.iterative_ns),
                        borderColor: '#f59e0b',
                        backgroundColor: 'rgba(245, 158, 11, 0.1)',
                        fill: true,
                        tension: 0.4
                    },
                    {
                        label: 'Matrix (ns)',
                        data: data.map(row => row.matrix_ns),
                        borderColor: '#10b981',
                        backgroundColor: 'rgba(16, 185, 129, 0.1)',
                        fill: true,
                        tension: 0.4
                    }
                ]
            },
            options: {
                responsive: true,
                maintainAspectRatio: false,
                plugins: {
                    legend: {
                        labels: { color: '#f1f5f9' }
                    },
                    title: {
                        display: true,
                        text: 'Execution Time vs Input Size',
                        color: '#f1f5f9'
                    }
                },
                scales: {
                    x: {
                        title: { display: true, text: 'n', color: '#94a3b8' },
                        ticks: { color: '#94a3b8' },
                        grid: { color: '#334155' }
                    },
                    y: {
                        title: { display: true, text: 'Time (ns)', color: '#94a3b8' },
                        ticks: { color: '#94a3b8' },
                        grid: { color: '#334155' }
                    }
                }
            }
        });

        document.getElementById('complexityData').innerHTML = 
            `<strong>Data Points:</strong> ${data.length} measurements from n=10 to n=1000`;
    } catch (error) {
        console.error('Failed to load complexity data:', error);
        document.getElementById('complexityData').innerHTML = 
            '<em>Data file not found. Run generate_report script first.</em>';
    }
}

async function loadBinetData() {
    try {
        const response = await fetch('data/binet_accuracy.csv');
        const text = await response.text();
        const data = parseCSV(text);
        
        const ctx = document.getElementById('binetChart').getContext('2d');
        new Chart(ctx, {
            type: 'line',
            data: {
                labels: data.map(row => row.n),
                datasets: [
                    {
                        label: 'Relative Error',
                        data: data.map(row => parseFloat(row.rel_error) || 0),
                        borderColor: '#ef4444',
                        backgroundColor: 'rgba(239, 68, 68, 0.1)',
                        fill: true,
                        tension: 0.4
                    }
                ]
            },
            options: {
                responsive: true,
                maintainAspectRatio: false,
                plugins: {
                    legend: {
                        labels: { color: '#f1f5f9' }
                    },
                    title: {
                        display: true,
                        text: 'Binet Formula Relative Error',
                        color: '#f1f5f9'
                    }
                },
                scales: {
                    x: {
                        title: { display: true, text: 'n', color: '#94a3b8' },
                        ticks: { color: '#94a3b8' },
                        grid: { color: '#334155' }
                    },
                    y: {
                        title: { display: true, text: 'Relative Error', color: '#94a3b8' },
                        ticks: { color: '#94a3b8' },
                        grid: { color: '#334155' },
                        type: 'logarithmic'
                    }
                }
            }
        });

        const accurateCount = data.filter(row => parseFloat(row.rel_error) === 0).length;
        document.getElementById('binetData').innerHTML = 
            `<strong>Accuracy:</strong> Binet formula is exact for ${accurateCount} values (n ≤ 78)`;
    } catch (error) {
        console.error('Failed to load Binet data:', error);
        document.getElementById('binetData').innerHTML = 
            '<em>Data file not found. Run generate_report script first.</em>';
    }
}

async function loadGoldenData() {
    try {
        const response = await fetch('data/golden_ratio_convergence.csv');
        const text = await response.text();
        const data = parseCSV(text);
        
        const ctx = document.getElementById('goldenChart').getContext('2d');
        new Chart(ctx, {
            type: 'line',
            data: {
                labels: data.map(row => row.n),
                datasets: [
                    {
                        label: 'F(n+1)/F(n)',
                        data: data.map(row => parseFloat(row.ratio)),
                        borderColor: '#6366f1',
                        backgroundColor: 'rgba(99, 102, 241, 0.1)',
                        fill: true,
                        tension: 0.4
                    },
                    {
                        label: 'φ (Golden Ratio)',
                        data: data.map(() => 1.618033988749895),
                        borderColor: '#10b981',
                        borderDash: [5, 5],
                        fill: false
                    }
                ]
            },
            options: {
                responsive: true,
                maintainAspectRatio: false,
                plugins: {
                    legend: {
                        labels: { color: '#f1f5f9' }
                    },
                    title: {
                        display: true,
                        text: 'Convergence to Golden Ratio',
                        color: '#f1f5f9'
                    }
                },
                scales: {
                    x: {
                        title: { display: true, text: 'n', color: '#94a3b8' },
                        ticks: { color: '#94a3b8' },
                        grid: { color: '#334155' }
                    },
                    y: {
                        title: { display: true, text: 'Ratio', color: '#94a3b8' },
                        ticks: { color: '#94a3b8' },
                        grid: { color: '#334155' },
                        min: 1.5,
                        max: 2.0
                    }
                }
            }
        });

        const lastError = data[data.length - 1]?.error_from_phi;
        document.getElementById('goldenData').innerHTML = 
            `<strong>Convergence:</strong> Error from φ at n=50: ${parseFloat(lastError).toExponential(4)}`;
    } catch (error) {
        console.error('Failed to load golden ratio data:', error);
        document.getElementById('goldenData').innerHTML = 
            '<em>Data file not found. Run generate_report script first.</em>';
    }
}

function parseCSV(text) {
    const lines = text.trim().split('\n');
    const headers = lines[0].split(',');
    
    return lines.slice(1).map(line => {
        const values = line.split(',');
        const obj = {};
        headers.forEach((header, i) => {
            obj[header.trim()] = values[i]?.trim();
        });
        return obj;
    });
}
