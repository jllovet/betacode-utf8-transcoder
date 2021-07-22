package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jllovet/betacode-utf8-transcoder/transcoder"
	"github.com/jllovet/betacode-utf8-transcoder/trie"
)

func main() {
	f, err := os.OpenFile("transcoder.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	t := transcoder.InitBetaToUniTrie()
	// dfs(t.Root, "")

	beta := os.Args[1]
	t.LongestPrefix(beta)
	s, err := transcoder.BetaToUni(beta)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(beta, "becomes: ", s)

	// fmt.Printf("%+v\n", t.Root.Children)
}

func dfs(n *trie.Node, padding string) {
	for c := range n.Children {
		fmt.Printf("%s%+v\n", padding, n.Children[c])
		dfs(n.Children[c], padding+"  ")
	}
}
