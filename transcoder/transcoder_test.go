package transcoder_test

import (
	"testing"

	"github.com/jllovet/betacode-utf8-transcoder/transcoder"
	"golang.org/x/text/unicode/norm"
)

// TestUniToBeta tests that the result of converting unicode is betacode.
// Comparison is done via the NFKC normalization for unicode.
// norm.NFKC.String
func TestUniToBeta(t *testing.T) {
	t.Parallel()
	tables := []struct {
		uni         string
		beta        string
		description string
	}{
		{``, ``, "Empty"},
		{`α`, `a`, "Single Simple Letter"},
		{`αβ`, `ab`, "Simple Conversion"},
		{`βίον τέχνης καὶ εὐδαιμονίας.`, `bi/on te/xnhs kai\ eu)daimoni/as.`, "Multi-Word"},
		{`Ἔφορος καὶ ἄλλοι`, `*)/eforos kai\ a)/lloi`, "Multiple Accents"},
		{`πλείους: ἔτι δὲ οἱ μετὰ`, `plei/ous: e)/ti de\ oi( meta\\`, "Colon Punctuation"},
		{`Let's see if we can convert this: ἔτι δὲ οἱ`, `Let's see if we can convert this: e)/ti de\ oi(`, "Mixed Conversion / Multi-Language"},
	}
	for _, table := range tables {
		u2b := transcoder.UniToBeta(table.uni) // Unicode converted to Betacode

		ucbn := norm.NFKC.String(u2b)      // Unicode converted to Betacode and Normalized
		bn := norm.NFKC.String(table.beta) // Betacode Normalized

		if ucbn != bn {
			t.Errorf(
				"Case %s failed.\nConversion of %q to %q was incorrect.\nGot: %q\nWant: %q.",
				table.description, table.uni, table.beta, ucbn, bn,
			)
		} else {
			t.Logf("Case %s passed", table.description)
		}
	}
}

// TestBetaToUni tests that the result of converting betacode is unicode.
// Comparison is done via the NFKC normalization for unicode.
// norm.NFKC.String
func TestBetaToUni(t *testing.T) {
	t.Parallel()
	tables := []struct {
		beta        string
		uni         string
		description string
	}{
		{``, ``, "Empty"},
		{`t`, `τ`, "Single Simple Letter"},
		{`tou=`, `τοῦ`, "Simple Conversion"},
		{`th=s`, `τῆς`, "Final Sigma"},
		{`th=s2`, `τῆς`, "Numeric Sigma Id"},
		{`th=s3 tou=`, `τῆϲ τοῦ`, "Keep Non-Final Sigma Numeric"},
		{`th=s tou=`, `τῆς τοῦ`, "Final Sigma Word"},
		{`th=s\ttou=`, `τῆς\tτοῦ`, "Final Sigma Whitespace"},
		{`th=s; tou=`, `τῆς; τοῦ`, "Final Sigma Punctuation"},
		{`th=s\' tou=`, `τῆσ’ τοῦ`, "Final Sigma Apostrophe"},
		{`analabo/ntes de\ kaq\' e(/kaston`, `αναλαβόντες δὲ καθ’ ἕκαστον`, "Multi Word"},
		{`e)/oiken h)\ dida/skonti; nh\\`, `ἔοικεν ἢ διδάσκοντι; νὴ`, "Punctuation Semicolon"},
		{`dh=lon: oi(/ te`, `δῆλον· οἵ τε`, "Punctuation Colon"},
		{`e/)oiken h\) dida/skonti; nh\\ a=|)i+\\`, `ἔοικεν ἢ διδάσκοντι; νὴ ᾆῒ`, "Out of Order"},
		{`*)/eforos ka*)/ei\ a/)lloi`, `Ἔφορος καἜὶ ἄλλοι`, "Capital Out of Order"},
		{`*)/eforos ka*)/ei\ a/)lloi *)h\|`, `Ἔφορος καἜὶ ἄλλοι ᾛ`, "Capital Out of Order with Iota"},
		{`e)n d\' e)\pes\' w)keanw=|`, `ἐν δ’ ἒπεσ’ ὠκεανῷ`, "Strict Correct"},
		{`e)n d\' e)\pes\' w)keanw|=`, `ἐν δ’ ἒπεσ’ ὠκεανῳ=`, "Strict Incorrect"},
		{`e)n d\' e)\pes\' w)keanw|=`, `ἐν δ’ ἒπεσ’ ὠκεανῷ`, "Unstrict"},
		{`*)e/foros ka*e)/i\ a/)lloi *)\h|`, `Ἔφορος καἜὶ ἄλλοι ᾛ`, "Unstrict Capitalization"},
	}
	for _, table := range tables {
		b2u := transcoder.BetaToUni(table.beta) // Betacode converted to Unicode
		bcun := norm.NFKC.String(b2u)           // Betacode converted to Unicode and Normalized
		un := norm.NFKC.String(table.uni)       // Unicode Normalized

		if bcun != un {
			t.Errorf(
				"Case %s failed.\nConversion of %q to %q was incorrect.\nGot: %q\nWant: %q.",
				table.description, table.beta, table.uni, bcun, un,
			)
		} else {
			t.Logf("Case %s passed", table.description)
		}
	}
}
