package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/aileks/pomodoro-timer/pkg/timer"
)

func main() {
	workMin := flag.Int("work", 25, "work duration in minutes")
	breakMin := flag.Int("break", 5, "break duration in minutes")
	flag.Parse()

	fmt.Printf("Work set to: %d minutes\n", *workMin)
	fmt.Printf("Break set to: %d minutes\n", *breakMin)

	workDuration := time.Duration(*workMin) * time.Minute
	t := timer.New(workDuration)
	t.Start()

	// Display loop
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Printf("\r[%s]", timer.FormatDuration(t.Remaining()))

		if t.IsFinished() {
			fmt.Println("\nTime's up!")
			break
		}
	}
}
