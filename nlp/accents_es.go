package nlp

import (
	set "gopkg.in/fatih/set.v0"
	"regexp"
)

type AccentsES struct {
	llanaAcc        *regexp.Regexp
	agudaMal        *regexp.Regexp
	monosil         *regexp.Regexp
	lastVowelPutAcc *regexp.Regexp
	lastVowelNotAcc *regexp.Regexp
	anyVowelAcc     *regexp.Regexp

	withAcc    map[string]string
	withoutAcc map[string]string
}

func NewAccentsES() *AccentsES {
	return &AccentsES{
		withAcc:    make(map[string]string),
		withoutAcc: make(map[string]string),
	}
}

func (this *AccentsES) FixAccentuation(candidates *set.Set, suf *sufrule) {
	LOG.Trace("ES accentuation. Candidates " + candidates.String())
}
