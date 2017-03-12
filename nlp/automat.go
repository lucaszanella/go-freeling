package nlp

import (
	set "gopkg.in/fatih/set.v0"
)

const AUTOMAT_MAX_STATES = 100
const AUTOMAT_MAX_TOKENS = 50

type AutomatStatus struct {
	ShiftBegin int
}

type Automat struct {
	InitialState int
	StopState    int
	Trans        [AUTOMAT_MAX_STATES][AUTOMAT_MAX_TOKENS]int
	Final        *set.Set
}
