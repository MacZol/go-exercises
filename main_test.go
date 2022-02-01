package main

import (
	"github.com/matryer/is"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("Test adding 2 numbers", func(t *testing.T) {
		is := is.New(t)

		// Got
		got := add(2, 3)

		// Want
		want := 5

		// Assert
		is.Equal(want, got)
	})

	t.Run("Test adding 3 numbers", func(t *testing.T) {
		is := is.New(t)

		// Got
		got := add(1, 2, 3)

		// Want
		want := 6

		// Assert
		is.Equal(want, got)
	})

	t.Run("Adding negative numbers", func (t *testing.T) {
		is := is.New(t)

		// Got
		got := add(2, -3)

		// Want
		want := -1

		is.Equal(got, want)
	})

	t.Run("Adding one big int", func (t *testing.T) {
		is := is.New(t)

		// Got
		got := add(1, 100000000)

		// Want
		want := 100000001

		is.Equal(got, want)
	})

	t.Run("", func (t *testing.T) {
		is := is.New(t)

		// Got
		got := add(100000000000, 100000000000)

		// Want
		want := 200000000000

		is.Equal(got, want)
	})

}

func TestLoadInput(t *testing.T) {
	t.Run("Load input file", func(t *testing.T) {
		is := is.New(t)
		// Got
		got := loadInputFromFile("input.txt")

		// Want
		want := []int {
			4,
			5,
			32,
			100,
			867543,
		}

		// Assert
		is.Equal(want, got)

	})
}
