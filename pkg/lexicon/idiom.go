//ff:type feature=lexicon type=model
//ff:what Idiom — 관용 틀 한 항목(명사, 붙은 조사들, 지배 서술어 앞부분들)
package lexicon

// Idiom is one idiom frame: a noun (or "*"), the particles it may carry
// ("-" = none, "*" = any) and governor predicate prefixes.
type Idiom struct {
	Noun      string
	Particles []string
	Governors []string
}
