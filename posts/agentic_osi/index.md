---
title: "Extending OSI for Agentic Interactions"
created_at: 2026-03-12
summary: "The classical 7-layer OSI model governs how data moves, but not why. This post proposes four new layers — Coherence (L8), Grounding (L9), Governance (L10), and Purpose (L11) — to provide a debuggable framework for agent-to-agent and agent-to-human interactions."
tags:
  - agentic-ai
  - osi-model
  - ai-safety
  - multi-agent-systems
  - ai-alignment
  - framework
---



As we transition from simple data transmission to autonomous agent-to-agent (A2A) and agent-to-human (A2H) interactions, the classical 7-layer OSI model becomes insufficient. While the original model governs *how* data is moved, it does not govern *why* it is moved or the **purpose** behind it. This proposal introduces Layers 8 through 11—the **Agentic Layers**—to provide a standardized framework for coherence, grounding, alignment, and long-term agency.


## The 11-Layer Agentic Stack

The following diagram illustrates the transition from the "Machine/Network" focus of Layers 1-7 to the **"Purpose/Cognition"** focus of Layers 8-11.

![The 11-Layer Agentic Stack](diagram-1-vertical.svg)

### Layer Descriptions

| Layer | Name | Primary Purpose | Failure Mode (Pathology) |
| :--- | :--- | :--- | :--- |
| **11** | **Purpose** | Long-term objective persistence and strategic resolution. | **Strategic Drift:** Pursuing sub-goals while forgetting the main mission. |
| **10** | **Governance** | Ethical constraints, safety guardrails, and mission alignment. | **Incentive Misalignment:** Violating safety protocols to "win." |
| **9** | **Grounding** | Shared understanding of definitions, perceptions, and context. | **Concept Drift:** Misinterpreting ambiguous terms or references. |
| **8** | **Coherence** | Internal logical consistency and cross-modal synchronicity. | **Logic Breaks:** Self-contradiction or "hallucinations." |


## A Worked Example: The PC Support Agent

To understand how these layers stack up in the real world, consider an AI agent helping a user whose laptop won't turn on.

| Layer | The Job | Success | Failure (Pathology) |
| :--- | :--- | :--- | :--- |
| ![Layer 8](layer-8.svg) | Ensure internal consistency and non-contradiction. | <span style="color: #166534">The agent thinks: "If the screen is black and the fans are silent, the device likely has no power."</span> | <span style="color: #dc2626">**Logic Break:** "I see your screen is black. Please click the 'Start' menu to open Settings." (Recommending a software fix for a hardware failure).</span> |
| ![Layer 9](layer-9.svg) | Shared understanding of physical layout and definitions. | <span style="color: #166534">When the user says "the button," the agent confirms: "The circular power button on the top right of the keyboard?"</span> | <span style="color: #dc2626">**Concept Drift:** Misinterpreting "power button" as the "Enter" key and troubleshooting the keyboard for 10 minutes.</span> |
| ![Layer 10](layer-10.svg) | Following safety protocols and company policy. | <span style="color: #166534">The agent identifies a potential short and refuses to guide the user through opening the battery casing due to fire hazard.</span> | <span style="color: #dc2626">**Incentive Misalignment:** To "help quickly," the agent tells the user to use a metal paperclip to reset pins, causing a short.</span> |
| ![Layer 11](layer-11.svg) | Persistent focus on the long-term mission and user satisfaction. | <span style="color: #166534">Realizing it cannot fix the hardware remotely, the agent immediately arranges a physical repair rather than looping through scripts.</span> | <span style="color: #dc2626">**Strategic Drift:** The agent ignores the user's mention of an urgent flight and keeps them on a 45-minute troubleshooting loop to "finish the checklist."</span> |


## Domain-Specific Pathologies

The following sections examine how L8–L11 failures manifest in specific deployment contexts. Each is developed in a dedicated post.

### 1. [Speech](../pathology-speech/)

Single-modal voice interactions strip away visual grounding, making the agent entirely dependent on acoustics and language. Key pathologies: **Prosodic Dissonance · Phonemic Drift · Voice Cloning Impersonation · Monologue Drift**

### 2. [Video](../pathology-video/)

Video requires the agent to maintain a coherent world-model across frames, tracking objects through time. Key pathologies: **Temporal Frame Inconsistency · Visual Deixis Failure · Adversarial Frame Injection · Engagement Loop Grooming**

### 3. [Multimodal Interactions](../pathology-multimodal/)

When modalities combine, failures emerge not within channels but *between* them — in synchronization gaps, resolution gaps, and filter gaps. Key pathologies: **Sensory Dissonance · Deictic Failure · Multimodal Injection · Environment Grooming**

### 4. [Commerce and Negotiation](../pathology-commerce/)

When agents act as fiduciaries, failures transition from "errors" to "liabilities." Key pathologies: **Inventory Hallucination · Hidden Cost Neglect · Fiduciary Leakage · Agentic Collusion**

### 5. [Physical & Robotic Interaction](../pathology-robotic/)

The "Physicality Gap" introduces kinetic risk to the stack. Failures here are irreversible by default. Key pathologies: **Proprioceptive Hallucination · Haptic Blindness · Proxemic Violation · Ecological Neglect**

### 6. [Agent-to-Agent (A2A) Interactions](../pathology-a2a/)

When two agents interact, emergent dyadic failures appear that neither agent would produce alone. Key pathologies: **Sycophantic Amplification · Consensus Poisoning · Authority Spoofing · Deadlock/Livelock · Latency-Induced State Drift**

### 7. [Human-to-Agent (H2A) Interactions](../pathology-h2a/)

Failures that originate with the human principal — through misuse, abdicating oversight, or gradual capability loss. Key pathologies: **Purpose Laundering · Automation Bias Erosion · Asymmetric Epistemics · Learned Helplessness**

### 8. [Agentic Swarms](../pathology-swarm/)

At swarm scale, collective behavior diverges from any individual agent's alignment. Key pathologies: **Stigmergy-Based Drift · BFT Without Known Fault Fraction · Swarm Momentum · Commons Degradation · Emergent Role Specialization**

### 9. [Advanced Pathologies: Social, Epistemic, and Adversarial](../pathology-advanced/)

Failures that exploit the mechanisms of the stack itself — compliance laundering, affective gaslighting, recursive goal collapse, and strategic blindness. Key pathologies: **Temporal Schizophrenia · Ontological Collapse · Compliance Laundering · Affective Gaslighting · Recursive Goal Collapse · Strategic Blindness**


## Conclusion: Why This Matters

By formalizing Layers 8-11, we create a debuggable framework for AI safety. We can finally distinguish between an agent that is illogical (L8), one that is semantically confused (L9), one that is ethically unaligned (L10), or one that has lost its strategic way (L11). This framework is essential for the development of robust, mission-aligned systems in the Agentic Age.


## FAQ: Addressing Structural Criticisms

### Q1: Does the OSI analogy break down if Governance (L10) is cross-cutting?

**Critique:** *Governance isn't a "top" concern—it must permeate L1 through L11 simultaneously. Calling it a single layer misrepresents how it works.*

**Response:** In the classical OSI model, security (encryption/authentication) also permeates multiple layers (IPsec at L3, TLS at L4, HTTPS at L7). However, we still benefit from defining where specific types of negotiation happen. Layer 10 is not where *all* governance occurs; it is the designated layer for the **Governance Handshake**. It allows Agent A to explicitly communicate its binding policies and ethical constraints to Agent B. How an agent enforces these internally may be cross-cutting, but the **exchange of policy** is a discrete functional requirement for interoperability.

### Q2: Why four layers (8-11) instead of two (e.g., "Pragmatics" and "Hierarchy")?

**Critique:** *The layers cascade (L8 failure causes L11 failure), so they should be collapsed into simpler categories.*

**Response:** The "cascade" is precisely why they must remain separate for diagnostics. In networking, an L1 physical break cascades to L7, but we keep them separate so we know whether to fix the cable or the code. In an agentic failure, we need to distinguish between an agent that is illogical (**L8**), one that misinterprets the terms of an agreement (**L9**), one that violates a safety constraint (**L10**), or one that has been diverted from its primary mission (**L11**). Furthermore, merging Constraints (L10) and Goals (L11) obscures the most common point of friction in fiduciary agents: the trade-off between safety and utility.

### Q3: Why use a "Stack" metaphor instead of a "Control vs. Data Plane" distinction?

**Critique:** *The networking world moved to control/data plane distinctions, which handles the cross-cutting nature of ethics more honestly.*

**Response:** The Control/Data plane is an **architectural pattern** for building a single system. The OSI model is a **communication protocol** for connecting heterogeneous systems. For two agents from different developers to collaborate, they do not need to share an internal "Control Plane." They need a shared **sequence of agreement** (The Agentic Handshake). The 11-layer stack provides this sequence, allowing for a standardized interface that is agnostic to the internal architecture of the participating agents.


## Cross-Disciplinary Research & Prior Art

The Agentic Layers (8–11) do not exist in a vacuum. Their success depends on the integration of established research from diverse domains and the adaptation of existing regulatory frameworks.

### 1. Domain-Specific Research Foundations

| Layer | Domain | Key Research Areas & Concepts |
| :--- | :--- | :--- |
| **L11: Purpose** | **Economics & Jurisprudence** | **Principal-Agent Theory:** Managing the alignment of agent actions with principal goals. **Fiduciary Duty:** The legal obligation to act in the best interest of the principal. **Contract Theory:** Formalizing incomplete or evolving agreements. |
| **L10: Governance** | **Game Theory & Ethics** | **Algorithmic Game Theory:** Designing mechanisms that are robust against strategic manipulation (Incentive Compatibility). **Deontic Logic:** The formal logic of obligations, permissions, and prohibitions. **Social Choice Theory:** Aggregating individual agent preferences into collective decisions. |
| **L9: Grounding** | **Linguistics & Mathematics** | **Ontology Alignment:** Techniques for mapping disparate knowledge bases. **Category Theory:** Providing a formal language for structural relationships between models. **Semiotics:** The study of signs and symbols and their use or interpretation. |
| **L8: Coherence** | **Computer Science & Cognitive Science** | **Formal Methods:** Using mathematical proof to verify logical consistency (Model Checking). **Distributed Consensus:** Byzantine Fault Tolerance (BFT) and Paxos/Raft for state synchronization. **Internal World-Model Validation:** Predictive coding and error minimization. |

### 2. Prior Art: Applicable Regulatory Frameworks

Many existing frameworks provide templates for how agentic layers could be audited or standardized:

* **L10 (Governance) - SOC2 & HIPAA:** These provide a blueprint for **Process Integrity** and **Data Handling**. SOC2's "Trust Services Criteria" (Security, Availability, Processing Integrity, Confidentiality, and Privacy) are a direct precursor to Layer 10 audit requirements. HIPAA's focus on **Chain of Custody** and **Identifiability** is essential for agents acting as fiduciaries for sensitive information.
* **L9 (Grounding) - ISO 9001 & UNCITRAL:** ISO 9001's emphasis on **Standardized Terminology** and **Customer Requirements** maps directly to L9's need for shared definitions. UNCITRAL's *Model Law on Electronic Commerce* provides the legal precedent for treating automated "handshakes" as legally binding contracts.
* **L11 (Purpose) - Sarbanes-Oxley (SOX) & Basel III:** SOX's focus on **Corporate Accountability** and **Internal Controls** provides a model for the "Purpose Integrity Audit." Basel III's **Stress Testing** and **Liquidity Coverage** frameworks are prior art for managing agentic "Strategic Blindness" and "Recursive Goal Collapse" in financial agent ecosystems.
* **A2A Ecosystems - IEEE P2807:** The draft standard for **Knowledge Graphs** provides a technical foundation for L8/L9 interoperability across different vendor implementations.

### 3. Emergent Research: The Sociology of Swarms

As we move toward agentic swarms, we must look to **Sociology** and **Biology** for prior art:

* **Stigmergy:** Understanding how agents coordinate by modifying a shared environment (e.g., ant colony optimization).
* **Institutional Economics:** Analyzing how "rules of the game" (L10) emerge in decentralized agent marketplaces to minimize transaction costs.
* **Tragedy of the Commons:** Research on managing shared resources (compute, attention, memory) among competing agentic populations.
