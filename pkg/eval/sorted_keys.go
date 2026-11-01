//ff:func feature=eval type=helper control=iteration dimension=1
//ff:what sortedKeys — item.sent 열쇠를 표본·문장 번호 순으로
package eval

import (
	"fmt"
	"sort"

	"github.com/geul-org/koparse/pkg/analyze"
)

// sortedKeys returns the keys of outs ordered by item then sentence number.
func sortedKeys(outs map[string]analyze.SentenceOut) []string {
	keys := make([]string, 0, len(outs))
	for k := range outs {
		keys = append(keys, k)
	}
	num := func(k string) [2]int {
		var a, b int
		fmt.Sscanf(k, "%d.%d", &a, &b)
		return [2]int{a, b}
	}
	sort.Slice(keys, func(i, j int) bool {
		x, y := num(keys[i]), num(keys[j])
		return x[0] < y[0] || (x[0] == y[0] && x[1] < y[1])
	})
	return keys
}
