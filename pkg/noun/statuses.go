package noun

// Result statuses.
const (
	StatusAdopted   = "adopted"
	StatusDropped   = "dropped"
	StatusUndecided = "undecided"
)

// Reasons of dropped and undecided results.
const (
	ReasonNoCandidate = "후보 없음"
	ReasonProperty    = "성질"
	ReasonIdiom       = "관용"
	ReasonBelow       = "문턱 미달"
	ReasonAmbiguous   = "애매"
)
