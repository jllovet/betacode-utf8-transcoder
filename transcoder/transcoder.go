package transcoder

import (
	"bufio"
	"errors"
	"io"
	"log"
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
				log.Fatal(err)
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

	// TODO: Implement non-strict parsing of diacritics: Reference python:
	// t = pygtrie.CharTrie()

	// for beta, uni in _map.BETACODE_MAP.items():
	//     if strict:
	//         t[beta] = uni
	//     else:
	//         # The order of accents is very strict and weak. Allow for many orders of
	//         # accents between asterisk and letter or after letter. This does not
	//         # introduce ambiguity since each betacode token only has one letter and
	//         # either starts with a asterisk or a letter.
	//         diacritics = beta[1:]

	//         perms = itertools.permutations(diacritics)
	//         for perm in perms:
	//             perm_str = beta[0] + ''.join(perm)
	//             t[perm_str.lower()] = uni
	//             t[perm_str.upper()] = uni

	// return t
}

var BetaToUniTrie trie.Trie = InitBetaToUniTrie()

// BetaToUni converts a betacode string to a unicode string
func BetaToUni(beta string) (uni string, err error) {
	uni, ok := BetaToUniTrie.Search(beta)
	if !ok {
		return "", errors.New("could not parse betacode")
	}
	return uni, nil
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
