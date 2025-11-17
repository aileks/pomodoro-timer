package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"golang.org/x/term"

	"github.com/aileks/pomodoro-timer/pkg/timer"
)

func runPhase(duration time.Duration, label string, commandChan chan rune) {
	fmt.Println("Commands:")
	fmt.Println("q: Quit | p: Pause")
	fmt.Printf("\n%s\n", label)

	t := timer.New(duration)
	t.Start()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		// Workaround to display the time immediately
		fmt.Printf("\r[%s]", timer.FormatDuration(t.Remaining()))

		select {
		case <-ticker.C:
			fmt.Printf("\r[%s]", timer.FormatDuration(t.Remaining()))
			if t.IsFinished() {
				fmt.Println("\nTime's up!")
				return
			}
		case cmd := <-commandChan:
			switch cmd {
			case 'p':
				t.Pause()
				fmt.Println("\nPaused")
			case 'r':
				t.Resume()
				fmt.Println("\nResumed")
			case 'q':
				fmt.Println("\nQuitting...")
				return
			}
		}
	}
}

func listenForInput(commandChan chan rune) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return
	}

	defer func() {
		if err := term.Restore(int(os.Stdin.Fd()), oldState); err != nil {
			fmt.Fprintf(os.Stderr, "failed to restore terminal: %v\n", err)
		}
	}()

	buf := make([]byte, 1)
	for {
		if _, err := os.Stdin.Read(buf); err != nil {
			return
		}

		ch := rune(buf[0])
		if ch == 'p' || ch == 'r' || ch == 'q' || ch == 'y' || ch == 'n' {
			commandChan <- ch
		}
	}
}

func main() {
	workMin := flag.Int("work", 25, "work duration in minutes")
	breakMin := flag.Int("break", 5, "break duration in minutes")
	flag.Parse()

	workDuration := time.Duration(*workMin) * time.Minute
	breakDuration := time.Duration(*breakMin) * time.Minute

	commandChan := make(chan rune)
	go listenForInput(commandChan)

	session := 1
	for {
		if runPhase(workDuration, fmt.Sprintf("Session %d: Work time!", session), commandChan) {
			return
		}
		if runPhase(breakDuration, fmt.Sprintf("Session %d: Break time!", session), commandChan) {
			return
		}

		fmt.Print("\nContinue? (y/n): ")
		for {
			cmd := <-commandChan
			if cmd == 'y' {
				session++
				break
			} else if cmd == 'n' {
				fmt.Println("Done!")
				return
			} else {
				fmt.Print("Not a valid input. ")
			}
		}
	}
}
