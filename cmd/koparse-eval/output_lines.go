//ff:func feature=cli type=command control=iteration dimension=1
//ff:what outputLines — 문장 출력을 열쇠 순으로 JSON 줄로({"key":…, "out":…})
package main

import (
	"encoding/json"
	"sort"
	"strings"

	"github.com/geul-org/koparse/pkg/analyze"
)

// outputLines renders every output sentence as a JSON line, sorted by key.
func outputLines(outs map[string]analyze.SentenceOut) string {
	keys := make([]string, 0, len(outs))
	for k := range outs {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var b strings.Builder
	for _, k := range keys {
		line, _ := json.Marshal(map[string]any{"key": k, "out": outs[k]})
		b.Write(line)
		b.WriteByte('\n')
	}
	return b.String()
}
