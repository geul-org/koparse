//ff:type feature=lexicon type=model
//ff:what Lexicon — 자체 목록 묶음(접미 목록, 사슬 예외, 성질·틀 명사, 관용 틀, 분류)
package lexicon

// Lexicon holds koparse's own lists.
type Lexicon struct {
	// Suffixes are ordered longest form first.
	Suffixes []Suffix
	// Allow lists slot-order exceptions as "앞\t뒤".
	Allow    map[string]bool
	Property map[string]bool
	Idioms   []Idiom
	Proper   map[string]bool
	Internal map[string]bool
}
