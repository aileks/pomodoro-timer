package main

import (
	"github.com/aileks/pomodoro-timer/pkg/timer"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "p":
			if m.state == "running" {
				m.timer.Pause()
				m.state = "paused"
			}
		case "r":
			if m.state == "paused" {
				m.timer.Resume()
				m.state = "running"
			}
		case "q":
			return m, tea.Quit
		case "y":
			if m.state == "prompt" {
				m.session++
				m.phase = "work"
				m.state = "running"
				m.timer = timer.New(m.workDuration)
				m.timer.Start()
			}
		case "n":
			if m.state == "prompt" {
				return m, tea.Quit
			}
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tickMsg:
		if m.timer.IsFinished() {
			if m.phase == "work" {
				m.phase = "break"
				m.timer = timer.New(m.breakDuration)
				m.timer.Start()
			} else {
				m.state = "prompt"
			}
		}

		return m, m.tickCmd()
	}

	return m, nil
}
