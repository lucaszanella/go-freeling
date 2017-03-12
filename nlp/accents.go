package nlp

import (
	set "gopkg.in/fatih/set.v0"
)

type Accent struct {
	who AccentsModule
}

//Create the appropriate accents module (according to received options), and create a wrapper to access it.
func NewAccent(lang string) *Accent {
	this := Accent{}
	if lang == "es" {
		//Create spanish accent handler
		who := NewAccentsES()
		this.who = who
	} else {
		//Create Default (null) accent handler. Ok for English.
		who := NewAccetsDefault()
		this.who = who
	}
	return &this
}

//Wrapper methods: just call the wrapped accents module.
func (this *Accent) FixAccentutation(candidates *set.Set, suf *sufrule) {
	this.who.FixAccentuation(candidates, suf)
}

type AccentsModule interface {
	FixAccentuation(*set.Set, *sufrule)
}

type AccentsDefault struct {
}

func NewAccetsDefault() *AccentsDefault {
	LOG.Trace("Create default accent handler")
	return &AccentsDefault{}
}

func (this *AccentsDefault) FixAccentuation(candidates *set.Set, suf *sufrule) {
	LOG.Trace("Default accentuation. Candidates " + candidates.String())
}