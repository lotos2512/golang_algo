package bloom_filter

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBloomFilter(t *testing.T) {
	bloom := NewBloomFilter(1000, 3)

	// Добавление элементов
	bloom.Add([]byte("apple"))
	bloom.Add([]byte("banana"))
	bloom.Add([]byte("cherry"))

	// Проверка наличия добавленных элементов
	testCases := []struct {
		element    string
		shouldBeIn bool
	}{
		{"apple", true},
		{"banana", true},
		{"cherry", true},
		{"orange", false},
	}

	for _, tc := range testCases {
		exists := bloom.Contains([]byte(tc.element))
		if exists != tc.shouldBeIn {
			t.Errorf("Expected element '%s' existence: %v, got: %v", tc.element, tc.shouldBeIn, exists)
		}
	}
}

func TestBloomFilterFalsePositives(t *testing.T) {
	size := uint(1000)
	numHashFuncs := uint(3)

	bloom := NewBloomFilter(size, numHashFuncs)

	// Добавляем элементы
	bloom.Add([]byte("apple"))
	bloom.Add([]byte("banana"))
	bloom.Add([]byte("cherry"))

	// Генерируем случайные элементы, которые не были добавлены
	var r = 0
	for i := 0; i < 1000; i++ {
		randomElement := fmt.Sprintf("random%d", i)
		belongs := bloom.Contains([]byte(randomElement))
		if belongs {
			r++
		}
	}
	assert.Equal(t, 4, r)
}
