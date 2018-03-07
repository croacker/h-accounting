package main

import (
	"fmt"

	"./filewatcher"
)

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println("File event", msg)
	}
}

func main() {
	var c = make(chan string)
	go printer(c)
	defer filewatcher.Watch("/home/alex/tmp", c).Close()
	doWait()
}

func doWait() {
	done := make(chan bool)
	<-done
}
