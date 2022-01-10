package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, 6, Add(1, 2, 3))
}

func TestAddInt(t *testing.T) {
	assert.Equal(t, 5, Add(2, 3))
}

// Positive

// Negative

// Big numbers

// 2 Really big numbers - how big