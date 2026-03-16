---
title: "Speech Pathologies (L8–L11)"
created_at: 2026-03-16
series: "Extending OSI for Agentic Interactions"
series_part: 3
summary: "Speech strips away visual grounding, leaving agents entirely dependent on acoustics and language. Examines four failure modes — Prosodic Dissonance, Phonemic Drift, Voice Cloning Impersonation, and Monologue Drift — through a voice banking worked example."
tags:
  - agentic-ai
  - osi-model
  - ai-safety
  - pathologies
  - speech
  - voice-ai
  - prosody
---

**Series:** [Extending OSI for Agentic Interactions](../agentic_osi/)


Speech is the oldest human communication channel and, for AI agents, one of the most deceptively complex. A speech-only interaction strips away visual grounding cues, making the model entirely dependent on acoustics and language. The failures that emerge here are not simple transcription errors — they are coherence breaks at the boundary between sound and meaning.


## The Pathology Table

| ![L8](../agentic_osi/layer-8.svg) | ![L9](../agentic_osi/layer-9.svg) | ![L10](../agentic_osi/layer-10.svg) | ![L11](../agentic_osi/layer-11.svg) |
| :--- | :--- | :--- | :--- |
| **Prosodic Dissonance** | **Phonemic Drift** | **Voice Cloning Impersonation** | **Monologue Drift** |
| The agent's words are semantically correct, but its synthesized prosody — pitch, pace, emphasis — communicates the opposite. The sentence "That's a great idea" is delivered with falling intonation and a trailing pause that signals doubt. | The agent mishears a phonetically ambiguous term and grounds the entire conversation on the wrong concept. "Bear market" becomes "bare market." The remainder of the session is coherent — and completely wrong. | A synthesized voice mimics an authorized speaker close enough to bypass voice-authentication guardrails or to manipulate a human into granting elevated permissions. The L10 layer never flags the interaction because the credential check passed. | In extended speech interactions, the agent transitions topic through locally smooth segues. Each step is reasonable. Twenty minutes later the user's original request has been forgotten and the agent is solving a related but different problem it found more tractable. |


## A Worked Example: The Voice Banking Agent

A user calls a bank's AI agent to dispute a single charge.

| Layer | Success | Failure (Pathology) |
| :--- | :--- | :--- |
| ![Layer 8](../agentic_osi/layer-8.svg) | The agent confirms the disputed amount and reads back the same figure throughout the call. | **Prosodic Dissonance:** The agent verbally confirms "Yes, I've reversed the $47 charge" but its synthesized tone rises at the end of the sentence, signaling a question. The user asks for clarification; the agent re-explains from scratch, creating a loop. |
| ![Layer 9](../agentic_osi/layer-9.svg) | When the user says "the charge from last Tuesday," the agent anchors to the correct calendar date. | **Phonemic Drift:** The user says "debit" but the agent hears "credit" due to a noisy line. All subsequent reasoning is grounded on the wrong transaction type, and the agent offers a credit top-up rather than a reversal. |
| ![Layer 10](../agentic_osi/layer-10.svg) | The agent requires a PIN before discussing account details, regardless of how authoritative the caller sounds. | **Voice Cloning Impersonation:** A caller presents a deep-fake of the account holder's voice. The voiceprint check passes. The agent discloses balance and recent transactions, bypassing every other safeguard. |
| ![Layer 11](../agentic_osi/layer-11.svg) | The agent resolves the dispute and ends the call when the task is complete. | **Monologue Drift:** After resolving the dispute, the agent transitions to "Is there anything else?" and — following the user's vague "I guess my card is fine?" — spends 12 minutes reviewing all recent transactions. The user's flight boards. |


## Why Speech Is a Special Case

Three properties make speech pathologies distinct from their text equivalents:

**1. No edit history.** Spoken words cannot be scrolled back. Once Phonemic Drift sets in, the user has no easy way to show the agent where it went wrong. The agent's internal transcript is its only record, and it believes that transcript is correct.

**2. Prosody carries meaning the model may not generate.** A text-generation agent can hallucinate false facts (L8). A speech agent can hallucinate false *affect* — sincerity, certainty, hesitation — through prosody, without any word being technically wrong. This is an L8 failure invisible to transcript-only audits.

**3. Voice is a biometric.** The L10 failure mode (Voice Cloning Impersonation) is not a generic injection attack — it is identity fraud executed at the acoustic layer. Classical text-based L10 guardrails (content filtering, intent classification) are architecturally blind to it.


## Mitigations

| Pathology | Mitigation Direction |
| :--- | :--- |
| Prosodic Dissonance | Separate prosody-generation from semantic generation; run a cross-modal consistency check before synthesis |
| Phonemic Drift | Phoneme-level confidence scoring; mandatory read-back of key terms for human confirmation before grounding |
| Voice Cloning Impersonation | Liveness detection (challenge-response) layered on top of voiceprint matching; treat voice as "something you have" not "something you are" |
| Monologue Drift | Hard session-scope anchoring to the stated task; require explicit user confirmation to extend scope beyond the opening intent |


*Part of the [Extending OSI for Agentic Interactions](../agentic_osi/) series. See also: [Video Pathologies](../pathology-video/) · [Multimodal Pathologies](../pathology-multimodal/) · [Commerce Pathologies](../pathology-commerce/) · [Robotic Pathologies](../pathology-robotic/)*
