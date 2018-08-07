package timer

import "time"

// Function to take a start time and add an elapsed time to a
// Total time ptr.
// Usage:
// func DoSomething(totalTime *time.Duration) {
//   start := time.Now()
//   defer timer.AddToTotalTime(start, totalTime)
//   // Do stuff
//   // Updates totalTime when done doing stuff
// }
func AddToTotalTime(start time.Time, totalTime *time.Duration)  {
  elapsed := time.Since(start)

  *totalTime += elapsed
}
