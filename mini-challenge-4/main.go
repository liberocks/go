package main

import (
	"fmt"
	"time"
)

 

func main() {
	fmt.Println("Unordered Execution")
	unorderedExecution()

	time.Sleep(300 * time.Millisecond)
	fmt.Println("===================")

	fmt.Println("Ordered Execution")
	orderedExecution()

	time.Sleep(300 * time.Millisecond)
}