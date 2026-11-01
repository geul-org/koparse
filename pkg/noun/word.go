//ff:type feature=noun type=model
//ff:what Word — 분리가 끝난 어절(표면, 어간, 접미 사슬, 뒤 쉼표류)
package noun

import "github.com/geul-org/koparse/pkg/lexicon"

// Word is an eojeol after the split decision.
type Word struct {
	Surface string
	Stem    string
	Chain   []lexicon.Suffix
	Break   bool
}
