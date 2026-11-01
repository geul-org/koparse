//ff:func feature=split type=builder control=sequence
//ff:what tryExtend — 접미 하나가 끝에 맞으면 후보로 넣고 더 앞으로 늘린다(서술격·용언화는 맨 끝에만)
package split

import (
	"strings"

	"github.com/geul-org/koparse/pkg/lexicon"
)

// tryExtend adds the candidate rest = stem + s + chain when s ends rest,
// then keeps extending the stem. Copula and verbal suffixes (slot 9) may only
// be the last element of a chain.
func tryExtend(lex *lexicon.Lexicon, rest string, s lexicon.Suffix, chain []lexicon.Suffix, depth int, out []Candidate, seen map[string]bool) []Candidate {
	stem, ok := strings.CutSuffix(rest, s.Form)
	if !ok || stem == "" || (len(chain) > 0 && s.Slot == 9) {
		return out
	}
	c := Candidate{Stem: stem, Suffixes: append([]lexicon.Suffix{s}, chain...)}
	if seen[c.Key()] {
		return out
	}
	seen[c.Key()] = true
	out = append(out, c)
	return extend(lex, stem, c.Suffixes, depth-1, out, seen)
}
