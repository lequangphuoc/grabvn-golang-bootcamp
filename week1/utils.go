package main

import (
	"fmt"
	"strings"
	"strconv"
)

func add(x, y float64) float64 {
	return x + y
}

func subtract(x, y float64) float64 {
	return x - y
}

func multiply(x, y float64) float64 {
	return x * y
}

func divide(x, y float64) (result float64, err error) {
	if y == 0 {
		err = fmt.Errorf("divide by zero")
		return
	}

	result = float64(x) / float64(y)
	return
}

func parseInput(input string) (x float64 , y float64 , op string, err error) {
	s := strings.Fields(input)
	if len(s) < 3 {
		err = fmt.Errorf("input not correct format")
		return
	}

	op = s[1]

	x, err = strconv.ParseFloat(s[0], 64)
	if err != nil {
		err = fmt.Errorf("%v is not a number", s[0])
		return
	}

	y, err = strconv.ParseFloat(s[2], 64)
	if err != nil {
		err = fmt.Errorf("%v is not a number", s[2])
		return 
	}

	return
}

func calc(x, y float64, op string) {
	var result float64
	var err error

	switch op {
	case "+":
		result = add(x, y)
	case "-":
		result = subtract(x, y)
	case "*":
		result = multiply(x, y)
	case "/":
		result, err = divide(x, y)
	default:
		fmt.Println("Operation not allowed")
		return
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(x, op, y, "=", result)
}