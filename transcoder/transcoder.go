package transcoder

import (
	"golang.org/x/text/unicode/norm"
)

func UniToBeta(uni string) (beta string) {
	beta = string(norm.NFC.String(UNICODE_MAP[uni]))
	return beta
}

func BetaToUni(beta string) (uni string) {
	uni = string(BETACODE_MAP[beta])
	return uni
}
