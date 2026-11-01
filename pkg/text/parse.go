//ff:func feature=text type=parser control=iteration dimension=1
//ff:what Parse — 정규화 → 문장 분리 → 괄호 곁가지 떼기 → 어절 분할
package text

// Parse turns raw text into sentences of eojeols.
func Parse(input string) []Sentence {
	var out []Sentence
	for _, s := range SplitSentences(Normalize(input)) {
		body, asides := StripAsides(s)
		out = append(out, Sentence{Text: s, Body: body, Eojeols: Tokenize(body), Asides: asides})
	}
	return out
}
