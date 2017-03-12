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
	return &AccentsPT{
		withAcc:    make(map[string]string),
		withoutAcc: make(map[string]string),
	}
}

func (this *AccentsPT) FixAccentuation(candidates *set.Set, suf *sufrule) {
	LOG.Trace("ES accentuation. Candidates " + candidates.String())
}