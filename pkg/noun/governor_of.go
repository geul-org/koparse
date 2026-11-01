//ff:func feature=noun type=helper control=iteration dimension=1
//ff:what governorOf — 관용 틀 검사용 지배어: 어절 자신의 서술격·용언화 접미, 없으면 다음 어절 표면
package noun

import "github.com/geul-org/koparse/pkg/lexicon"

// governorOf returns the eojeol's own copula/verbal suffix, else the next
// eojeol's surface, else "".
func governorOf(words []Word, last int) string {
	for _, s := range words[last].Chain {
		if s.Kind != lexicon.KindParticle {
			return s.Form
		}
	}
	if last+1 < len(words) {
		return words[last+1].Surface
	}
	return ""
}
