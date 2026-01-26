package main

import (
	"os"
	"time"

	"github.com/aileks/pomodoro-timer/pkg/timer"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	workDuration      time.Duration
	breakDuration     time.Duration
	longBreakDuration time.Duration
	timer             *timer.PomoTimer
	totalSessions     int
	session           int
	width             int
	height            int
	cycleCount        int    // tracks pomodori
	phase             string // "work" or "break"
	state             string // "running" or "prompt"
	awaitingInput     string
}

type tickMsg time.Time

func (m *Model) Init() tea.Cmd {
	m.timer = timer.New(m.workDuration)
	m.timer.Start()

	cmds := []tea.Cmd{m.tickCmd()}
	if os.Getenv("TMUX") == "" {
		cmds = append(cmds, tea.EnterAltScreen)
	}
	return tea.Batch(cmds...)
}

func (m *Model) tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
