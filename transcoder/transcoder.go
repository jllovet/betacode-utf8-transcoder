package transcoder

import (
	"bufio"
	"io"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/jllovet/betacode-utf8-transcoder/trie"
	// "golang.org/x/text/unicode/norm"
)

// UniToBeta converts a unicode string to a betacode string
func UniToBeta(uni string) (beta string, err error) {
	u := UnicodeMap

	t := []string{} // transformed characters in string
	b := bufio.NewReader(strings.NewReader(uni))
	for {
		if r, _, err := b.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				return "", err
			}
		} else {
			// check whether betacode value in map
			if b, ok := u[string(r)]; ok {
				// found in unicode map, transform value
				t = append(t, b)
			} else {
				// not found in unicode map, keep value
				t = append(t, string(r))
			}
		}
	}
	beta = strings.Join(t, "")
	return beta, err
}

// InitBetaToUniTrie creates a trie for converting from betacode to unicode.
// It takes a bool value to indicate whether the betacode parsing is strict.
// If true, diacritics must be in the officially specified order:
// breathing, accent, iota subscript / dieresis
// If false, diacritic parsing is more generous, allowing any order after the
// initial letter or asterisk (which still must be in the first position).
func InitBetaToUniTrie() (t trie.Trie) {
	t = *trie.Init()
	for beta, uni := range BetacodeMap {
		t.Update(beta, uni)
	}
	return t
}

var BetaToUniTrie trie.Trie = InitBetaToUniTrie()

// BetaToUni converts a betacode string to a unicode string
func BetaToUni(beta string) (uni string, err error) {
	t := []string{}
	// var i int
	var pwb bool // possible word boundary

	i := 0
	for i < utf8.RuneCountInString(beta) {
		if pwb && penultimateSigmaWordFinal(strings.Join(t, "")) {
			t[len(t)-2] = finalLowerCaseSigma
		}
		k, v := BetaToUniTrie.LongestPrefix(beta[i:findFinalIndex(beta, i)])
		if err != nil {
			return "", err
		}
		if k != "" {
			pwb = betaPunctuation[string(beta[i])] // checks for presence of current character in set
			t = append(t, v)
			i += len(k)
		} else {
			pwb = true
			t = append(t, string(beta[i]))
			i += 1
		}
	}

	// Check one last time in case there is some whitespace or punctuation at the
	// end and check if the last character is a sigma.
	if pwb && penultimateSigmaWordFinal(strings.Join(t, "")) {
		t[len(t)-2] = finalLowerCaseSigma
	} else if len(t) > 0 && t[len(t)-1] == medialLowerCaseSigma {
		t[len(t)-1] = finalLowerCaseSigma
	}
	uni = strings.Join(t, "")
	return uni, nil
}

func findFinalIndex(beta string, idx int) int {
	var finalIndex int
	if idx+maxBetaTokenLen > utf8.RuneCountInString(beta) {
		finalIndex = utf8.RuneCountInString(beta)
	} else {
		finalIndex = idx + maxBetaTokenLen
	}
	return finalIndex
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

var maxBetaTokenLen int = findLongestBetaTokenLen(BetacodeMap)

// Special characters that need their own references to rewrite with
var finalLowerCaseSigma string = `ς`  // `s2`, \u03c2
var medialLowerCaseSigma string = `σ` // `s`,  \u03c3
var betaApostrophe string = `’`       // `\'` \u2019
var betaPunctuation = map[string]bool{
	`:`: true, // `·`,  \u00b7
	`'`: true, // `’`, \u2019
	`-`: true, // `‐`,  \u2010
	`_`: true, // `—`,  \u2014
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
