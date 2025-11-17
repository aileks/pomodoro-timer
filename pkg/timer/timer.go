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

func (t *PomoTimer) Start() {
	if t.state == StateRunning {
		return
	}

	t.deadline = time.Now().Add(t.remaining)
	t.state = StateRunning
	t.timer = time.NewTimer(t.remaining)
}

func (t *PomoTimer) Pause() {
	if t.state == StatePaused {
		return
	}

	t.timer.Stop()
	t.remaining = time.Until(t.deadline)
	t.state = StatePaused
}

func (t *PomoTimer) Resume() {
	if t.state != StatePaused {
		return
	}

	t.Start()
}

func (t *PomoTimer) Remaining() time.Duration {
	if t.state == StateRunning {
		return time.Until(t.deadline)
	}

	return t.remaining
}

func (t *PomoTimer) IsFinished() bool {
	if t.state == StateRunning {
		return time.Now().After(t.deadline)
	}

	return t.remaining <= 0
}
