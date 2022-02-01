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
		data := loadInputFromFile(filelocation)
		return NumsFromStrings(data)
	}
	if len(os.Args) == 1 {
		data := loadInputFromFile("input.txt")
		return NumsFromStrings(data)
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

func makeInt(str string) int{
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

func loadInputFromFile(fileName string) string {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func check(e error) {
	if e !=nil {
		log.Panic(e)
	}
}

//func readCsvFile(filePath string) [][]string {
//	f, err := os.Open(filePath)
//	if err != nil {
//		log.Fatal("Unable to read input file " + filePath, err)
//	}
//	defer f.Close()
//
//	csvReader := csv.NewReader(f)
//	records, err := csvReader.ReadAll()
//	if err != nil {
//		log.Fatal("Unable to parse file as CSV for " + filePath, err)
//	}
//
//	return records
//}

func NumsFromStrings(input string) []int {
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