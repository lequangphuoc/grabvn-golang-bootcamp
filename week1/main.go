package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		x, y, op, err := parseInput(scanner.Text())
		if err != nil {
			fmt.Println(err)
			continue
		}

		calc(x, y, op)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}