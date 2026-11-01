//ff:func feature=text type=parser control=sequence
//ff:what makeEojeol — 조각에서 앞뒤 문장부호를 떼어 어절을 만든다(안쪽 따옴표는 지우고, 앞뒤가 모두 따옴표면 인용 표시, 쉼표류 기록)
package text

import "strings"

const quoteRunes = "\"'“”‘’「」『』《》〈〉«»"

// makeEojeol strips edge punctuation from piece and drops inner quotation
// marks ("빛"을 → 빛을). Quoted is set only when the token is wrapped in
// quotes on both sides. ok is false when nothing is left.
func makeEojeol(piece string) (Eojeol, bool) {
	core := strings.TrimFunc(piece, isTrimRune)
	i := strings.Index(piece, core)
	lead, trail := piece[:i], piece[i+len(core):]
	if core == "" {
		lead, trail = "", piece
	}
	e := Eojeol{
		Surface: strings.Map(dropQuote, core),
		Quoted:  strings.ContainsAny(lead, quoteRunes) && strings.ContainsAny(trail, quoteRunes),
		Break:   strings.ContainsAny(trail, ",;:"),
	}
	return e, core != ""
}
