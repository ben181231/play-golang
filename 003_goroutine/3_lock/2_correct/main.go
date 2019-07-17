package main

import (
	"log"
	"sync"
)

type counter struct {
	mutex sync.RWMutex
	value uint
}

func (m *counter) Increase() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.value++
}

func (m *counter) Value() uint {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	return m.value
}

func main() {
	wg := &sync.WaitGroup{}
	m := &counter{}

	for idx := 0; idx < 1000; idx++ {
		wg.Add(1)
		go doSomething(m, wg)
	}

	wg.Wait()

	log.Printf("Counter value: %d", m.Value())
}

func doSomething(ctr *counter, wg *sync.WaitGroup) {
	defer wg.Done()

	ctr.Increase()
}
