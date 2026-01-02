# SPECIFICATION PRODUIT : Auto-Gemini-CLI
## Requirements Fonctionnels & Non-Fonctionnels

**Version**: 1.0  
**Date**: 2 janvier 2026  
**Produit**: Auto-Gemini-CLI - Autonomous AI Coding Agent (CLI-First)  
**Fondé sur**: Adaptation d'Auto-Claude pour Gemini API  
**Public cible**: Développeurs, équipes d'engineering, entreprises tech (CLI-first workflow)  

---

## TABLE DES MATIÈRES

1. [Vision Produit](#vision-produit)
2. [Requirements Fonctionnels](#requirements-fonctionnels)
3. [Requirements Non-Fonctionnels](#requirements-non-fonctionnels)
4. [Contraintes & Limitations](#contraintes--limitations)
5. [Cas d'Usage](#cas-dusage)
6. [Architecture Services](#architecture-services)
7. [Intégrations Externes](#intégrations-externes)
8. [Roadmap Produit](#roadmap-produit)

---

## VISION PRODUIT

### Énoncé de Vision

**Auto-Gemini-CLI** est un agent IA autonome CLI-first pour la programmation et l'ingénierie logicielle. Il automatise les tâches de coding, maintient le contexte entre sessions, exécute du code réel dans des sandboxes sécurisés, et s'intègre nativement à l'écosystème Unix/Git/GitHub.

**Slogan**: *"Your autonomous CLI coding companion - faster, cheaper, context-aware."*

### Caractéristiques Clés vs Auto-Claude

| Aspect | Auto-Gemini-CLI | Auto-Claude |
|--------|-----------------|-------------|
| **Interface** | CLI-first (Oclif) | Desktop app (Electron) |
| **API** | Gemini 2.0 Flash | Claude 3.5 Sonnet |
| **Context Window** | 1M tokens | 200K tokens |
| **Rate Limiting** | 1000 req/24h free | Unlimited (paid) |
| **Cost Structure** | Free tier available | Claude Code $200/mo |
| **Performance** | Faster (Flash model) | More capable (Sonnet) |
| **Deployment** | Lightweight CLI | Desktop app |
| **Platform** | Unix-first (Linux, macOS) | Cross-platform |

### Principes Fondamentaux

- ✅ **CLI-First**: Native to Unix philosophy (piping, scripting)
- ✅ **Cost-Effective**: Gemini free tier for learning/small teams
- ✅ **Large Context**: 1M token window for complex tasks
- ✅ **Fast Execution**: Gemini 2.0 Flash optimized for speed
- ✅ **Transparent**: Comprehensive logging & debugging
- ✅ **Modular**: Easily extensible via MCP servers

---

## REQUIREMENTS FONCTIONNELS

### RF.1 Agent Core & Reasoning

#### RF.1.1 Task Execution Engine
- [ ] **Agent doit exécuter des tâches de code de bout en bout**
  - Entrée: Texte de tâche + contexte (fichiers, historique)
  - Processus:
    1. Parse task description
    2. Analyze context (project structure, files)
    3. Decompose into sub-tasks
    4. Generate code via Gemini API
    5. Execute in sandbox
    6. Validate results
    7. Commit if successful
  - Sortie: Code généré, résultats exécution, logs complets
  - Acceptation: Tâche complète sans erreurs non-recouvrables

#### RF.1.2 Multi-Step Reasoning (Chain-of-Thought)
- [ ] **Agent doit décomposer tâches complexes**
  - Entrée: "Build a REST API with authentication and pagination"
  - Logique: Utiliser Gemini thinking pour:
    1. Analyser les besoins
    2. Identifier les étapes
    3. Estimer la complexité
    4. Planifier l'ordre d'exécution
    5. Identifier les dépendances
  - Sortie: Plan structuré avec timing estimé
  - Acceptation: Plan exécutable, étapes claires

#### RF.1.3 Error Handling & Recovery
- [ ] **Agent doit gérer erreurs et tenter récupération**
  - Erreurs gérées:
    - Syntax errors (parse errors)
    - Runtime errors (undefined, null pointer)
    - File not found / permission errors
    - Timeout errors (>30s execution)
    - API errors (rate limit, auth failure)
    - Network errors (connection timeout)
  - Stratégies:
    - Analyser erreur avec Gemini
    - Générer fix suggestions
    - Réessayer avec limite (max 3 attempts)
    - Fallback strategies
  - Max retries: 3 par étape
  - Acceptation: Erreur documentée avec recovery tentées

#### RF.1.4 Context Persistence & Reuse
- [ ] **Agent doit maintenir contexte entre sessions**
  - Contexte inclut:
    - Project structure snapshot
    - Task history avec résultats
    - Variables d'état & configuration
    - Décisions prises (pour éviter re-exploration)
    - Patterns réussis (pour réutilisation)
  - Storage: SQLite database (`~/.auto-gemini/sessions.db`)
  - Loading: Automatique au démarrage session
  - Compression: Auto-summarize old context (>100K tokens)
  - Acceptation: Context restauré, pas de perte de state

#### RF.1.5 Learning from History
- [ ] **Agent doit apprendre des patterns réussis**
  - Méchanisme: Stocker patterns dans pattern database
  - Exemple: "If error X occurs, solution Y works"
  - Réutilisation: Proposer solution avant re-generating
  - Adaptation: Modifier patterns basé sur results
  - Storage: JSON file à `~/.auto-gemini/patterns.json`
  - Acceptation: Agent utilise patterns dans nouvelles tâches

---

### RF.2 Code Generation & Execution

#### RF.2.1 Code Generation via Gemini API
- [ ] **Agent doit générer du code via Gemini**
  - Modèle: `gemini-2.0-flash` (default) ou `gemini-2.0-pro` (complex)
  - Paramètres:
    - Temperature: 0.5-0.7 (balance)
    - Max tokens: Adaptatif (max 8000 per call)
    - System prompt: Role-specific instructions
    - Think tokens: Optionnel pour complex reasoning
  - Langages supportés: Python, JavaScript/TypeScript, Go, Rust, Bash, SQL
  - Fallback: Si Gemini rate-limit, queue et retry après 1h
  - Acceptation: Code généré, utile et exécutable

#### RF.2.2 Code Execution Sandbox
- [ ] **Agent doit exécuter code dans sandbox sécurisé**
  - Modes:
    - Node.js: TypeScript/JavaScript via Node runtime
    - Python: Via subprocess with timeout
    - Bash: POSIX shell (whitelist commands)
    - Docker: Optionnel pour isolation max
  - Restrictions (sécurité):
    - File system: Read-only outside project directory
    - Network: Disabled (sauf whitelist)
    - Process: Single-threaded per execution
    - Time: 30 secondes max
    - Memory: 256MB limit
  - Output: Capture stdout, stderr, exit code
  - State: Sandboxes are ephemeral (clean each time)
  - Acceptation: Code exécuté isolated, sécurisé

#### RF.2.3 Code Testing & Validation
- [ ] **Agent doit tester code généré**
  - Types de tests:
    - Unit tests: Via Jest, Vitest, Pytest (if exists)
    - Integration tests: Avec mocks
    - Manual validation: Via script si no test framework
  - Exigence: Min 1 test per task (créer si needed)
  - Rapport: Format JSON avec pass/fail per test
  - Thresholds:
    - All tests must pass before acceptance
    - Min coverage: 70% (warn if <70%)
  - Acceptation: Test report generated, tous les tests passing

#### RF.2.4 Code Quality Checks
- [ ] **Agent doit analyser qualité du code**
  - Outils utilisés:
    - Linting: ESLint (JS/TS), Pylint (Python)
    - Type checking: TypeScript strict mode, mypy
    - Code style: Prettier (JS), Black (Python)
    - Complexity: Cyclomatic complexity check
  - Métriques rapportées:
    - Errors count (must be 0)
    - Warnings count
    - Complexity score
    - Duplication percentage
  - Seuils:
    - 0 errors: Mandatory
    - <5 warnings: Acceptable
    - Complexity <10: Preferred
  - Acceptation: Code passe quality gates

#### RF.2.5 Git Integration & Version Control
- [ ] **Agent doit intégrer avec Git pour versioning**
  - Opérations supportées:
    - Clone repository
    - Create feature branches
    - Stage & commit changes
    - Push to remote origin
    - Create pull requests (via GitHub API)
    - Merge après approval (optionnel)
  - Comportement:
    - Chaque task = feature branch distinct
    - Commit message: Auto-generated avec context
    - PR description: Include task description + test results
    - Squash option: Combine commits si optionnel
  - Branch naming: `feature/task-<id>-<description>`
  - Acceptation: Commits dans Git, PRs sur GitHub

#### RF.2.6 Multi-Language Support
- [ ] **Agent doit supporter multiple langages**
  - Python 3.9+: Full support
  - JavaScript/TypeScript: Node 18+ support
  - Go 1.18+: Compilation & execution
  - Rust 1.60+: Cargo support
  - Bash/Shell: POSIX compatibility
  - SQL: Via sqlite3 CLI
  - Detection: Auto-detect from file extension ou task context
  - Acceptation: Code généré & exécuté pour chaque langage

---

### RF.3 Session Management

#### RF.3.1 Session Lifecycle
- [ ] **Agent doit créer et maintenir sessions**
  - Session structure:
    ```json
    {
      "id": "sess-uuid",
      "createdAt": "ISO8601",
      "updatedAt": "ISO8601",
      "workspacePath": "/path/to/project",
      "tasks": ["task-id1", "task-id2"],
      "context": { "files": {}, "state": {} },
      "config": { "model": "gemini-2.0-flash" },
      "metadata": { "taskCount": 5, "tokensUsed": 45000 }
    }
    ```
  - Storage: SQLite (`sessions.db`)
  - Auto-save: After each task completion
  - Lifecycle: Create → Run tasks → Close → Archive
  - Acceptation: Session persisted, recoverable

#### RF.3.2 Multi-Session Concurrency
- [ ] **Agent doit supporter plusieurs sessions en parallèle**
  - Isolation: Chaque session = working directory isolé
  - Locking: SQLite transaction-based (no race conditions)
  - Max sessions: Configurable (default: 10)
  - Resource management:
    - Monitor memory per session
    - Kill stale sessions (>24h inactive)
    - Graceful shutdown on resource limit
  - Concurrency: True parallel execution (not sequential)
  - Acceptation: Sessions en parallèle sans corruption

#### RF.3.3 Session History & Recovery
- [ ] **Agent doit permettre replay & recovery**
  - Fonctionnalité:
    - View complete execution history
    - Replay specific task
    - Resume interrupted session
    - Export session as JSON
  - History format: JSON lines (1 event per line)
  - Recovery: Auto-resume on agent restart
  - Acceptation: Replay produit résultats identiques

#### RF.3.4 Session Summary & Export
- [ ] **Agent doit générer résumés**
  - Contenu du résumé:
    - Tasks executed: count, success rate
    - Code generated: LOC, files modified
    - Errors encountered: types, recovery attempts
    - Time spent: per task, total
    - Tokens used: input/output breakdown
    - Files generated/modified: list
  - Formats: JSON, Markdown, CSV
  - Export command: `auto-gemini session export <id> [--format]`
  - Acceptation: Résumé généré, exportable

---

### RF.4 CLI Interface

#### RF.4.1 Command-Line Interface (Oclif)
- [ ] **Agent doit avoir une CLI complète**
  - Framework: Oclif (Salesforce)
  - Commandes principales:
    ```bash
    auto-gemini init [--name <name>]           # Initialize workspace
    auto-gemini task new <description>         # Create task
    auto-gemini task run <id> [--watch]        # Execute task
    auto-gemini task list [--status]           # List tasks
    auto-gemini task cancel <id>               # Cancel running task
    auto-gemini session list [--limit <n>]     # List sessions
    auto-gemini session show <id>              # Session details
    auto-gemini session export <id> [--format] # Export session
    auto-gemini session clear <id>             # Delete session
    auto-gemini chat [--session <id>]          # Interactive REPL mode
    auto-gemini status [--watch]               # System status
    auto-gemini config set <key> <value>       # Set configuration
    auto-gemini config get <key>               # Get configuration
    auto-gemini config show                    # Show all config
    auto-gemini logs [--tail] [--grep <pat>]   # View logs
    auto-gemini help [<command>]               # Help
    ```
  - Flags globals: `--verbose`, `--debug`, `--config-dir`
  - Output format: Pretty-print (default) ou JSON (`--json`)
  - Acceptation: Toutes commandes fonctionnelles

#### RF.4.2 Interactive Mode (REPL)
- [ ] **Agent doit supporter mode interactif**
  - Mode d'invocation: `auto-gemini chat` ou `auto-gemini task run --interactive`
  - Features:
    - Prompt: "You> " pour user input
    - Context: Maintained between commands
    - History: Command history (arrow keys)
    - Completion: Auto-complete suggestions
    - Inspection: `$inspect` pour voir contexte
    - Modify: `$set <key> <value>` pour changer variables
  - Session: Auto-create session per chat instance
  - Save: Auto-save session on exit
  - Acceptation: REPL responsive, contexte maintained

#### RF.4.3 Real-time Logging & Monitoring
- [ ] **Agent doit afficher logs en temps réel**
  - Features:
    - Live streaming de stdout/stderr
    - Color-coded output (errors: red, success: green, info: blue)
    - Timestamps: [HH:MM:SS]
    - Searchable logs: `--grep` pattern matching
    - Tail mode: `--tail <n>` dernières N lignes
    - Export: `auto-gemini logs > output.log`
  - Performance: <100ms latency pour log display
  - Levels: debug, info, warn, error
  - Acceptation: Logs affichés en temps réel

#### RF.4.4 Progress Indicators
- [ ] **Agent doit afficher progression**
  - Indicators:
    - Spinner animation pour tâches longues
    - Progress bar pour tasks avec % known
    - Time estimation: "~30s remaining"
    - Current step display: "Step 3/5: Running tests"
  - Update frequency: Every 100ms
  - Performance: No UI lag
  - Acceptation: Progress affiché smooth

---

### RF.5 Integrations

#### RF.5.1 GitHub Integration (Full)
- [ ] **Agent doit intégrer avec GitHub**
  - Capacités:
    - Authenticate via PAT (Personal Access Token)
    - Clone repositories
    - Create branches & commits
    - Push to remote
    - Create pull requests avec auto-generated descriptions
    - Comment on PRs / Issues
    - Read issues & projects
    - Manage labels & milestones
  - Authentication: PAT via env var `GITHUB_TOKEN`
  - API: Utiliser Octokit.js pour requests
  - Rate limits: Respecter limites GitHub (60 req/hr public, 1000 private)
  - Error handling: Rate limit → queue requests, retry avec exponential backoff
  - Acceptation: PRs créées, commits pushés, issues commentées

#### RF.5.2 npm/Package Manager Integration
- [ ] **Agent doit gérer les dépendances**
  - Capacités:
    - `npm install` / `npm ci`
    - `npm update` avec version management
    - Run npm scripts: `npm run <script>`
    - Analyze `package.json` & `package-lock.json`
    - Support yarn, pnpm (auto-detect)
    - Audit dependencies: `npm audit`
  - Utilisation: Auto-install si dépendances manquantes
  - Lock files: Always commit lock files
  - Version management: Respect semver
  - Acceptation: Dépendances installées, lock files updated

#### RF.5.3 MCP (Model Context Protocol) Support
- [ ] **Agent doit supporter les MCP servers**
  - Fonctionnalité: Intégrer MCP servers pour étendre capacités
  - Configuration:
    ```json
    {
      "mcp": {
        "github": { "url": "mcp://github-server" },
        "database": { "url": "mcp://postgres-server" }
      }
    }
    ```
  - Usage: Disponible dans Gemini context (tools)
  - Protocol: JSON-RPC 2.0 over stdio/socket
  - Fallback: Utiliser direct API si MCP unavailable
  - Acceptation: MCP tools disponibles dans contexte

#### RF.5.4 API & HTTP Integration
- [ ] **Agent doit supporter HTTP calls**
  - Capacités:
    - HTTP methods: GET, POST, PUT, DELETE, PATCH
    - Headers: Authorization (Bearer, API key)
    - Body: JSON, form-data, raw
    - Parsing: Auto-parse JSON responses
    - Error handling: Retry with backoff
  - Whitelist: Configurable domains/endpoints
  - Rate limiting: Respecter API rate limits
  - Timeout: 30 secondes max per request
  - Acceptation: API calls exécutés, résultats reçus

#### RF.5.5 Docker Integration (Optionnel)
- [ ] **Agent peut utiliser Docker pour isolation**
  - Utilisation: Pour complex projects (Java, Go compilation)
  - Docker support:
    - Auto-detect if Docker available
    - Run code in container
    - Mount project as volume
    - Cleanup containers after execution
  - Fallback: Sandbox sans Docker (limited)
  - Acceptation: Code isolé via Docker si available

---

### RF.6 Configuration & Customization

#### RF.6.1 Configuration System
- [ ] **Agent doit permettre configuration complète**
  - Config file: `~/.auto-gemini/config.json`
  - Schéma:
    ```json
    {
      "api": {
        "gemini": {
          "apiKey": "string",
          "model": "gemini-2.0-flash",
          "temperature": 0.5,
          "maxTokens": 8000
        }
      },
      "github": {
        "token": "string",
        "organization": "string (optional)"
      },
      "sandbox": {
        "timeout": 30,
        "memoryLimit": 256,
        "useDocker": false
      },
      "workspace": {
        "defaultPath": "/path/to/workspace",
        "maxSessions": 10
      },
      "logging": {
        "level": "info",
        "format": "json"
      }
    }
    ```
  - Sources: File + env vars + CLI flags (priority: CLI > env > file)
  - Validation: JSON schema validation on startup
  - Acceptation: Configuration appliquée correctement

#### RF.6.2 Environment Variables
- [ ] **Agent doit supporter env variables**
  - Variables requises:
    - `GEMINI_API_KEY`: API key pour Gemini
  - Variables optionnelles:
    - `GITHUB_TOKEN`: GitHub PAT
    - `AUTO_GEMINI_WORKSPACE`: Default workspace
    - `GEMINI_MODEL`: Model override (default: gemini-2.0-flash)
    - `LOG_LEVEL`: debug, info, warn, error
    - `SANDBOX_TIMEOUT`: En secondes
  - Acceptation: Env vars lues et appliquées

#### RF.6.3 Custom Prompts
- [ ] **Agent doit permettre custom system prompts**
  - Fichiers:
    - `~/.auto-gemini/prompts/system.md`: System prompt personnalisé
    - `~/.auto-gemini/prompts/task-template.md`: Task template
    - `~/.auto-gemini/prompts/review.md`: Code review prompt
  - Usage: Intégré dans appels Gemini
  - Variables: Support `{{variable}}` substitution
  - Acceptation: Custom prompts utilisés dans générations

#### RF.6.4 Workspace Management
- [ ] **Agent doit gérer multiple workspaces**
  - Commande: `auto-gemini workspace <cmd>`
  - Sub-commands:
    - `workspace init <path>`: Initialize workspace
    - `workspace list`: List all workspaces
    - `workspace switch <path>`: Switch workspace
    - `workspace config show`: Show workspace config
  - Each workspace: Isolated config & sessions
  - Acceptation: Workspaces isolés, switchable

---

### RF.7 Monitoring & Observability

#### RF.7.1 Metrics Collection
- [ ] **Agent doit collecter métriques**
  - Métriques:
    - Tasks executed: count, success rate, avg duration
    - Code generated: LOC per language, files modified
    - Errors: count by type, recovery rate
    - API usage: Gemini calls, tokens used, estimated cost
    - Session stats: Duration, context size, parallel count
  - Storage: SQLite metrics table
  - Retention: 90 jours (configurable)
  - Export: JSON, CSV formats
  - Acceptation: Métriques collectées & exportables

#### RF.7.2 Performance Profiling
- [ ] **Agent doit profiler les performances**
  - Métriques:
    - API latency: Gemini calls (p50, p95, p99)
    - Code execution time: Per language
    - Memory usage: Per session, peak
    - Disk usage: Sessions database size
  - Alertes: Si anomalies détectées
    - API latency > 10s → warn
    - Memory usage > 500MB → warn
  - Dashboard: Via `auto-gemini status --watch`
  - Acceptation: Performance data collectée

#### RF.7.3 Error Tracking & Analysis
- [ ] **Agent doit tracker et analyser erreurs**
  - Data:
    - Error type (SyntaxError, RuntimeError, etc.)
    - Message & stack trace
    - Context (file, line, task)
    - Recovery attempts & results
    - Frequency & patterns
  - Storage: Error database
  - Reports: Exportable avec `auto-gemini logs --grep error`
  - Analysis: Auto-suggest solutions
  - Acceptation: Erreurs tracées & analysables

#### RF.7.4 Token Usage Tracking
- [ ] **Agent doit tracker utilisation de tokens**
  - Tracking:
    - Input tokens per call
    - Output tokens per call
    - Estimated cost (based on Gemini pricing)
    - Cumulative usage per session
    - Projected monthly usage
  - Display: `auto-gemini status` shows usage
  - Alerts: If approaching quota limits
  - Export: Token usage report
  - Acceptation: Token usage suivi & reportable

---

### RF.8 Security & Privacy

#### RF.8.1 API Key Management
- [ ] **Agent doit gérer API keys sécurisément**
  - Stockage:
    - Jamais en clair dans logs
    - Jamais en fichiers de config versionned
    - Masqué lors d'affichage (show last 4 chars only)
  - Configuration:
    - Env var `GEMINI_API_KEY` recommandé
    - Config file alternative (but warn)
    - Prompt si missing (interactive)
  - Rotation: Support pour changer la key
  - Acceptation: Keys protégées, jamais exposées

#### RF.8.2 Sandbox Isolation & Security
- [ ] **Agent doit isoler l'exécution de code**
  - Isolation:
    - Separate process per execution
    - No access to parent process
    - No access to system files outside project
  - Restrictions:
    - File system: Read-only outside project directory
    - Network: Disabled (sauf whitelist)
    - Environment: Filtered env vars
    - Processes: Single process, no spawning
  - Timeouts: 30 secondes max
  - Memory: 256MB limit (hard kill at limit)
  - Acceptation: Code exécuté isolé & sécurisé

#### RF.8.3 Input Validation & Sanitization
- [ ] **Agent doit valider toutes les entrées**
  - Validation:
    - CLI arguments: Type checking, bounds
    - Config values: Schema validation
    - API responses: JSON schema validation
    - User input: Escape special characters
  - Sanitization:
    - Remove potential injection vectors
    - Escape for shell commands
    - Validate file paths (no traversal)
  - Acceptation: Injections impossible

#### RF.8.4 Audit Logging
- [ ] **Agent doit logger toutes les actions importantes**
  - Events loggés:
    - Task creation & execution
    - Code generation & execution
    - Git operations (commits, pushes)
    - API calls (Gemini, GitHub)
    - Configuration changes
    - Authentication events
  - Format: JSON with timestamp, actor, action, result
  - Retention: 1 year (configurable)
  - Acceptance: Audit logs complets & queryable

#### RF.8.5 Data Privacy
- [ ] **Agent doit respecter la vie privée**
  - Politique:
    - No data sent to external services (except APIs configured)
    - No telemetry by default
    - User code never shared with Anthropic/Google
    - Local storage only
  - Opt-in: User can enable analytics if willing
  - GDPR: Support for data export/deletion
  - Acceptation: Privacy policy implémentée

---

## REQUIREMENTS NON-FONCTIONNELS

### RNF.1 Performance

#### RNF.1.1 API Response Time
- [ ] Gemini API calls: <5 secondes (p95)
- [ ] Token counting: <1 seconde
- [ ] Code execution: <30 secondes (timeout)
- [ ] CLI commands (non-blocking): <2 secondes
- [ ] Task startup: <500ms
- Measurement: Via timing logs, `auto-gemini status --watch`

#### RNF.1.2 CLI Responsiveness
- [ ] Command execution: <500ms (for UI feedback)
- [ ] Log display: <100ms for new entries
- [ ] Session list: <1 seconde (100+ sessions)
- [ ] Task switching: <200ms
- Measurement: Using time command, profiling

#### RNF.1.3 Memory Usage
- [ ] Idle CLI: <50MB
- [ ] Running task: <300MB
- [ ] 10 concurrent sessions: <1GB
- [ ] Memory leak tests: Pass (24h stability test)
- Measurement: Via process monitor, memory profiling

#### RNF.1.4 Startup Time
- [ ] CLI initialization: <500ms
- [ ] Session recovery: <1 seconde (from disk)
- [ ] Project analysis: <5 secondes (100K LOC project)
- Measurement: Time from shell invocation to ready

---

### RNF.2 Reliability & Availability

#### RNF.2.1 Uptime & Resilience
- [ ] Target uptime: 99% (agent availability)
- [ ] Mean Time Between Failures: >7 days
- [ ] Mean Time To Recovery: <5 minutes
- [ ] Graceful degradation: Works offline (queues requests)
- Measurement: Production monitoring

#### RNF.2.2 Data Persistence
- [ ] No session loss on crash
- [ ] Database backup: Daily auto-backup
- [ ] Recovery Point Objective (RPO): <1 hour
- [ ] Data integrity: No corruption tests pass
- Measurement: Failure scenario testing

#### RNF.2.3 Error Rate & Recovery
- [ ] Unhandled exceptions: <0.01%
- [ ] Failed tasks due to system: <5%
- [ ] API errors (rate limit, auth): <2%
- [ ] Recovery success rate: >95%
- Measurement: Error logs analysis

#### RNF.2.4 Network Resilience
- [ ] Timeout handling: Exponential backoff (3 retries)
- [ ] Offline mode: Queue requests, sync when online
- [ ] Connection loss: Graceful degradation
- [ ] Rate limit: Queue and retry after reset
- Measurement: Network failure testing

---

### RNF.3 Scalability

#### RNF.3.1 Concurrent Sessions
- [ ] Min support: 10 sessions simultaneous
- [ ] Optimal: 50+ sessions (with tuning)
- [ ] Performance degradation: <10% (at 50 sessions)
- [ ] Resource usage: Linear or better
- Measurement: Load testing

#### RNF.3.2 Data Growth
- [ ] Database: Support 10,000+ sessions
- [ ] Logs: Support 100GB without performance issues
- [ ] Query time: <1s (even with large dataset)
- [ ] Storage: No runaway growth
- Measurement: Database benchmarking

#### RNF.3.3 Project Size
- [ ] Support projects >500K LOC
- [ ] Analysis time: <10 secondes
- [ ] Memory: Scales linearly with project size
- [ ] No hitting limits
- Measurement: Test with large projects

---

### RNF.4 Maintainability

#### RNF.4.1 Code Quality
- [ ] Type safety: 100% (TypeScript strict mode)
- [ ] Test coverage: >80% (unit + integration)
- [ ] Documentation: JSDoc on all public APIs
- [ ] Code style: Automated (ESLint + Prettier)
- [ ] Complexity: Cyclomatic <10 per function
- Measurement: Coverage reports, linter output

#### RNF.4.2 Modularity & Architecture
- [ ] Clear separation of concerns
- [ ] Service-based architecture
- [ ] Dependency injection used
- [ ] Easy to extend (plugins, MCP)
- [ ] Low coupling between modules
- Measurement: Architecture review

#### RNF.4.3 Logging & Debugging
- [ ] Structured logging (JSON format)
- [ ] Debug mode: Verbose logging available
- [ ] Source maps: For stack traces
- [ ] Execution replay: Complete trace
- [ ] Log levels: debug, info, warn, error
- Measurement: Log completeness

---

### RNF.5 Security

#### RNF.5.1 Authentication & Authorization
- [ ] Gemini API: Require valid API key
- [ ] GitHub: OAuth2 or PAT validation
- [ ] No hardcoded credentials
- [ ] Secure secret storage
- [ ] Credential rotation support
- Measurement: Security audit

#### RNF.5.2 Data Protection
- [ ] Encryption at rest: For sensitive data (optional)
- [ ] No sensitive data in logs
- [ ] Secure file deletion (shred)
- [ ] GDPR compliance (data export/deletion)
- [ ] No telemetry by default
- Measurement: Data protection audit

#### RNF.5.3 Vulnerability Management
- [ ] Dependency scanning: npm audit, Snyk
- [ ] Security updates: Monthly
- [ ] CVE tracking: Automated
- [ ] Incident response: Plan documented
- [ ] Penetration testing: Annual (for commercial)
- Measurement: Regular scans

---

### RNF.6 Usability

#### RNF.6.1 User Onboarding
- [ ] `auto-gemini init`: Interactive wizard
- [ ] Help text: For all commands (`--help`)
- [ ] Examples: Example tasks & templates provided
- [ ] Documentation: Comprehensive (>100 pages)
- [ ] Error messages: Clear & actionable
- Measurement: User feedback, success rate

#### RNF.6.2 Error Messages
- [ ] Clear & non-technical language
- [ ] Actionable suggestions for fixes
- [ ] Links to relevant documentation
- [ ] Code examples where applicable
- [ ] Avoid jargon
- Measurement: UX testing

#### RNF.6.3 Accessibility
- [ ] CLI: Full keyboard navigation
- [ ] Color: Not only color for information
- [ ] Screen reader: Compatible (for tools that use CLI)
- [ ] Contrast: WCAG AA standards
- [ ] Font size: Configurable
- Measurement: Accessibility audit

---

### RNF.7 Compatibility

#### RNF.7.1 Platform Support
- [ ] Linux: Ubuntu 20.04+, Debian 11+, Fedora 35+
- [ ] macOS: 11.x (Big Sur)+
- [ ] Windows: 10/11 (with WSL2 recommended)
- [ ] Node.js: 18.x LTS or higher
- Measurement: Test on all platforms

#### RNF.7.2 Language Support
- [ ] Python: 3.9, 3.10, 3.11, 3.12
- [ ] JavaScript/TypeScript: Node 18+
- [ ] Go: 1.18+
- [ ] Rust: 1.60+
- [ ] Bash: POSIX compatible
- Measurement: Test generation per language

#### RNF.7.3 Dependency Compatibility
- [ ] Support latest LTS versions
- [ ] Backward compatibility: 2 major versions
- [ ] Minimal transitive dependencies
- [ ] Regular updates: Monthly security reviews
- [ ] Lock files: Always committed
- Measurement: Dependency audit

---

### RNF.8 Deployment & Operations

#### RNF.8.1 Installation
- [ ] npm install: `npm install -g auto-gemini-cli`
- [ ] Homebrew: `brew install auto-gemini-cli`
- [ ] Direct download: Binary packages available
- [ ] Zero-config: Works out of box with defaults
- [ ] Self-updates: Check & notify for updates
- Measurement: Installation success rate

#### RNF.8.2 Configuration
- [ ] Zero-config default: Sensible defaults
- [ ] Easy customization: Config file
- [ ] Validation: On startup
- [ ] Error messages: Clear if invalid
- [ ] Documentation: Config reference complete
- Measurement: Configuration error rate

#### RNF.8.3 Monitoring & Observability
- [ ] Health check: `auto-gemini status`
- [ ] Metrics export: Prometheus format (optional)
- [ ] Structured logging: JSON format
- [ ] Performance profiling: Built-in tools
- [ ] Error tracking: Error database
- Measurement: Monitoring completeness

#### RNF.8.4 Updates & Patches
- [ ] Auto-check: Weekly for updates
- [ ] Non-breaking: Install instantly
- [ ] Breaking changes: Clear migration guide
- [ ] Rollback: Previous version available
- [ ] Changelog: Maintained & clear
- Measurement: Update deployment success rate

---

## CONTRAINTES & LIMITATIONS

### Contraintes Techniques

| Contrainte | Valeur | Justification |
|-----------|--------|---------------|
| **API Rate Limit** | 1000 req/day (free) | Gemini free tier |
| **Max Token Window** | 1M tokens | Gemini API limit |
| **Max Execution Time** | 30 secondes | Security & resource |
| **Sandbox Memory** | 256MB | Prevent DoS |
| **Max File Size** | 100MB | Processing limit |
| **Session Retention** | 1 year | Storage management |
| **Max Concurrent Sessions** | 50 (configurable) | Resource constraint |
| **CLI Response Time** | 2 secondes | UX expectation |

### Limitations Fonctionnelles

1. **Pas de GUI**: CLI-only (par design, pas limitation)
2. **Pas de real-time collab**: Single-user focus
3. **Pas de GPU**: CPU-only operations
4. **Pas de database direct**: SQLite seulement (no cloud DB)
5. **Pas de deployment auto**: User gère déploiement
6. **Pas de web interface**: CLI/API seulement
7. **Pas de live debugging**: Post-execution analysis seulement

### Limitations de Performance

- Tâches complexes (>1000 LOC): 10-30 secondes
- Large context (>500K tokens): +latency
- 50+ sessions: Dégradation ~10%
- Gros projects (>1M LOC): Analyse lente

### Limitations du Modèle

- Gemini Flash: Moins puissant que Sonnet (pour complex reasoning)
- Rate limits: 1000 req/24h free tier (paid: higher)
- Context: 1M max (some issues need more)
- Token pricing: Incrementé avec usage

---

## CAS D'USAGE

### UC.1 Quick Code Refactoring
**Acteur**: Solo Developer  
**Timeline**: <5 minutes  
**Flux**:
```bash
auto-gemini task new "Refactor getUserById to async/await"
auto-gemini task run <id> --watch
```
**Résultat**: Code refactorisé, tests passent, PR créée  
**Acceptation**: PR fonctionnelle, code review possible

---

### UC.2 Bug Fix with Testing
**Acteur**: Developer  
**Timeline**: 10-20 minutes  
**Flux**:
```bash
auto-gemini task new "Fix failing test in auth.test.ts"
auto-gemini task run <id>  # Auto-runs test, fixes, repeats
auto-gemini session show <id> --summary
```
**Résultat**: Test passing, no regressions  
**Acceptation**: CI/CD green, no new failures

---

### UC.3 API Documentation Generation
**Acteur**: Tech Lead  
**Timeline**: 15-30 minutes  
**Flux**:
```bash
auto-gemini task new "Generate API documentation with examples"
auto-gemini task run <id>
# Auto-generates README, OpenAPI spec, examples
```
**Résultat**: Docs complete, publishable  
**Acceptation**: All endpoints documented, examples work

---

### UC.4 Learning New Codebase
**Acteur**: Junior Developer  
**Timeline**: 30-60 minutes  
**Flux**:
```bash
auto-gemini chat  # Interactive mode
You> Explain the architecture of this project
Agent> [Analysis & explanation]
You> Add a new feature that follows this pattern
Agent> [Implementation following patterns]
```
**Résultat**: Understands codebase, feature added  
**Acceptation**: Code review passes, follows conventions

---

### UC.5 Multi-Task Development
**Acteur**: Team  
**Timeline**: 2-4 hours  
**Flux**:
```bash
# Session A: Backend models
auto-gemini task new "Create database models"
auto-gemini task run task-a &

# Session B: Frontend components (parallel)
auto-gemini session switch workspace-b
auto-gemini task new "Create React components"
auto-gemini task run task-b &

# Session C: Integration
auto-gemini task new "Integrate frontend & backend"
auto-gemini task run task-c  # Uses results from A & B
```
**Résultat**: Modules intégrés, système fonctionnel  
**Acceptation**: All modules working together

---

## ARCHITECTURE SERVICES

### Service 1: Agent Core Engine
- **Responsabilité**: Task reasoning, decomposition, orchestration
- **Inputs**: Task description, project context
- **Outputs**: Execution trace, results
- **Key methods**:
  - `analyzeTask()`: Parse & understand task
  - `decomposeTask()`: Break into steps
  - `executeTask()`: Orchestrate execution
  - `validateResults()`: Check outputs
- **Dependencies**: Gemini API, Sandbox, Session Manager
- **Protocol**: Async/event-based

### Service 2: Code Generation Service
- **Responsabilité**: Generate code via Gemini
- **Inputs**: Requirements, context, examples
- **Outputs**: Code string, metadata
- **Key methods**:
  - `generateCode()`: Call Gemini API
  - `validateCode()`: Check syntax
  - `explainCode()`: Generate explanation
- **Dependencies**: Gemini API, Config
- **Protocol**: REST to Gemini API

### Service 3: Code Execution Service
- **Responsabilité**: Safe code execution
- **Inputs**: Code, language, timeout
- **Outputs**: stdout, stderr, exit code
- **Key methods**:
  - `executeCode()`: Run in sandbox
  - `monitorExecution()`: Track progress
  - `captureOutput()`: Collect results
- **Dependencies**: Node.js, Python, Bash, Docker
- **Protocol**: Subprocess/IPC

### Service 4: Session Manager
- **Responsabilité**: Session lifecycle & persistence
- **Inputs**: Session commands (create, load, save)
- **Outputs**: Session data
- **Key methods**:
  - `createSession()`: Initialize
  - `loadSession()`: Recover from disk
  - `saveSession()`: Persist to DB
  - `listSessions()`: Query all
- **Dependencies**: SQLite, File system
- **Protocol**: Sync queries

### Service 5: Git Integration Service
- **Responsabilité**: Git & GitHub operations
- **Inputs**: Git commands (clone, commit, push)
- **Outputs**: Git results, GitHub API responses
- **Key methods**:
  - `cloneRepo()`: Get project
  - `createBranch()`: Feature branch
  - `commitChanges()`: Stage & commit
  - `createPR()`: GitHub pull request
- **Dependencies**: Git CLI, GitHub API
- **Protocol**: Shell + REST

### Service 6: CLI Interface (Oclif)
- **Responsabilité**: User-facing commands
- **Inputs**: CLI commands & arguments
- **Outputs**: Formatted output, logs
- **Key methods**:
  - `parseCommand()`: Parse CLI
  - `formatOutput()`: Pretty-print results
  - `handleErrors()`: User-friendly errors
  - `showProgress()`: Progress indicators
- **Dependencies**: Agent Core, Session Manager
- **Protocol**: Command dispatch

### Service 7: Configuration & Logging
- **Responsabilité**: Config & observability
- **Inputs**: Config file, env vars, CLI flags
- **Outputs**: Logger instance, config object
- **Key methods**:
  - `loadConfig()`: Merge sources
  - `validateConfig()`: Check schema
  - `createLogger()`: Initialize logging
  - `exportMetrics()`: Collect metrics
- **Dependencies**: File system, logger library
- **Protocol**: Sync access

---

## INTÉGRATIONS EXTERNES

### INT.1 Gemini API (Google)
- **Purpose**: Code generation & reasoning
- **Endpoint**: `https://generativelanguage.googleapis.com/v1beta/models`
- **Authentication**: API key in header
- **Models**:
  - `gemini-2.0-flash`: Default (fast, cheap)
  - `gemini-2.0-pro`: Complex reasoning
- **Rate Limits**: 1000 req/day (free tier), 100K/month (paid)
- **Fallback**: Queue requests, retry with exponential backoff
- **Pricing**: Free for limited usage, $0.075/M input, $0.30/M output tokens

### INT.2 GitHub API
- **Purpose**: Repository operations, PRs, issues
- **Endpoint**: `https://api.github.com`
- **Authentication**: PAT (Personal Access Token)
- **Operations**:
  - Clone repos
  - Create/manage branches
  - Create PRs with descriptions
  - Comment on issues
  - Manage labels, milestones
- **Rate Limits**: 1000 req/hr (public), 5000 (authenticated)
- **Fallback**: Local git only, sync when API available

### INT.3 npm Registry
- **Purpose**: Dependency management
- **Endpoint**: `https://registry.npmjs.org`
- **Operations**:
  - Search packages
  - Install dependencies
  - Update packages
  - Audit vulnerabilities
- **Fallback**: Use cached versions

### INT.4 MCP Servers
- **Purpose**: Extended functionality via Model Context Protocol
- **Protocol**: JSON-RPC 2.0
- **Examples**:
  - GitHub MCP server
  - Database MCP server
  - Slack MCP server
- **Fallback**: Direct API calls if MCP unavailable

### INT.5 Docker (Optional)
- **Purpose**: Advanced code isolation
- **Usage**: For complex projects needing full OS isolation
- **Detection**: Auto-detect if Docker available
- **Fallback**: Use subprocess sandbox if Docker unavailable

---

## ROADMAP PRODUIT

### Phase 1: MVP (v1.0 - May 2026)
**Features**:
- ✅ Task execution with Gemini
- ✅ Code generation & testing
- ✅ Session persistence
- ✅ Git integration (CLI commands)
- ✅ CLI interface (Oclif)
- ✅ Basic error handling
- ✅ Configuration system
- ✅ Logging & monitoring

**Timeline**: 14 weeks (from Jan 6, 2026)  
**Metrics**: >80% test coverage, p95 latency <2s, zero critical bugs

---

### Phase 2: Polish & Optimization (v1.1 - July 2026)
**Features**:
- [ ] Advanced reasoning (chain-of-thought)
- [ ] Performance optimization (caching, streaming)
- [ ] Web-based dashboard (optional)
- [ ] Team collaboration features
- [ ] Plugin system foundation
- [ ] Comprehensive error recovery
- [ ] Enhanced documentation
- [ ] Community examples

**Timeline**: 8 weeks  
**Metrics**: p99 latency <5s, 50+ concurrent sessions support

---

### Phase 3: Enterprise (v2.0 - Q4 2026)
**Features**:
- [ ] Self-hosted deployment
- [ ] SSO & SAML
- [ ] Advanced audit logging
- [ ] SLA monitoring
- [ ] Multi-tenant support
- [ ] Custom model support
- [ ] Analytics dashboards
- [ ] Data export/deletion

**Timeline**: 12 weeks  
**Target**: Enterprise customers

---

### Phase 4: Ecosystem (v3.0 - 2027+)
**Features**:
- [ ] Multi-agent orchestration
- [ ] Specialized agents (frontend, backend, DevOps)
- [ ] Learning system (improve over time)
- [ ] Plugin marketplace
- [ ] Open-source ecosystem
- [ ] Community contributions
- [ ] Advanced workflows

**Timeline**: Ongoing  
**Vision**: Agentic AI ecosystem

---

## DATA MODELS

### Task
```json
{
  "id": "task-uuid",
  "sessionId": "session-uuid",
  "description": "string",
  "status": "pending|running|completed|failed|cancelled",
  "createdAt": "ISO8601",
  "startedAt": "ISO8601",
  "completedAt": "ISO8601",
  "executionTrace": ["step1", "step2", "step3"],
  "output": {
    "code": "string",
    "results": {}
  },
  "errors": [
    {
      "type": "string",
      "message": "string",
      "recovery": "string"
    }
  ],
  "artifacts": [
    {
      "path": "file path",
      "type": "code|test|docs",
      "content": "string"
    }
  ],
  "metrics": {
    "duration": 1234,
    "tokensUsed": { "input": 1000, "output": 500 },
    "apiCalls": 3
  }
}
```

### Session
```json
{
  "id": "session-uuid",
  "workspacePath": "/path/to/project",
  "createdAt": "ISO8601",
  "updatedAt": "ISO8601",
  "lastActivityAt": "ISO8601",
  "tasks": ["task-id1", "task-id2"],
  "context": {
    "projectStructure": {},
    "recentFiles": [],
    "variables": {}
  },
  "config": {
    "model": "gemini-2.0-flash",
    "temperature": 0.5
  },
  "metadata": {
    "totalDuration": 3600,
    "taskCount": 5,
    "successCount": 4,
    "totalTokens": 150000
  }
}
```

### Execution Event
```json
{
  "timestamp": "ISO8601",
  "level": "debug|info|warn|error",
  "message": "string",
  "context": {
    "taskId": "string",
    "sessionId": "string",
    "phase": "string"
  },
  "sourceFile": "string",
  "lineNumber": 123
}
```

---

## GLOSSAIRE

| Terme | Définition |
|-------|-----------|
| **Agent** | Autonomous AI system using Gemini for reasoning & generation |
| **Task** | High-level objective given to the agent |
| **Session** | Execution context with persistent state |
| **Artifact** | Generated file, code, or documentation |
| **Sandbox** | Isolated environment for safe code execution |
| **MCP** | Model Context Protocol for tool integration |
| **PAT** | Personal Access Token for GitHub authentication |
| **Token** | Atomic unit of text (word or subword) |
| **Rate Limit** | API call quota (1000 req/24h for Gemini free) |
| **Context Window** | Max tokens per API call (1M for Gemini) |

---

## CONTACT & SUPPORT

**Repository**: https://github.com/agbruneau/auto-gemini-cli  
**Issues**: GitHub Issues  
**Discussions**: GitHub Discussions  
**Documentation**: `/docs` directory  
**License**: MIT  

---

**Document Version**: 1.0  
**Last Updated**: 2 janvier 2026  
**Status**: Complete & Ready for Implementation  
**Approval**: Architecture Review Completed