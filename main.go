package main

import (
	"github.com/AtinAgnihotri/gokedex/utils"
)

func main() {
	signal := make(chan int)
	go utils.Repl(signal)
	<-signal
}
