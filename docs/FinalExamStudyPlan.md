# Final Exam Study Plan

The final exam coverage in `README.md` lists the following topics:

```
- Prolog: state space approach, heuristic search with A*
  - ppt for Weeks 13-15 on Teams
- Concurrency:
  - paper "Concepts and Notations for Concurrent Programming", Section 1, Section 3.1-3.2
  - Go book Section 8.1-8.5, 8.7, 9.1-9.2
```
(See README lines 1075-1081.)

Below is an intensive review schedule that ensures all materials are covered.

## Day 1: Prolog – State Space and A\* Search
- **Reading**: `materials/76_73_heuristic_search_UI1.pdf` (47 pages)
- **Goals**:
  - Understand state space representation
  - Learn how A\* evaluates nodes (`f(n) = g(n) + h(n)`)
  - Recognize conditions for admissibility
  - See examples of heuristic design
- **Estimated time**: 3–4 hours
- **Checklist**:
  - [ ] Review definitions of state, successor relation `s(X,Y,C)`
  - [ ] Work through A\* algorithm steps and the admissibility theorem
  - [ ] Summarize optimistic vs. pessimistic heuristics
  - [ ] Attempt simple exercises (e.g., 8-puzzle)
- **Unit review / exam tips**:
  - Practice drawing small search trees and computing `f`, `g`, `h` values.
  - Focus on how heuristics guide the search; be ready to explain admissibility.

## Day 2: Concurrency Concepts (Theory)
- **Reading**: `materials/Concepts and Notations for Concurrent Programming.pdf`
  - Section 1: Concurrent Programs—Processes
  - Section 3.1–3.2: Busy-Waiting and Semaphores
- **Estimated time**: 2–3 hours
- **Checklist**:
  - [ ] Identify key abstractions of processes, communication and synchronization
  - [ ] Understand busy-waiting vs. semaphore primitives
  - [ ] Write notes comparing language notations discussed in the paper
- **Unit review / exam tips**:
  - Prepare short definitions (process, critical section, semaphore, etc.).
  - Consider how mutual exclusion can be expressed in different languages.

## Day 3: Go Concurrency Basics (Go Book §8.1–8.5)
- **Reading**: `The.Go.Programming.Language.pdf`
- **Estimated time**: 4 hours
- **Checklist**:
  - [ ] Section 8.1 Goroutines – understand go statements
  - [ ] Section 8.2 Channels – unbuffered and buffered usage
  - [ ] Section 8.3 Looping – typical patterns with channels
  - [ ] Section 8.4 Select – multiplexing channel operations
  - [ ] Section 8.5 Looping in Parallel – running loops concurrently
  - [ ] Review examples in README: spinner, clock1/clock2, netcat1–3, reverb1–2, pipeline1–3
- **Unit review / exam tips**:
  - Code small examples to reinforce syntax.
  - Remember when to close channels and how select statements work.

## Day 4: Go Concurrency Advanced (Go Book §8.7, §9.1–9.2)
- **Reading**: continue with `The.Go.Programming.Language.pdf`
- **Estimated time**: 3 hours
- **Checklist**:
  - [ ] Section 8.7 Cancellation – using done channels
  - [ ] Section 9.1 Share by Communicating – pipeline patterns
  - [ ] Section 9.2 Concurrency with Shared Variables – mutexes, sync primitives
- **Unit review / exam tips**:
  - Understand the difference between sharing memory by communicating vs. communication by sharing memory.
  - Recall the standard pattern for goroutine cancellation.

## Day 5: Consolidation and Practice
- **Activities**:
  - Revisit key slides from Weeks 13–15 (Teams PPTs)
  - Solve sample Prolog search problems
  - Implement small Go programs using goroutines, channels and mutex
  - Review notes from the concurrency paper and Go book
- **Estimated time**: 3 hours
- **Checklist**:
  - [ ] Summarize each topic in your own words
  - [ ] Write potential exam questions and answer them
  - [ ] Ensure code examples compile and run
- **Final tips**:
  - Rehearse definitions out loud to aid memorization.
  - Sleep well before the exam; avoid last-minute cramming.

