package main

import (
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

