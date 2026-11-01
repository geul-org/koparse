//ff:func feature=lexicon type=parser control=iteration dimension=1
//ff:what dataLines — 주석(#)·빈 줄을 뺀 줄들을 탭으로 나눈 필드로 돌려준다(줄 끝 # 주석 제거)
package lexicon

import "strings"

// dataLines splits s into lines, drops blank and comment lines and trailing
// "# ..." comments, and returns each line's tab-separated fields.
func dataLines(s string) [][]string {
	var out [][]string
	for _, line := range strings.Split(s, "\n") {
		if i := strings.Index(line, "#"); i >= 0 {
			line = line[:i]
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fields := strings.Split(line, "\t")
		for i := range fields {
			fields[i] = strings.TrimSpace(fields[i])
		}
		out = append(out, fields)
	}
	return out
}
