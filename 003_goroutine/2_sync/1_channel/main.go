package main

import "log"

type printingEnvironment string

const (
	mainEnvironment      = printingEnvironment("main")
	goroutineEnvironment = printingEnvironment("goroutine")
)

func main() {
	done := make(chan bool, 2)

	printSomething(mainEnvironment, done)

	go printSomething(goroutineEnvironment, done)
	go printSomething(goroutineEnvironment, done)

	printSomething(mainEnvironment, done)

	for idx := 0; idx < 2; idx++ {
		select {
		case <-done:
		}
	}
}

func printSomething(env printingEnvironment, done chan<- bool) {
	log.Printf("[%s] Print something", env)
	if env != mainEnvironment {
		done <- true
	}
}
