package trie

import (
	"unicode/utf8"
)

// Trie represents a trie and contains a pointer to the root node
type Trie struct {
	Root *Node
}

// Node represents a single node in a trie, usually representing a character
type Node struct {
	Data     string
	Prefix   string
	IsEnd    bool
	Children map[string]*Node
	Key      string
	Val      string
}

// Init creates a Trie pre-initialized with a pointer to the root Node
func Init() *Trie {
	t := &Trie{
		Root: &Node{
			Data:     "",
			Prefix:   "",
			IsEnd:    false,
			Children: make(map[string]*Node),
			Key:      "",
			Val:      "",
		},
	}
	return t
}

// Update adds a key, value pair to the Trie
// It functions similarly to a key, value pair in a map
// Once all of the letters of the key have been exhausted, the Val attribute of
// the leaf node in the trie is updated with the value v provided.
func (t *Trie) Update(k string, v string) {
	currentNode := t.Root
	l := utf8.RuneCountInString(k)
	for c := 0; c < l; c++ {
		s := string(k[c])
		if _, ok := currentNode.Children[s]; !ok {
			child := &Node{
				Data:     s,
				Prefix:   currentNode.Prefix + currentNode.Data,
				IsEnd:    false,
				Children: make(map[string]*Node),
				Key:      currentNode.Prefix + currentNode.Data,
				Val:      "",
			}
			currentNode.Children[s] = child
		}
		currentNode = currentNode.Children[s]
	}
	currentNode.Key = currentNode.Prefix + currentNode.Data
	currentNode.Val = v
	currentNode.IsEnd = true
}

func (t *Trie) Search(k string) (s string, ok bool) {
	currentNode := t.Root
	l := utf8.RuneCountInString(k)
	for c := 0; c < l; c++ {
		s := string(k[c])
		if _, ok := currentNode.Children[s]; ok {
			currentNode = currentNode.Children[s]
		} else {
			return "", false
		}
	}
	currentNode.Key = currentNode.Prefix + currentNode.Data
	if currentNode.IsEnd {
		s = currentNode.Val
	}
	return s, true
}
