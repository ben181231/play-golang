package main

import (
	"log"
	"sync"
)

type printingEnvironment string

const (
	mainEnvironment      = printingEnvironment("main")
	goroutineEnvironment = printingEnvironment("goroutine")
)

func main() {
	wg := &sync.WaitGroup{}

	printSomething(mainEnvironment, wg)

	for idx := 0; idx < 5; idx++ {
		wg.Add(1)
		go printSomething(goroutineEnvironment, wg)
	}

	printSomething(mainEnvironment, wg)

	wg.Wait()
}

func printSomething(env printingEnvironment, wg *sync.WaitGroup) {
	log.Printf("[%s] Print something", env)
	if env != mainEnvironment {
		wg.Done()
	}
}
