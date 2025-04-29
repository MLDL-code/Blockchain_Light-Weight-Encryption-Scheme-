package cpt

import (
	"log"
)

type PatriciaTrie struct {
	children map[byte]*PatriciaTrie
	isEnd    bool
	data     string
}

func NewPatriciaTrie() *PatriciaTrie {
	return &PatriciaTrie{
		children: make(map[byte]*PatriciaTrie),
	}
}

func (pt *PatriciaTrie) Insert(key, data string) {
	node := pt
	for i := 0; i < len(key); i++ {
		if _, exists := node.children[key[i]]; !exists {
			node.children[key[i]] = NewPatriciaTrie()
		}
		node = node.children[key[i]]
	}
	node.isEnd = true
	node.data = data
	log.Println("Patricia Trie inserted data")
}

func (pt *PatriciaTrie) Search(key string) (string, bool) {
	node := pt
	for i := 0; i < len(key); i++ {
		if _, exists := node.children[key[i]]; !exists {
			return "", false
		}
		node = node.children[key[i]]
	}
	if node.isEnd {
		return node.data, true
	}
	return "", false
}
