package main

import (
	"fmt"
	"log"
	"strconv"
)

func check(e error) {
	if e !=nil {
		log.Panic(e)
	}
}

func MakeInt(str string) int{
	intVal, err := strconv.Atoi(str)
	check(err)
	return intVal
}

func AddInts(int ...int) int {
	sum := 0
	for _, num := range int {
		sum += num
	}
	return sum
}

func main() {
	fmt.Println(AddInts(1, 2, 3, 4))
}