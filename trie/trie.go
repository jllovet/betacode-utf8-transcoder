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
	trie := &Trie{
		Root: &Node{
			Data:     "",
			Prefix:   "",
			IsEnd:    false,
			Children: make(map[string]*Node),
			Key:      "",
			Val:      "",
		},
	}
	return trie
}

// Update adds a key, value pair to the Trie
// It functions similarly to a key, value pair in a map
// Once all of the letters of the key have been exhausted, the Val attribute of
// the leaf node in the trie is updated with the value v provided.
func (trie *Trie) Update(k string, v string) {
	currentNode := trie.Root
	lenRunes := utf8.RuneCountInString(k)
	for idx := 0; idx < lenRunes; idx++ {
		s := string(k[idx])
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

// Search attempts to retrieve the corresponding value for the key provided
// It provides a similar experience to referencing a key, value pair in a map
// where both the string value and a bool are returned if the key is found;
// otherwise, it returns an empty string and false.
func (trie *Trie) Search(k string) (s string, ok bool) {
	currentNode := trie.Root
	lenRunes := utf8.RuneCountInString(k)
	for idx := 0; idx < lenRunes; idx++ {
		s := string(k[idx])
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

// longestPrefixNode returns the node that contains the longest prefix
// available in the trie of the argument k
// It provides a similar experience to referencing a key, value pair in a map
// where both the node value and a bool are returned if the key is found;
// otherwise, it returns an empty string and false.
func (trie *Trie) longestPrefixNode(k string) (n *Node, ok bool) {
	lenRunes := utf8.RuneCountInString(k)
	if lenRunes < 1 {
		return &Node{}, false
	}
	currentNode := trie.Root
	for idx := 0; idx < lenRunes; idx++ {
		s := string(k[idx])
		if _, ok := currentNode.Children[s]; ok {
			currentNode = currentNode.Children[s]
		} else if idx == 0 {
			return &Node{}, false
		} else {
			return currentNode, true
		}
	}
	return currentNode, true
}

// LongestPrefix returns the key, value pair of the longest prefix
// available in the trie of the argument s.
// If there is no path in the trie that contains a prefix of the
// argument s, then empty strings are returned for the key and value
func (trie *Trie) LongestPrefix(s string) (k string, v string) {
	if n, ok := trie.longestPrefixNode(s); !ok {
		return "", ""
	} else {
		return n.Key, n.Val
	}
}
