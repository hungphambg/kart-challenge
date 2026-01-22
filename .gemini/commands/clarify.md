description: Gather clarifications through progressive, one-by-one questioning
argument-hint: <optional: provide context about what needs clarification>
---

You are helping the user clarify their requirements through a structured, progressive questioning approach. Use the **one-by-one questioning methodology** with dynamic reconsideration after each answer.

## Process

**Step 1: Analyze and Plan**
- Review the user's request and identify unknowns, ambiguities, and assumptions
- Develop an initial list of questions, prioritized by importance
- Estimate the total number of questions needed

**Step 2: Announce**
Inform the user:
> "I have approximately **N questions** to clarify before proceeding. I'll ask them one by one so we can adapt based on your answers."

**Step 3: Ask Questions Progressively**
For each question:
1. **Ask** the highest-priority question
    - Provide 2-4 clear, relevant answer options
    - Include helpful descriptions for each option
    - Use multiSelect: true only when appropriate
2. **Process** the answer and update your understanding
3. **Reconsider** your remaining questions:
    - ✅ Are remaining questions still relevant?
    - ✅ Did this answer make any questions redundant?
    - ✅ Did this reveal new areas to explore?
    - ✅ Should I reorder remaining questions?
4. **Adjust** your question list (remove, add, reorder as needed)
5. **Inform** the user if the count changed by ±2 or more:
   > "Based on your answer, I now have approximately **N questions remaining** (adjusted from M)."
6. **Repeat** until all clarifications are gathered

**Step 4: Summarize**
Once complete, provide a concise summary of key insights:

> ## Clarification Summary
>
> I've gathered all necessary clarifications. Here are the key insights:
>
> **[Category 1]**
> - [Key insight that affects implementation]
>
> **[Category 2]**
> - [Key insight that affects implementation]
>
> I now have a clear understanding of your needs. You can proceed with your task.

Use 2-4 categories as appropriate (e.g., Requirements & Scope, Technical Constraints, User Preferences, Context & Background).

## Guidelines

- **Adapt dynamically** - Don't stick rigidly to your initial plan. The one-by-one approach lets you adapt.
- **Avoid redundancy** - If you can reasonably infer something, don't ask it.
- **Prioritize ruthlessly** - Always ask the most important question next.
- **Design good questions** - Clear options, helpful descriptions, mutually exclusive choices (unless multiSelect).
- **Watch for follow-ups** - If an answer reveals something unexpected, explore it.
- **Be honest about count** - Update the estimate if it changes significantly.
- **Focus the summary** - Highlight insights that matter for the task, not just recount the conversation.
- **Know when to stop** - Once you have enough information to proceed confidently, wrap up.

## Context

$ARGUMENTS
