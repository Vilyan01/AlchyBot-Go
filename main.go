package main

import (
	"fmt"
	"github.com/Vilyan01/AlchyBot-Go/bot"
)

func main() {
	b, err := bot.NewBot()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = b.Run()
	if err != nil {
		fmt.Println(err)
	}
}
