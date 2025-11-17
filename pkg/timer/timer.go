// Package timer provides a countdown timer with pause/resume capabilities.
package timer

import "time"

const (
	StateRunning = "running"
	StatePaused  = "paused"
	StateStopped = "stopped"
)

type PomoTimer struct {
	deadline  time.Time
	state     string
	remaining time.Duration
	timer     *time.Timer
}

func New(duration time.Duration) *PomoTimer {
	return &PomoTimer{
		state:     StateStopped,
		remaining: duration,
		timer:     nil,
	}
}

func Start() {
	// TODO
}

func Pause() {
	// TODO
}

func Resume() {
	// TODO
}
