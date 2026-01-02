# Planification Exhaustive d'Exécution - Scénario 1

## Développement From Scratch : Auto Gemini CLI

**Version**: 1.0  
**Date**: 2 janvier 2026  
**Horizon**: 19 semaines (4.5 mois)  
**Timeline**: 9 janvier 2026 → 22 mai 2026  
**Équipe**: 2.5 FTE (1 Lead Architect, 1 Senior Developer, 0.5 DevOps/QA)  
**Budget Estimé**: $102-152K

---

## I. PHASE 0: INITIALISATION (Semaine 0 | 2-8 janvier 2026)

### Objectif

Validation budgétaire, constitution de l'équipe et définition de l'architecture "Clean Slate".

### 0.1 Tâche: Validation & Kickoff

**Responsable**: Lead Architect  
**Durée**: 2 jours
**Livrables**:

- [ ] Budget signé ($150K réservé)
- [ ] Équipe assignée (2.5 FTE)
- [ ] Repo GitHub initialisé (vide)

### 0.2 Tâche: Architecture Design System

**Responsable**: Lead Architect  
**Durée**: 3 jours
**Livrables**:

- [ ] ADR #1: Architecture CLI (Oclif vs Ink)
- [ ] ADR #2: State Management (SQLite schema)
- [ ] ADR #3: Gemini Integration Pattern (Streaming/Tools)
- [ ] Diagramme de séquence (Core Loop)

---

## II. PHASE 1: FONDATION (Semaines 1-5 | 9 janvier - 13 février 2026)

### Objectif

Mettre en place le squelette de l'application, le wrapper API Gemini robuste et le stockage local.

### 1.1 Semaine 1: Setup TypeScript + Oclif

**Focus**: Infrastructure de base  
**Actions**:

- Init Oclif project (TypeScript)
- Configurer ESLint, Prettier, Husky
- Setup CI/CD (GitHub Actions) de base (Build/Test)
- Implémenter le système de Logging centralisé

### 1.2 Semaine 2: Wrapper Gemini API Complet

**Focus**: Abstraction API  
**Actions**:

- Implémenter `GeminiClient` (Google Generative AI SDK)
- Gérer le Streaming (`streamGenerateContent`)
- Implémenter le Token Counting (Gestion context 1M)
- Tests unitaires mocks pour l'API

### 1.3 Semaine 3: Event System & Architecture

**Focus**: Découplage  
**Actions**:

- Créer le bus d'événements (`EventEmitter`)
- Définir les types d'événements (User, System, Error)
- Implémenter le pattern Command/Query
- Architecture plugin-ready

### 1.4 Semaine 4: SQLite Session Store

**Focus**: Persistance  
**Actions**:

- Intégrer `better-sqlite3`
- Schema DB: Sessions, Messages, Contexts
- Repository Pattern pour l'accès données
- Tests de migration et intégrité

### 1.5 Semaine 5: Basic CLI & Context Manager

**Focus**: MVP Interface  
**Actions**:

- Commandes de base: `init`, `status`, `help`
- Gestionnaire de contexte (LRU cache, priority queue)
- Intégration initiale `GeminiClient` + `CLI`
- **Jalon**: "Hello World" conversationnel fonctionnel

---

## III. PHASE 2: CORE AGENT (Semaines 6-11 | 16 février - 27 mars 2026)

### Objectif

Développer le cerveau de l'agent: exécution parallèle, gestion de code et résilience.

### 2.1 Semaine 6: Streaming Response Handling

**Focus**: UX Interactif  
**Actions**:

- Parser le stream Gemini en temps réel
- Affichage TUI dynamique (buffer)
- Gestion des interruptions utilisateur (Ctrl+C)
- Recovery sur coupure réseau

### 2.2 Semaine 7: Parallel Worker Coordination

**Focus**: Performance  
**Actions**:

- Implémenter `piscina` pour workers threads
- Pool d'exécution pour tâches lourdes
- Coordination des sous-tâches agent
- Gestion de concurrence (Locks si SQLite)

### 2.3 Semaine 8: Code Sandbox & Execution

**Focus**: Sécurité & Capacité  
**Actions**:

- Design de l'environnement d'exécution (Sandboxing)
- Exécution de scripts TS/JS à la volée
- Capture stdout/stderr vers le contexte
- Sécurisation (permissions fichiers)

### 2.4 Semaine 9: Git Integration

**Focus**: Workflow Dev  
**Actions**:

- Wrapper `simple-git`
- Commandes: diff, commit, log, status, branch
- Analyse automatique des diffs par l'agent
- Gestion des conflits (basique)

### 2.5 Semaine 10: Rate Limiting & Resilience

**Focus**: Stabilité Gemini  
**Actions**:

- Implémenter Token Bucket / Leaky Bucket
- Gestion des `429 Too Many Requests`
- Stratégies de Backoff exponentiel
- Queue prioritaire des requêtes

### 2.6 Semaine 11: Persistence & Recovery

**Focus**: Robustesse  
**Actions**:

- Reprise sur crash (State restoration)
- Historique de session navigable
- Export/Import de sessions
- **Jalon**: Agent capable de réaliser une tâche complexe multi-étapes

---

## IV. PHASE 3: INTEGRATIONS (Semaines 12-15 | 30 mars - 24 avril 2026)

### Objectif

Connecter l'agent au monde extérieur et polir l'expérience utilisateur terminal.

### 3.1 Semaine 12: MCP Server Design (Custom)

**Focus**: Extensibilité  
**Actions**:

- Implémenter le protocole Model Context Protocol (MCP)
- Créer un serveur MCP interne
- Connecteurs standards (Filesystem, Bravia/Search?)
- API pour tools tiers

### 3.2 Semaine 13: GitHub Integration

**Focus**: Collaboration  
**Actions**:

- API GitHub (Octokit)
- Gestion des PRs (Create, Review, Merge)
- Analyse des Issues
- Lier commit à issue

### 3.3 Semaine 14: Advanced Terminal UI (Ink)

**Focus**: "Wow" Effect  
**Actions**:

- Réécriture UI avec React Ink
- Composants: Spinners, Tables, Syntax Highlighting
- Dashboards interactifs dans le terminal
- Support thèmes (Dark/Light/System)

### 3.4 Semaine 15: Config & Profiles

**Focus**: Personnalisation  
**Actions**:

- Système de configuration global/local (`.agbrc`)
- Gestion des profils (Work, Personal)
- Secrets management (Keyring integration)
- **Jalon**: Feature Code Complete

---

## V. PHASE 4: QUALITY & RELEASE (Semaines 16-19 | 27 avril - 22 mai 2026)

### Objectif

Stabiliser, tester massivement et préparer le lancement 1.0.

### 4.1 Semaine 16: Unit Testing Framework

**Focus**: Qualité Code  
**Actions**:

- Vitest à 100% sur Core et Utils
- Mocking intensif de Gemini API
- Tests de régression
- Code Coverage Report (>80% required)

### 4.2 Semaine 17: Integration & E2E Tests

**Focus**: Qualité Système  
**Actions**:

- Tests scénarisés complets (Real Gemini API - Dev env)
- Tests de charge (Memory leaks ?)
- Tests multi-plateformes (Win, Mac, Linux)
- Security Audit (Deps, Permissions)

### 4.3 Semaine 18: Documentation

**Focus**: Adoption  
**Actions**:

- Site documentation (Docusaurus ou Markdown)
- Tutoriels "Getting Started"
- API Reference (Typedoc)
- Contributor Guide

### 4.4 Semaine 19: Build, Packaging & Release

**Focus**: Distribution  
**Actions**:

- Build optimisés (pkg / ncc)
- Installers (npm, brew, choco)
- Release Notes
- **LIVRAISON FINALE v1.0**

---

## VI. STRUCTURE DE COÛTS DÉTAILLÉE

| Catégorie               | Estimation   | Détails                                      |
| ----------------------- | ------------ | -------------------------------------------- |
| **Ressources Humaines** | $142,500     | 2.5 FTE (Lead $80/h, Senior $60/h) x 760h    |
| **Infrastructure Dev**  | $1,500       | Serveurs CI/CD, instances test, stockage     |
| **Logiciels/SaaS**      | $2,000       | GitHub Enterprise, Tuple.app, Cloud services |
| **Réserve Aléas**       | $6,000       | ~4% du budget                                |
| **TOTAL**               | **$152,000** | Plafond haute fourchette                     |

---

## VII. GESTION DES RISQUES (SPÉCIFIQUE SCÉNARIO 1)

| Risque                               | Impact   | Probabilité | Mitigation                                      |
| ------------------------------------ | -------- | ----------- | ----------------------------------------------- |
| **Évolution API Gemini**             | Critique | Haute       | Wrapper stricte, veille technique hebdomadaire. |
| **Explosion du Scope**               | Élevé    | Moyenne     | Backlog strict, validation PM chaque semaine.   |
| **Courbe d'apprentissage Oclif/Ink** | Moyen    | Moyenne     | Pair programming, Spike en S0.                  |
| **Performance Node.js (Memory)**     | Moyen    | Faible      | Profiling dès S7, Workers threads.              |

---

## VIII. COMPARATIF RAPIDE VS SCÉNARIO 2

> [!WARNING]
> Ce plan "From Scratch" représente un investissement **4x supérieur** au Scénario 2 (Adaptation).
> Il ne doit être choisi que si l'objectif est une **propriété intellectuelle totale** et une **architecture sur-mesure** sans aucun compromis.

- **Délai**: 19 semaines (vs 14)
- **Coût**: ~$150K (vs ~$34K)
- **Dette Technique**: 0% (vs ~15%)
- **Liberté Architecture**: 100% (vs 20%)
