package apf

import (
	"hash/fnv"
	"log"
)

type BloomFilter struct {
	bitset []bool
	size   int
}

func NewBloomFilter(size int) *BloomFilter {
	return &BloomFilter{
		bitset: make([]bool, size),
		size:   size,
	}
}

func (bf *BloomFilter) Add(data string) {
	hash := fnv.New32a()
	hash.Write([]byte(data))
	index := hash.Sum32() % uint32(bf.size)
	bf.bitset[index] = true
	log.Println("Bloom filter added record")
}

func (bf *BloomFilter) Check(data string) bool {
	hash := fnv.New32a()
	hash.Write([]byte(data))
	index := hash.Sum32() % uint32(bf.size)
	return bf.bitset[index]
}
