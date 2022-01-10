package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeInt(t *testing.T) {
	assert.Equal(t, 123, MakeInt("123"))
}

func TestAddInt(t *testing.T) {
	assert.Equal(t, 5, AddInt(2,3))
}

func TestAddIntNumber(t *testing.T) {

}