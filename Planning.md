# 🔬 Fibonacci Performance Benchmark Suite

## 📋 Planification d'Implémentation Exhaustive

> **Version**: 1.0.0  
> **Dernière mise à jour**: Janvier 2026  
> **Statut**: Phase 1 - Fondation ✅ Complétée

---

## 📑 Table des matières

1. [Vue d'ensemble](#-vue-densemble)
2. [État actuel du projet](#-état-actuel-du-projet)
3. [Architecture technique](#-architecture-technique)
4. [Phases de développement](#-phases-de-développement)
5. [Planification détaillée par sprint](#-planification-détaillée-par-sprint)
6. [Dépendances et ordre d'exécution](#-dépendances-et-ordre-dexécution)
7. [Risques et mitigations](#-risques-et-mitigations)
8. [Métriques de succès](#-métriques-de-succès)
9. [Ressources et outils](#-ressources-et-outils)
10. [Annexes techniques](#-annexes-techniques)

---

## 🎯 Vue d'ensemble

### Mission du projet

Créer un **écosystème complet de benchmarking** des algorithmes Fibonacci en Rust, servant à la fois de :
- 🎓 **Ressource pédagogique** pour l'apprentissage algorithmique
- 📊 **Outil de référence** pour les comparaisons de performance
- 🏆 **Projet portfolio** démontrant l'expertise en systèmes critiques

### Objectifs principaux

| Objectif | Description | Priorité |
|----------|-------------|----------|
| **Performance** | Implémenter 5+ algorithmes avec complexités variées | P0 |
| **Benchmarking** | Mesures précises avec Criterion.rs | P0 |
| **Documentation** | Théorie mathématique complète | P1 |
| **Comparaison** | Rust vs Go avec FFI | P2 |
| **Visualisation** | Graphiques et rapports automatisés | P2 |
| **CI/CD** | Pipeline de test et benchmark automatisé | P1 |

### Livrables finaux

1. ✅ **fib-core** - Bibliothèque d'algorithmes Fibonacci
2. ✅ **fib-cli** - Interface ligne de commande complète
3. ✅ **fib-profiler** - Outil de profiling de performance
4. ✅ **fib-viz** - Générateur de visualisations
5. ⬜ **fib-compare-go** - Bridge FFI Rust/Go (futur)
6. ✅ **Documentation** - README, guides, théorie mathématique

---

## 📊 État actuel du projet

### Composants complétés ✅

```
✅ Structure workspace Cargo
✅ fib-core avec 5 algorithmes
   ├── recursive.rs (O(2^n) + memoization O(n))
   ├── iterative.rs (O(n) + branchless + cache + iterator)
   ├── matrix.rs (O(log n) + modulo + doubling)
   └── closed_form.rs (O(1) Binet + analyse d'erreur)
✅ fib-cli avec 6 commandes
   ├── calc, compare, bench
   └── info, sequence, binet-analysis
✅ fib-profiler (base)
✅ fib-viz (base)
✅ Benchmarks Criterion (6 groupes)
✅ Tests unitaires (25) + doc-tests (18)
✅ Documentation
   ├── README.md
   ├── MATHEMATICS.md
   ├── BENCHMARKS.md
   └── docs/ (5 fichiers)
✅ CI/CD GitHub Actions
```

### Composants en cours 🔄

```
🔄 Scripts d'automatisation
🔄 Génération de rapports HTML
🔄 Intégration flamegraph (Unix)
```

### Composants planifiés ⬜

```
⬜ fib-compare-go (bridge FFI)
⬜ simd.rs (optimisation SIMD)
⬜ Interface web (optionnel)
⬜ Publication crates.io
```

---

## 🏗️ Architecture technique

### Structure complète du projet

```
fibonacci-benchmark/
├── Cargo.toml                          # Workspace root
├── Cargo.lock
├── rust-toolchain.toml
├── README.md                           # Guide complet
├── BENCHMARKS.md                       # Résultats et analyses
├── MATHEMATICS.md                      # Théorie mathématique
├── LICENSE                             # MIT
├── .gitignore
│
├── crates/
│   ├── fib-core/                       # 🧮 Bibliothèque principale
│   │   ├── Cargo.toml
│   │   ├── src/
│   │   │   ├── lib.rs                  # Point d'entrée + FibMethod enum
│   │   │   ├── recursive.rs            # O(2^n) + O(n) mémorisé
│   │   │   ├── iterative.rs            # O(n) + cache + iterator
│   │   │   ├── matrix.rs               # O(log n) + modulo + doubling
│   │   │   ├── closed_form.rs          # O(1) Binet + analyse
│   │   │   └── simd.rs                 # [FUTUR] SIMD optimisé
│   │   └── benches/
│   │       └── fib_benchmarks.rs       # Criterion benchmarks
│   │
│   ├── fib-cli/                        # 🖥️ Interface CLI
│   │   ├── Cargo.toml
│   │   └── src/
│   │       ├── main.rs
│   │       └── commands/
│   │           ├── mod.rs
│   │           ├── calc.rs
│   │           ├── compare.rs
│   │           ├── bench.rs
│   │           ├── info.rs
│   │           ├── sequence.rs
│   │           ├── binet_analysis.rs
│   │           ├── profile.rs          # [FUTUR]
│   │           └── report.rs           # [FUTUR]
│   │
│   ├── fib-profiler/                   # 📊 Profiling
│   │   ├── Cargo.toml
│   │   └── src/
│   │       ├── main.rs
│   │       ├── flamegraph.rs           # [FUTUR] Unix only
│   │       ├── memory.rs               # [FUTUR]
│   │       └── allocator.rs            # [FUTUR]
│   │
│   ├── fib-viz/                        # 📈 Visualisations
│   │   ├── Cargo.toml
│   │   └── src/
│   │       ├── main.rs
│   │       ├── chart_generator.rs      # [FUTUR]
│   │       └── data_parser.rs          # [FUTUR]
│   │
│   └── fib-compare-go/                 # [FUTUR] Bridge FFI Go
│       ├── Cargo.toml
│       ├── build.rs
│       ├── src/
│       │   ├── lib.rs
│       │   └── go_bridge.rs
│       └── go-src/
│           └── fib.go
│
├── docs/
│   ├── math/
│   │   ├── fibonacci_theory.md
│   │   ├── matrix_method.md
│   │   └── binet_formula.md
│   ├── performance/
│   │   ├── optimization_techniques.md
│   │   ├── rust_vs_go.md               # [FUTUR]
│   │   └── memory_analysis.md          # [FUTUR]
│   └── usage/
│       ├── getting_started.md
│       └── advanced_profiling.md       # [FUTUR]
│
├── results/                            # Généré, gitignored
│   ├── csv/
│   ├── flamegraphs/
│   └── reports/
│
├── scripts/
│   ├── run_all_benchmarks.sh           # [FUTUR]
│   ├── generate_report.sh              # [FUTUR]
│   ├── setup_go_env.sh                 # [FUTUR]
│   └── ci_pipeline.sh                  # [FUTUR]
│
└── .github/
    ├── workflows/
    │   ├── rust-check.yml              # ✅ CI tests
    │   ├── benchmark.yml               # ✅ CI benchmarks
    │   └── release.yml                 # [FUTUR]
    └── CODEOWNERS                      # [FUTUR]
```

### Diagramme de dépendances des crates

```
                    ┌─────────────┐
                    │  fib-core   │ (bibliothèque)
                    └──────┬──────┘
                           │
           ┌───────────────┼───────────────┐
           │               │               │
           ▼               ▼               ▼
    ┌────────────┐  ┌────────────┐  ┌────────────┐
    │  fib-cli   │  │fib-profiler│  │  fib-viz   │
    └────────────┘  └────────────┘  └────────────┘
           │
           ▼
    ┌─────────────────┐
    │ fib-compare-go  │ (optionnel)
    └─────────────────┘
```

---

## 📅 Phases de développement

### Phase 1: Fondation ✅ COMPLÉTÉE

**Durée**: 1 semaine  
**Statut**: ✅ 100% complété

| Tâche | Statut | Temps estimé | Temps réel |
|-------|--------|--------------|------------|
| Initialiser workspace Cargo | ✅ | 1h | 30min |
| Implémenter recursive.rs | ✅ | 2h | 1h |
| Implémenter iterative.rs | ✅ | 2h | 1.5h |
| Tests unitaires de base | ✅ | 2h | 1h |
| Setup Criterion benchmarks | ✅ | 3h | 2h |
| README initial | ✅ | 2h | 1h |
| **Total Phase 1** | ✅ | **12h** | **7h** |

### Phase 2: Algorithmes avancés ✅ COMPLÉTÉE

**Durée**: 1 semaine  
**Statut**: ✅ 100% complété

| Tâche | Statut | Temps estimé | Temps réel |
|-------|--------|--------------|------------|
| Implémenter matrix.rs | ✅ | 4h | 3h |
| Implémenter closed_form.rs | ✅ | 3h | 2h |
| FibMethod enum avec traits | ✅ | 2h | 1h |
| FibonacciCache + Iterator | ✅ | 2h | 1.5h |
| Tests exhaustifs | ✅ | 3h | 2h |
| Documentation mathématique | ✅ | 4h | 3h |
| **Total Phase 2** | ✅ | **18h** | **12.5h** |

### Phase 3: CLI & Outils ✅ COMPLÉTÉE

**Durée**: 1 semaine  
**Statut**: ✅ 100% complété

| Tâche | Statut | Temps estimé | Temps réel |
|-------|--------|--------------|------------|
| CLI structure avec clap | ✅ | 2h | 1.5h |
| Commande calc | ✅ | 1h | 30min |
| Commande compare | ✅ | 2h | 1h |
| Commande info | ✅ | 1h | 30min |
| Commande sequence | ✅ | 1h | 30min |
| Commande binet-analysis | ✅ | 2h | 1h |
| fib-profiler base | ✅ | 2h | 1.5h |
| fib-viz base | ✅ | 2h | 1.5h |
| **Total Phase 3** | ✅ | **13h** | **8h** |

### Phase 4: Documentation & CI 🔄 EN COURS

**Durée**: 1 semaine  
**Statut**: 🔄 80% complété

| Tâche | Statut | Temps estimé | Dépendance |
|-------|--------|--------------|------------|
| MATHEMATICS.md complet | ✅ | 4h | - |
| BENCHMARKS.md | ✅ | 2h | Phase 1 |
| docs/math/*.md | ✅ | 4h | - |
| docs/usage/getting_started.md | ✅ | 2h | Phase 3 |
| docs/performance/optimization.md | ✅ | 2h | Phase 2 |
| GitHub Actions CI | ✅ | 2h | - |
| GitHub Actions benchmarks | ✅ | 2h | Phase 1 |
| Cleanup et polish | 🔄 | 3h | Tout |
| **Total Phase 4** | 🔄 | **21h** | - |

### Phase 5: Profiling avancé ⬜ PLANIFIÉE

**Durée**: 1-2 semaines  
**Statut**: ⬜ Non démarrée  
**Prérequis**: Phase 4 complétée

| Tâche | Priorité | Temps estimé | Dépendance |
|-------|----------|--------------|------------|
| Intégration flamegraph (Unix) | P1 | 4h | pprof |
| Memory allocator instrumentation | P2 | 6h | - |
| Commande CLI profile | P1 | 3h | flamegraph |
| Commande CLI memory | P2 | 2h | allocator |
| Custom allocator tracking | P2 | 4h | - |
| docs/usage/advanced_profiling.md | P1 | 3h | - |
| **Total Phase 5** | ⬜ | **22h** | - |

### Phase 6: Visualisations ⬜ PLANIFIÉE

**Durée**: 1-2 semaines  
**Statut**: ⬜ Non démarrée  
**Prérequis**: Phase 4 complétée

| Tâche | Priorité | Temps estimé | Dépendance |
|-------|----------|--------------|------------|
| chart_generator.rs avec Plotly | P1 | 6h | plotly |
| data_parser.rs pour CSV | P1 | 3h | - |
| Génération SVG des graphiques | P2 | 4h | - |
| Rapport HTML automatisé | P1 | 5h | - |
| Comparaison visuelle algorithmes | P1 | 3h | - |
| Intégration avec benchmark CI | P2 | 3h | Phase 4 |
| **Total Phase 6** | ⬜ | **24h** | - |

### Phase 7: Comparaison Go ⬜ PLANIFIÉE

**Durée**: 2 semaines  
**Statut**: ⬜ Non démarrée  
**Prérequis**: Phase 5 complétée

| Tâche | Priorité | Temps estimé | Dépendance |
|-------|----------|--------------|------------|
| Implémentations Go (fib.go) | P2 | 4h | Go installé |
| FFI bridge avec cgo | P2 | 8h | - |
| Build script (build.rs) | P2 | 4h | - |
| Benchmarks comparatifs | P2 | 4h | - |
| Commande CLI compare-go | P2 | 3h | - |
| docs/performance/rust_vs_go.md | P2 | 4h | - |
| Scripts setup Go environment | P3 | 2h | - |
| **Total Phase 7** | ⬜ | **29h** | - |

### Phase 8: SIMD & Optimisations ⬜ OPTIONNELLE

**Durée**: 1-2 semaines  
**Statut**: ⬜ Non démarrée  
**Prérequis**: Phase 6 complétée

| Tâche | Priorité | Temps estimé | Dépendance |
|-------|----------|--------------|------------|
| simd.rs avec std::simd | P3 | 8h | nightly |
| Batch processing SIMD | P3 | 4h | - |
| Benchmarks SIMD | P3 | 3h | - |
| AVX2/AVX512 variants | P3 | 6h | - |
| Documentation SIMD | P3 | 3h | - |
| **Total Phase 8** | ⬜ | **24h** | - |

### Phase 9: Publication & Release ⬜ FINALE

**Durée**: 1 semaine  
**Statut**: ⬜ Non démarrée  
**Prérequis**: Phases 1-6 complétées

| Tâche | Priorité | Temps estimé | Dépendance |
|-------|----------|--------------|------------|
| Audit sécurité (cargo-audit) | P1 | 2h | - |
| Licence vérification | P1 | 1h | - |
| README final polish | P1 | 2h | - |
| CHANGELOG.md | P1 | 2h | - |
| Version tagging | P1 | 1h | - |
| Publication crates.io | P1 | 2h | - |
| GitHub Release | P1 | 2h | - |
| Annonce communauté | P3 | 1h | - |
| **Total Phase 9** | ⬜ | **13h** | - |

---

## 📆 Planification détaillée par sprint

### Sprint 1 (Semaine 1) ✅ COMPLÉTÉ

**Objectif**: Fondation solide

```
Jour 1-2: Setup & Structure
├── [x] Créer workspace Cargo
├── [x] Configurer Cargo.toml (workspace, profiles)
├── [x] Structure des crates
├── [x] .gitignore, LICENSE, rust-toolchain.toml
└── [x] GitHub Actions base

Jour 3-4: Algorithmes de base
├── [x] recursive.rs (naïf + mémoisation)
├── [x] iterative.rs (standard + branchless)
├── [x] Tests unitaires
└── [x] Doc comments

Jour 5: Benchmarks initiaux
├── [x] Setup Criterion
├── [x] Benchmark complexity_comparison
├── [x] README initial
└── [x] Premier commit fonctionnel
```

### Sprint 2 (Semaine 2) ✅ COMPLÉTÉ

**Objectif**: Algorithmes avancés

```
Jour 1-2: Méthode matricielle
├── [x] Matrix2x2 struct
├── [x] Fast exponentiation
├── [x] fib_matrix_modulo
├── [x] fib_doubling
└── [x] Tests et benchmarks

Jour 3: Formule de Binet
├── [x] fib_binet_f64
├── [x] Analyse d'erreur
├── [x] Constantes (PHI, PSI, SQRT_5)
└── [x] Tests de précision

Jour 4-5: Utilitaires
├── [x] FibMethod enum
├── [x] FibonacciCache
├── [x] FibonacciIterator
├── [x] count_recursive_calls
└── [x] Tests intégration
```

### Sprint 3 (Semaine 3) ✅ COMPLÉTÉ

**Objectif**: CLI & Outils

```
Jour 1-2: CLI fib-bench
├── [x] Structure clap
├── [x] Commande calc
├── [x] Commande compare
├── [x] Commande info
├── [x] Commande sequence
└── [x] Commande binet-analysis

Jour 3-4: Outils
├── [x] fib-profiler main.rs
├── [x] fib-viz main.rs
├── [x] Génération CSV
└── [x] Profiling basique

Jour 5: Documentation
├── [x] getting_started.md
├── [x] BENCHMARKS.md
└── [x] Tests E2E CLI
```

### Sprint 4 (Semaine 4) 🔄 EN COURS

**Objectif**: Documentation mathématique & polish

```
Jour 1-2: Documentation math
├── [x] MATHEMATICS.md complet
├── [x] matrix_method.md
├── [x] binet_formula.md
└── [x] fibonacci_theory.md

Jour 3-4: Performance docs
├── [x] optimization_techniques.md
├── [ ] memory_analysis.md
├── [ ] Résultats de benchmark réels
└── [ ] Graphiques de comparaison

Jour 5: Polish final
├── [ ] Relecture complète
├── [ ] Correction typos
├── [ ] Tests finaux
└── [ ] Tag v0.1.0
```

### Sprint 5 (Semaine 5) ⬜ PLANIFIÉ

**Objectif**: Profiling avancé (Unix)

```
Jour 1-2: Flamegraph
├── [ ] Intégration pprof
├── [ ] flamegraph.rs module
├── [ ] Commande CLI profile
└── [ ] Tests Unix only

Jour 3-4: Memory analysis
├── [ ] memory.rs module
├── [ ] allocator.rs custom
├── [ ] Tracking allocations
└── [ ] Rapport mémoire

Jour 5: Documentation
├── [ ] advanced_profiling.md
├── [ ] Exemples d'utilisation
└── [ ] Tests intégration
```

### Sprint 6 (Semaine 6) ⬜ PLANIFIÉ

**Objectif**: Visualisations

```
Jour 1-2: Chart generator
├── [ ] chart_generator.rs
├── [ ] Intégration Plotly
├── [ ] Templates graphiques
└── [ ] Export SVG/PNG

Jour 3-4: Data processing
├── [ ] data_parser.rs
├── [ ] Lecture CSV Criterion
├── [ ] Agrégation données
└── [ ] Rapport HTML

Jour 5: Intégration
├── [ ] CLI report command
├── [ ] CI benchmark artifacts
└── [ ] GitHub Pages deploy
```

### Sprint 7-8 (Semaines 7-8) ⬜ OPTIONNEL

**Objectif**: Bridge Go + SIMD

```
Semaine 7: Go FFI
├── [ ] go-src/fib.go
├── [ ] Build script CGO
├── [ ] go_bridge.rs
├── [ ] Benchmarks comparatifs
└── [ ] rust_vs_go.md

Semaine 8: SIMD
├── [ ] simd.rs (nightly)
├── [ ] Batch SIMD processing
├── [ ] Benchmarks SIMD
└── [ ] Documentation
```

---

## 🔗 Dépendances et ordre d'exécution

### Graphe de dépendances

```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│   Phase 1 ──────► Phase 2 ──────► Phase 3 ──────► Phase 4      │
│   (Fondation)     (Algos)         (CLI)          (Docs)         │
│       │              │               │              │           │
│       │              │               │              ▼           │
│       │              │               │         ┌─────────┐      │
│       │              │               │         │ Phase 9 │      │
│       │              │               │         │(Release)│      │
│       │              │               │         └─────────┘      │
│       │              │               │              ▲           │
│       │              │               ▼              │           │
│       │              │         ┌──────────┐        │           │
│       │              └────────►│ Phase 5  │────────┤           │
│       │                        │(Profiling)│        │           │
│       │                        └──────────┘        │           │
│       │                              │              │           │
│       │                              ▼              │           │
│       │                        ┌──────────┐        │           │
│       └───────────────────────►│ Phase 6  │────────┤           │
│                                │  (Viz)   │        │           │
│                                └──────────┘        │           │
│                                      │              │           │
│                                      ▼              │           │
│                                ┌──────────┐        │           │
│                                │ Phase 7  │────────┘           │
│                                │  (Go)    │                     │
│                                └──────────┘                     │
│                                      │                          │
│                                      ▼                          │
│                                ┌──────────┐                     │
│                                │ Phase 8  │ (Optionnel)         │
│                                │ (SIMD)   │                     │
│                                └──────────┘                     │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Matrice de dépendances

| Phase | Dépend de | Bloque |
|-------|-----------|--------|
| 1 - Fondation | - | 2, 3, 4, 5, 6 |
| 2 - Algos | 1 | 3, 5, 6, 8 |
| 3 - CLI | 1, 2 | 4, 5, 6 |
| 4 - Docs | 1, 2, 3 | 9 |
| 5 - Profiling | 3 | 7 |
| 6 - Viz | 4 | 7, 9 |
| 7 - Go | 5, 6 | 9 |
| 8 - SIMD | 6 | 9 |
| 9 - Release | 4, (6), (7), (8) | - |

### Chemin critique

```
Phase 1 → Phase 2 → Phase 3 → Phase 4 → Phase 9
   │         │         │         │
   └─────────┴─────────┴─────────┴── MVP RELEASE
   
Temps total chemin critique: ~5 semaines
```

---

## ⚠️ Risques et mitigations

### Risques techniques

| Risque | Probabilité | Impact | Mitigation |
|--------|-------------|--------|------------|
| **pprof incompatible Windows** | ✅ Confirmé | Moyen | Conditionner compilation Unix |
| **Overflow u128 grands n** | Faible | Faible | BigInt optionnel, doc limites |
| **Binet perte précision** | ✅ Confirmé | Faible | Documentation claire, n ≤ 78 |
| **CGO complexité** | Moyenne | Moyen | Phase optionnelle, isolation |
| **SIMD nightly only** | Haute | Faible | Feature flag, phase optionnelle |
| **Plotly breaking changes** | Faible | Moyen | Fixer version, tests CI |

### Risques projet

| Risque | Probabilité | Impact | Mitigation |
|--------|-------------|--------|------------|
| **Scope creep** | Moyenne | Élevé | Phases optionnelles clairement identifiées |
| **Perfectionnisme** | Moyenne | Moyen | MVP first, itérer ensuite |
| **Dépendances obsolètes** | Faible | Faible | cargo-outdated en CI |
| **Tests insuffisants** | Faible | Moyen | Coverage > 80% cible |

### Plan de contingence

```
Si Phase 5 (Profiling) bloquée:
└── Documenter limitations Windows
└── Fournir instructions manuelles perf

Si Phase 7 (Go) bloquée:
└── Phase optionnelle, skip pour v1.0
└── Documenter alternative benchmarks externes

Si Phase 8 (SIMD) bloquée:
└── Garder comme "future work"
└── Feature flag disabled par défaut
```

---

## 📏 Métriques de succès

### KPIs techniques

| Métrique | Cible | Statut actuel |
|----------|-------|---------------|
| Tests passants | 100% | ✅ 100% (43/43) |
| Couverture code | > 80% | 🔄 À mesurer |
| Warnings clippy | 0 | ✅ 0 |
| Doc coverage | 100% public | ✅ 100% |
| Benchmarks Criterion | 6 groupes | ✅ 6 groupes |
| Temps CI | < 5 min | ✅ ~2 min |

### KPIs fonctionnels

| Métrique | Cible | Statut actuel |
|----------|-------|---------------|
| Algorithmes implémentés | 5+ | ✅ 5 |
| Commandes CLI | 6+ | ✅ 6 |
| Fichiers documentation | 10+ | ✅ 11 |
| Exemples de code | 20+ | ✅ 25+ |

### Critères de release v1.0

```
✅ Tous les tests passent
✅ Documentation complète
✅ README avec exemples
✅ CHANGELOG à jour
✅ Licence MIT valide
⬜ cargo publish --dry-run réussi
⬜ Tag Git signé
⬜ GitHub Release créée
```

---

## 🛠️ Ressources et outils

### Stack technique

| Catégorie | Outil | Version |
|-----------|-------|---------|
| Langage | Rust | 1.70+ (stable) |
| Build | Cargo | 1.70+ |
| Benchmark | Criterion | 0.5 |
| CLI | clap | 4.4 |
| Sérialisation | serde + serde_json | 1.0 |
| Visualisation | plotly | 0.8 |
| BigInt | num-bigint | 0.4 |
| Profiling | pprof | 0.13 (Unix) |
| Tests property | proptest | 1.4 |

### Outils de développement

```bash
# Formatage
rustfmt

# Linting
clippy

# Benchmarks
cargo bench

# Tests
cargo test

# Documentation
cargo doc

# Audit sécurité
cargo audit

# Dépendances obsolètes
cargo outdated

# Coverage (optionnel)
cargo tarpaulin
```

### Commandes fréquentes

```bash
# Build complet
cargo build --release --all

# Tests avec output
cargo test -- --nocapture

# Benchmarks
cargo bench

# Générer documentation
cargo doc --open

# Exécuter CLI
cargo run --bin fib-bench -- --help

# Vérifier avant commit
cargo fmt && cargo clippy && cargo test
```

---

## 📎 Annexes techniques

### A. Spécifications des algorithmes

#### A.1 Récursif naïf

```
Entrée: n ∈ ℕ
Sortie: F(n)
Complexité: O(2^n) temps, O(n) espace (pile)
Limite pratique: n ≤ 35
```

#### A.2 Récursif mémorisé

```
Entrée: n ∈ ℕ
Sortie: F(n)
Complexité: O(n) temps, O(n) espace (cache)
Limite pratique: n ≤ 100,000 (stack)
```

#### A.3 Itératif

```
Entrée: n ∈ ℕ
Sortie: F(n)
Complexité: O(n) temps, O(1) espace
Limite pratique: n ≤ 186 (overflow u128)
```

#### A.4 Matriciel

```
Entrée: n ∈ ℕ
Sortie: F(n)
Complexité: O(log n) temps, O(1) espace
Limite pratique: n ≤ 186 (overflow u128)
```

#### A.5 Binet

```
Entrée: n ∈ ℕ
Sortie: F(n) (approximation f64)
Complexité: O(1) temps, O(1) espace
Limite précision: n ≤ 78
```

### B. Limites techniques

| Type | Valeur max | F(n) correspondant |
|------|------------|---------------------|
| u64 | 2^64 - 1 | F(93) |
| u128 | 2^128 - 1 | F(186) |
| f64 précision | ~10^15 | F(78) |
| Stack recursif | ~10^5 | Dépend OS |

### C. Valeurs de référence

```
F(10)  = 55
F(20)  = 6765
F(50)  = 12586269025
F(78)  = 8944394323791464 (limite Binet)
F(93)  = 12200160415121876738 (limite u64)
F(100) = 354224848179261915075
F(186) = (limite u128)
```

### D. Checklist pré-commit

```
[ ] cargo fmt --check
[ ] cargo clippy -- -D warnings
[ ] cargo test
[ ] cargo doc --no-deps
[ ] Pas de TODO/FIXME oubliés
[ ] CHANGELOG mis à jour (si release)
```

### E. Checklist release

```
[ ] Version bumped dans Cargo.toml
[ ] CHANGELOG.md à jour
[ ] Tests passent sur toutes les plateformes CI
[ ] Documentation générée
[ ] cargo publish --dry-run réussi
[ ] Tag Git créé et signé
[ ] GitHub Release créée avec notes
[ ] Annonce sur r/rust (optionnel)
```

---

## 📝 Historique des révisions

| Date | Version | Changements |
|------|---------|-------------|
| 2026-01-03 | 1.0.0 | Création initiale de la planification |
| - | 1.1.0 | Phases 1-3 complétées |
| - | 1.2.0 | Phase 4 en cours |

---

<p align="center">
<strong>🦀 Fibonacci Performance Benchmark Suite</strong><br>
<em>Un projet démontrant l'excellence en ingénierie Rust</em>
</p>
