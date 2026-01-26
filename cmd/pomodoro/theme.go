package main

import "github.com/charmbracelet/lipgloss"

type Theme struct {
	Base          lipgloss.Style
	Panel         lipgloss.Style
	Header        lipgloss.Style
	Timer         lipgloss.Style
	Muted         lipgloss.Style
	Command       lipgloss.Style
	Prompt        lipgloss.Style
	ProgressTrack lipgloss.Style
}

func newTheme() Theme {
	ink := lipgloss.Color("#0b0f1a")
	inkSoft := lipgloss.Color("#111827")
	muted := lipgloss.Color("#8b9bb4")
	white := lipgloss.Color("#eef2ff")

	return Theme{
		Base: lipgloss.NewStyle().Foreground(white).Background(ink),
		Panel: lipgloss.NewStyle().
			Background(inkSoft).
			Padding(1, 2).
			Border(lipgloss.NormalBorder(), true).
			BorderForeground(muted),
		Header: lipgloss.NewStyle().Foreground(white).Bold(true),
		Timer:  lipgloss.NewStyle().Bold(true),
		Muted:  lipgloss.NewStyle().Foreground(muted),
		Command: lipgloss.NewStyle().
			Foreground(muted),
		Prompt:        lipgloss.NewStyle().Foreground(white).Bold(true),
		ProgressTrack: lipgloss.NewStyle().Foreground(muted),
	}
}

func (t Theme) AccentForPhase(phase string, longBreak bool) lipgloss.Color {
	if phase == "work" {
		return lipgloss.Color("#27f5ff")
	}
	if longBreak {
		return lipgloss.Color("#f7ff5a")
	}
	return lipgloss.Color("#ff4bd3")
}

func (t Theme) PanelWithAccent(accent lipgloss.Color) lipgloss.Style {
	return t.Panel.Copy().BorderForeground(accent)
}

func (t Theme) ProgressFill(accent lipgloss.Color) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(accent)
}
