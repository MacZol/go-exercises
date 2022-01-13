package main

import (
	"bufio"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"log"
	"os"
	"strconv"
	"flag"
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

func strToInt(s string) int {
	value, err := strconv.Atoi(s)
	check(err)
	return value
}

func Add(int ...int) int {
	sum := 0
	for _, num := range int {
		sum += num
	}

	return sum
}

func CollectArgs() []int {
	var numbers []int

	// Collect arguments and convert to number array
	for _, s := range os.Args[1:] {
		i:= strToInt(s)
		numbers = append(numbers, i)
	}
	return numbers
}

func LoadInput(fileName string) []int {
	input := []int{}
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		block := scanner.Text()
		input = append(input, MakeInt(block))
	}

	return input
}

func main() {
	// Print total
	p := message.NewPrinter(language.English)

	arguments := len(os.Args)

	if arguments == 1 {
		inputFile := LoadInput("input.txt")
		p.Printf("%d\n", Add(inputFile...))
	} else {
		p.Printf("%d\n", Add(CollectArgs()...))
	}

}