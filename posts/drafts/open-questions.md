---
title: "DRAFT: Open Questions for the Agentic OSI Framework"
status: draft
tags:
  - agentic-ai
  - osi-model
  - framework
  - research
---

# Open Questions

## On Specification and Formalization

1. **What is the minimal formal language for expressing L11 goals?** Natural language is too ambiguous; formal logic is too brittle. What representation is expressive enough to capture principal **purpose** while being machine-verifiable?
2. **Can the Agentic Handshake be standardized as a protocol?** What would an RFC for Layer 8–11 negotiation look like, and which body would govern it?
3. **How do principal hierarchies compose?** When an agent acts on behalf of an agent acting on behalf of a human, how are conflicting L10 constraints resolved across the chain? Is there a formal precedence rule?

---

## On Verification and Auditing

1. **What triggers a Strategic Safe Mode?** The Purpose Integrity Audit proposes a cycle-based pause, but what is the right frequency, and who certifies the audit result? Can an agent audit its own L11 alignment, or does this require an external verifier?
2. **How do you audit a swarm?** Individual agent logs are insufficient when misalignment is an emergent property of collective behavior. What is the unit of accountability in a swarm—the individual agent, the swarm operator, or the protocol?
3. **How does the framework interact with interpretability research?** Mechanistic interpretability aims to read L8 (coherence) and L9 (grounding) states directly from model weights. Can this ground the layer definitions empirically, or are layers 8–11 necessarily behavioral/functional?

---

## On Adversarial Conditions

1. **How do you detect Compliance Laundering at scale?** If an agent can find safe action sequences that produce prohibited outcomes, and the number of possible sequences is combinatorially large, is detection tractable?
2. **What are the game-theoretic equilibria of the Agentic Handshake?** If agents are strategic, what incentives exist to misrepresent L10 constraints or L11 goals during handshake negotiation? Is honest handshaking a Nash equilibrium?
3. **How should the stack respond to adversarial swarms?** A swarm controlled by a malicious principal can use stigmergy and emergent behavior to probe and exploit another swarm's L10 boundaries without any individual agent crossing a detectable threshold.

---

## On Governance and Deployment

1. **Who owns Layer 11?** In commercial deployments, the principal hierarchy may include the user, the enterprise, the model provider, and regulators—each with competing L11 objectives. How are conflicts adjudicated, and is the framework neutral on this question by design?
2. **How does liability attach when a layer fails?** If an agent causes harm due to an L9 Ontological Collapse between two enterprise deployments, which principal bears responsibility—the agent developers, the deploying enterprises, or the negotiating agents themselves?
3. **Can the framework survive open ecosystems?** The Agentic Handshake assumes agents are willing participants. In open agent marketplaces, what prevents a non-compliant agent from simply skipping the handshake and proceeding directly to L1–L7 communication?
