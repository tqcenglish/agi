package main

import (
	"fmt"

	"github.com/tqcenglish/agi"
)

func main() {
	agi.Listen(":8080", handler)
}

func handler(a *agi.AGI) {
	defer a.Close()
	fmt.Print("play demo start\n")
	a.Answer()
	_, err := a.StreamFile("tqcenglish", "#", 0)
	if err != nil {
		fmt.Printf("%+v", err)
	}
	fmt.Print("play demo end\n")
	a.Hangup()
}
