package main

import (
	"fmt"

	"github.com/jllovet/betacode-utf8-transcoder/transcoder"
)

func main() {
	cp := string([]rune(transcoder.BETACODE_MAP["s2"]))
	fmt.Println(cp)
}
