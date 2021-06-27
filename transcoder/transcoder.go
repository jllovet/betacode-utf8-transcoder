package transcoder

func UniToBeta(uni string) (beta string) {
	return uni
}

func BetaToUni(beta string) (uni string) {
	uni = string(BETACODE_MAP[beta])
	return uni
}
