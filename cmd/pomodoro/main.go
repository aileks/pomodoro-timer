// A pomodoro CLI utility
package main

import (
	"flag"
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	workMin := flag.Int("work", 25, "work duration in minutes")
	breakMin := flag.Int("break", 5, "break duration in minutes")
	longBreakMin := flag.Int("long-break", 15, "long break duration in minutes")
	sessions := flag.Int("sessions", 4, "number of sessions before exiting (0 for infinite)")
	flag.Parse()

	m := Model{
		workDuration:      time.Duration(*workMin) * time.Minute,
		breakDuration:     time.Duration(*breakMin) * time.Minute,
		longBreakDuration: time.Duration(*longBreakMin) * time.Minute,
		totalSessions:     *sessions,
		session:           1,
		phase:             "work",
		state:             "running",
	}

	p := tea.NewProgram(&m)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
