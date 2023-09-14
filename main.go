package main

import "github.com/AtinAgnihotri/gokedex/repl"

func main() {
	signal := make(chan int)
	go repl.Repl(signal)
	<-signal
}
