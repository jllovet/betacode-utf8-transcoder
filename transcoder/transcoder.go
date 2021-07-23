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

// UniToBeta converts a unicode string to a betacode string.
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

// initBetaToUniTrie creates a trie for converting from betacode to unicode.
// It takes a bool value to indicate whether the betacode parsing is strict.
// If true, diacritics must be in the officially specified order:
// breathing, accent, iota subscript / dieresis.
// If false, diacritic parsing is more generous, allowing any order after the
// initial letter or asterisk (which still must be in the first position).
func initBetaToUniTrie() (t trie.Trie) {
	t = *trie.Init()
	for beta, uni := range BetacodeMap {
		t.Update(beta, uni)
	}
	return t
}

var BetaToUniTrie trie.Trie = initBetaToUniTrie()

// BetaToUni converts a betacode string to a unicode string.
func BetaToUni(beta string) (uni string, err error) {
	t := []string{}
	var pwb bool // possible word boundary

	i := 0
	for i < utf8.RuneCountInString(beta) {
		if pwb && convertFinalSigma(strings.Join(t, "")) {
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
	if pwb && convertFinalSigma(strings.Join(t, "")) {
		t[len(t)-2] = finalLowerCaseSigma
	} else if len(t) > 0 && t[len(t)-1] == medialLowerCaseSigma {
		t[len(t)-1] = finalLowerCaseSigma
	}
	uni = strings.Join(t, "")
	return uni, nil
}

// findFinalIndex determines the final index of the string that can be examined
// when taking a slice from a string by indices. This is a helper function to
// avoid going out of bounds when reading in beta code characters from a string.
func findFinalIndex(beta string, idx int) int {
	var finalIndex int
	if idx+maxBetaTokenLen > utf8.RuneCountInString(beta) {
		finalIndex = utf8.RuneCountInString(beta)
	} else {
		finalIndex = idx + maxBetaTokenLen
	}
	return finalIndex
}

// findLongestBetaTokenLen returns the maximum length of a single betacode token.
func findLongestBetaTokenLen(betacodeMap map[string]string) (maxBetaTokenLen int) {
	maxBetaTokenLen = -1
	for beta := range betacodeMap {
		if utf8.RuneCountInString(beta) > maxBetaTokenLen {
			maxBetaTokenLen = utf8.RuneCountInString(beta)
		}
	}
	return
}

var maxBetaTokenLen int = findLongestBetaTokenLen(BetacodeMap)

// Special characters that need their own references to rewrite with
var finalLowerCaseSigma string = `ς`  // `ς`, `s2`, \u03c2
var medialLowerCaseSigma string = `σ` // `σ`, `s`,  \u03c3
var betaApostrophe string = `’`       // `’`, `'` \u2019
var betaPunctuation = map[string]bool{
	`:`: true, // `·`, `:`,  \u00b7
	`'`: true, // `’`, `'`,  \u2019
	`-`: true, // `‐`, `-`,  \u2010
	`_`: true, // `—`, `_`,  \u2014
	` `: true,
}

// convertFinalSigma returns true if the penultimate character is a
// sigma that should be transformed into a final sigma.
// Generally, a sigma in the last position should be converted into
// a the final sigma form 'ς'. However, if the last letter is followed
// by an apostrophe, as in contractions, then the sigma should remain in
// its medial form 'σ'.
func convertFinalSigma(text string) (p bool) {
	l := utf8.RuneCountInString(text) // length in runes

	var ms bool // Last letter is medial lower-case sigma
	var ll bool // Last character is a letter
	var ba bool // Last character is a beta apostrophe
	n := 1      // index counting in runes
	for _, r := range text {
		if n == l-1 {
			if string(r) == medialLowerCaseSigma {
				ms = true // Last letter is medial lower-case sigma
			}
		}
		if n == l {
			if unicode.IsLetter(r) {
				ll = true // Last character is a letter
			}
			if string(r) == betaApostrophe {
				ba = true // Last character is a beta apostrophe
			}
		}
		n += 1
	}
	p = l > 1 &&
		ms &&
		!ll &&
		!ba

	return p
}
