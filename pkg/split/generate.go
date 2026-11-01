//ff:func feature=split type=builder control=sequence
//ff:what Generate — 어절 끝에서 조사·어미를 긴 것부터 맞춰 가능한 분리를 모두 후보로 만든다(통째 후보가 0번)
package split

import "github.com/geul-org/koparse/pkg/lexicon"

// maxChain is the longest suffix chain tried (선수+들+에게+는 needs 3).
const maxChain = 3

// Generate lists every split of surface into a non-empty stem and a chain of
// up to maxChain listed suffixes. Index 0 is always the whole candidate.
func Generate(lex *lexicon.Lexicon, surface string) []Candidate {
	out := []Candidate{{Stem: surface}}
	return extend(lex, surface, nil, maxChain, out, map[string]bool{})
}
