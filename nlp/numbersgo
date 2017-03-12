package nlp

type NumbersStatus struct {
	AutomatStatus
    /// partial value of partially build number expression
    bilion,milion,units int64
    block int
    iscode int

    // These are used only in NUMBERS_PT_it. !! unify process with other languages !! 
    hundreds int64   // this is additional.
    thousands int64  // this is additional.
    floatUnits int64 // "e tre quarto". Count of how many "halfs", "quartrs" we have
}

func NewNumbersStatus() *NumbersStatus {
	return &NumbersStatus{

	}
}