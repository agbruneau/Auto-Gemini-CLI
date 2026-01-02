# Planification Exhaustive d'Exécution

## Adaptation d'Auto-Claude → Auto Gemini CLI

**Version**: 1.0  
**Date**: 2 janvier 2026  
**Horizon**: 14 semaines (3.5 mois)  
**Timeline**: 9 janvier 2026 → 28 avril 2026  
**Équipe**: 1.5 FTE (1 Lead Architect + 1 Full-Stack Dev)  
**Budget**: $13-34K

---

## I. PHASE 0: PRÉ-LANCEMENT (Semaine 0 | 2-8 janvier 2026)

### Objectif

Valider la faisabilité, assembler l'équipe, préparer l'infrastructure.

### 0.1 Tâche: Approbation Budgétaire & Stakeholder Alignment

**Responsable**: Lead Architect (vous)  
**Durée**: 2 jours (2-3 janvier)

**Livrables**:

- [ ] Deck présentation (3 scénarios + recommandation)
- [ ] Budget approval document ($13-34K)
- [ ] Timeline commitment (14 weeks)
- [ ] Risk mitigation plan
- [ ] Signoff de Finance, IT, Security

**Actions**:

```
Jour 1 (2 jan):
├─ Préparer slides executives
├─ Identifier stakeholders clés (Finance, IT, Security)
├─ Prévoir réunion décision
└─ Chiffrer détails budgétaires

Jour 2 (3 jan):
├─ Présenter recommandation Scénario 2
├─ Répondre objections
├─ Obtenir signoff written
└─ Lancer Phase 0 officieusement
```

**Critères de Succès**:

- ✅ Budget approuvé par Finance
- ✅ IT commit sur ressources
- ✅ Security cleared architecture
- ✅ Executive alignment

---

### 0.2 Tâche: Constitution de l'Équipe & Skillset Validation

**Responsable**: Lead Architect  
**Durée**: 1 jour (4 janvier)

**Livrables**:

- [ ] RACI matrix (responsabilités claires)
- [ ] Skill inventory checklist
- [ ] Training plan (si gaps identifiés)
- [ ] Communication channels setup
- [ ] Pair programming schedule

**Détails Équipe**:

```
LEAD ARCHITECT (1.0 FTE)
├─ Role: Architecture, decisions, reviews, Gemini integration
├─ Skills Required:
│  ├─ TypeScript/Node.js (expert)
│  ├─ API design (proven)
│  ├─ System architecture (demonstrated)
│  ├─ Gemini API (learning sprint 2 days)
│  └─ CLI/TUI frameworks (Oclif + Ink basics)
├─ Hours/Week: 40
└─ Person: André-Guy Bruneau (agbruneau)

FULL-STACK DEVELOPER (0.5 FTE - ramp to 1.0 FTE Week 2+)
├─ Role: Implementation, testing, integrations
├─ Skills Required:
│  ├─ TypeScript/Node.js (intermediate+)
│  ├─ Git workflows (strong)
│  ├─ Testing frameworks (Vitest)
│  └─ SQLite + persistence layers
├─ Hours/Week: 20 Week 1, 40 Week 2+
└─ Person: TBD (hire or internal reassign)

OPTIONAL: QA/DevOps (0.2 FTE starting Week 4)
├─ Role: Testing orchestration, build/release
├─ Triggered if: Testing bottlenecks emerge
└─ Hours/Week: 8-10 (part-time)
```

**Actions**:

```
- [ ] Post job description (if hiring) or reassign internally
- [ ] Interview + skill assessment (if new person)
- [ ] Setup onboarding plan
- [ ] Create RACI chart
- [ ] Schedule kickoff meeting
```

**Critères de Succès**:

- ✅ Team fully assigned (no gaps)
- ✅ All critical skills covered
- ✅ RACI matrix signed
- ✅ Communication protocols established

---

### 0.3 Tâche: Infrastructure & Repository Setup

**Responsable**: Full-Stack Dev (with Lead guidance)  
**Durée**: 2 jours (5-6 janvier)

**Livrables**:

- [ ] GitHub fork: `agbruneau/Auto-Gemini-CLI`
- [ ] CI/CD pipeline (.github/workflows)
- [ ] Development environment (local + Docker)
- [ ] Secret management (.env, API keys)
- [ ] Monitoring setup (optional: Sentry, LogRocket)
- [ ] Communication channels (Slack, GitHub Discussions)

**Détails Technique**:

```
REPOSITORY STRUCTURE
auto-gemini-cli/
├── .github/
│   ├── workflows/
│   │   ├── test.yml          # Run tests on push
│   │   ├── build.yml         # Build artifacts
│   │   └── release.yml       # Release to GitHub, npm
│   └── ISSUE_TEMPLATE/
├── src/
│   ├── index.ts              # CLI entry point
│   ├── core/
│   │   ├── agent.ts          # Agent orchestrator
│   │   ├── gemini-client.ts  # Gemini API wrapper
│   │   ├── context-manager.ts
│   │   ├── session-manager.ts
│   │   └── event-bus.ts
│   ├── executors/
│   ├── integrations/
│   ├── ui/
│   ├── storage/
│   └── utils/
├── tests/
│   ├── unit/
│   ├── integration/
│   └── e2e/
├── docs/
│   ├── ARCHITECTURE.md
│   ├── API-REFERENCE.md
│   ├── EXTENDING.md
│   └── DEVELOPMENT.md
├── package.json
├── tsconfig.json
├── jest.config.js (or vitest.config.ts)
├── .env.example
├── .gitignore
└── README.md

CI/CD PIPELINE
├── On push to main:
│  ├─ Run linter (eslint)
│  ├─ Run type check (tsc)
│  ├─ Run unit tests (vitest)
│  └─ Upload coverage
├── On PR:
│  ├─ Run same checks
│  ├─ Block merge if tests fail
│  └─ Require 1 approval
└── On tag release:
   ├─ Run full test suite
   ├─ Build artifacts (Windows, macOS, Linux)
   ├─ Create GitHub Release
   └─ Publish to npm
```

**Actions**:

```
Day 1 (5 jan):
├─ Fork Auto-Claude repository
├─ Create new branch: main (protected)
├─ Setup .env.example with Gemini API key placeholder
├─ Create initial package.json with dependencies
└─ Test local build

Day 2 (6 jan):
├─ Setup GitHub Actions workflows
├─ Configure code scanning (Dependabot)
├─ Setup npm publishing credentials (saved in Secrets)
├─ Create development.md guide
├─ Test CI/CD pipeline with dummy commit
```

**Dependencies to Install**:

```json
{
  "dependencies": {
    "@google/generative-ai": "^0.4.0",
    "oclif": "^4.0.0",
    "ink": "^4.4.0",
    "ink-spinner": "^5.0.0",
    "simple-git": "^3.20.0",
    "eventemitter3": "^5.0.0",
    "better-sqlite3": "^9.0.0",
    "piscina": "^4.4.0",
    "dotenv": "^16.3.0",
    "chalk": "^5.3.0",
    "ora": "^8.0.0",
    "yargs": "^17.7.0",
    "commander": "^11.0.0"
  },
  "devDependencies": {
    "typescript": "^5.3.0",
    "vitest": "^1.0.0",
    "@types/node": "^20.0.0",
    "eslint": "^8.50.0",
    "@typescript-eslint/parser": "^6.5.0",
    "ts-node": "^10.9.0",
    "esbuild": "^0.19.0"
  }
}
```

**Critères de Succès**:

- ✅ Repository fork functional
- ✅ CI/CD pipeline green (all checks pass)
- ✅ Environment variables configured
- ✅ Team can clone + develop locally
- ✅ First successful test build

---

### 0.4 Tâche: Codebase Analysis & Architecture Audit

**Responsable**: Lead Architect  
**Durée**: 2 jours (7-8 janvier)

**Livrables**:

- [ ] Code audit report (Auto-Claude analysis)
- [ ] Dependency map (what's coupled, what's independent)
- [ ] API difference matrix (Claude vs Gemini)
- [ ] Migration strategy document (step-by-step)
- [ ] Risk assessment (what could break)
- [ ] ADR #1: Fork Strategy

**Actions**:

```
Day 1 (7 jan):
├─ Clone Auto-Claude locally
├─ Read codebase:
│  ├─ Agent orchestrator logic
│  ├─ Claude API integration points
│  ├─ Session management (multi-session)
│  ├─ Event-driven system
│  └─ Electron UI layers
├─ Document architecture patterns
└─ Identify "Claude SDK usage" (all occurences)

Day 2 (8 jan):
├─ Map dependencies (3rd party)
├─ Identify tech debt areas
├─ Create API differences matrix
│  ├─ Request format (Claude → Gemini)
│  ├─ Response handling (streaming, errors)
│  ├─ Token counting
│  └─ Context window management
├─ Write migration strategy ADR
└─ Validate no blockers found
```

**Key Audit Findings Template**:

```
CODEBASE ANALYSIS REPORT

1. PROJECT STATISTICS
   ├─ Lines of Code: ~X,XXX
   ├─ Number of files: XXX
   ├─ Main languages: TypeScript/JavaScript
   ├─ Test coverage: X%
   └─ Documentation: X pages

2. ARCHITECTURE LAYERS
   ├─ UI Layer (Electron)
   │  └─ Findings: [Will be replaced with Oclif CLI]
   ├─ Agent Layer (core logic)
   │  └─ Findings: [Can mostly stay same]
   ├─ API Layer (Claude SDK)
   │  └─ Findings: [PRIMARY swap target]
   ├─ Storage Layer (SQLite)
   │  └─ Findings: [Already there or can add]
   └─ Integration Layer (Git, etc)
       └─ Findings: [Can stay mostly same]

3. CLAUDE SDK USAGE
   ├─ File: src/core/gemini-client.ts
   │  └─ Usage: API calls, message formatting, token counting
   ├─ File: src/agent/orchestrator.ts
   │  └─ Usage: Request/response handling
   ├─ File: src/storage/cache.ts
   │  └─ Usage: Response caching (format adaptation needed)
   └─ Total: ~15 files with Claude-specific code

4. COUPLED COMPONENTS
   ├─ Electron UI ↔ Agent logic (MODERATE coupling)
   │  └─ Solution: Abstraction layer (EventEmitter)
   ├─ Claude SDK ↔ Agent logic (HIGH coupling)
   │  └─ Solution: Wrapper pattern (GeminiClient)
   └─ Session Manager ↔ In-memory store (LOW coupling)
       └─ Solution: Keep pattern, swap storage

5. RISK ASSESSMENT
   ├─ HIGH: API response format differences
   │  └─ Mitigation: Comprehensive test suite Phase 1
   ├─ MEDIUM: Rate limiting (200K context → 1M token handling)
   │  └─ Mitigation: Queue system Phase 3
   ├─ MEDIUM: Streaming behavior differences
   │  └─ Mitigation: Adapter pattern
   └─ LOW: Event system (should work unchanged)
       └─ Validation: Tests will confirm

6. REUSABLE COMPONENTS (80% estimate)
   ├─ Session Manager (90% reusable)
   ├─ Event Bus (95% reusable)
   ├─ Git integration (85% reusable)
   ├─ Code execution (80% reusable)
   ├─ Test framework (90% reusable)
   └─ Utilities (95% reusable)

7. COMPONENTS TO REPLACE (20%)
   ├─ Electron UI → Oclif CLI (100% new)
   ├─ Claude client → Gemini client (100% new)
   ├─ Token counter → Gemini token counting (100% new)
   └─ Response parser → Gemini parser (100% new)
```

**Critères de Succès**:

- ✅ Full codebase understood
- ✅ API differences documented
- ✅ No blockers identified (or mitigation planned)
- ✅ Migration path clear
- ✅ Team confident in approach

---

### 0.5 Tâche: Gemini API Learning Sprint

**Responsable**: Lead Architect + Full-Stack Dev  
**Durée**: 2 jours (7-8 janvier, parallel to audit)

**Livrables**:

- [ ] Gemini API playground scripts (working examples)
- [ ] Token counting validation
- [ ] Streaming behavior tests
- [ ] Error handling patterns documented
- [ ] Rate limiting logic POC
- [ ] ADR #2: Gemini API Patterns

**Actions**:

```
Day 1 (7 jan):
├─ Read Gemini API documentation
├─ Setup API credentials (free tier account)
├─ Create playground project:
│  ├─ basic-prompt.ts (simple text generation)
│  ├─ streaming.ts (streaming responses)
│  ├─ token-counting.ts (validate countTokens API)
│  └─ error-handling.ts (test all error codes)
├─ Test rate limits (make 100 calls, observe behavior)
└─ Document findings

Day 2 (8 jan):
├─ Test context window (load full Auto-Claude codebase)
├─ Streaming large responses (test chunking)
├─ Compare response format vs Claude
├─ Test JSON mode (structured output)
├─ Benchmark latency vs Claude
└─ Create decision matrix (patterns to follow)
```

**Sample Playground Scripts**:

```typescript
// playground/01-basic-prompt.ts
import { GoogleGenerativeAI } from "@google/generative-ai";

async function basicPrompt() {
  const genAI = new GoogleGenerativeAI(process.env.GEMINI_API_KEY!);
  const model = genAI.getGenerativeModel({ model: "gemini-2.5-pro" });

  const prompt = "Hello, Gemini! What is 2+2?";
  const result = await model.generateContent(prompt);
  const text = result.response.text();

  console.log("Response:", text);
}

basicPrompt().catch(console.error);
```

```typescript
// playground/02-token-counter.ts
import { GoogleGenerativeAI } from "@google/generative-ai";

async function tokenCounter() {
  const genAI = new GoogleGenerativeAI(process.env.GEMINI_API_KEY!);
  const model = genAI.getGenerativeModel({ model: "gemini-2.5-pro" });

  const prompt = "This is a test prompt to count tokens.";
  const tokenCount = await model.countTokens(prompt);

  console.log("Tokens:", tokenCount);
  // Expected: { totalTokens: X }
}

tokenCounter().catch(console.error);
```

**Critères de Succès**:

- ✅ Basic prompt works
- ✅ Token counting understood
- ✅ Streaming tested
- ✅ Error patterns documented
- ✅ Team confident in API

---

### Phase 0 Summary Checklist

```
✓ Phase 0: PRÉ-LANCEMENT (by 8 Jan 2026)

Budget & Approvals:
- [ ] Finance approval ($13-34K)
- [ ] IT commitment
- [ ] Security review passed
- [ ] Executive signoff

Team & Skills:
- [ ] Lead Architect assigned (agbruneau confirmed)
- [ ] Full-Stack Dev assigned (hired or reassigned)
- [ ] RACI matrix created
- [ ] Skill gaps identified
- [ ] Training plan (if needed)

Infrastructure:
- [ ] GitHub fork created
- [ ] CI/CD pipeline functional
- [ ] Environment variables setup
- [ ] Team can develop locally
- [ ] Secrets management secure

Codebase Understanding:
- [ ] Auto-Claude audit complete
- [ ] API differences documented
- [ ] Migration strategy written
- [ ] Risks identified + mitigations planned
- [ ] 80/20 rule validated (80% reuse)

Gemini API Knowledge:
- [ ] API playground scripts working
- [ ] Token counting validated
- [ ] Streaming tested
- [ ] Error patterns documented
- [ ] Team confident

GO/NO-GO DECISION: Phase 0 Exit Gate
→ If ALL checkboxes ✓: PROCEED TO PHASE 1 (Week 1)
→ If ANY blocker: PIVOT or ESCALATE
```

---

## II. PHASE 1: API SWAP (Semaines 1-3 | 9-27 janvier 2026)

### Objectif

Remplacer Claude SDK par Gemini SDK. 80% du code devrait fonctionner avec minimal changes.

### Timeline

```
Week 1 (9-13 jan):  Fork analysis + initial SDK swap
Week 2 (16-20 jan): Fix compilation + basic functionality
Week 3 (23-27 jan): Testing + validation
```

---

### 1.1 Tâche: SDK Integration Layer Design

**Responsable**: Lead Architect  
**Durée**: 1 jour (9 janvier)

**Livrables**:

- [ ] `src/core/gemini-client.ts` (abstraction layer)
- [ ] `src/types/gemini.ts` (TypeScript interfaces)
- [ ] `src/adapters/response-mapper.ts` (format conversion)
- [ ] ADR #3: Gemini SDK Abstraction

**Détails**:

```typescript
// src/core/gemini-client.ts
import { GoogleGenerativeAI, GenerativeModel } from "@google/generative-ai";

export interface GeminiClientConfig {
  apiKey: string;
  modelName?: string; // default: "gemini-2.5-pro"
  temperature?: number; // default: 0.3
  maxOutputTokens?: number; // default: 4096
}

export interface GenerateContentRequest {
  prompt: string;
  systemInstruction?: string;
  history?: Array<{ role: "user" | "model"; content: string }>;
  streaming?: boolean;
}

export interface GenerateContentResponse {
  text: string;
  finishReason: string;
  tokenCount?: number;
  usageMetadata?: {
    promptTokenCount: number;
    candidatesTokenCount: number;
    totalTokenCount: number;
  };
}

export class GeminiClient {
  private client: GoogleGenerativeAI;
  private model: GenerativeModel;
  private config: GeminiClientConfig;

  constructor(config: GeminiClientConfig) {
    this.config = config;
    this.client = new GoogleGenerativeAI(config.apiKey);
    this.model = this.client.getGenerativeModel({
      model: config.modelName || "gemini-2.5-pro",
      systemInstruction: undefined, // Set per-request if needed
      generationConfig: {
        temperature: config.temperature ?? 0.3,
        maxOutputTokens: config.maxOutputTokens ?? 4096,
        topP: 1,
        topK: 1,
      },
    });
  }

  async generateContent(
    request: GenerateContentRequest
  ): Promise<GenerateContentResponse> {
    try {
      // Build contents array (handle history)
      const contents = [];

      if (request.history && request.history.length > 0) {
        for (const msg of request.history) {
          contents.push({
            role: msg.role,
            parts: [{ text: msg.content }],
          });
        }
      }

      contents.push({
        role: "user",
        parts: [{ text: request.prompt }],
      });

      const result = await this.model.generateContent({
        contents,
        systemInstruction: request.systemInstruction,
      });

      const response = result.response;
      return {
        text: response.text(),
        finishReason: response.candidates?.[0]?.finishReason || "STOP",
        usageMetadata: response.usageMetadata,
      };
    } catch (error) {
      throw new GeminiError(`Failed to generate content: ${error}`, error);
    }
  }

  async *generateContentStream(
    request: GenerateContentRequest
  ): AsyncGenerator<string> {
    try {
      const contents = [];

      if (request.history && request.history.length > 0) {
        for (const msg of request.history) {
          contents.push({
            role: msg.role,
            parts: [{ text: msg.content }],
          });
        }
      }

      contents.push({
        role: "user",
        parts: [{ text: request.prompt }],
      });

      const result = await this.model.generateContentStream({
        contents,
        systemInstruction: request.systemInstruction,
      });

      for await (const chunk of result.stream) {
        yield chunk.text();
      }
    } catch (error) {
      throw new GeminiError(`Streaming failed: ${error}`, error);
    }
  }

  async countTokens(prompt: string): Promise<number> {
    try {
      const result = await this.model.countTokens(prompt);
      return result.totalTokens;
    } catch (error) {
      throw new GeminiError(`Token count failed: ${error}`, error);
    }
  }

  async countTokensForRequest(
    request: GenerateContentRequest
  ): Promise<number> {
    try {
      const contents = [];

      if (request.history) {
        for (const msg of request.history) {
          contents.push({
            role: msg.role,
            parts: [{ text: msg.content }],
          });
        }
      }

      contents.push({
        role: "user",
        parts: [{ text: request.prompt }],
      });

      const result = await this.model.countTokens({
        contents,
      });

      return result.totalTokens;
    } catch (error) {
      throw new GeminiError(`Token count for request failed: ${error}`, error);
    }
  }

  isRateLimited(error: any): boolean {
    // Detect rate limit errors from Gemini API
    return (
      error?.status === 429 ||
      error?.code === "RATE_LIMIT_EXCEEDED" ||
      error?.message?.includes("too many requests")
    );
  }

  isContextWindowExceeded(error: any): boolean {
    // Detect context window exceeded
    return (
      error?.message?.includes("context length") ||
      error?.message?.includes("token limit")
    );
  }
}

export class GeminiError extends Error {
  constructor(message: string, public originalError: any) {
    super(message);
    this.name = "GeminiError";
  }
}
```

**Interface Types**:

```typescript
// src/types/gemini.ts
export interface Message {
  role: "user" | "model";
  content: string;
  timestamp: Date;
  tokenCount?: number;
}

export interface ChatHistory {
  messages: Message[];
  totalTokens: number;
  modelName: string;
}

export interface AgentPlan {
  goal: string;
  steps: PlanStep[];
  estimatedTokens: number;
  confidence: number;
}

export interface PlanStep {
  id: string;
  description: string;
  action: "code_generation" | "code_review" | "test_write" | "refactor";
  inputs?: Record<string, any>;
  expectedOutput?: string;
}

export interface ExecutionResult {
  success: boolean;
  output?: string;
  error?: string;
  executionTime: number;
  tokenUsed: number;
}
```

**Critères de Succès**:

- ✅ GeminiClient class compiles
- ✅ All methods have proper error handling
- ✅ TypeScript types complete
- ✅ No external code needed yet
- ✅ Design reviewed by team

---

### 1.2 Tâche: Replace Claude SDK → Gemini SDK in Agent

**Responsable**: Full-Stack Dev (with Lead review)  
**Durée**: 2 jours (10-11 janvier)

**Livrables**:

- [ ] Agent class updated (use GeminiClient)
- [ ] All Claude-specific code replaced
- [ ] Compilation passes
- [ ] Basic unit tests for Agent

**Actions**:

```
Day 1 (10 jan):
├─ Update Agent class to use GeminiClient
├─ Replace all claudeClient.xxx with geminiClient.xxx
├─ Update request formatting (Claude → Gemini)
├─ Update response parsing (Claude → Gemini)
├─ Test compilation
└─ Fix any TypeScript errors

Day 2 (11 jan):
├─ Update token estimation (use Gemini counter)
├─ Update error handling (Gemini error codes)
├─ Update rate limiting logic
├─ Update context window checks (200K → 1M)
├─ Write unit tests for Agent.generatePlan()
└─ Validate existing tests adapt
```

**Example: Before vs After**

```typescript
// BEFORE (Claude)
import Anthropic from "@anthropic-ai/sdk";

class Agent {
  private anthropic = new Anthropic({
    apiKey: process.env.ANTHROPIC_API_KEY,
  });

  async generatePlan(task: string): Promise<AgentPlan> {
    const response = await this.anthropic.messages.create({
      model: "claude-3-5-sonnet",
      max_tokens: 4096,
      messages: [{ role: "user", content: task }],
    });

    const plan = JSON.parse(response.content[0].text);
    return plan;
  }
}

// AFTER (Gemini)
import { GeminiClient } from "./core/gemini-client";

class Agent {
  private gemini: GeminiClient;

  constructor(apiKey: string) {
    this.gemini = new GeminiClient({ apiKey });
  }

  async generatePlan(task: string): Promise<AgentPlan> {
    const response = await this.gemini.generateContent({
      prompt: task,
      systemInstruction: "You are an AI agent planner...",
    });

    const plan = JSON.parse(response.text);
    return plan;
  }
}
```

**Critères de Succès**:

- ✅ All Claude SDK imports removed
- ✅ Code compiles without errors
- ✅ Agent.generatePlan() works
- ✅ Unit tests pass (80%+ of original)
- ✅ No external dependencies blocking

---

### 1.3 Tâche: Token Counter Migration

**Responsable**: Full-Stack Dev  
**Durée**: 1 jour (12 janvier)

**Livrables**:

- [ ] `src/utils/token-counter.ts` (Gemini-based)
- [ ] Token estimation tests
- [ ] Validation vs countTokens() API

**Détails**:

```typescript
// src/utils/token-counter.ts
import { GeminiClient } from "../core/gemini-client";

export class TokenCounter {
  constructor(private gemini: GeminiClient) {}

  /**
   * Estimate tokens for a single prompt
   */
  async estimate(prompt: string): Promise<number> {
    // Use Gemini's countTokens API for accuracy
    return await this.gemini.countTokens(prompt);
  }

  /**
   * Estimate tokens for a multi-turn conversation
   */
  async estimateConversation(messages: Message[]): Promise<number> {
    const prompt = messages.map((m) => `${m.role}: ${m.content}`).join("\n");
    return await this.gemini.countTokens(prompt);
  }

  /**
   * Check if adding more content would exceed limit
   * Gemini: 1M tokens (1,000,000)
   * Reserve 10% for safety (100K)
   */
  canFitMore(currentTokens: number, addedTokens: number): boolean {
    const maxTokens = 900_000; // Reserve 100K for safety
    return currentTokens + addedTokens <= maxTokens;
  }

  /**
   * Get percentage of context window used
   */
  getUsagePercentage(usedTokens: number): number {
    return (usedTokens / 1_000_000) * 100;
  }

  /**
   * Truncate content to fit within token limit
   * Returns truncated text and token count
   */
  async truncateToFit(
    text: string,
    maxTokens: number
  ): Promise<{ text: string; tokens: number }> {
    let currentText = text;

    while (true) {
      const tokens = await this.estimate(currentText);

      if (tokens <= maxTokens) {
        return { text: currentText, tokens };
      }

      // Remove last 10% of content and retry
      currentText = currentText.slice(0, Math.floor(currentText.length * 0.9));

      if (currentText.length === 0) break;
    }

    return { text: "", tokens: 0 };
  }
}
```

**Tests**:

```typescript
// tests/unit/token-counter.test.ts
import { describe, it, expect } from "vitest";
import { TokenCounter } from "../../src/utils/token-counter";
import { GeminiClient } from "../../src/core/gemini-client";

describe("TokenCounter", () => {
  let counter: TokenCounter;

  beforeEach(() => {
    const gemini = new GeminiClient({
      apiKey: process.env.GEMINI_API_KEY!,
    });
    counter = new TokenCounter(gemini);
  });

  it("should estimate tokens for a prompt", async () => {
    const prompt = "Hello, world!";
    const tokens = await counter.estimate(prompt);
    expect(tokens).toBeGreaterThan(0);
  });

  it("should check if content fits", () => {
    const result = counter.canFitMore(500_000, 400_000);
    expect(result).toBe(true);
  });

  it("should return false if exceeds limit", () => {
    const result = counter.canFitMore(900_000, 50_000);
    expect(result).toBe(false);
  });

  it("should calculate usage percentage", () => {
    const percentage = counter.getUsagePercentage(500_000);
    expect(percentage).toBe(50);
  });
});
```

**Critères de Succès**:

- ✅ Token counter uses Gemini API
- ✅ Tests validate accuracy
- ✅ Context window logic correct (1M)
- ✅ No external token counting lib needed

---

### 1.4 Tâche: Session Manager Persistence Setup

**Responsable**: Full-Stack Dev  
**Durée**: 1.5 jours (13-14 janvier)

**Livrables**:

- [ ] `src/storage/session-store.ts` (SQLite)
- [ ] Schema migrations
- [ ] Session CRUD operations
- [ ] Tests

**Détails**:

```typescript
// src/storage/session-store.ts
import Database from "better-sqlite3";
import { randomUUID } from "crypto";

export interface SessionRecord {
  id: string;
  taskDescription: string;
  status: "pending" | "running" | "completed" | "failed";
  context: {
    files: Record<string, string>;
    history: Message[];
    workdir: string;
  };
  plan?: AgentPlan;
  subTasks?: SubTask[];
  gitCommit?: string;
  createdAt: Date;
  updatedAt: Date;
  completedAt?: Date;
}

export class SessionStore {
  private db: Database.Database;

  constructor(dbPath: string = "./sessions.db") {
    this.db = new Database(dbPath);
    this.initializeSchema();
  }

  private initializeSchema() {
    this.db.exec(`
      CREATE TABLE IF NOT EXISTS sessions (
        id TEXT PRIMARY KEY,
        taskDescription TEXT NOT NULL,
        status TEXT NOT NULL DEFAULT 'pending',
        context TEXT NOT NULL,
        plan TEXT,
        subTasks TEXT,
        gitCommit TEXT,
        createdAt INTEGER NOT NULL,
        updatedAt INTEGER NOT NULL,
        completedAt INTEGER
      );

      CREATE TABLE IF NOT EXISTS messages (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        sessionId TEXT NOT NULL,
        role TEXT NOT NULL,
        content TEXT NOT NULL,
        timestamp INTEGER NOT NULL,
        tokenCount INTEGER,
        FOREIGN KEY (sessionId) REFERENCES sessions(id) ON DELETE CASCADE
      );

      CREATE INDEX IF NOT EXISTS idx_sessions_status ON sessions(status);
      CREATE INDEX IF NOT EXISTS idx_sessions_createdAt ON sessions(createdAt);
      CREATE INDEX IF NOT EXISTS idx_messages_sessionId ON messages(sessionId);
    `);
  }

  async createSession(
    taskDescription: string,
    context: any
  ): Promise<SessionRecord> {
    const id = randomUUID();
    const now = Date.now();

    const stmt = this.db.prepare(`
      INSERT INTO sessions (id, taskDescription, status, context, createdAt, updatedAt)
      VALUES (?, ?, ?, ?, ?, ?)
    `);

    stmt.run(id, taskDescription, "pending", JSON.stringify(context), now, now);

    return {
      id,
      taskDescription,
      status: "pending",
      context,
      createdAt: new Date(now),
      updatedAt: new Date(now),
    };
  }

  async getSession(id: string): Promise<SessionRecord | null> {
    const stmt = this.db.prepare(`
      SELECT * FROM sessions WHERE id = ?
    `);

    const row = stmt.get(id) as any;
    if (!row) return null;

    return {
      ...row,
      context: JSON.parse(row.context),
      plan: row.plan ? JSON.parse(row.plan) : undefined,
      subTasks: row.subTasks ? JSON.parse(row.subTasks) : undefined,
      createdAt: new Date(row.createdAt),
      updatedAt: new Date(row.updatedAt),
      completedAt: row.completedAt ? new Date(row.completedAt) : undefined,
    };
  }

  async updateSession(
    id: string,
    updates: Partial<SessionRecord>
  ): Promise<void> {
    const now = Date.now();
    const fields: string[] = [];
    const values: any[] = [];

    if (updates.status) {
      fields.push("status = ?");
      values.push(updates.status);
    }
    if (updates.context) {
      fields.push("context = ?");
      values.push(JSON.stringify(updates.context));
    }
    if (updates.plan) {
      fields.push("plan = ?");
      values.push(JSON.stringify(updates.plan));
    }
    if (updates.gitCommit) {
      fields.push("gitCommit = ?");
      values.push(updates.gitCommit);
    }

    fields.push("updatedAt = ?");
    values.push(now);
    values.push(id);

    const stmt = this.db.prepare(`
      UPDATE sessions SET ${fields.join(", ")} WHERE id = ?
    `);

    stmt.run(...values);

    if (updates.status === "completed" || updates.status === "failed") {
      const stmt2 = this.db.prepare(`
        UPDATE sessions SET completedAt = ? WHERE id = ?
      `);
      stmt2.run(now, id);
    }
  }

  async listSessions(
    status?: string,
    limit: number = 50
  ): Promise<SessionRecord[]> {
    let query = `SELECT * FROM sessions`;
    const values: any[] = [];

    if (status) {
      query += ` WHERE status = ?`;
      values.push(status);
    }

    query += ` ORDER BY createdAt DESC LIMIT ?`;
    values.push(limit);

    const stmt = this.db.prepare(query);
    const rows = stmt.all(...values) as any[];

    return rows.map((row) => ({
      ...row,
      context: JSON.parse(row.context),
      plan: row.plan ? JSON.parse(row.plan) : undefined,
      subTasks: row.subTasks ? JSON.parse(row.subTasks) : undefined,
      createdAt: new Date(row.createdAt),
      updatedAt: new Date(row.updatedAt),
      completedAt: row.completedAt ? new Date(row.completedAt) : undefined,
    }));
  }

  async deleteSession(id: string): Promise<void> {
    const stmt = this.db.prepare(`DELETE FROM sessions WHERE id = ?`);
    stmt.run(id);
  }

  close(): void {
    this.db.close();
  }
}
```

**Critères de Succès**:

- ✅ SQLite schema created
- ✅ CRUD operations work
- ✅ Transactions safe
- ✅ Tests pass
- ✅ Sessions persist across restarts

---

### 1.5 Tâche: Basic CLI Commands

**Responsable**: Full-Stack Dev  
**Durée**: 2 jours (16-17 janvier)

**Livrables**:

- [ ] `src/commands/task.ts` (main command)
- [ ] `src/commands/status.ts` (check progress)
- [ ] `src/commands/history.ts` (view sessions)
- [ ] Working CLI binary

**Détails**:

```typescript
// src/commands/task.ts
import { Command, Flags } from "oclif";
import { Agent } from "../core/agent";
import { GeminiClient } from "../core/gemini-client";
import { SessionStore } from "../storage/session-store";

export default class TaskCommand extends Command {
  static description = "Create and execute an autonomous task";

  static flags = {
    help: Flags.help({ char: "h" }),
    project: Flags.string({
      char: "p",
      description: "Project directory",
      default: process.cwd(),
    }),
    watch: Flags.boolean({
      char: "w",
      description: "Watch mode (re-run on changes)",
      default: false,
    }),
  };

  static args = [
    {
      name: "task",
      required: true,
      description: "Task description or goal",
    },
  ];

  async run(): Promise<void> {
    const { args, flags } = await this.parse(TaskCommand);
    const taskDescription = args.task;

    this.log(`🚀 Starting task: ${taskDescription}`);
    this.log(`📁 Project: ${flags.project}`);

    const gemini = new GeminiClient({
      apiKey: process.env.GEMINI_API_KEY!,
    });

    const store = new SessionStore();
    const agent = new Agent(gemini, store);

    try {
      const session = await agent.executeTask(taskDescription, flags.project);

      this.log(`✅ Task completed: ${session.id}`);
      this.log(`📝 Commit: ${session.gitCommit}`);
    } catch (error) {
      this.error(`❌ Task failed: ${error}`);
    }

    store.close();
  }
}
```

**Critères de Succès**:

- ✅ CLI runs without errors
- ✅ At least 3 commands work
- ✅ Help text displays
- ✅ User can execute first task

---

### 1.6 Tâche: Compilation & Test Suite Baseline

**Responsable**: Full-Stack Dev  
**Durée**: 1 jour (18 janvier)

**Livrables**:

- [ ] Full compilation succeeds
- [ ] 80%+ of original tests pass
- [ ] No blocking errors
- [ ] GitHub Actions workflow green

**Actions**:

```
├─ Run: npm run build
├─ Fix all TypeScript errors
├─ Run: npm test
├─ Verify test coverage (target: 70%+)
├─ Run: npm run lint
├─ Fix linting issues
└─ Push to GitHub + verify CI/CD green
```

**Critères de Succès**:

- ✅ `npm run build` succeeds
- ✅ `npm test` passes (80%+ tests)
- ✅ `npm run lint` succeeds
- ✅ GitHub Actions green
- ✅ No unresolved dependencies

---

### 1.7 Tâche: Validate Basic Functionality

**Responsable**: Lead Architect + Full-Stack Dev  
**Durée**: 2 jours (19-20 janvier)

**Livrables**:

- [ ] E2E test: Single task execution
- [ ] Validation: Session persists
- [ ] Validation: Output matches expected
- [ ] ADR #4: API Swap Validation Results

**Actions**:

```
Day 1 (19 jan):
├─ Create simple test task
│  └─ "Write a Hello World in Python"
├─ Execute via CLI
├─ Verify output generated
├─ Verify session saved to DB
└─ Check token counting accurate

Day 2 (20 jan):
├─ Test multi-file generation
├─ Test error handling
├─ Test rate limit graceful degradation
├─ Verify streaming works
└─ Document any issues found
```

**Expected Results**:

```
Input:  "Write a Hello World function in Python"
Output:
  ├─ Generated Python code (hello_world.py)
  ├─ Test code (test_hello.py)
  ├─ Session record in DB
  ├─ Git commit created
  └─ Total tokens: ~1,200

Validate:
- ✅ Code is valid Python
- ✅ Can be executed
- ✅ Session recoverable
- ✅ No token miscounting
```

**Critères de Succès**:

- ✅ Simple task executes start-to-finish
- ✅ Output correct
- ✅ Session saved
- ✅ No data loss on restart
- ✅ Team confident in foundation

---

### 1.8 Tâche: Code Review & Milestone Gate

**Responsable**: Lead Architect  
**Durée**: 1 jour (21 janvier)

**Livrables**:

- [ ] Pull request review (Phase 1 code)
- [ ] Architecture sign-off
- [ ] ADR #5: Phase 1 Learnings
- [ ] Phase 1 Milestone: COMPLETE

**Actions**:

```
├─ Review all Phase 1 PRs
├─ Check:
│  ├─ Code quality (style, patterns)
│  ├─ Test coverage (target: 70%+)
│  ├─ Documentation (comments, docstrings)
│  ├─ No regressions vs original
│  └─ Architecture decisions sound
├─ Merge to main
└─ Tag: v1.0.0-phase1
```

**Phase 1 Exit Criteria**:

```
✅ API Swap Complete
├─ Claude SDK fully replaced with Gemini
├─ 80% of original tests passing
├─ Basic functionality working
├─ Single task execution works
├─ Session persistence functional
└─ Team confident in foundation

✅ No Blockers Identified
├─ Gemini API behaves as expected
├─ Token counting accurate
├─ Streaming works
├─ Error handling acceptable
└─ Rate limiting understood

✅ Ready for Phase 2
└─ UI migration can proceed
```

**GO/NO-GO for Phase 2**:

- ✅ If all criteria met → PROCEED to Phase 2 (Week 4)
- ❌ If any blocker → ESCALATE before continuing

---

### Phase 1 Summary Checklist

```
✓ PHASE 1: API SWAP (Weeks 1-3, 9-27 Jan 2026)

Week 1 (9-13 Jan):
- [ ] SDK Integration layer designed
- [ ] Agent updated to use GeminiClient
- [ ] Compilation passes

Week 2 (16-20 Jan):
- [ ] Token counter implemented (Gemini)
- [ ] Session manager (SQLite) working
- [ ] Basic CLI commands functional
- [ ] 80%+ tests passing

Week 3 (23-27 Jan):
- [ ] Single task executes end-to-end
- [ ] Session persists + recovers
- [ ] GitHub Actions green
- [ ] Code reviewed + merged

EXIT GATE (28 Jan):
- [ ] All original tests adapted (80%+)
- [ ] Basic functionality validated
- [ ] No blockers for UI migration
- [ ] Team confident

→ MILESTONE REACHED: "API Swap Complete"
→ DECISION: Proceed to Phase 2 ✅
```

---

## III. PHASE 2: UI MIGRATION (Semaines 4-6 | 30 jan - 10 février 2026)

### Objectif

Migrer de Electron (Desktop) vers Oclif (CLI). Adapter Terminal UI avec Ink.

### Timeline

```
Week 4 (30 jan - 3 feb):  Remove Electron, setup Oclif
Week 5 (6-10 feb):        Build Terminal UI components
Week 6 (13-17 feb):       Polish + validation
```

_(Phase 2 details would continue with similar exhaustive breakdown)_

---

## IV. PHASE 3: OPTIMIZATION (Semaines 7-10 | 20 feb - 16 mars 2026)

### Objectif

Rate limiting, context window optimization, performance tuning.

_(Phase 3 details would continue)_

---

## V. PHASE 4: POLISH & RELEASE (Semaines 11-14 | 19 mars - 28 avril 2026)

### Objectif

Testing, documentation, v1.0 release.

_(Phase 4 details would continue)_

---

## PROJECT MANAGEMENT ARTIFACTS

### A. Weekly Stand-up Template

```markdown
## Weekly Stand-up Report

**Week**: X (Date Range)
**Attendees**: [names]
**Status**: On Track / At Risk / Off Track

### Completed This Week

- [ ] Item 1
- [ ] Item 2
- [ ] Item 3

### In Progress

- [ ] Item (% complete)
- [ ] Item (% complete)

### Blockers

- [ ] Blocker 1 (Impact, Mitigation)
- [ ] Blocker 2 (Impact, Mitigation)

### Next Week Goals

- [ ] Goal 1
- [ ] Goal 2

### Metrics

- Lines of code added: XXX
- Test coverage: XX%
- Performance (latency): Xms
- Issues closed: X
- PRs merged: X
```

### B. Milestones & Gate Criteria

| Milestone            | Date   | Criteria                                 | Owner      |
| -------------------- | ------ | ---------------------------------------- | ---------- |
| **Phase 0 Complete** | 8 Jan  | All pre-launch checklist done            | Lead       |
| **Phase 1 Complete** | 28 Jan | API swap validated, basic task works     | Lead + Dev |
| **Phase 2 Complete** | 17 Feb | CLI UI working, MVP ready                | Dev        |
| **Phase 3 Complete** | 16 Mar | Optimized, rate limiting works           | Dev + Lead |
| **v1.0 Released**    | 28 Apr | Tests pass, docs complete, npm published | Lead       |

### C. Risk Register & Escalation

```
RISK REGISTER

ID  | Risk | Impact | Prob | Mitigation | Status
----|------|--------|------|-----------|--------
R1  | Gemini SDK breaking change | HIGH | MED | Pin version, monitor releases | Active
R2  | Rate limit hitting prod | MED | MED | Queue system, early tests | Active
R3  | Team unavailable | HIGH | LOW | Cross-training, documentation | Watch
R4  | Performance regression | MED | MED | Profiling, benchmarks | Active
R5  | Security vulnerability | HIGH | LOW | Code review, dependencies scan | Active

ESCALATION PATH:
├─ Dev → Lead Architect (daily blockers)
├─ Lead Architect → Product Manager (timeline/scope)
├─ Product Manager → Exec Sponsor (budget/resources)
└─ All → ESCALATION if risk materializes
```

### D. Definition of Done (DoD)

```
FOR EACH TASK:
- [ ] Code written (typed, formatted)
- [ ] Tests written (unit + integration)
- [ ] Code reviewed (1 approval minimum)
- [ ] Tests passing (CI/CD green)
- [ ] Documentation updated (README, ADRs, inline)
- [ ] No new technical debt introduced
- [ ] Performance validated (no regression)
- [ ] Security review passed (if relevant)
- [ ] Merged to main branch
- [ ] Closed in issue tracker
```

---

## APPENDIX: USEFUL COMMANDS

```bash
# Development
npm install
npm run build
npm run dev
npm test
npm run lint
npm run type-check

# Testing
npm run test:unit
npm run test:integration
npm run test:e2e
npm run test:coverage

# Build & Release
npm run build:release
npm run package    # Create executables
npm run publish    # Publish to npm

# Gemini API
export GEMINI_API_KEY="your_key_here"

# Database
sqlite3 sessions.db ".schema"
sqlite3 sessions.db "SELECT * FROM sessions;"
```

---

**Document**: Planification d'Exécution Exhaustive  
**Project**: Auto Gemini CLI (Scénario 2)  
**Status**: Ready for Implementation  
**Next Step**: Execute Phase 0 starting 2 Jan 2026
