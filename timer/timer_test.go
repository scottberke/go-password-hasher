package timer

import (
  "testing"
  "time"
)

func TestTimer1Second(t *testing.T) {
  var delay     time.Duration = 1
  var totalTime time.Duration

  func (delay time.Duration, totalTime *time.Duration) {
    start := time.Now()
    defer AddToTotalTime(start, totalTime)

    time.Sleep(delay * time.Second)
  }(delay, &totalTime)

  // Just test that the totalTime recorded is greater than timer delay
  // since theres some execution time included in the totalTime
  if (delay * time.Second) > (totalTime * time.Second) {
    t.Errorf("delay == %q, want %q", delay * time.Second, totalTime)
  }

}

func TestTimer3Seconds(t *testing.T) {
  var delay     time.Duration = 3
  var totalTime time.Duration

  func (delay time.Duration, totalTime *time.Duration) {
    start := time.Now()
    defer AddToTotalTime(start, totalTime)

    time.Sleep(delay * time.Second)
  }(delay, &totalTime)

  // Just test that the totalTime recorded is greater than timer delay
  // since theres some execution time included in the totalTime
  if (delay * time.Second) > (totalTime * time.Second) {
    t.Errorf("delay == %q, want %q", delay * time.Second, totalTime)
  }

}
