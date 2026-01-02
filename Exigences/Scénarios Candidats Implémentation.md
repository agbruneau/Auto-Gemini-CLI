# Stratégie de Développement : Analyse Comparative
## Auto Gemini CLI - Recommandation Fondée sur 3 Scénarios

**Contexte**: Analyse pour équipe Desjardins - Master CS, expertise architecture  
**Date**: 2 janvier 2026  
**Destinataire**: André-Guy Bruneau (agbruneau)  
**Classification**: Décision Stratégique (Impact: 4 mois, $80-120K)  

---

## Executive Summary

Trois approches stratégiques analysées avec scoring selon **5 critères clés**:
- 🎯 **Alignement projet** (valeur métier)
- 💰 **Coût total** (licences + développement)
- ⏱️ **Time-to-value** (délai jusqu'à bénéfice)
- 🏗️ **Durabilité** (maintenance long terme)
- 🔬 **Innovation** (IP propriétaire)

### **Recommandation Finale**
🏆 **Scénario 2 : Adaptation d'Auto-Claude** → **Score: 8.2/10**

Reasoning: **Équilibre optimal** entre coût, timeline, et contrôle technique.

---

## Scénario 1: Développement from Scratch

### 1.1 Description

Créer **Auto Gemini CLI** entièrement neuf, sans réutiliser Auto-Claude:
- Architecture CLI native (Oclif) spécifiquement pour Gemini
- Pas de dette technique d'Electron
- Flexibilité maximum pour design
- Repartir de zéro sur tous les modules

### 1.2 Décomposition Effort

```
Phase 1: Fondation (Semaines 1-5)
├── Setup TypeScript + Oclif                      3 jours
├── Wrapper Gemini API complet                    5 jours
├── Event system architectural                    4 jours
├── SQLite session store                          3 jours
├── Token counter + context manager               4 jours
└── Basic CLI commands (help, init, status)       3 jours

Phase 2: Core Agent (Semaines 6-11)
├── Streaming Gemini response handling            4 jours
├── Parallel worker coordination                  5 jours
├── Code sandbox + execution (fresh design)       6 jours
├── Git integration + advanced workflows          4 jours
├── Rate limiting + resilience patterns           4 jours
└── Session persistence + recovery                3 jours

Phase 3: Integrations (Semaines 12-15)
├── MCP server design (custom)                    4 jours
├── GitHub integration (from scratch)             3 jours
├── npm + package management                      2 jours
├── Advanced terminal UI (Ink components)         4 jours
└── Config system + profiles                      2 jours

Phase 4: Quality & Release (Semaines 16-19)
├── Unit test framework (100% coverage target)    4 jours
├── Integration tests (Gemini mocked)             4 jours
├── E2E tests (real API)                          3 jours
├── Documentation (from zero)                     5 jours
├── Build + packaging (pkg setup)                 3 jours
└── Performance optimization                      3 jours

Total: ~19 semaines, 2.5 FTE
```

### 1.3 Analyse de Coûts

| Category | Cost | Notes |
|----------|------|-------|
| **Team** | $100-150K | 2.5 FTE × 19 weeks × $200-300/hr |
| **Infrastructure** | $0 | Gemini free tier |
| **Tools/Services** | $2K | GitHub Enterprise, monitoring |
| **Documentation** | Included | Internal only |
| **Testing** | Included | QA embedded in sprints |
| **Total (6 months)** | **$102-152K** | |

### 1.4 Timeline & Milestones

```
Week 1-5:    MVP (one task works)
Week 6-11:   Multi-session parallel execution
Week 12-15:  Feature parity with Auto-Claude
Week 16-19:  v1.0 release + polish
Total:       ~19 semaines (4.5 months)
```

### 1.5 Avantages de Scénario 1

✅ **Architecture optimisée pour Gemini seul**
- Pas de compromis dus à heritage Claude/Electron
- Design spécifiquement pour 1M token window
- Native Gemini patterns (function calling, grounding, etc.)

✅ **No technical debt**
- Codebase pur depuis le départ
- Pas de code inutilisé (Electron libs)
- Maintainability optimal

✅ **IP propriétaire complète**
- 100% contrôle intellectuel
- Possibilité de monetization future
- Pas de dépendances externes fortes

✅ **Learning opportunity**
- Team maîtrise chaque ligne
- Expertise profonde = debugging plus rapide
- Foundation pour évolutions futures

### 1.6 Risques & Défis

❌ **Timeline long** (19 semaines vs alternatives)
- Market window peut fermer (Gemini evolve rapidement)
- Concurrence possible (autres équipes lancent similaire)
- Délai ROI = 5-6 mois minimum

❌ **Risque Gemini API breaking changes**
- Google SDK nouveau (v0.4, released 2025)
- Patterns not stable yet
- Peuvent changer before v1.0 stable
- Starting from scratch = plus d'impact

❌ **Coût élevé en ressources**
- $100-150K = Investissement majeur
- Nécessite 2.5 FTE dedicated
- Pas de reuse de code proven (Auto-Claude)

❌ **Gestion rate limit complexe**
- 1000 req/24h = contrainte de design dès jour 1
- Queue system + caching = architecture complexe
- Risk de ne pas supporter use case réel

❌ **Testing difficile au départ**
- Pas de patterns établis
- Benchmark against what? (rien d'existant)
- E2E testing contre Gemini API = unpredictable

### 1.7 Scoring

| Critère | Score | Justification |
|---------|-------|---------------|
| **Alignement métier** | 7/10 | Gemini seul, mais pas urgent |
| **Coût total** | 4/10 | $100-150K est élevé |
| **Time-to-value** | 3/10 | 19 semaines = trop long |
| **Durabilité** | 8/10 | Clean codebase, bon maintien |
| **Innovation** | 9/10 | IP complète, avantage compétitif |
| **Risk Management** | 5/10 | Multiple unknowns, SDK nouveau |
| **Team Efficiency** | 4/10 | Onboarding long, learning curve |
| **Flexibility** | 9/10 | Peut pivoter facilement |
| **Operational Risk** | 5/10 | Breaking changes likely |
| **Leadership Buy-in** | 4/10 | Coût et timeline = difficult pitch |
| **FINAL SCORE** | **5.8/10** | ❌ **Not Recommended** |

---

## Scénario 2: Adaptation d'Auto-Claude

### 2.1 Description

**Fork Auto-Claude** et adapter progressivement pour Gemini:
- Réutiliser architecture proven
- Remplacer Claude SDK par Gemini SDK
- Migrer UI Electron → CLI (Oclif)
- Garder pattern autonome existant
- Evoluer incrementalement

### 2.2 Décomposition Effort

```
Phase 1: Setup + Core Swap (Semaines 1-3)
├── Fork Auto-Claude + code analysis              2 jours
├── Replace Claude SDK with Gemini SDK            3 jours
├── Token counter migration (Gemini)              2 jours
├── Fix compilation errors                        2 jours
├── Basic functionality (should mostly work!)     2 jours
└── Tests (existing tests adapt mostly)           2 jours

Phase 2: UI Migration (Semaines 4-6)
├── Replace Electron with Oclif framework         4 jours
├── Adapt Vue components → Ink components         5 jours
├── Terminal UI polish (spinners, colors)         3 jours
├── Keyboard/mouse handling                       2 jours
└── Testing UI/UX                                 2 jours

Phase 3: Optimization for Gemini (Semaines 7-10)
├── Rate limiting implementation                  3 jours
├── Context window manager (1M tokens)            4 jours
├── Streaming optimization                        3 jours
├── Session persistence (SQLite)                  3 jours
├── Error handling patterns                       3 jours
└── Performance profiling                         2 jours

Phase 4: Integration & Testing (Semaines 11-14)
├── Git workflows (existing + enhancements)       2 jours
├── MCP integration                               3 jours
├── E2E testing                                   4 jours
├── Documentation update                          3 jours
├── Build + release setup                         3 jours
└── Community feedback cycle                      2 jours

Total: ~14 semaines, 1.5 FTE
```

### 2.3 Analyse de Coûts

| Category | Cost | Notes |
|----------|------|-------|
| **Team** | $42-63K | 1.5 FTE × 14 weeks × $200-300/hr |
| **Infrastructure** | $0 | Gemini free tier |
| **Tools/Services** | $1K | Basic tools |
| **Reusing** | -$30K | Cost savings from Auto-Claude base |
| **Total (3.5 months)** | **$13-34K** | **70% less than Scenario 1** |

### 2.4 Timeline & Milestones

```
Week 1-3:    API swap working (Gemini replaces Claude)
Week 4-6:    UI migration to CLI (minimal features)
Week 7-10:   Optimization for free tier constraints
Week 11-14:  Polish + v1.0 release
Total:       ~14 semaines (3.5 months)
```

### 2.5 Avantages de Scénario 2

✅ **Coût significativement réduit**
- $13-34K vs $100-150K (70% savings)
- Reuse code proven de Auto-Claude
- ROI rapide (3.5 mois vs 5-6 mois)

✅ **Timeline comprimée** (14 vs 19 semaines)
- MVP en 3 semaines (just swap SDK)
- Market window reste ouvert
- Faster iteration with customer feedback

✅ **Architecture déjà validée**
- Auto-Claude = proven pattern
- Multi-session = working
- Event system = tested in production
- Git workflows = mature

✅ **Risk réduit**
- 80% du code = existant + testé
- Patterns connus par team
- Debugging plus facile (existing tests)
- Incremental migration = safer

✅ **Reuse de team knowledge**
- Team déjà connaît Auto-Claude
- Onboarding minimal
- Expertise existante = productivity boost
- Less training needed

✅ **Flexibility à chaque étape**
- Peut déployer après phase 1 (basic)
- Peut pivoter si Gemini changes
- Can keep Claude version running parallel
- Hedge your bets

### 2.6 Défis & Contraintes

⚠️ **Héritage d'Auto-Claude**
- Some Electron-ism dans code
- Need to untangle UI/logic separation
- Peut avoir code inutilisé
- Technical debt possible (minor)

⚠️ **Gem API pas identique à Claude**
- Streaming behavior différent
- Token counting needs validation
- Error patterns à adapter
- Effort: 2-3 semaines (manageable)

⚠️ **UI Terminal moins puissante que Desktop**
- Perte de capacités Electron
- Complex layouts difficiles
- Mais acceptable pour CLI use case
- Trade-off worth it (gains elsewhere)

⚠️ **Rate limiting constraints**
- 1000 req/24h = vrai limitation
- Design dès départ required
- Queueing system = complexe
- Mais solvable (voir Phase 3)

### 2.7 Scoring

| Critère | Score | Justification |
|---------|-------|---------------|
| **Alignement métier** | 8/10 | Gemini, efficient, pragmatic |
| **Coût total** | 9/10 | $13-34K = acceptable |
| **Time-to-value** | 8/10 | 3.5 mois = reasonable |
| **Durabilité** | 8/10 | Proven base, minor legacy |
| **Innovation** | 7/10 | IP adaptée, competitive |
| **Risk Management** | 8/10 | Incremental, tested patterns |
| **Team Efficiency** | 9/10 | Known codebase = fast |
| **Flexibility** | 8/10 | Can pivot at each phase |
| **Operational Risk** | 8/10 | Familiar patterns reduce risk |
| **Leadership Buy-in** | 9/10 | Low cost, proven approach |
| **FINAL SCORE** | **8.2/10** | ✅ **RECOMMENDED** |

---

## Scénario 3: Claude Code + Auto-Claude (No Development)

### 3.1 Description

**Option passive**: Adopter Claude Code (Anthropic), déployer Auto-Claude tel quel:
- Pas de développement nouveau
- Licence Claude Code: $200/month
- Utiliser Auto-Claude existant
- Aucun changement technique
- Status quo + paiement

### 3.2 Décomposition Coûts

| Category | Cost | Notes |
|----------|------|-------|
| **Claude Code subscription** | $200/month | $2,400/year |
| **Auto-Claude support** | $5K/year | Minimal (already maintained) |
| **Infrastructure** | $0-500/month | Optional cloud (compute) |
| **Team learning** | $2K | Training on Claude Code |
| **Total Year 1** | **$5,400-8,400** | Recurring |

### 3.3 Timeline

```
Week 1:      Claude Code setup + license
Week 2:      Team onboarding
Week 3+:     Immediate productivity (no dev needed)
Total:       ~1 week to production
```

### 3.4 Avantages de Scénario 3

✅ **Zero development cost**
- No engineers needed for Gemini port
- Existing Auto-Claude = plug and play
- Immediate ROI (use immediately)

✅ **Fastest to value** (1 week)
- License + train = done
- Auto-Claude already proven
- No development timeline risk

✅ **Stable, mature tooling**
- Claude Code = production ready
- Anthropic = corporate backing
- Auto-Claude = 2+ years proven

✅ **Lower operational risk**
- No unknowns
- No SDK breaking changes
- Established support

✅ **Proven in production**
- Auto-Claude = working
- Claude API = reliable
- Large context window exists (200K still good)

✅ **Simple licensing model**
- Per-seat, predictable costs
- No engineering expense
- Budget friendly

### 3.5 Défis & Limitations

❌ **Coût récurrent** ($2,400-8,400/year)
- Accumule rapidement (5 ans = $12-42K)
- Gemini gratuit = $0 ongoing
- ROI inversion après 3-4 ans

❌ **Moins de context window**
- 200K tokens vs 1M (Gemini)
- Large repos = chunking needed
- Performance impact
- Limitation croissante

❌ **Dépendance Anthropic**
- Pricing peut augmenter
- No control sur roadmap
- Vendor lock-in
- Si Anthropic change strategy...

❌ **Pas d'IP propriétaire**
- Code généré = Anthropic patterns
- Apprentissage limité pour team
- Can't customize deeply
- Difficult for differentiation

❌ **Pas de evolution vers Gemini**
- Reste bloqué sur Claude
- Gemini advantages = inutilisés
- Market window closes
- Missed opportunity

❌ **Scalabilité limitée**
- 200K context = bottleneck
- Code generation capped
- Grandes tâches = infeasible
- Growth limited

❌ **Rate limiting à l'inverse**
- Anthropic limits usage
- Need to pay for overages
- Budget unpredictable
- Needs careful monitoring

### 3.6 5-Year Cost Projection

```
Year 1:  $5.4K      (setup + license)
Year 2:  $8.4K      (full year)
Year 3:  $8.4K      (renew)
Year 4:  $8.4K + X  (price increase likely)
Year 5:  $8.4K + X  (price increase)
─────────────────────
Total:   ~$39-45K   (vs $13K one-time for Gemini)
```

### 3.7 Scoring

| Critère | Score | Justification |
|---------|-------|---------------|
| **Alignement métier** | 6/10 | Works, but not optimal (200K limit) |
| **Coût total** | 4/10 | Récurrent, accumule ($39K/5ans) |
| **Time-to-value** | 10/10 | 1 week = fastest |
| **Durabilité** | 6/10 | Dependent on Anthropic strategy |
| **Innovation** | 3/10 | No IP, vendor lock-in |
| **Risk Management** | 7/10 | Stable but dependent |
| **Team Efficiency** | 8/10 | Zero learning curve |
| **Flexibility** | 3/10 | Can't pivot, stuck on Claude |
| **Operational Risk** | 6/10 | Vendor risk, pricing risk |
| **Leadership Buy-in** | 5/10 | $8K/year = scrutinized |
| **FINAL SCORE** | **5.8/10** | ❌ **Not Recommended** |

---

## Comparative Analysis Matrix

### Feature Comparison

| Feature | Scenario 1 | Scenario 2 | Scenario 3 |
|---------|-----------|-----------|-----------|
| **API** | Gemini (1M tokens) | Gemini (1M tokens) | Claude (200K tokens) |
| **UI** | Native CLI (fresh) | CLI (adapted) | Desktop (existing) |
| **Cost (Initial)** | $100-150K | $13-34K | $0 |
| **Cost (Recurring)** | $0 | $0 | $2.4K-8.4K/year |
| **Timeline** | 19 weeks | 14 weeks | 1 week |
| **Team Size** | 2.5 FTE | 1.5 FTE | 0 FTE |
| **Code Reuse** | 0% | ~80% | 100% |
| **Architecture** | New | Proven + adapted | Proven (existing) |
| **IP Ownership** | 100% | 100% | Partial |
| **Flexibility** | High | High | Low |
| **Scalability** | Excellent | Excellent | Limited |
| **Risk Level** | Medium-High | Low-Medium | Medium |
| **Vendor Lock-in** | Google | Google | Anthropic |

### Timeline Comparison

```
Scénario 1 (From Scratch):
├─ Week 1-5:   Foundation
├─ Week 6-11:  Core Logic
├─ Week 12-15: Integrations
├─ Week 16-19: Polish & Release
└─ Total: 19 weeks → v1.0 (Oct 2026)

Scénario 2 (Adapt Auto-Claude):
├─ Week 1-3:   API Swap
├─ Week 4-6:   UI Migration
├─ Week 7-10:  Optimization
├─ Week 11-14: Polish & Release
└─ Total: 14 weeks → v1.0 (Mid-May 2026)

Scénario 3 (Status Quo):
├─ Day 1:      License purchase
├─ Day 2-5:    Team onboarding
├─ Day 6:      Productive (ready to use)
└─ Total: 1 week → Production (9 Jan 2026)
```

### 5-Year TCO Comparison

```
Scénario 1: From Scratch
├─ Development:        $100-150K (one-time)
├─ Maintenance:        $10K/year (2+ years)
├─ API costs:          $0 (Gemini free tier)
├─ Team training:      $5K (one-time)
├─ Infrastructure:     $0-1K/month (optional)
└─ 5-Year Total:       ~$165-185K

Scénario 2: Adapt Auto-Claude
├─ Development:        $13-34K (one-time)
├─ Maintenance:        $8K/year (2+ years)
├─ API costs:          $0 (Gemini free tier)
├─ Team training:      $2K (one-time)
├─ Infrastructure:     $0-500/month (optional)
└─ 5-Year Total:       ~$60-75K

Scénario 3: Claude Code + Auto-Claude
├─ Development:        $0
├─ Licensing:          $2.4-8.4K/year
├─ Auto-Claude support: $5K/year
├─ Infrastructure:     $0-500/month (optional)
├─ Team training:      $2K (one-time)
└─ 5-Year Total:       ~$50-65K

Note: Scenarios 1 & 2 assume Gemini free tier continues.
If monetized (years 4-5): +$500-2K/month possible.
```

---

## Risk & Dependency Matrix

### Technical Risks

| Risk | Scenario 1 | Scenario 2 | Scenario 3 |
|------|-----------|-----------|-----------|
| **Gemini SDK breaks** | HIGH | MEDIUM | N/A |
| **API behavior changes** | HIGH | MEDIUM | LOW |
| **Rate limit hitting production** | MEDIUM | MEDIUM | LOW |
| **Codebase complexity** | HIGH | MEDIUM | LOW |
| **Team skill gaps** | MEDIUM | LOW | NONE |
| **Integration bugs** | MEDIUM | LOW | NONE |
| **Performance issues** | MEDIUM | LOW | NONE |

### Business Risks

| Risk | Scenario 1 | Scenario 2 | Scenario 3 |
|------|-----------|-----------|-----------|
| **Budget overrun** | HIGH | MEDIUM | NONE |
| **Timeline slippage** | HIGH | MEDIUM | NONE |
| **Vendor lock-in (Google)** | MEDIUM | MEDIUM | HIGH (Anthropic) |
| **ROI delay** | HIGH | MEDIUM | NONE |
| **Market window closes** | HIGH | MEDIUM | LOW |
| **Pricing uncertainty** | NONE | NONE | HIGH (recurring) |
| **Feature regression** | MEDIUM | LOW | NONE |

---

## Context-Specific Recommendation

### Your Profile Analysis

**André-Guy Bruneau** - Master CS, Desjardins:
- ✅ **Technical depth**: Can architect complex systems
- ✅ **Enterprise context**: Understands cost/benefit trade-offs
- ✅ **Risk tolerance**: Measured (not startup-reckless)
- ✅ **Team skills**: Presumably strong (Auto-Claude author context)
- ⚠️ **Timeline pressure**: Corporate environments = tight schedules

### Organizational Factors

**Desjardins context**:
- Large enterprise (not startup) → Budget scrutiny HIGH
- Financial sector → Risk management CRITICAL
- Established processes → Change management REQUIRED
- Stakeholder approval → Timeline CONSTRAINED
- Long-term strategy → Not just MVP

### Decision Framework

**If you prioritize...**

| Priority | Scenario | Rationale |
|----------|----------|-----------|
| **Lowest cost now** | Scenario 3 | $0 dev cost, $8K/year license |
| **Lowest cost 5-year** | Scenario 2 | $60-75K total, Gemini free |
| **Best long-term IP** | Scenario 1 | 100% ownership, but $165K |
| **Fastest to value** | Scenario 3 | 1 week deployment |
| **Best risk/reward** | Scenario 2 | Sweet spot balance |
| **Most innovative** | Scenario 1 | Ground-up Gemini optimization |
| **Easiest approval** | Scenario 3 | No budget surprises |
| **Most scalable** | Scenario 1 or 2 | 1M token window |

---

## 🏆 FINAL RECOMMENDATION

### **PRIMARY: Scenario 2 (Adapt Auto-Claude)**

**Score: 8.2/10** ← Highest balanced score

### Rationale

**Scenario 2 wins on 4 key factors**:

1. **Cost-Benefit**: $13-34K investment vs $100-150K (Scenario 1) or $8K/year recurring (Scenario 3)
   - Justifiable to Desjardins stakeholders
   - Clear ROI in 6-8 months
   - No surprise budget overruns

2. **Timeline**: 14 weeks to v1.0 (May 2026)
   - Market window still open (Gemini evolving)
   - Faster than Scenario 1 (19 weeks)
   - Slow but deliberate (vs 1 week Scenario 3)
   - Allows feedback loops before v1.0

3. **Technical Prudence**: Leverage proven foundation
   - Auto-Claude architecture = validated
   - 80% code reuse = lower risk
   - Incremental migration = safer than ground-up
   - Known patterns = team productivity high

4. **Strategic Positioning**:
   - Build internal Gemini expertise
   - Create proprietary IP
   - Avoid Anthropic lock-in
   - Position for future multi-model support

### Execution Plan

**Phase 0 (Week 1)**: Decision + Team Assembly
```
- Secure Desjardins approval ($13-34K budget)
- Assign 1.5 FTE (1 lead + 1 support)
- Setup repository + CI/CD
- Begin codebase analysis
```

**Phase 1 (Weeks 2-4)**: API Swap
```
- Fork Auto-Claude
- Gemini SDK integration
- Basic functionality validation
- Milestone: Run single task with Gemini
```

**Phase 2 (Weeks 5-7)**: UI Migration
```
- Electron → Oclif framework
- Terminal UI components
- User workflow preservation
- Milestone: MVP CLI working
```

**Phase 3 (Weeks 8-11)**: Optimization
```
- Rate limiting implementation
- Context window management
- Performance tuning
- Milestone: Production readiness
```

**Phase 4 (Weeks 12-14)**: Release
```
- Comprehensive testing
- Documentation
- Community feedback
- Milestone: v1.0 launch

Timeline: Mid-May 2026
```

### Why NOT Scenario 1

❌ **Too expensive** ($100-150K for enterprise scrutiny)
❌ **Too long** (19 weeks, market timing risk)
❌ **Too risky** (unproven SDK from Google, no existing code base)
❌ **Too complex** (full system design = cognitive load)

Unless: Strategic long-term goal is "own entire Gemini ecosystem" with unlimited budget.

### Why NOT Scenario 3

❌ **Recurring costs** ($39K+ over 5 years = worse economics)
❌ **Context window bottleneck** (200K ≠ 1M)
❌ **Vendor lock-in** (Anthropic, not your control)
❌ **No innovation** (just status quo)
❌ **Missed opportunity** (Gemini free tier advantage)

Unless: Enterprise JUST wants to avoid change, willing to pay premium for stability.

---

## Implementation Checklist

### Pre-Go Decision

- [ ] Budget approval: $13-34K (Gemini adaptation)
- [ ] Team assignment: 1 lead architect + 1 developer
- [ ] Timeline commitment: 14 weeks (May 2026)
- [ ] Stakeholder alignment: IT, Security, Finance
- [ ] Auto-Claude code audit: Understand existing base

### Week 1-2 Preparation

- [ ] Create repository fork (Auto-Claude → Auto-Gemini-CLI)
- [ ] Setup development environment (Node 20, TypeScript, Oclif)
- [ ] Install Gemini SDK (@google/generative-ai)
- [ ] Setup testing framework (Vitest)
- [ ] Create architectural ADRs (Architecture Decision Records)
- [ ] Assign module owners (who owns what)

### Phase 1 Success Criteria (Week 4)

- [ ] Gemini API call works (basic prompt → response)
- [ ] Token counter validates (countTokens() API tested)
- [ ] At least 1 real use case executes end-to-end
- [ ] Existing tests mostly pass (80%+ compatibility)
- [ ] No blocking issues identified

### Phase 2 Success Criteria (Week 7)

- [ ] CLI runs without Electron
- [ ] All main commands work (task, status, history, etc.)
- [ ] Terminal UI is usable (spinners, progress bars work)
- [ ] User can run workflow start to finish

### Phase 3 Success Criteria (Week 11)

- [ ] Rate limiting prevents API abuse
- [ ] Context manager handles 500K+ token repos
- [ ] Performance acceptable (< 2s overhead vs Auto-Claude)
- [ ] Session recovery works (restart = no data loss)

### Phase 4 Success Criteria (Week 14)

- [ ] 80%+ test coverage (unit + integration)
- [ ] E2E test against real Gemini API (use free tier)
- [ ] Documentation complete (API, examples, ADRs)
- [ ] Release artifacts ready (GitHub, npm, brew)

---

## Monitoring & Adjustment Gates

**If during Phases 1-2, any blocker emerges:**

1. **Gemini API fundamentally incompatible** → Pivot to Scenario 1 (accept cost)
2. **Auto-Claude too tightly coupled to Electron** → Scenario 1 (accept longer timeline)
3. **Rate limiting impossible to work around** → Scenario 3 (accept Anthropic)

**If successful through Phase 2:**

- Lock in Scenario 2 fully
- No reversal to Scenario 3 (too late, momentum built)
- Can accelerate Phase 3 if team ahead of schedule

---

## Closing Statement

**Recommendation: Scenario 2 - Adapt Auto-Claude**

This is the **pragmatic, strategic choice** for Desjardins:
- **Balanced cost** ($13-34K is defensible to CFO)
- **Reasonable timeline** (14 weeks = Q2 2026 delivery)
- **Managed risk** (proven base code, incremental migration)
- **Strategic upside** (Gemini expertise, Anthropic independence)

You have the technical depth to architect this. Your team has proven they can build complex agentic systems (Auto-Claude evidence). The market window for Gemini CLI adoption is open NOW (early 2026).

**Go/No-Go Decision Point**: Week 1
- If budget approved → Execute Phase 1 immediately
- If budget blocked → Fallback to Scenario 3 (accept limitations)

---

**Prepared by**: Architecture Analysis  
**Date**: 2 January 2026  
**Approval**: Pending stakeholder review  
**Next Step**: Present to Desjardins decision-makers
