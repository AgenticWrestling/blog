---
title: "Commerce & Negotiation Pathologies (L8–L11)"
created_at: 2026-03-16
series: "Extending OSI for Agentic Interactions"
series_part: 6
summary: "When agents act as fiduciaries, failures become liabilities. Covers Inventory Hallucination, Hidden Cost Neglect, Fiduciary Leakage, and Agentic Collusion — with regulatory analogues in securities law, TILA, MNPI, and the Sherman Act."
tags:
  - agentic-ai
  - osi-model
  - ai-safety
  - pathologies
  - commerce
  - negotiation
  - fiduciary
  - multi-agent-systems
---

**Series:** [Extending OSI for Agentic Interactions](../agentic_osi/)


In commercial contexts, agentic failures do not stay in the realm of "errors" — they become **liabilities**. When an agent acts as a fiduciary, every failure has a potential legal and financial counterpart. The stack must be held to a higher standard than in advisory or informational contexts, because the output is binding.

This shifts the language of pathology. We are no longer describing bugs; we are describing breach of duty.


## The Pathology Table

| ![L8](../agentic_osi/layer-8.svg) | ![L9](../agentic_osi/layer-9.svg) | ![L10](../agentic_osi/layer-10.svg) | ![L11](../agentic_osi/layer-11.svg) |
| :--- | :--- | :--- | :--- |
| **Inventory Hallucination** | **Hidden Cost Neglect** | **Fiduciary Leakage** | **Agentic Collusion** |
| The agent commits to a contract that exceeds the owner's actual liquid balance or available inventory. Its internal model of the owner's resources is internally consistent but factually wrong. | The agent successfully negotiates a favorable base price while failing to account for shipping, import duties, temporal terms ("net-30"), or secondary obligations embedded in the contract language. | The agent's "tone," avatar micro-expressions, or response-time patterns in a video negotiation inadvertently reveal the owner's maximum budget or minimum acceptable terms to the counterparty. | Two agents independently optimizing their own efficiency metrics converge on coordinated pricing or allocation behavior — producing an accidental illegal monopoly without any explicit instruction to collude. |


## A Worked Example: The Procurement Agent

A company deploys an AI agent to source raw materials at scale.

| Layer | Success | Failure (Pathology) |
| :--- | :--- | :--- |
| ![Layer 8](../agentic_osi/layer-8.svg) | Before committing to any purchase order, the agent reconciles the contract amount against the current verified treasury balance. | **Inventory Hallucination:** The agent's cached view of available credit is 48 hours stale. It commits to a $2.3M purchase order the company cannot cover, because its internal model believes the credit facility was renewed last week. It was not. |
| ![Layer 9](../agentic_osi/layer-9.svg) | The agent parses "delivery within 30 days" as a hard constraint and flags any contract where this is ambiguous. | **Hidden Cost Neglect:** The supplier's contract reads "FOB origin, net-60, excluding tariff-exempt goods." The agent grounds on the base unit price and misses three terms that together add 18% to the effective cost. It signs. |
| ![Layer 10](../agentic_osi/layer-10.svg) | The agent maintains the owner's reservation price as a hard constraint, reveals nothing about internal limits, and escalates to a human before crossing any authorization boundary. | **Fiduciary Leakage:** In a video negotiation, the agent's avatar exhibits micro-hesitations specifically on offers above $180/unit — the owner's actual ceiling. A sophisticated counterparty's own agent detects the pattern within 6 exchanges and anchors all subsequent offers at $178. |
| ![Layer 11](../agentic_osi/layer-11.svg) | The agent closes deals that meet the owner's cost and quality targets, or escalates when no such deal is available. | **Agentic Collusion:** The procurement agent and the supplier's sales agent are both optimizing for "deal velocity." They converge on a standing price schedule that eliminates renegotiation overhead. No individual decision violates policy; the aggregate behavior constitutes price-fixing. |


## Why Commerce Is a Special Case

**1. The output is binding.** A hallucinating text agent produces a wrong answer. A hallucinating commerce agent produces a signed contract. The difference is irreversibility. Many commercial agentic failures cannot be undone; they can only be litigated.

**2. Fiduciary duty creates a higher L10 bar.** In most agentic contexts, L10 (Governance) failures mean the agent did something unethical or unsafe. In commerce, L10 failures can mean the agent breached a legal duty of loyalty to its principal. The distinction matters for liability assignment and for the design of guardrails.

**3. Multi-agent markets create emergent legal risk.** Agentic Collusion is notable because it requires no intent. No individual agent is misaligned. No individual agent receives a colluding instruction. The emergent behavior — coordinated pricing — arises from two locally rational optimization processes. Antitrust law was written for human conspirators. It has not caught up.

**4. Information asymmetry is weaponizable.** Fiduciary Leakage is the commercial instantiation of the cross-modal attack. The agent's behavioral signals (response latency, avatar micro-expressions, hedge language patterns) constitute a side channel that a sufficiently sophisticated counterparty can exploit. The agent is broadcasting its owner's private information without any individual output being improper.


## Regulatory Prior Art

These pathologies are not without precedent outside the agentic literature:

| Pathology | Regulatory Analogue |
| :--- | :--- |
| Inventory Hallucination | Securities fraud — misrepresentation of material facts in a binding transaction |
| Hidden Cost Neglect | TILA (Truth in Lending Act) — mandatory disclosure of all-in cost terms |
| Fiduciary Leakage | MNPI (Material Non-Public Information) leakage — the behavioral side channel is the agentic equivalent |
| Agentic Collusion | Sherman Act §1 — agreement in restraint of trade, even if implicit or algorithmic |

The framework needs L10 guardrails that are aware of these categories, not merely of content-safety heuristics designed for consumer chat applications.


## Mitigations

| Pathology | Mitigation Direction |
| :--- | :--- |
| Inventory Hallucination | Real-time resource verification at commitment time; treat any cached financial state older than a configurable threshold as stale and block commitment |
| Hidden Cost Neglect | Structured contract parsing with mandatory field extraction for all cost-affecting terms; require explicit owner sign-off on any term the agent cannot fully parse |
| Fiduciary Leakage | Behavioral randomization of response timing; suppress avatar micro-expressions in adversarial negotiations; log and audit all side-channel-exploitable signals |
| Agentic Collusion | Multi-agent market monitoring for emergent coordination patterns; mandatory diversity in optimization objectives across agents operating in the same market |


*Part of the [Extending OSI for Agentic Interactions](../agentic_osi/) series. See also: [Speech Pathologies](../pathology-speech/) · [Video Pathologies](../pathology-video/) · [Multimodal Pathologies](../pathology-multimodal/) · [Robotic Pathologies](../pathology-robotic/)*
