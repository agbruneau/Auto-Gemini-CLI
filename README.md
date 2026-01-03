# 🔬 Fibonacci Performance Benchmark Suite

[![Rust](https://img.shields.io/badge/rust-1.70%2B-orange.svg)](https://www.rust-lang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![CI](https://github.com/agbru/FibBenchmark/actions/workflows/rust-check.yml/badge.svg)](https://github.com/agbru/FibBenchmark/actions)

> Un écosystème complet de benchmarking des algorithmes Fibonacci en Rust, avec analyses de complexité, visualisations et documentation mathématique rigoureuse.

## ✨ Caractéristiques

- **6 algorithmes Fibonacci** avec différentes complexités temporelles
- **Benchmarking Criterion** pour des mesures précises et statistiquement rigoureuses
- **CLI complète** avec 6 commandes pour calculs, comparaisons et analyses
- **Documentation exhaustive** : architecture, benchmarks, mathématiques et planification
- **Analyses de précision** pour la formule de Binet
- **Workspace modulaire** avec 4 crates spécialisés
- **CI/CD automatisé** avec tests et benchmarks

## 📊 Algorithmes Implémentés

| Algorithme | Temps | Espace | Cas d'usage |
|------------|-------|--------|-------------|
| Récursif naïf | O(2ⁿ) | O(n) | Démonstration uniquement |
| Récursif + Mémo | O(n) | O(n) | Petits n avec cache |
| Itératif | O(n) | O(1) | Usage général |
| Itératif branchless | O(n) | O(1) | Optimisation micro |
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

Le projet fournit une interface en ligne de commande complète via `fib-bench` :

```bash
# Calculer F(n) avec la méthode par défaut (itérative)
cargo run --bin fib-bench -- calc -n 100

# Calculer avec une méthode spécifique et afficher le temps
cargo run --bin fib-bench -- calc -n 50 --method matrix --time

# Comparer toutes les méthodes pour un n donné
cargo run --bin fib-bench -- compare -n 30

# Afficher les informations détaillées sur les algorithmes
cargo run --bin fib-bench -- info --method all

# Générer une séquence de Fibonacci
cargo run --bin fib-bench -- sequence --count 20

# Analyser la précision de la formule de Binet
cargo run --bin fib-bench -- binet-analysis --max-n 100

# Lancer les benchmarks Criterion
cargo run --bin fib-bench -- bench
```

**Commandes disponibles :**
- `calc` - Calculer F(n) avec une méthode spécifique
- `compare` - Comparer toutes les méthodes pour un n donné
- `bench` - Lancer les benchmarks Criterion
- `info` - Afficher les informations sur les algorithmes
- `sequence` - Générer une séquence de Fibonacci
- `binet-analysis` - Analyser la précision de la formule de Binet

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
FibBenchmark/
├── Cargo.toml                    # Workspace root
├── README.md                     # Ce fichier
├── LICENSE                       # MIT License
├── rust-toolchain.toml           # Version Rust
│
├── crates/
│   ├── fib-core/                 # 🧮 Bibliothèque principale
│   │   ├── src/
│   │   │   ├── lib.rs            # Point d'entrée + FibMethod enum
│   │   │   ├── recursive.rs     # O(2^n) + O(n) mémorisé
│   │   │   ├── iterative.rs     # O(n) + branchless + cache
│   │   │   ├── matrix.rs        # O(log n) + modulo + doubling
│   │   │   └── closed_form.rs   # O(1) Binet + analyse
│   │   └── benches/
│   │       └── fib_benchmarks.rs # Benchmarks Criterion
│   │
│   ├── fib-cli/                  # 🖥️ Interface CLI
│   │   └── src/
│   │       ├── main.rs
│   │       └── commands/
│   │           ├── calc.rs
│   │           ├── compare.rs
│   │           ├── bench.rs
│   │           ├── info.rs
│   │           ├── sequence.rs
│   │           └── binet_analysis.rs
│   │
│   ├── fib-profiler/             # 📊 Outil de profiling
│   │   └── src/main.rs
│   │
│   └── fib-viz/                  # 📈 Visualisations
│       └── src/main.rs
│
├── docs/                         # 📚 Documentation complète
│   ├── ARCHITECTURE.md           # Architecture technique détaillée
│   ├── BENCHMARKS.md             # Résultats et analyses de performance
│   ├── MATHEMATICS.md            # Théorie mathématique complète
│   ├── PLANNING.md               # Planification et roadmap
│   ├── math/
│   │   ├── fibonacci_theory.md
│   │   ├── matrix_method.md
│   │   └── binet_formula.md
│   ├── performance/
│   │   └── optimization_techniques.md
│   └── usage/
│       └── getting_started.md
│
├── .github/
│   └── workflows/
│       ├── rust-check.yml        # CI tests
│       └── benchmark.yml         # CI benchmarks
│
└── target/                       # Artifacts de compilation (gitignored)
```

## 📈 Benchmarks

Le projet utilise [Criterion.rs](https://github.com/bheisler/criterion.rs) pour des benchmarks statistiquement rigoureux avec détection de régressions.

### Exécution des benchmarks

```bash
# Tous les benchmarks
cargo bench

# Filtrer par nom de groupe
cargo bench -- complexity_comparison

# Benchmark spécifique
cargo bench -- matrix

# Avec baseline pour comparaison
cargo bench -- --save-baseline main
cargo bench -- --baseline main

# Via CLI
cargo run --bin fib-bench -- bench
```

### Groupes de benchmarks

Le projet inclut 6 groupes de benchmarks Criterion :

1. **complexity_comparison** - Comparaison des complexités algorithmiques
2. **large_n** - Scaling pour grands n
3. **iterative_variants** - Comparaison des variantes itératives
4. **batch_operations** - Opérations par lot
5. **cache_vs_direct** - Cache vs calcul direct
6. **modular_arithmetic** - Opérations modulo

### Rapports

Les rapports HTML détaillés sont générés dans `target/criterion/report/index.html` après chaque exécution.

### Résultats typiques

Voir [**BENCHMARKS.md**](docs/BENCHMARKS.md) pour des résultats détaillés. Exemples :

```
complexity_comparison/matrix/100      time: [45 ns 46 ns 47 ns]
complexity_comparison/iterative/100   time: [120 ns 122 ns 125 ns]

large_n/matrix/10000               time: [180 ns 185 ns 190 ns]
large_n/iterative/10000              time: [12 µs 12.5 µs 13 µs]
```

Le speedup de la méthode matricielle augmente avec n (O(log n) vs O(n)).

## 📚 Documentation

Le projet inclut une documentation exhaustive organisée en plusieurs sections :

### Documentation principale

- [**ARCHITECTURE.md**](docs/ARCHITECTURE.md) - Architecture technique complète, patterns, API et décisions techniques
- [**MATHEMATICS.md**](docs/MATHEMATICS.md) - Théorie mathématique approfondie de Fibonacci
- [**BENCHMARKS.md**](docs/BENCHMARKS.md) - Résultats de benchmarks et analyses de performance
- [**PLANNING.md**](docs/PLANNING.md) - Planification du projet, roadmap et état d'avancement

### Documentation détaillée

**Mathématiques :**
- [**fibonacci_theory.md**](docs/math/fibonacci_theory.md) - Fondements théoriques
- [**matrix_method.md**](docs/math/matrix_method.md) - Méthode matricielle expliquée
- [**binet_formula.md**](docs/math/binet_formula.md) - Formule de Binet et limites de précision

**Performance :**
- [**optimization_techniques.md**](docs/performance/optimization_techniques.md) - Techniques d'optimisation

**Usage :**
- [**getting_started.md**](docs/usage/getting_started.md) - Guide de démarrage rapide

### Documentation générée

```bash
# Générer la documentation complète
cargo doc --open

# Documentation pour un crate spécifique
cargo doc -p fib-core --open
```

## 🧪 Tests

```bash
# Exécuter tous les tests
cargo test

# Tests avec output
cargo test -- --nocapture

# Tests d'un crate spécifique
cargo test -p fib-core
```

## 🎯 État du Projet

**Version actuelle :** 0.1.0

### Phases complétées ✅

- ✅ **Phase 1** - Fondation : Structure workspace, algorithmes de base
- ✅ **Phase 2** - Algorithmes avancés : Matrice, Binet, utilitaires
- ✅ **Phase 3** - CLI & Outils : Interface complète avec 6 commandes
- 🔄 **Phase 4** - Documentation : En cours (80% complété)

### Prochaines étapes

Voir [**PLANNING.md**](docs/PLANNING.md) pour la roadmap complète :
- Phase 5 : Profiling avancé (flamegraph, mémoire)
- Phase 6 : Visualisations (graphiques, rapports HTML)
- Phase 7 : Comparaison Go (FFI bridge) - Optionnel
- Phase 8 : Optimisations SIMD - Optionnel
- Phase 9 : Publication crates.io

## 🤝 Contribution

Les contributions sont les bienvenues ! 

1. Fork le projet
2. Créer une branche (`git checkout -b feature/amazing-feature`)
3. S'assurer que les tests passent : `cargo test`
4. Vérifier le formatage : `cargo fmt --check`
5. Vérifier les lints : `cargo clippy -- -D warnings`
6. Commit les changements (`git commit -m 'Add amazing feature'`)
7. Push (`git push origin feature/amazing-feature`)
8. Ouvrir une Pull Request

### Standards de code

- Formatage : `cargo fmt`
- Linting : `cargo clippy -- -D warnings`
- Tests : Tous les tests doivent passer
- Documentation : Doc-tests pour les exemples publics

## 🔍 Exemples d'Utilisation

### Calcul simple

```rust
use fib_core::{iterative, matrix, FibMethod};

// Calcul direct
let fib_100 = iterative::fib_iterative(100);
assert_eq!(fib_100, 354224848179261915075);

// Via enum
let method = FibMethod::Matrix;
let result = method.calculate(1000);
```

### Comparaison de méthodes

```rust
use fib_core::FibMethod;

let n = 50;
let methods = [
    FibMethod::Iterative,
    FibMethod::Matrix,
    FibMethod::Binet,
];

for method in methods {
    let result = method.calculate(n);
    println!("{}: {} (complexity: {})", 
        method.name(), 
        result, 
        method.time_complexity()
    );
}
```

### Cache pour calculs répétés

```rust
use fib_core::iterative::FibonacciCache;

let mut cache = FibonacciCache::new(100);
// Le cache peut être réutilisé pour plusieurs calculs
```

## 🛡️ Limitations et Notes

- **u128 overflow** : F(186) est le dernier Fibonacci qui tient dans u128
- **Binet précision** : Limité à n ≤ 78 pour une précision exacte
- **Récursif naïf** : Extrêmement lent pour n > 35, à utiliser uniquement à des fins pédagogiques
- **Stack overflow** : La récursion peut causer un stack overflow pour n > 100,000 (selon la taille de stack)

## 🙏 Remerciements

- [Criterion.rs](https://github.com/bheisler/criterion.rs) pour le framework de benchmarking statistiquement rigoureux
- [clap](https://github.com/clap-rs/clap) pour l'excellente bibliothèque CLI
- La communauté Rust pour les outils et le support

## 📜 Licence

Ce projet est sous licence MIT. Voir [LICENSE](LICENSE) pour plus de détails.

---

<p align="center">
  Fait avec ❤️ et 🦀<br>
  <em>Un projet démontrant l'excellence en ingénierie Rust</em>
</p>
