package main

import (
	"time"

	"github.com/aileks/pomodoro-timer/pkg/timer"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	workDuration  time.Duration
	breakDuration time.Duration
	session       int
	phase         string // "work" or "break"
	state         string // "running" or "prompt"
	timer         *timer.PomoTimer
}

type tickMsg time.Time

func (m *Model) Init() tea.Cmd {
	m.timer = timer.New(m.workDuration)
	m.timer.Start()

	return tea.Batch(
		tea.EnterAltScreen,
		m.tickCmd(),
	)
}

func (m *Model) tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
