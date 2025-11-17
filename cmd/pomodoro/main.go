package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/aileks/pomodoro-timer/pkg/timer"
)

func runPhase(duration time.Duration, label string, commandChan chan rune) bool {
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
				return false // goes to next work/break timer
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
				return true // completely exits program
			}
		}
	}
}

func listenForInput(commandChan chan rune, done chan struct{}) {
	reader := bufio.NewReader(os.Stdin)
	for {
		select {
		case <-done:
			return
		default:
			ch, _, err := reader.ReadRune()
			if err != nil {
				return
			}
			if ch == 'p' || ch == 'r' || ch == 'q' || ch == 'y' || ch == 'n' {
				commandChan <- ch
			}
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
	done := make(chan struct{})
	go listenForInput(commandChan, done)

	session := 1
	for {
		if runPhase(workDuration, fmt.Sprintf("Session %d: Work time!", session), commandChan) {
			close(done)
			return
		}
		if runPhase(breakDuration, fmt.Sprintf("Session %d: Break time!", session), commandChan) {
			close(done)
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
				close(done)
				return
			} else {
				fmt.Print("Not a valid input. ")
			}
		}
	}
}
