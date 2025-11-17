package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"golang.org/x/term"

	"github.com/aileks/pomodoro-timer/pkg/timer"
)

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
		if ch == 'p' || ch == 'r' || ch == 'q' {
			commandChan <- ch
		}
	}
}

func main() {
	workMin := flag.Int("work", 25, "work duration in minutes")
	breakMin := flag.Int("break", 5, "break duration in minutes")
	flag.Parse()

	fmt.Printf("%d\n", *breakMin)

	workDuration := time.Duration(*workMin) * time.Minute
	t := timer.New(workDuration)
	t.Start()

	commandChan := make(chan rune)
	go listenForInput(commandChan)

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
