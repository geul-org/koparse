//ff:func feature=eval type=helper control=iteration dimension=1
//ff:what sortedMapKeys — 맵 열쇠를 정렬해 돌려준다
package eval

import "sort"

// sortedMapKeys returns the sorted keys of m.
func sortedMapKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
