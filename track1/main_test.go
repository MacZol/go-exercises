package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, 6, Add(1, 2, 3))
}

func TestAdd2Int(t *testing.T) {
	assert.Equal(t, 5, Add(2, 3))
}

func TestAddNegativeInt(t *testing.T) {
	assert.Equal(t, -1, Add(2, -3))
}

func TestAddBigInt(t *testing.T) {
	assert.Equal(t, 100000001, Add(1, 100000000))
}

func TestAdd2BigInt(t *testing.T) {
	assert.Equal(t, 200000000000, Add(100000000000, 100000000000))
}

// how big