package main

import (
	"fmt"

	"github.com/jllovet/betacode-utf8-transcoder/transcoder"
)

func main() {
	fmt.Println(string(transcoder.BetaToUni("s2")))
	fmt.Println(string(transcoder.UniToBeta("α")))
}
