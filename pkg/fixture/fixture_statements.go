//ff:func feature=fixture type=adapter control=iteration dimension=1
//ff:what Statements — 개체 진술 값(속성 목록이 nil이면 전부). 진술이 없는 개체는 빈 맵
package fixture

import "context"

// Statements returns the statement values of qid for props (nil = all).
func (f *Fixture) Statements(_ context.Context, qid string, props []string) (map[string][]string, error) {
	all := f.statements[qid]
	out := map[string][]string{}
	if props == nil {
		for p, v := range all {
			out[p] = v
		}
		return out, nil
	}
	for _, p := range props {
		if v, ok := all[p]; ok {
			out[p] = v
		}
	}
	return out, nil
}
