package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatrix(t *testing.T) {
	originMatrix := [][]int{
		{2, 6},
		{3, 7},
		{4, 8},
		{5, 9},
	}
	updateMatrix := transposition(originMatrix)
	assert.Equal(t, updateMatrix, [][]int{
		{2, 3, 4, 5},
		{6, 7, 8, 9},
	})
}

func TestMatrix2(t *testing.T) {
	originMatrix := [][]int{
		{2, 3},
		{4, 5},
	}
	updateMatrix := transposition(originMatrix)
	assert.Equal(t, updateMatrix, [][]int{
		{2, 4},
		{3, 5},
	})
}
