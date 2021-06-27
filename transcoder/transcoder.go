package transcoder

import (
	"bufio"
	"io"
	"log"
	"strings"
	"unicode"
	"unicode/utf8"
	// "github.com/golang-collections/collections/trie"
	// "golang.org/x/text/unicode/norm"
)

// UniToBeta converts a unicode string to a betacode string
func UniToBeta(uni string) (beta string, err error) {
	u := UNICODE_MAP

	transform := []string{}
	b := bufio.NewReader(strings.NewReader(uni))
	for {
		if r, _, err := b.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
				return "", err
			}
		} else {
			// check whether betacode value in map
			if betaValue, ok := u[string(r)]; ok {
				// found in unicode map, replace with betacode value
				transform = append(transform, betaValue)
			} else {
				// not found in unicode map, keep value
				transform = append(transform, string(r))
			}
		}
	}
	beta = strings.Join(transform, "")
	return beta, err
}

// BetaToUni converts a betacode string to a unicode string
func BetaToUni(beta string) (uni string, err error) {
	uni = string(BETACODE_MAP[beta])
	return uni, err
}

// findLongestBetaTokenLen returns the maximum length of a single betacode token
func findLongestBetaTokenLen(BETACODE_MAP map[string]string) (maxBetaTokenLen int) {
	maxBetaTokenLen = -1
	for beta := range BETACODE_MAP {
		if utf8.RuneCountInString(beta) > maxBetaTokenLen {
			maxBetaTokenLen = utf8.RuneCountInString(beta)
		}
	}
	return
}

var maxBetaTokenLen int = findLongestBetaTokenLen(BETACODE_MAP)

// Special characters that need their own references to rewrite with
var finalLowerCaseSigma string = `ς`  // `s2`, \u03c2
var medialLowerCaseSigma string = `σ` // `s`,  \u03c3
var betaApostrophe string = `’`       // `\'` \u2019
var betaPunctuation []string = []string{
	`·`, // `:`,  \u00b7
	`’`, // `\'`, \u2019
	`‐`, // `-`,  \u2010
	`—`, // `_`,  \u2014
}

// penultimateSigmaWordFinal returns whether the penultimate character is a sigma
// and is in a context where it should be treated as a medial sigma
func penultimateSigmaWordFinal(text string) (p bool) {
	l := utf8.RuneCountInString(text) // length in runes
	p = utf8.RuneCountInString(text) > 1 &&
		string(text[l-2]) == medialLowerCaseSigma &&
		!unicode.IsLetter(rune(text[l-1])) &&
		string(text[l-1]) != betaApostrophe
	return p
}
