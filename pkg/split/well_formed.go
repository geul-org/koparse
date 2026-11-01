//ff:func feature=split type=helper control=sequence
//ff:what wellFormed — 후보의 사슬 순서와 받침 호응이 모두 맞는지
package split

// wellFormed reports whether c has a valid chain and fitting allomorphs.
func wellFormed(f *facts, c Candidate) bool {
	return f.lex.ChainValid(c.Suffixes) && batchimMismatchAt(c) == ""
}
