package redact

import (
	"encoding/json"
	"sync"
)

// TrieNode represents a node in the Trie
type TrieNode struct {
	Children    map[rune]*TrieNode `json:"children"`
	IsEndOfWord bool               `json:"is_end_of_word"`

	mutex sync.RWMutex `json:"-"`
}

// Trie represents the entire Trie structure
type Trie struct {
	Root *TrieNode `json:"root"`
}

// NewTrieNode initializes a new TrieNode
func NewTrieNode() *TrieNode {
	return &TrieNode{
		Children: make(map[rune]*TrieNode),
	}
}

// NewTrie initializes a new Trie
//
// The Trie data structure [1] is used to efficiently store and
// search for single and composed words like "Company" and "Company X"
//
// [1] https://en.wikipedia.org/wiki/Trie
func NewTrie(expressions ...string) *Trie {
	trie := &Trie{
		Root: NewTrieNode(),
	}
	for _, expression := range expressions {
		trie.Insert(expression)
	}
	return trie
}

func (tree *Trie) Insert(expression string) {
	if len(expression) == 0 {
		return
	}

	node := tree.Root

	for _, c := range expression {
		node.mutex.Lock()
		if _, exists := node.Children[c]; !exists {
			node.Children[c] = NewTrieNode()
		}
		node.mutex.Unlock()
		node = node.Children[c]
	}

	node.mutex.Lock()
	node.IsEndOfWord = true
	node.mutex.Unlock()
}

func (tree *Trie) Search(expression string) bool {
	node := tree.Root

	for _, c := range expression {
		node.mutex.RLock()
		next, found := node.Children[c]
		node.mutex.RUnlock()

		if !found {
			return false
		}

		node = next
	}
	node.mutex.RLock()
	defer node.mutex.RUnlock()
	return node.IsEndOfWord
}

// SerializeTrie serializes the Trie to JSON format
func (t *Trie) SerializeTrie() ([]byte, error) {
	return json.Marshal(t)
}

// DeserializeTrie deserializes JSON data into a Trie
func DeserializeTrie(data []byte) (*Trie, error) {
	var trie Trie
	err := json.Unmarshal(data, &trie)
	if err != nil {
		return nil, err
	}
	return &trie, nil
}
