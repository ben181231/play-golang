package main

import "log"

func main() {
	printSomething("main")
	go printSomething("goroutine")
	printSomething("main")
}

func printSomething(env string) {
	log.Printf("[%s] Print something", env)
}
