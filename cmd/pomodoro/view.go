package main

import (
	"fmt"
	"strings"

	"github.com/aileks/pomodoro-timer/pkg/timer"
	"github.com/charmbracelet/lipgloss"
	"github.com/common-nighthawk/go-figure"
)

func (m *Model) View() string {
	// ASCII art for timer
	remaining := timer.FormatDuration(m.timer.Remaining())
	fig := figure.NewFigure(remaining, "banner", true)
	asciiTimer := fig.String()

	// Progress bar
	total := int64(m.timer.Remaining())
	if m.phase == "work" {
		total = int64(m.workDuration)
	} else {
		total = int64(m.breakDuration)
	}

	progress := int(float64(total-int64(m.timer.Remaining())) / float64(total) * 50)
	bar := "[" + strings.Repeat("█", progress) + strings.Repeat("░", 50-progress) + "]"

	// Header
	status := m.state
	if status == "paused" {
		status = "PAUSED"
	}
	header := fmt.Sprintf("Session %d: %s (%s)", m.session, strings.ToUpper(m.phase), status)

	// Commands
	commands := "q: Quit | p: Pause | r: Resume"

	// Combine with centering
	style := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Foreground(lipgloss.Color("3"))

	if m.state == "prompt" {
		prompt := "Continue? (y/n)"
		return style.Render(header + "\n\n" + asciiTimer + "\n" + bar + "\n\n" + prompt)
	}

	return style.Render(header + "\n\n" + asciiTimer + "\n" + bar + "\n\n" + commands)
}
