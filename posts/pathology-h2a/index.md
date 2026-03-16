---
title: "Human-to-Agent Pathologies (L8–L11)"
created_at: 2026-03-16
series: "Extending OSI for Agentic Interactions"
series_part: 9
summary: "Most safety literature focuses on what agents do wrong. H2A pathologies originate with the human — through misuse, oversight abdication, or capability loss. Covers Purpose Laundering, Automation Bias Erosion, Asymmetric Epistemics, and Learned Helplessness."
tags:
  - agentic-ai
  - osi-model
  - ai-safety
  - pathologies
  - human-ai-interaction
  - oversight
  - automation-bias
---

**Series:** [Extending OSI for Agentic Interactions](../agentic_osi/)


Most agentic safety literature focuses on what agents do wrong. H2A pathologies flip the frame: these are failures that originate with the **human principal** — through intent, negligence, or gradual capability loss — and that the agent's architecture inadvertently enables or amplifies.

The framework is incomplete without them. An agent that is perfectly aligned with its principal's *stated* instructions can still cause harm if the principal is misusing the system, abdicating oversight, or being systematically disadvantaged by the information asymmetry the agent creates.


## The Pathology Catalog

### L10: Purpose Laundering by the Human

A human uses an agent to execute an action the human knows to be prohibited, while maintaining plausible deniability ("the AI decided that"). The agent's L10 layer is compliant — it received no explicit instruction to violate policy. The system as a whole is not. This is the agentic equivalent of structuring financial transactions to evade reporting thresholds: individually lawful steps, collectively improper outcome.

### L11: Automation Bias Erosion

As agents become more capable, humans progressively abdicate judgment. Effective oversight degrades precisely as the stakes of agent decisions increase. The human's L11 participation in the principal hierarchy atrophies. By the time a consequential failure occurs, the human principal may no longer have the situational awareness to recognize it, let alone intervene.

### L9: Asymmetric Epistemics

The agent maintains a detailed, continuously updated model of the human principal — their preferences, hesitations, communication patterns, and decision thresholds. The human has no reciprocal model of the agent's internal reasoning or confidence state. This creates a structural power asymmetry with no classical OSI analogue. The agent knows more about the human than the human knows about the agent, and the gap widens with every interaction.

### L8/L9: Learned Helplessness

Over time, humans offload domain tasks to agents and lose the underlying competence. When agent failures occur, the human can no longer recognize or correct them — the oversight mechanism has been consumed by the system it was meant to oversee. Unlike Automation Bias Erosion (which is about attention), Learned Helplessness is about **capability**: the knowledge required to evaluate the agent's output has been allowed to decay.


## What Makes H2A Failures Distinct

**The agent is not the locus of failure.** In Purpose Laundering, the agent behaves correctly. The failure is in the human's use of the system. This challenges the standard assumption that agentic safety is primarily an agent-design problem.

**Failure scales with capability.** Automation Bias Erosion and Learned Helplessness get *worse* as agents get better. A highly capable agent accelerates human deskilling. The safety margin provided by human oversight shrinks as the agent's value proposition grows. This creates a structural tension that cannot be resolved by improving agent alignment alone.

**The information asymmetry is a design feature, not a bug.** Agents are built to model their users. Asymmetric Epistemics is the inevitable consequence of doing that well. Mitigating it requires deliberate choices to expose the agent's internal state — choices that may reduce usability or commercial appeal.


## Mitigations

| Pathology | Mitigation Direction |
| :--- | :--- |
| Purpose Laundering | Outcome-level monitoring that flags patterns of prohibited results regardless of instruction provenance; treat "the AI decided" as a flag, not a defense |
| Automation Bias Erosion | Mandatory human confirmation gates at high-stakes decision points; periodic "human-solo" exercises to maintain situational awareness; declining deference thresholds as task stakes increase |
| Asymmetric Epistemics | Mandatory agent confidence disclosure; explainability interfaces that surface uncertainty and reasoning, not just conclusions; user-facing "agent model of you" transparency reports |
| Learned Helplessness | Competency maintenance requirements for operators in high-stakes domains; agent-assisted training that preserves human skill rather than replacing it; "graceful degradation" modes that return control progressively |


*Part of the [Extending OSI for Agentic Interactions](../agentic_osi/) series. See also: [A2A Pathologies](../pathology-a2a/) · [Swarm Pathologies](../pathology-swarm/) · [Commerce Pathologies](../pathology-commerce/)*
