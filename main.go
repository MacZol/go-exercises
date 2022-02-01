package main

import (
	"flag"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	numbs := loadNumbers()
	total := add(numbs...)
	printTotal(total)
}

func loadNumbers() []int {
	var filelocation string
	flag.StringVar(&filelocation, "input-file", "", "Allows you to open add values of specific file")
	flag.Parse()

	if filelocation != "" {
		total := []int{}
		for _, osArg := range os.Args[2:] {
			loadedFile := loadInputFromFile(osArg)
			sum := add(loadedFile...)
			total = append(total, sum)
		}
		return total
	}
	if len(os.Args) == 1 {
		return loadInputFromFile("input.txt")
	}

	return collectCliArgs()
}

func add(int ...int) int {
	sum := 0
	for _, num := range int {
		sum += num
	}

	return sum
}

func printTotal(total int) {
	p := message.NewPrinter(language.English)
	p.Println(total)
}

func makeInt(str string) int {
	intVal, err := strconv.Atoi(str)
	check(err)
	return intVal
}

func collectCliArgs() []int {
	var numbers []int

	// Collect arguments and convert to number array
	for _, s := range os.Args[1:] {
		i:= makeInt(s)
		numbers = append(numbers, i)
	}
	return numbers
}

func loadInputFromFile(fileName string) []int {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	file := string(bytes)
	return numsFromStrings(file)
}

func check(e error) {
	if e !=nil {
		log.Panic(e)
	}
}

func numsFromStrings(input string) []int {
	trimmedInput:= strings.TrimSpace(input)
	if len(trimmedInput) == 0 {
		return []int{}
	}
	noSpaces:= strings.ReplaceAll(trimmedInput, " ", "")
	replaceCommasWithNewLines := strings.ReplaceAll(noSpaces, ",", "\n")
	strs := strings.Split(replaceCommasWithNewLines, "\n")
	nums := []int{}
	for _, s := range strs {
		value, err := strconv.Atoi(s)
		if err != nil {
			log.Println("not an int: ", s)
		}
		nums =  append(nums, value)
	}
	return nums
}