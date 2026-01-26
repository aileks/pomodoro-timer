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
	ink := lipgloss.Color("#1b120d")
	inkSoft := lipgloss.Color("#2a1e17")
	muted := lipgloss.Color("#cbb7a0")
	white := lipgloss.Color("#f6e7d0")

	return Theme{
		Base: lipgloss.NewStyle().Foreground(white).Background(ink),
		Panel: lipgloss.NewStyle().
			Background(inkSoft).
			Padding(1, 3),
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
		return lipgloss.Color("#ff8a5b")
	}
	if longBreak {
		return lipgloss.Color("#f2c36b")
	}
	return lipgloss.Color("#d96b74")
}

func (t Theme) PanelWithAccent(accent lipgloss.Color) lipgloss.Style {
	return t.Panel.Copy()
}

func (t Theme) ProgressFill(accent lipgloss.Color) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(accent)
}
