package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeInt(t *testing.T) {
	assert.Equal(t, 123, MakeInt("123"))
}

func TestAddInts(t *testing.T) {
	assert.Equal(t, 6, AddInts(1, 2, 3))
}