//ff:func feature=split type=rule control=sequence
//ff:what chainInvalid — 반박: 허용되지 않는 조사 순서(은+들)면 공격
package split

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// chainInvalid rebuts a split whose suffix order breaks the slot grammar.
func chainInvalid(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	f := factsOf(ctx)
	if f.lex.ChainValid(f.cands[candOf(specs)].Suffixes) {
		return -1, nil
	}
	return 1, nil
}
