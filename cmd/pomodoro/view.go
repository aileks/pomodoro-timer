package main

import (
	"fmt"
	"strings"

	"github.com/aileks/pomodoro-timer/pkg/timer"
	"github.com/charmbracelet/lipgloss"
)

func sevenSegmentDisplay(digit rune) []string {
	segments := map[rune][]string{
		'0': {" _  ", "| | ", "|_| "},
		'1': {"    ", "  | ", "  | "},
		'2': {" _  ", " _| ", "|_  "},
		'3': {" _  ", " _| ", " _| "},
		'4': {"    ", "|_| ", "  | "},
		'5': {" _  ", "|_  ", " _| "},
		'6': {" _  ", "|_  ", "|_| "},
		'7': {" _  ", "  | ", "  | "},
		'8': {" _  ", "|_| ", "|_| "},
		'9': {" _  ", "|_| ", " _| "},
		':': {"  ", "• ", "• "},
	}
	if seg, ok := segments[digit]; ok {
		return seg
	}
	return []string{"   ", "   ", "   "}
}

func buildTimerDisplay(timeStr string, phase string) string {
	var lines [3]string

	for _, ch := range timeStr {
		segs := sevenSegmentDisplay(ch)
		for i := range 3 {
			lines[i] += segs[i]
		}
	}

	display := strings.Join(lines[:], "\n")

	var color string
	if phase == "work" {
		color = "1" // red
	} else {
		color = "2" // green
	}

	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(color)).
		Bold(true).
		Render(display)
}

func (m *Model) View() string {
	remaining := timer.FormatDuration(m.timer.Remaining())
	timerDisplay := buildTimerDisplay(remaining, m.phase)

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

	progress := int(float64(total-int64(m.timer.Remaining())) / float64(total) * 50)
	bar := "[" + strings.Repeat("█", progress) + strings.Repeat("░", 50-progress) + "]"

	status := m.state
	if status == "paused" {
		status = "| [PAUSED]"
	} else {
		status = ""
	}

	header := fmt.Sprintf("Session: %d/%d | %s %s", m.session, m.totalSessions, strings.ToUpper(m.phase), status)

	commands := "q: Quit | p: Pause | r: Resume"

	var content string
	if m.awaitingInput != "" {
		prompt := fmt.Sprintf("Add more sessions? Current input: %s (press Enter to confirm)", m.awaitingInput)
		content = header + "\n\n" + timerDisplay + "\n\n\n" + bar + "\n\n" + prompt
	} else if m.state == "prompt" {
		content = "All sessions complete!\n\nHow many additional sessions? (0 to quit, or enter a number):"
	} else {
		content = header + "\n\n" + timerDisplay + "\n\n\n" + bar + "\n\n" + commands
	}

	return lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Align(lipgloss.Center).
		AlignVertical(lipgloss.Center).
		Render(content)
}
