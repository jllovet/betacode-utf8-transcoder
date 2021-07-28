package trie

import (
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name string
		want *Trie
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Init(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Init() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrie_Update(t *testing.T) {
	type fields struct {
		Root *Node
	}
	type args struct {
		k string
		v string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trie := &Trie{
				Root: tt.fields.Root,
			}
			trie.Update(tt.args.k, tt.args.v)
		})
	}
}

func TestTrie_Search(t *testing.T) {
	type fields struct {
		Root *Node
	}
	type args struct {
		k string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantS  string
		wantOk bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trie := &Trie{
				Root: tt.fields.Root,
			}
			gotS, gotOk := trie.Search(tt.args.k)
			if gotS != tt.wantS {
				t.Errorf("Trie.Search() gotS = %v, want %v", gotS, tt.wantS)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Trie.Search() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestTrie_longestPrefixNode(t *testing.T) {
	type fields struct {
		Root *Node
	}
	type args struct {
		k string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantN  *Node
		wantOk bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trie := &Trie{
				Root: tt.fields.Root,
			}
			gotN, gotOk := trie.longestPrefixNode(tt.args.k)
			if !reflect.DeepEqual(gotN, tt.wantN) {
				t.Errorf("Trie.longestPrefixNode() gotN = %v, want %v", gotN, tt.wantN)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Trie.longestPrefixNode() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestTrie_LongestPrefix(t *testing.T) {
	type fields struct {
		Root *Node
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantK  string
		wantV  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trie := &Trie{
				Root: tt.fields.Root,
			}
			gotK, gotV := trie.LongestPrefix(tt.args.s)
			if gotK != tt.wantK {
				t.Errorf("Trie.LongestPrefix() gotK = %v, want %v", gotK, tt.wantK)
			}
			if gotV != tt.wantV {
				t.Errorf("Trie.LongestPrefix() gotV = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}
