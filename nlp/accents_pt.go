package nlp

import (
	set "gopkg.in/fatih/set.v0"
	"regexp"
)

type AccentsPT struct {
	llanaAcc        *regexp.Regexp
	agudaMal        *regexp.Regexp
	monosil         *regexp.Regexp
	lastVowelPutAcc *regexp.Regexp
	lastVowelNotAcc *regexp.Regexp
	anyVowelAcc     *regexp.Regexp

	withAcc    map[string]string
	withoutAcc map[string]string
}

func NewAccentsPT() *AccentsPT {
	withoutAcc := make(map[string]string)
	withAcc := make(map[string]string)
	withoutAcc["á"] = "a"
	withoutAcc["é"] = "e"
	withoutAcc["í"] = "i"
	withoutAcc["ó"] = "o"
	withoutAcc["ú"] = "u"
	withAcc["a"] = "á"
	withAcc["e"] = "é"
	withAcc["i"] = "í"
	withAcc["o"] = "ó"
	withAcc["u"] = "ú"
	return &AccentsPT{
		withAcc:    withAcc,
		withoutAcc: withoutAcc,
	}
}

func (this *AccentsPT) FixAccentuation(candidates *set.Set, suf *sufrule) {
	LOG.Trace("PT accentuation. Candidates " + candidates.String())
}