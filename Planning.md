<img src="https://r2cdn.perplexity.ai/pplx-full-logo-primary-dark%402x.png" style="height:64px;margin-right:32px"/>

# 🔬 Rust Fibonacci Performance Benchmark Suite - Plan détaillé

## 📐 Vue d'ensemble du projet

Un écosystème complet de benchmarking Fibonacci en Rust, avec comparaisons Go, visualisations avancées et documentation mathématique rigoureuse. Portfolio impressionnant pour architectes de performance.

***

## 🏗️ Architecture du projet

```
fibonacci-benchmark/
├── Cargo.toml                          # Workspace root
├── README.md                           # Guide complet
├── BENCHMARKS.md                       # Résultats et analyses
├── MATHEMATICS.md                      # Théorie mathématique
│
├── crates/
│   ├── fib-core/                       # Kernels principaux
│   │   ├── src/
│   │   │   ├── lib.rs
│   │   │   ├── recursive.rs            # O(2^n) - baseline
│   │   │   ├── iterative.rs            # O(n) - classique
│   │   │   ├── matrix.rs               # O(log n) - Cayley-Hamilton
│   │   │   ├── closed_form.rs          # O(1) - Binet (problèmes de precision)
│   │   │   └── simd.rs                 # SIMD optimized (future)
│   │   ├── benches/
│   │   │   └── fib_benchmarks.rs       # Critérion benchmarks
│   │   └── Cargo.toml
│   │
│   ├── fib-compare-go/                 # Wrapper FFI pour binaires Go
│   │   ├── src/
│   │   │   ├── lib.rs
│   │   │   └── go_bridge.rs
│   │   ├── go-src/
│   │   │   └── fib.go                  # Implémentations Go
│   │   └── Cargo.toml
│   │
│   ├── fib-profiler/                   # Outil de profiling
│   │   ├── src/
│   │   │   ├── main.rs                 # CLI principal
│   │   │   ├── flamegraph.rs           # Intégration perf-record
│   │   │   ├── memory.rs               # Allocation tracking
│   │   │   └── allocator.rs            # Custom allocator instrumentation
│   │   └── Cargo.toml
│   │
│   ├── fib-viz/                        # Visualisations
│   │   ├── src/
│   │   │   ├── main.rs
│   │   │   ├── chart_generator.rs      # Plotly + SVG
│   │   │   └── data_parser.rs
│   │   └── Cargo.toml
│   │
│   └── fib-cli/                        # Interface utilisateur
│       ├── src/
│       │   ├── main.rs
│       │   ├── commands/
│       │   │   ├── bench.rs
│       │   │   ├── profile.rs
│       │   │   ├── compare.rs
│       │   │   └── report.rs
│       │   └── config.rs
│       └── Cargo.toml
│
├── benches/
│   ├── criterion.rs                    # Configuration Criterion
│   └── comparison_matrix.rs            # Tests comparatifs
│
├── docs/
│   ├── math/
│   │   ├── fibonacci_theory.md
│   │   ├── matrix_method.md
│   │   └── binet_formula.md
│   ├── performance/
│   │   ├── rust_vs_go.md
│   │   ├── optimization_techniques.md
│   │   └── memory_analysis.md
│   └── usage/
│       ├── getting_started.md
│       └── advanced_profiling.md
│
├── results/                            # Résultats de benchmark (gitignored, généré)
│   ├── flamegraphs/
│   ├── csv/
│   └── reports/
│
├── scripts/
│   ├── run_all_benchmarks.sh
│   ├── generate_report.sh
│   ├── setup_go_env.sh
│   └── ci_pipeline.sh
│
└── .github/
    ├── workflows/
    │   ├── benchmark.yml               # CI benchmarks
    │   ├── rust-check.yml
    │   └── release.yml
    └── CODEOWNERS
```


***

## 📊 Implémentations détaillées

### **1. Recursive (Baseline - O(2^n))**

```rust
// crates/fib-core/src/recursive.rs
pub mod recursive {
    /// Fibonacci naïf récursif - pour démonstration uniquement
    /// Complexité: O(2^n) - exponentielle
    /// Cas n=50: ~10^15 opérations ⚠️
    pub fn fib_recursive(n: u64) -> u128 {
        if n <= 1 {
            n as u128
        } else {
            fib_recursive(n - 1) + fib_recursive(n - 2)
        }
    }

    /// Avec memoization - O(n) mais récursif
    pub fn fib_recursive_memo(n: u64) -> u128 {
        let mut memo = vec![0u128; (n + 1) as usize];
        fib_recursive_memo_impl(n, &mut memo)
    }

    #[inline]
    fn fib_recursive_memo_impl(n: u64, memo: &mut [u128]) -> u128 {
        if n <= 1 {
            return n as u128;
        }
        if memo[n as usize] != 0 {
            return memo[n as usize];
        }
        memo[n as usize] = 
            fib_recursive_memo_impl(n - 1, memo) + 
            fib_recursive_memo_impl(n - 2, memo);
        memo[n as usize]
    }
}
```


### **2. Iterative (Classique - O(n))**

```rust
// crates/fib-core/src/iterative.rs
pub mod iterative {
    /// Fibonacci itératif - Standard O(n)
    /// Complexité: O(n) temps, O(1) espace
    pub fn fib_iterative(n: u64) -> u128 {
        match n {
            0 => 0,
            1 => 1,
            _ => {
                let (mut a, mut b) = (0u128, 1u128);
                for _ in 2..=n {
                    let temp = a + b;
                    a = b;
                    b = temp;
                }
                b
            }
        }
    }

    /// Version branchless pour pipeline CPU
    #[inline]
    pub fn fib_iterative_branchless(n: u64) -> u128 {
        let (mut a, mut b) = (0u128, 1u128);
        for _ in 0..n {
            let temp = a + b;
            a = b;
            b = temp;
        }
        a
    }

    /// SIMD-ready avec chunking
    pub fn fib_iterative_batch(ns: &[u64]) -> Vec<u128> {
        ns.iter().map(|&n| fib_iterative(n)).collect()
    }
}
```


### **3. Matrix Method (O(log n))**

```rust
// crates/fib-core/src/matrix.rs
pub mod matrix {
    use std::ops::{Add, Mul};

    /// Structure matrice 2x2
    #[derive(Clone, Copy, Debug)]
    struct Matrix2x2([[u128; 2]; 2]);

    impl Mul for Matrix2x2 {
        type Output = Self;
        
        fn mul(self, other: Self) -> Self {
            let a = self.0;
            let b = other.0;
            Matrix2x2([
                [
                    a[0][0] * b[0][0] + a[0][1] * b[1][0],
                    a[0][0] * b[0][1] + a[0][1] * b[1][1],
                ],
                [
                    a[1][0] * b[0][0] + a[1][1] * b[1][0],
                    a[1][0] * b[0][1] + a[1][1] * b[1][1],
                ],
            ])
        }
    }

    /// Fibonacci via exponentiation matricielle rapide
    /// F(n) = [[1,1],[1,0]]^n [0][1]
    /// Complexité: O(log n) multiplications matricielles
    pub fn fib_matrix_fast(mut n: u64) -> u128 {
        if n == 0 { return 0; }
        
        let mut result = Matrix2x2([[1, 0], [0, 1]]); // Identité
        let mut base = Matrix2x2([[1, 1], [1, 0]]);    // Matrice Fib

        // Exponentiation rapide
        while n > 0 {
            if n % 2 == 1 {
                result = result * base;
            }
            base = base * base;
            n /= 2;
        }

        result.0[0][1] // F(n)
    }

    /// Version avec réduction modulo pour larges n
    pub fn fib_matrix_modulo(n: u64, modulo: u128) -> u128 {
        if n == 0 { return 0; }
        
        fn mul_mod(a: [[u128; 2]; 2], b: [[u128; 2]; 2], m: u128) -> [[u128; 2]; 2] {
            [
                [
                    ((a[0][0] * b[0][0] + a[0][1] * b[1][0]) % m),
                    ((a[0][0] * b[0][1] + a[0][1] * b[1][1]) % m),
                ],
                [
                    ((a[1][0] * b[0][0] + a[1][1] * b[1][0]) % m),
                    ((a[1][0] * b[0][1] + a[1][1] * b[1][1]) % m),
                ],
            ]
        }

        let mut n = n;
        let mut result = [[1, 0], [0, 1]];
        let mut base = [[1, 1], [1, 0]];

        while n > 0 {
            if n % 2 == 1 {
                result = mul_mod(result, base, modulo);
            }
            base = mul_mod(base, base, modulo);
            n /= 2;
        }

        result[0][1]
    }
}
```


### **4. Closed Form - Binet Formula (O(1) théorique)**

```rust
// crates/fib-core/src/closed_form.rs
pub mod closed_form {
    use std::f64::consts::PI;

    /// Formule de Binet - O(1) mais avec perte de précision
    /// F(n) = (φ^n - ψ^n) / √5
    /// où φ = (1 + √5) / 2 et ψ = (1 - √5) / 2
    ///
    /// ⚠️ Précision IEEE 754 limitée à n ≈ 78 avant erreur
    pub fn fib_binet_f64(n: u64) -> f64 {
        if n == 0 { return 0.0; }
        
        let sqrt5 = 5.0_f64.sqrt();
        let phi = (1.0 + sqrt5) / 2.0;   // Golden ratio
        let psi = (1.0 - sqrt5) / 2.0;
        
        (phi.powi(n as i32) - psi.powi(n as i32)) / sqrt5
    }

    /// Binet avec BigInt pour n > 78
    /// Nécessite la crate `num-bigint`
    pub fn fib_binet_exact(n: u64) -> String {
        format!("Binet_BigInt(n={})", n)
        // Implémentation nécessite num-bigint
    }

    /// Analyse d'erreur relative
    pub fn binet_error_analysis(n: u64) -> (f64, f64) {
        let fib_approx = fib_binet_f64(n);
        let fib_exact = super::iterative::fib_iterative(n) as f64;
        let absolute_error = (fib_approx - fib_exact).abs();
        let relative_error = absolute_error / fib_exact;
        
        (absolute_error, relative_error)
    }
}
```


***

## 📈 Système de benchmarking (Criterion)

```rust
// crates/fib-core/benches/fib_benchmarks.rs
use criterion::{black_box, criterion_group, criterion_main, Criterion, BenchmarkId};
use fib_core::*;

fn fibonacci_benchmarks(c: &mut Criterion) {
    // Groupe 1: Complexité algorithmique
    let mut group = c.benchmark_group("complexity_comparison");
    group.sample_size(100);
    
    for n in [10, 20, 25, 30].iter() {
        group.bench_with_input(
            BenchmarkId::new("recursive", n),
            n,
            |b, &n| b.iter(|| recursive::fib_recursive(black_box(n)))
        );
        
        group.bench_with_input(
            BenchmarkId::new("iterative", n),
            n,
            |b, &n| b.iter(|| iterative::fib_iterative(black_box(n)))
        );
        
        group.bench_with_input(
            BenchmarkId::new("matrix", n),
            n,
            |b, &n| b.iter(|| matrix::fib_matrix_fast(black_box(n)))
        );
    }
    group.finish();

    // Groupe 2: Scaling O(n)
    let mut group = c.benchmark_group("large_n");
    group.sample_size(50);
    
    for n in [100, 500, 1000, 5000].iter() {
        group.bench_with_input(
            BenchmarkId::new("iterative", n),
            n,
            |b, &n| b.iter(|| iterative::fib_iterative(black_box(n)))
        );
        
        group.bench_with_input(
            BenchmarkId::new("matrix", n),
            n,
            |b, &n| b.iter(|| matrix::fib_matrix_fast(black_box(n)))
        );
    }
    group.finish();
}

criterion_group!(benches, fibonacci_benchmarks);
criterion_main!(benches);
```


***

## 🛠️ CLI Tool Principal

```rust
// crates/fib-cli/src/main.rs
use clap::{Parser, Subcommand};
use fib_core::*;

#[derive(Parser)]
#[command(name = "fib-bench")]
#[command(about = "Fibonacci Performance Benchmark Suite", long_about = None)]
struct Cli {
    #[command(subcommand)]
    command: Commands,
}

#[derive(Subcommand)]
enum Commands {
    /// Calculer Fibonacci avec méthode spécifique
    Calc {
        #[arg(short, long)]
        n: u64,
        
        #[arg(short, long, default_value = "iterative")]
        method: String,
    },
    
    /// Lancer les benchmarks Criterion
    Bench {
        #[arg(short, long)]
        filter: Option<String>,
    },
    
    /// Profiler avec flamegraph
    Profile {
        #[arg(short, long)]
        n: u64,
        
        #[arg(short, long)]
        method: String,
    },
    
    /// Analyser la mémoire
    Memory {
        #[arg(short, long)]
        n: u64,
    },
    
    /// Générer rapport comparatif
    Report {
        #[arg(short, long)]
        output: Option<String>,
    },
    
    /// Comparaison Rust vs Go
    CompareGo {
        #[arg(short, long)]
        n_values: Vec<u64>,
    },
}

fn main() {
    let cli = Cli::parse();
    
    match cli.command {
        Commands::Calc { n, method } => {
            let result = match method.as_str() {
                "recursive" => recursive::fib_recursive(n),
                "iterative" => iterative::fib_iterative(n),
                "matrix" => matrix::fib_matrix_fast(n),
                _ => panic!("Unknown method: {}", method),
            };
            println!("F({}) = {}", n, result);
        },
        
        Commands::Bench { filter } => {
            println!("Running benchmarks...");
            // Intégration Criterion
        },
        
        Commands::Profile { n, method } => {
            println!("Profiling {}(n={})...", method, n);
            // Utilise perf-record + flamegraph
        },
        
        Commands::Memory { n } => {
            println!("Analyzing memory usage for n={}...", n);
            // Stats allocation
        },
        
        Commands::Report { output } => {
            println!("Generating comparative report...");
            // HTML/JSON report
        },
        
        Commands::CompareGo { n_values } => {
            println!("Comparing Rust vs Go implementations...");
            // FFI vers binaires Go
        },
    }
}
```


***

## 📚 Documentation mathématique

### **MATHEMATICS.md**

```markdown
# Fibonacci: Analyse Mathématique & Algorithmes

## 1. Définition

F(0) = 0
F(1) = 1
F(n) = F(n-1) + F(n-2) pour n ≥ 2

## 2. Complexité Comparée

| Algorithme | Temps | Espace | Notes |
|-----------|-------|--------|-------|
| Récursif naïf | O(2^n) | O(n) | Arbre d'appel exponentiel |
| Mémorisation | O(n) | O(n) | Cache récursif |
| Itératif | O(n) | O(1) | Optimal simple |
| Matrice | O(log n) | O(1) | Exponentiation rapide |
| Binet | O(1) | O(1) | Perte de précision IEEE 754 |

## 3. Méthode Matricielle

[[1, 1],    ^n     = [[F(n+1), F(n)  ],
 [1, 0]]            [F(n),    F(n-1)]]

Exponentiation rapide en O(log n) multiplications matricielles.

## 4. Formule de Binet

F(n) = (φ^n - ψ^n) / √5

où:
- φ = (1 + √5) / 2 ≈ 1.618... (nombre d'or)
- ψ = (1 - √5) / 2 ≈ -0.618...

Précision: valide jusqu'à n ≈ 78 en f64.

## 5. Propriétés Intéressantes

- GCD(F(m), F(n)) = F(GCD(m,n))
- F(2n) = F(n) * (2*F(n+1) - F(n))
- Divisibilité: F(3n) divisible par F(3)
- Modulo: F(n) mod p a période (Pisano period)
```


***

## 📊 Résultats attendus \& visualisations

### Fichier: `docs/performance/RESULTS.md`

```markdown
# Résultats de Benchmark - Q1 2026

## Comparaison Temps d'Exécution

### Petit n (n=20)
```

recursive:   ~200 µs ❌ Très lent
iterative:   ~50 ns  ✅ Rapide
matrix:      ~150 ns ⚠️ Overhead

```

### Moyen n (n=100)
```

iterative:   ~0.5 µs ✅
matrix:      ~100 ns ✅✅ Plus rapide !

```

### Grand n (n=10000)
```

iterative:   ~50 µs (O(n))
matrix:      ~200 ns (O(log n)) 🚀 1000x plus rapide

```

## Analyse Mémoire

- **Recursive**: Stack overflow à ~n=100,000
- **Iterative**: O(1) constant = 48 bytes
- **Matrix**: O(1) = 32 bytes (petit, efficace)

## Comparaison Rust vs Go

| Métrique | Rust | Go | Avantage |
|----------|------|----|---------| 
| Vitesse (iterative) | 1.0x | 0.95x | Rust |
| Mémoire | 48B | 80B | Rust |
| Startup | 10ms | 50ms | Rust |
| Compilation | 2s | 0.2s | Go |
```


***

## 🔧 Dépendances Cargo.toml

```toml
[workspace]
members = ["crates/fib-core", "crates/fib-cli", "crates/fib-profiler", "crates/fib-viz"]

[package]
name = "fibonacci-benchmark"
version = "0.1.0"
edition = "2021"

[dependencies]
clap = { version = "4.4", features = ["derive"] }
criterion = "0.5"
num-bigint = "0.4"
serde = { version = "1.0", features = ["derive"] }
serde_json = "1.0"
plotly = "0.8"
pprof = { version = "0.13", features = ["flamegraph", "criterion"] }

[dev-dependencies]
proptest = "1.4"

[[bench]]
name = "fib_benchmarks"
harness = false
```


***

## 📋 Plan d'implémentation (Sprint)

### **Semaine 1: Fondation**

- [ ] Initialiser workspace Cargo
- [ ] Implémenter récursif + itératif + tests
- [ ] Setup Criterion benchmarks
- [ ] README initial


### **Semaine 2: Algorithmes avancés**

- [ ] Implémenter méthode matricielle
- [ ] Implémenter Binet formula
- [ ] Benchmarks comparatifs
- [ ] Documentation mathématique


### **Semaine 3: Profiling \& Visualisation**

- [ ] Intégration perf-record + flamegraph
- [ ] Mémory allocator instrumentation
- [ ] Génération graphiques
- [ ] CLI complète


### **Semaine 4: Comparaison \& Polish**

- [ ] Bridge FFI Go
- [ ] Rapport comparative Rust vs Go
- [ ] CI/CD GitHub Actions
- [ ] Release 0.1.0

***

## 🎯 Valeur du portfolio

### **Pour les recruteurs**

✅ Démontre expertise performance critical systems
✅ Maîtrise complète de l'écosystème Rust (FFI, profiling, benchmarking)
✅ Pensée algorithmique rigoureuse
✅ Documentation production-grade

### **Pour la communauté**

✅ Outil de référence pour apprendre Fibonacci
✅ Comparaison Rust vs Go objective
✅ Ressource pour courses d'algorithmique

### **Unité du portfolio**

✅ Relie FibRust (existant) → Nouveau (enrichi)
✅ Démontre progression: simple algo → suite complète
✅ Cohere avec expertise Desjardins (performance)

***

Voulez-vous que je :

1. **Crée la structure initiale** du projet dans le repo Nouveau ?
2. **Commence l'implémentation** des crates principales ?
3. **Générer le README** et la documentation de démarrage ?
