# 🔬 Fibonacci Performance Benchmark Suite

[![Rust](https://img.shields.io/badge/rust-1.70%2B-orange.svg)](https://www.rust-lang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![CI](https://github.com/agbru/FibBenchmark/actions/workflows/rust-check.yml/badge.svg)](https://github.com/agbru/FibBenchmark/actions)

> Un écosystème complet de benchmarking des algorithmes Fibonacci en Rust, avec analyses de complexité, visualisations et documentation mathématique rigoureuse.

## ✨ Caractéristiques

- **5 algorithmes Fibonacci** avec différentes complexités temporelles
- **Benchmarking Criterion** pour des mesures précises
- **CLI complète** pour calculs et comparaisons
- **Documentation mathématique** détaillée
- **Analyses de précision** pour la formule de Binet
- **Génération de données** pour visualisations

## 📊 Algorithmes Implémentés

| Algorithme | Temps | Espace | Cas d'usage |
|------------|-------|--------|-------------|
| Récursif naïf | O(2ⁿ) | O(n) | Démonstration uniquement |
| Récursif + Mémo | O(n) | O(n) | Petits n avec cache |
| Itératif | O(n) | O(1) | Usage général |
| Matriciel | O(log n) | O(1) | Grands n |
| Binet | O(1) | O(1) | Approximation (n ≤ 78) |

## 🚀 Installation

### Prérequis

- Rust 1.70+ ([rustup](https://rustup.rs/))
- Cargo (inclus avec Rust)

### Compilation

```bash
# Cloner le repository
git clone https://github.com/agbru/FibBenchmark.git
cd FibBenchmark

# Compiler en mode release
cargo build --release

# Exécuter les tests
cargo test

# Lancer les benchmarks
cargo bench
```

## 🛠️ Utilisation

### CLI Tool

```bash
# Calculer F(n) avec la méthode par défaut (itérative)
cargo run --bin fib-bench -- calc -n 100

# Calculer avec une méthode spécifique
cargo run --bin fib-bench -- calc -n 50 --method matrix --time

# Comparer toutes les méthodes
cargo run --bin fib-bench -- compare -n 30

# Afficher les informations sur les algorithmes
cargo run --bin fib-bench -- info --method all

# Générer une séquence de Fibonacci
cargo run --bin fib-bench -- sequence --count 20

# Analyser la précision de Binet
cargo run --bin fib-bench -- binet-analysis --max-n 100
```

### Comme bibliothèque

```rust
use fib_core::{iterative, matrix, FibMethod};

// Calcul simple
let fib_100 = iterative::fib_iterative(100);
assert_eq!(fib_100, 354224848179261915075);

// Méthode matricielle pour grands n
let fib_1000 = matrix::fib_matrix_fast(1000);

// Via l'enum FibMethod
let method = FibMethod::Matrix;
let result = method.calculate(100);
```

## 📁 Structure du Projet

```
fibonacci-benchmark/
├── Cargo.toml                    # Workspace root
├── README.md                     # Ce fichier
├── BENCHMARKS.md                 # Résultats de benchmark
├── MATHEMATICS.md                # Théorie mathématique
│
├── crates/
│   ├── fib-core/                 # Kernels principaux
│   │   ├── src/
│   │   │   ├── lib.rs            # Point d'entrée
│   │   │   ├── recursive.rs      # O(2^n) - baseline
│   │   │   ├── iterative.rs      # O(n) - classique
│   │   │   ├── matrix.rs         # O(log n) - exponentiation rapide
│   │   │   └── closed_form.rs    # O(1) - Binet
│   │   └── benches/              # Benchmarks Criterion
│   │
│   ├── fib-cli/                  # Interface utilisateur
│   │   └── src/
│   │       ├── main.rs
│   │       └── commands/
│   │
│   ├── fib-profiler/             # Outil de profiling
│   └── fib-viz/                  # Visualisations
│
├── docs/                         # Documentation étendue
│   ├── math/
│   ├── performance/
│   └── usage/
│
└── results/                      # Données générées (gitignored)
```

## 📈 Benchmarks

Exécuter les benchmarks Criterion :

```bash
# Tous les benchmarks
cargo bench

# Filtrer par nom
cargo bench -- complexity_comparison

# Avec profiling flamegraph
cargo bench --bench fib_benchmarks -- --profile-time 5
```

Les rapports HTML sont générés dans `target/criterion/report/index.html`.

### Résultats Typiques

```
complexity_comparison/matrix/100   time: [45 ns 46 ns 47 ns]
complexity_comparison/iterative/100 time: [120 ns 122 ns 125 ns]

large_n/matrix/10000               time: [180 ns 185 ns 190 ns]
large_n/iterative/10000            time: [12 µs 12.5 µs 13 µs]
```

## 📚 Documentation

- [**MATHEMATICS.md**](MATHEMATICS.md) - Théorie mathématique de Fibonacci
- [**BENCHMARKS.md**](BENCHMARKS.md) - Analyses de performance détaillées
- [**docs/usage/getting_started.md**](docs/usage/getting_started.md) - Guide de démarrage
- [**docs/math/matrix_method.md**](docs/math/matrix_method.md) - Explication de la méthode matricielle

## 🧪 Tests

```bash
# Exécuter tous les tests
cargo test

# Tests avec output
cargo test -- --nocapture

# Tests d'un crate spécifique
cargo test -p fib-core
```

## 🤝 Contribution

Les contributions sont les bienvenues ! Voir [CONTRIBUTING.md](CONTRIBUTING.md) pour les détails.

1. Fork le projet
2. Créer une branche (`git checkout -b feature/amazing-feature`)
3. Commit les changements (`git commit -m 'Add amazing feature'`)
4. Push (`git push origin feature/amazing-feature`)
5. Ouvrir une Pull Request

## 📜 Licence

Ce projet est sous licence MIT. Voir [LICENSE](LICENSE) pour plus de détails.

## 🙏 Remerciements

- [Criterion.rs](https://github.com/bheisler/criterion.rs) pour le framework de benchmarking
- La communauté Rust pour les excellents outils

---

<p align="center">
  Fait avec ❤️ et 🦀
</p>
