//ff:func feature=noun type=helper control=iteration dimension=1
//ff:what chainForms — 접미 사슬의 형태들
package noun

import "github.com/geul-org/koparse/pkg/lexicon"

// chainForms returns the forms of a suffix chain.
func chainForms(chain []lexicon.Suffix) []string {
	out := make([]string, 0, len(chain))
	for _, s := range chain {
		out = append(out, s.Form)
	}
	return out
}
