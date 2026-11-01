//ff:func feature=split type=helper control=sequence
//ff:what candOf — 규칙에 붙은 후보 번호
package split

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// candOf returns the candidate index of the rule's candSpec.
func candOf(specs toulmin.Specs) int {
	return specs.Find("cand").(candSpec).Idx
}
