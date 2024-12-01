package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDistance(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}

	result, err := distance(left, right)
	require.NoError(t, err)
	assert.Equal(t, result, 11)
}

func TestSimilarity(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}

	result := similarity(left, right)
	assert.Equal(t, result, 31)
}
