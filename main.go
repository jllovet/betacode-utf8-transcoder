package main

import (
	"fmt"

	"github.com/jllovet/betacode-utf8-transcoder/transcoder"
	"github.com/jllovet/betacode-utf8-transcoder/trie"
)

func main() {
	// fmt.Println(transcoder.FindLongestBetaTokenLen(transcoder.BETACODE_MAP))
	// text := "apple pie"
	// l := utf8.RuneCountInString(text) // length in runes
	// penultimateRune := string(text[l-2])
	// fmt.Println(penultimateRune)
	// fmt.Println(string(transcoder.BetaToUni("s2")))
	// fmt.Println(string(transcoder.UniToBeta("α")))
	t := transcoder.InitBetaToUniTrie(false)
	// dfs(t.Root, "")

	beta := `e)/`
	s, ok := t.Search(beta)
	if ok {
		fmt.Println(beta, "becomes: ", s)
	}

	// fmt.Printf("%+v\n", t.Root.Children)
}

func dfs(n *trie.Node, padding string) {
	for c := range n.Children {
		fmt.Printf("%s%+v\n", padding, n.Children[c])
		dfs(n.Children[c], padding+"  ")
	}
}
