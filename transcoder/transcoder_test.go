package transcoder_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/jllovet/betacode-utf8-transcoder/transcoder"
	"golang.org/x/text/unicode/norm"
)

// TestUniToBeta tests that the result of converting unicode is betacode.
// Comparison is done via the NFC normalization for unicode.
// norm.NFC.String
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
		{`πλείους: ἔτι δὲ οἱ μετὰ`, `plei/ous: e)/ti de\ oi( meta\`, "Colon Punctuation"},
		{`Let's see if we can convert this: ἔτι δὲ οἱ`, `Let's see if we can convert this: e)/ti de\ oi(`, "Mixed Conversion / Multi-Language"},
	}
	for _, table := range tables {
		u2b, err := transcoder.UniToBeta(table.uni) // Unicode converted to Betacode
		if err != nil {
			t.Errorf("Error during test %s: %s", table.description, err)
		}
		ucbn := norm.NFC.String(u2b)      // Unicode converted to Betacode and Normalized
		bn := norm.NFC.String(table.beta) // Betacode Normalized

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
// Comparison is done via the NFC normalization for unicode.
// norm.NFC.String
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
		{`th=s	tou=`, "τῆς\tτοῦ", "Final Sigma Whitespace"},
		{`th=s; tou=`, `τῆς; τοῦ`, "Final Sigma Punctuation"},
		{`th=s' tou=`, `τῆσ’ τοῦ`, "Final Sigma Apostrophe"},
		{`analabo/ntes de\ kaq' e(/kaston`, `αναλαβόντες δὲ καθ’ ἕκαστον`, "Multi Word"},
		{`e)/oiken h)\ dida/skonti; nh\`, `ἔοικεν ἢ διδάσκοντι; νὴ`, "Punctuation Semicolon"},
		{`dh=lon: oi(/ te`, `δῆλον· οἵ τε`, "Punctuation Colon"},
		{`e/)oiken h\) dida/skonti; nh\ a=|)i+\`, `ἔοικεν ἢ διδάσκοντι; νὴ ᾆῒ`, "Out of Order"},
		{`*)/eforos ka*)/ei\ a/)lloi`, `Ἔφορος καἜὶ ἄλλοι`, "Capital Out of Order"},
		{`*)/eforos ka*)/ei\ a/)lloi *)h\|`, `Ἔφορος καἜὶ ἄλλοι ᾛ`, "Capital Out of Order with Iota"},
		{`e)n d' e)\pes' w)keanw=|`, `ἐν δ’ ἒπεσ’ ὠκεανῷ`, "Strict Correct"},
		{`e)n d' e)\pes' w)keanw|=`, `ἐν δ’ ἒπεσ’ ὠκεανῷ`, "Unstrict"},
		{`*)e/foros ka*e)/i\ a/)lloi *)\h|`, `Ἔφορος καἜὶ ἄλλοι ᾛ`, "Unstrict Capitalization"},
	}
	for _, table := range tables {
		b2u, err := transcoder.BetaToUni(table.beta) // Betacode converted to Unicode
		if err != nil {
			t.Errorf("Error during test %s: %s", table.description, err)
		}
		bcun := norm.NFC.String(b2u)     // Betacode converted to Unicode and Normalized
		un := norm.NFC.String(table.uni) // Unicode Normalized

		if bcun != un {
			t.Errorf(
				"Case %s failed.\nConversion of \"%s\" to \"%s\" was incorrect.\nGot: \"%s\"\nWant: \"%s\".",
				table.description, table.beta, table.uni, bcun, un,
			)
		} else {
			t.Logf("Case %s passed", table.description)
		}
	}
}

func ExampleUniToBeta() {
	u := `ἄλφα`
	b, err := transcoder.UniToBeta(u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u, "becomes", b)
	// Output: ἄλφα becomes a)/lfa
}

func ExampleBetaToUni() {
	b := `a)/lfa`
	u, err := transcoder.BetaToUni(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b, "becomes", u)
	// Output: a)/lfa becomes ἄλφα
}
