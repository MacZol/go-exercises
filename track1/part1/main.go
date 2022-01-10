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

func AddInt(int1 int, int2 int) int {
	return int1 + int2
}

func main() {
	fmt.Println(AddInt(5,4))
}