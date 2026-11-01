//ff:func feature=fixture type=adapter control=iteration dimension=1
//ff:what Misses — 표에 없어 "후보 없음"으로 답한 표면(lang:표면)을 정렬해 돌려준다
package fixture

import "sort"

// Misses returns the "lang:surface" keys that were asked but are not in the table.
func (f *Fixture) Misses() []string {
	f.mu.Lock()
	defer f.mu.Unlock()
	out := make([]string, 0, len(f.misses))
	for k := range f.misses {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
