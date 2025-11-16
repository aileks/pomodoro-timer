package main

import (
	"flag"
	"fmt"
)

func main() {
	workMain := flag.Int("work", 25, "work duration in minutes")
	breakMain := flag.Int("break", 5, "break duration in minutes")
	flag.Parse()

	fmt.Printf("Work: %d min, Break: %d min\n", *workMain, *breakMain)
}
