package main

import "fmt"

func unorderedRoutine(id int, data interface{}) {
	fmt.Printf("%v %d\n", data, id)
}


func unorderedExecution() {
	arrayA := []string{"coba1", "coba2", "coba3"}
	arrayB := []string{"bisa1", "bisa2", "bisa3"}

	// zip setA and setB and run routine using goroutine
	for i := 1; i <=4; i++ {
		go unorderedRoutine(i, arrayA)
		go unorderedRoutine(i, arrayB)
	}
}
