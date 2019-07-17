package main

import (
	"log"
	"time"
)

func main() {
	printSomething("main")
	go printSomething("goroutine")
	printSomething("main")

	time.Sleep(300 * time.Millisecond)
}

func printSomething(env string) {
	log.Printf("[%s] Print something", env)
}
