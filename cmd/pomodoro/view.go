package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/aileks/pomodoro-timer/pkg/timer"
	"github.com/charmbracelet/lipgloss"
)

func clampInt(value int, min int, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func (m *Model) View() string {
	ui := newTheme()

	if m.timer == nil {
		return ui.Muted.Render("Starting...")
	}

	remaining := timer.FormatDuration(m.timer.Remaining())
	longBreak := m.phase == "break" && m.cycleCount%4 == 0
	accent := ui.AccentForPhase(m.phase, longBreak)

	availableWidth := 0
	if m.width > 0 {
		availableWidth = m.width - 6
	}
	renderer := NewRenderer()
	block := renderer.Render(remaining, availableWidth)
	timerDisplay := strings.Join(block.Lines, "\n")
	timerDisplay = ui.Timer.Copy().Foreground(accent).Render(timerDisplay)

	var total int64
	if m.phase == "work" {
		total = int64(m.workDuration)
	} else {
		breakDur := m.breakDuration
		if m.cycleCount%4 == 0 {
			breakDur = m.longBreakDuration
		}
		total = int64(breakDur)
	}

	elapsed := total - int64(m.timer.Remaining())
	progressPercent := float64(elapsed) / float64(total)

	barWidth := block.Width
	if barWidth > 2 {
		barWidth -= 2
	}
	filledWidth := int(float64(barWidth) * progressPercent)
	var bar strings.Builder
	bar.WriteString("[")
	for i := 0; i < barWidth; i++ {
		if i < filledWidth {
			bar.WriteString(ui.ProgressFill(accent).Bold(true).Render("█"))
		} else {
			bar.WriteString(ui.ProgressTrack.Render("·"))
		}
	}
	bar.WriteString("]")

	status := ""
	if m.state == "paused" {
		status = "PAUSED"
	}

	phaseLabel := "BREAK"
	if m.phase == "work" {
		phaseLabel = "WORK"
	}
	if longBreak {
		phaseLabel = "LONG BREAK"
	}

	sessionTotal := fmt.Sprintf("%d", m.totalSessions)
	if m.totalSessions == 0 {
		sessionTotal = "inf"
	}

	header := fmt.Sprintf("Session %d/%s  %s", m.session, sessionTotal, phaseLabel)
	if status != "" {
		header = fmt.Sprintf("%s  %s", header, status)
	}

	commands := ui.Command.Render("q quit  p pause  r resume")

	mainBlock := strings.Join([]string{
		ui.Header.Foreground(accent).Render(header),
		"",
		timerDisplay,
		"",
		bar.String(),
		"",
		commands,
	}, "\n")

	content := ui.PanelWithAccent(accent).Render(mainBlock)

	if m.state == "prompt" {
		promptHeader := ui.Prompt.Render("Add more sessions")
		inputLine := ui.Muted.Render("Enter a number (0 to quit)")
		entry := ui.Header.Render("> " + m.awaitingInput)
		promptBlock := strings.Join([]string{promptHeader, "", inputLine, "", entry}, "\n")
		content = ui.PanelWithAccent(accent).Render(promptBlock)
	}

	if m.width == 0 || m.height == 0 || os.Getenv("DEBUG_NO_PLACE") == "1" {
		return content
	}

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)
}
