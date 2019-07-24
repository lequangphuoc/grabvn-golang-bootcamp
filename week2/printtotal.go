package main

import (
	"fmt"
)

func calcTotal(a, b, c <-chan int) {
	var total int
	select {
		case val := <-a:
			total += val
		case val := <-b:
			total += val
		case val := <-c:
			total += val 
	}

	fmt.Printf("Total: %v", total)
}

func printTotal() {
	input := make(chan int)
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)
	 
	go func() {
		val := <- input
		a <- val
		b <- val
		c <- val
	}()

	go func() {
		val := <-a
		result := val*2
		fmt.Println(result)
		a <- result
	}()

	go func() {
		val := <-b
		result := val*3
		fmt.Println(result)
		b <- result
	}()

	go func() {
		val := <-c
		result := val*4
		fmt.Println(result)
		c <- result
	}()

	go calcTotal(a, b, c)

	input <- 1

    fmt.Scanln()
    fmt.Println("done")
}