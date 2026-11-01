//ff:func feature=lexicon type=parser control=iteration dimension=1
//ff:what parseSuffixes — 접미 목록과 @allow 사슬 예외를 읽고 긴 형태부터 정렬한다
package lexicon

import (
	"fmt"
	"sort"
	"unicode/utf8"
)

// parseSuffixes parses suffixes.tsv. Suffixes come back longest form first.
func parseSuffixes(s string) ([]Suffix, map[string]bool, error) {
	var list []Suffix
	allow := map[string]bool{}
	for _, f := range dataLines(s) {
		if f[0] == "@allow" && len(f) == 3 {
			allow[f[1]+"\t"+f[2]] = true
			continue
		}
		suf, err := parseSuffixLine(f)
		if err != nil {
			return nil, nil, fmt.Errorf("%v: %w", f, err)
		}
		list = append(list, suf)
	}
	sort.SliceStable(list, func(i, j int) bool {
		return utf8.RuneCountInString(list[i].Form) > utf8.RuneCountInString(list[j].Form)
	})
	return list, allow, nil
}
