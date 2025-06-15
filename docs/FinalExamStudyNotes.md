# Final Exam Study Notes

These notes summarize the key points for each section of the [FinalExamStudyPlan](FinalExamStudyPlan.md).

## Day 1: Prolog – State Space and A* Search
- **State Space Representation**: A problem is formulated as states connected by the successor relation `s(X,Y,C)` where `X` and `Y` are states and `C` the cost.
- **A\* Algorithm**: Evaluates nodes using `f(n) = g(n) + h(n)` where `g(n)` is the cost so far and `h(n)` an admissible heuristic. Low `f` values indicate promising nodes.
- **Admissibility**: A heuristic is admissible when it never overestimates the true cost to reach the goal. This guarantees optimal solutions.
- **Heuristic Design**: Consider optimistic vs. pessimistic heuristics; choose admissible ones such as Manhattan distance for sliding puzzles.
- **Practice**: Draw small search trees and compute `f`, `g`, `h` for each node. Try exercises like the 8‑puzzle to reinforce.

## Day 2: Concurrency Concepts (Theory)
- **Processes**: Independent threads of control that interact via communication and synchronization.
- **Busy-Waiting**: Spinning loops to wait for conditions can waste CPU cycles. Semaphores provide blocking alternatives.
- **Semaphores**: Two operations, `P` (wait) and `V` (signal), manage access to critical sections and enforce mutual exclusion.
- **Language Notations**: The paper surveys different notations for expressing concurrency, including shared variables and message passing.
- **Exam Reminder**: Be ready to define terms such as process, critical section, and semaphore, and describe how they relate to mutual exclusion.

## Day 3: Go Concurrency Basics (Go Book §8.1–8.5)
- **Goroutines**: Lightweight threads started with the `go` keyword.
- **Channels**: Typed conduits used to send and receive values; can be buffered or unbuffered.
- **Select Statement**: Waits on multiple channel operations, enabling multiplexing.
- **Looping in Parallel**: Launch goroutines inside loops to run iterations concurrently. Remember to close channels when producers are done.
- **Examples**: Review README programs such as `spinner`, `clock1`, `netcat1–3`, `reverb1–2`, and the `pipeline` series for idiomatic patterns.

## Day 4: Go Concurrency Advanced (Go Book §8.7, §9.1–9.2)
- **Cancellation**: Use a `done` channel to signal goroutines to stop early, preventing goroutine leaks.
- **Share by Communicating**: Prefer channels to share data safely without explicit locks.
- **Shared Variables**: When necessary, use `sync.Mutex` or other primitives to protect data. Keep critical sections small.
- **Exam Focus**: Understand when to use channels vs. mutexes and recall the cancellation pattern shown in the book.

## Day 5: Consolidation and Practice
- **Review Slides**: Revisit the Teams PPTs for Weeks 13–15 on Prolog and concurrency.
- **Practice Problems**: Solve sample Prolog searches and implement small Go programs using goroutines, channels, and mutexes.
- **Self‑Testing**: Summarize topics in your own words, draft possible exam questions, and ensure all code examples compile and run correctly.
- **Final Advice**: Rehearse definitions aloud and rest well before the exam.
