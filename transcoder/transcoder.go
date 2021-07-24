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
	t := []string{} // transformed text
	var pwb bool    // possible word boundary
	n := 0          // index
	for n < utf8.RuneCountInString(beta) {
		// change medial to final sigma if followed by punctuation/whitespace
		if pwb && isFinalSigma(strings.Join(t, "")) {
			t[len(t)-2] = finalLowerCaseSigma
		}
		// look up betacode character in converstion trie
		k, v := BetaToUniTrie.LongestPrefix(beta[n:findFinalIndex(beta, n)])
		if err != nil {
			return "", err
		}
		if k != "" {
			// check whether word ends in punctuation/whitespace
			pwb = betaPunctuation[string(beta[n])]
			t = append(t, v) // append unicode value found in trie
			n += len(k)      // move index ahead to next betacode character
		} else {
			pwb = true
			// no match found in conversion trie, keep current value
			t = append(t, string(beta[n]))
			n += 1
		}
	}

	// Check last charcter in slice of transformed strings.
	// If it is a medial lowercase sigma `σ`,
	// then it should be changed into a final-position lowercase sigma `ς`.
	if len(t) > 0 && t[len(t)-1] == medialLowerCaseSigma {
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

// isFinalSigma returns true if the penultimate character is a
// sigma that should be transformed into a final sigma.
// Generally, a sigma in the last position should be converted into
// a the final sigma form 'ς'. However, if the last letter is followed
// by an apostrophe, as in contractions, then the sigma should remain in
// its medial form 'σ'.
func isFinalSigma(text string) (p bool) {
	lenRunes := utf8.RuneCountInString(text) // length in runes

	var lastLetterIsMedialSigma bool  // Last letter is medial lower-case sigma
	var lastCharIsLetter bool         // Last character is a letter
	var lastCharIsBetaApostrophe bool // Last character is a beta apostrophe
	n := 1                            // index counting in runes
	for _, r := range text {
		if n == lenRunes-1 {
			if string(r) == medialLowerCaseSigma {
				lastLetterIsMedialSigma = true
			}
		}
		if n == lenRunes {
			if unicode.IsLetter(r) {
				lastCharIsLetter = true
			}
			if string(r) == betaApostrophe {
				lastCharIsBetaApostrophe = true
			}
		}
		n += 1
	}
	p = lenRunes > 1 &&
		lastLetterIsMedialSigma &&
		!lastCharIsLetter &&
		!lastCharIsBetaApostrophe

	return p
}
