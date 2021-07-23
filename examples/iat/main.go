package main

import (
	"fmt"

	"github.com/tqcenglish/agi"
)

func main() {
	fmt.Print("start 18081\n")
	agi.Listen(":18081", handler)
	fmt.Print("exit")
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

	a.StreamFile("/var/lib/asterisk/sounds/tqcenglish", "#", 0)

	// 开始转化
	result := wavToText("/var/lib/asterisk/sounds/tqcenglish.wav")
	fmt.Printf("get result %s\n", result)
	a.Hangup()
}
