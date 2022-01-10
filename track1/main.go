package main

func Add(int ...int) int {
	sum := 0
	for _, num := range int {
		sum += num
	}

	return sum
}