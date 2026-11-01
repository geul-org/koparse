//ff:func feature=split type=builder control=iteration dimension=1
//ff:what extend — 남은 앞부분 끝에 맞는 접미마다 후보를 늘린다(깊이 제한)
package split

import "github.com/geul-org/koparse/pkg/lexicon"

// extend tries every suffix at the end of rest, prepending it to chain.
func extend(lex *lexicon.Lexicon, rest string, chain []lexicon.Suffix, depth int, out []Candidate, seen map[string]bool) []Candidate {
	if depth == 0 {
		return out
	}
	for _, s := range lex.Suffixes {
		out = tryExtend(lex, rest, s, chain, depth, out, seen)
	}
	return out
}
