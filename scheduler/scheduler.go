package scheduler

import (
	"time"

	"github.com/go-co-op/gocron"
)

// Task is type for task functions
type Task func()

// Run schedules a new task with an interval.
// Interval can be an int, time.Duration or a string that
// parses with time.ParseDuration().
// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
func Run(interval interface{}, task Task) {
	s := gocron.NewScheduler(time.UTC)

	_, _ = s.Every(interval).Do(task)

	s.StartAsync()
}
