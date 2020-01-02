package utils

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleWorstCaseSort(t *testing.T) {
	els := []int{9, 8, 7, 6, 5}
	BubbleSort(els)
	assert.EqualValues(t, []int{5, 6, 7, 8, 9}, els)
}

func TestBubbleBestCaseSort(t *testing.T) {
	els := []int{5, 6, 7, 8, 9}
	BubbleSort(els)
	assert.EqualValues(t, []int{5, 6, 7, 8, 9}, els)
}

func TestBubbleSortNil(t *testing.T) {
	assert.NotPanics(t, func() { BubbleSort(nil) })
}

func TestBubbleNormalCaseSort(t *testing.T) {
	els := []int{5, 2, 6, 4, 3}
	BubbleSort(els)
	assert.EqualValues(t, []int{2, 3, 4, 5, 6}, els)
}

func getElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}
func BenchmarkBubbleSort10000Elements(b *testing.B) {
	for i := 0; i < b.N; i++ {
		els := getElements(1000)
		BubbleSort(els)
	}
}

func BenchmarkSort10000Elements(b *testing.B) {
	for i := 0; i < b.N; i++ {
		els := getElements(1000)
		sort.Ints(els)
	}
}
