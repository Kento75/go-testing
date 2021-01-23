package sort

import "testing"

// ベンチマーク
func BenchmarkBubbleSort(b *testing.B) {
	// Arrange
	elements := getElements(1000000)

	// Act
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

// ベンチマーク
func BenchmarkSort(b *testing.B) {
	// Arrange
	elements := getElements(1000000)
	// Act
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}

// test helper
func getElements(n int) []int {
	result := make([]int, n)
	j := 0
	for i := n; i > 0; i-- {
		result[j] = i
		j++
	}
	return result
}

func getElementsForASC(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = i + 1
	}
	return result
}
