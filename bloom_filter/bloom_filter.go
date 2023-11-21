package bloom_filter

import (
	"hash/fnv"
)

import (
	"math"
)

// BloomFilter представляет структуру фильтра Блума
type BloomFilter struct {
	bitArray []byte
	size     uint
	hashFunc []func(data []byte) uint
}

// NewBloomFilter создает новый фильтр Блума заданного размера и с количеством хеш-функций
func NewBloomFilter(size uint, numHashFuncs uint) *BloomFilter {
	byteSize := uint(math.Ceil(float64(size) / 8)) // Размер в байтах для хранения битов
	return &BloomFilter{
		bitArray: make([]byte, byteSize),
		size:     size,
		hashFunc: createHashFuncs(numHashFuncs),
	}
}

// createHashFuncs создает несколько хеш-функций на основе хеша FNV
func createHashFuncs(numHashFuncs uint) []func(data []byte) uint {
	hashFuncs := make([]func(data []byte) uint, numHashFuncs)
	for i := uint(0); i < numHashFuncs; i++ {
		seed := i
		hashFuncs[i] = func(data []byte) uint {
			hash := fnv.New32a()
			hash.Write(data)
			return uint(hash.Sum32()) + seed
		}
	}
	return hashFuncs
}

// Add добавляет элемент в фильтр Блума
func (bf *BloomFilter) Add(data []byte) {
	for _, hashFunc := range bf.hashFunc {
		index := hashFunc(data) % bf.size
		byteIndex := index / 8
		bitIndex := index % 8
		bf.bitArray[byteIndex] |= 1 << bitIndex
	}
}

// Contains проверяет наличие элемента в фильтре Блума
func (bf *BloomFilter) Contains(data []byte) bool {
	for _, hashFunc := range bf.hashFunc {
		index := hashFunc(data) % bf.size
		byteIndex := index / 8
		bitIndex := index % 8
		if bf.bitArray[byteIndex]&(1<<bitIndex) == 0 {
			return false
		}
	}
	return true
}
