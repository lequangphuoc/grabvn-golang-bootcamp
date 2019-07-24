package main

import(
	"fmt"
)

func pingpong() {
	var ping = make(chan int)
	var pong = make(chan int)

	go func () {
		for {
			<- ping
			fmt.Println("ping")
			pong <- 1
		}
	} ()

	go func () {
		for {
			<- pong
			fmt.Println("pong")
			ping <- 1
		}
	} ()

	ping <- 1

	fmt.Scanln()
    fmt.Println("done")
}