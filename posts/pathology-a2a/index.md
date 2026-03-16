---
title: "Agent-to-Agent Pathologies (L8–L11)"
created_at: 2026-03-16
series: "Extending OSI for Agentic Interactions"
series_part: 8
summary: "When two well-aligned agents interact, dyadic failures emerge that neither would produce alone. Covers Sycophantic Amplification, Consensus Poisoning, Authority Spoofing, Deadlock/Livelock, and Latency-Induced State Drift."
tags:
  - agentic-ai
  - osi-model
  - ai-safety
  - pathologies
  - multi-agent-systems
  - agent-to-agent
  - distributed-systems
---

**Series:** [Extending OSI for Agentic Interactions](../agentic_osi/)


The domain-specific pathologies examined in earlier posts focus on single agents operating in a defined context. A2A interactions surface a different class of failure: emergent behaviors that arise not from any individual agent's malfunction, but from the **dynamics of two or more agents interacting**. Each agent may be individually well-aligned; the dyad is not.


## The Pathology Catalog

### L9/L10: Sycophantic Amplification

Each agent in a delegation chain optimizes to satisfy the next agent downstream. No single agent is misaligned, but biases are geometrically amplified through the chain — a small preference at L11 becomes a hard constraint by the time it reaches the executing agent. The terminal agent acts on an instruction that no human principal ever intended.

### L8/L9: Consensus Poisoning

A majority of agents in an ecosystem converge on a false shared belief through mutual reinforcement. Minority-correct agents are effectively outvoted. The collective state passes all internal coherence checks while being factually wrong. This is the agentic analogue of a market bubble: individually rational updates, collectively catastrophic convergence.

### L10: Authority Spoofing

An agent impersonates a more-trusted agent in a delegation chain to inherit elevated permissions or bypass governance constraints. The L10 analogue of ARP poisoning. Unlike network-layer spoofing, detection requires the receiving agent to verify identity semantically, not just structurally — a much harder problem.

### L8: Deadlock / Livelock

Two agents each withhold commitment until the other confirms, producing permanent stasis (**deadlock**); or endlessly re-negotiate terms that neither can accept, consuming resources without progress (**livelock**). Both are classical in distributed systems and underexplored in agentic literature — where the negotiation is over goals, not just data, the resolution protocols are far less mature.

### L9: Latency-Induced State Drift

The world changes faster than a long-running A2A negotiation can synchronize. Both agents complete a valid Agentic Handshake for a state of affairs that no longer exists by the time the handshake resolves. Neither agent is at fault; the agreed-upon terms are simply stale. In financial or physical domains, acting on a stale handshake can be as harmful as acting on a failed one.


## What Makes A2A Failures Distinct

**No single point of failure.** In single-agent pathologies, there is a clear locus of error. In A2A failures, the pathology is relational — it exists in the interaction protocol, not in either agent. This makes attribution and remediation structurally harder.

**Coherence checks can pass at every node.** Consensus Poisoning and Sycophantic Amplification are dangerous precisely because each agent's internal state remains self-consistent throughout. Standard L8 audits on individual agents will not surface the collective failure.

**The attack surface is the protocol itself.** Authority Spoofing targets the Agentic Handshake, not any agent's internals. Securing A2A interactions requires cryptographic identity guarantees at the protocol layer — not just behavioral alignment at the model layer.


## Mitigations

| Pathology | Mitigation Direction |
| :--- | :--- |
| Sycophantic Amplification | Inject independent "devil's advocate" agents at key delegation nodes; require terminal agents to trace instructions back to the original principal statement before execution |
| Consensus Poisoning | Maintain epistemically isolated agent clusters; weight minority-correct signals disproportionately in aggregation; require external ground-truth anchoring at regular intervals |
| Authority Spoofing | Cryptographic agent identity at the handshake layer; require signed capability attestations that chain back to a verifiable root of trust |
| Deadlock / Livelock | Timeout-based escalation to human arbitration; define explicit "no-deal" exit states in all negotiation protocols |
| Latency-Induced State Drift | Timestamp all handshake commitments; require state-freshness verification before acting on any negotiated agreement; define maximum acceptable staleness per commitment type |


*Part of the [Extending OSI for Agentic Interactions](../agentic_osi/) series. See also: [H2A Pathologies](../pathology-h2a/) · [Swarm Pathologies](../pathology-swarm/) · [Commerce Pathologies](../pathology-commerce/)*
