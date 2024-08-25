package main

import (
	"fmt"
	"sync"
)

func orderedRoutine(id int, data interface{}, mu *sync.Mutex, wg *sync.WaitGroup) {
	mu.Lock()

	defer wg.Done()
	defer mu.Unlock()

	fmt.Printf("%v %d\n", data, id)
}

func orderedExecution() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	arrayA := []string{"coba1", "coba2", "coba3"}
	arrayB := []string{"bisa1", "bisa2", "bisa3"}

	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go orderedRoutine(i, arrayA, &mu, &wg)
		go orderedRoutine(i, arrayB, &mu, &wg)
		wg.Wait()
	}
}
