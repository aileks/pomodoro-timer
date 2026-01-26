package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/aileks/pomodoro-timer/pkg/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gen2brain/beeep"
)

func notifyOS(title, message string) {
	switch runtime.GOOS {
	case "darwin":
		notifyMacOS(title, message)
	case "linux":
		notifyLinux(title, message)
	case "windows":
		notifyWindows(title, message)
	}
}

func notifyMacOS(title, message string) {
	script := fmt.Sprintf(`display notification "%s" with title "%s"`, message, title)
	exec.Command("osascript", "-e", script).Run()
}

func notifyLinux(title, message string) {
	exec.Command("notify-send", title, message).Run()
}

// TODO: Check if this works on a Windows machine
func notifyWindows(title, message string) {
	script := fmt.Sprintf(`
	powershell -Command "Add-Type -AssemblyName PresentationFramework; [System.Windows.MessageBox]::Show('%s', '%s')"
	`, message, title)
	exec.Command("cmd", "/C", script).Run()
}

func parseSessionCount(s string) int {
	count := 0
	for _, ch := range s {
		count = count*10 + int(ch-'0')
	}
	return count
}

func trimLastRune(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	return string(runes[:len(runes)-1])
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.state == "prompt" || m.awaitingInput != "" {
			// Collecting additional sessions input
			switch msg.String() {
			case "q":
				return m, tea.Quit
			case "enter":
				if m.awaitingInput == "" {
					return m, nil
				}
				if count := parseSessionCount(m.awaitingInput); count > 0 {
					m.totalSessions += count
					m.awaitingInput = ""
					m.state = "running"
					m.timer = timer.New(m.workDuration)
					m.timer.Start()
					return m, nil
				}
				if m.awaitingInput == "0" {
					return m, tea.Quit
				}
			case "backspace", "ctrl+h", "delete":
				m.awaitingInput = trimLastRune(m.awaitingInput)
			default:
				if msg.String() >= "0" && msg.String() <= "9" {
					m.awaitingInput += msg.String()
				}
			}
			return m, nil
		}

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
		}

	case tickMsg:
		if m.timer == nil {
			m.timer = timer.New(m.workDuration)
			m.timer.Start()
			m.state = "running"
		}

		if m.timer.IsFinished() {
			if m.phase == "work" {
				m.cycleCount++

				// Determine break duration
				breakDur := m.breakDuration
				if m.cycleCount%4 == 0 {
					breakDur = m.longBreakDuration
				}

				notifyOS("Pomodoro", "Break time! Take a rest.")

				m.phase = "break"
				m.timer = timer.New(breakDur)
				m.timer.Start()
			} else {
				notifyOS("Pomodoro", "Break over! Time to work.")

				m.phase = "work"
				m.session++

				if m.totalSessions > 0 && m.session > m.totalSessions {
					m.state = "prompt"
					m.awaitingInput = ""
				} else {
					m.timer = timer.New(m.workDuration)
					m.timer.Start()
				}
				beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
			}
		}

		return m, m.tickCmd()

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		if os.Getenv("DEBUG") == "1" {
			log.Printf("window size: %dx%d", m.width, m.height)
		}
	}

	return m, nil
}
