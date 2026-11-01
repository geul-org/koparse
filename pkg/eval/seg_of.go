//ff:func feature=eval type=helper control=iteration dimension=1
//ff:what segOf — 출력 어절의 분리를 정답 형식(어간+접미, 분리 없으면 어간)으로
package eval

import "github.com/geul-org/koparse/pkg/analyze"

// segOf renders an output eojeol as "stem+suffixes" (or "stem").
func segOf(e analyze.EojeolOut) string {
	rest := ""
	for _, p := range e.Split[1:] {
		rest += p.Form
	}
	if rest == "" {
		return e.Split[0].Form
	}
	return e.Split[0].Form + "+" + rest
}
