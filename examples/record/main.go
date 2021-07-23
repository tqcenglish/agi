package main

import (
	"fmt"

	"github.com/tqcenglish/agi"
)

func main() {
	fmt.Printf("listen 18081")
	agi.Listen(":18081", handler)
}

func handler(a *agi.AGI) {
	defer a.Close()
	fmt.Print("record demo start\n")
	a.Answer()
	//保存路径 /var/lib/asterisk/sounds
	//按#号结束
	err := a.Record("tqcenglish", &agi.RecordOptions{})
	if err != nil {
		fmt.Printf("%+v", err)
	}
	fmt.Print("record demo end\n")
	a.Hangup()
}
