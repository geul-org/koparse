//ff:func feature=text type=parser control=iteration dimension=1
//ff:what StripAsides — 괄호 곁가지를 본문에서 떼어 기록한다(중첩 괄호 포함, 닫히지 않은 괄호는 본문에 남김)
package text

// StripAsides removes top-level (...) and [...] groups from s and records
// them with the word before. A space before a group is dropped, so
// "φῶς (...)와" becomes "φῶς와".
func StripAsides(s string) (string, []Aside) {
	rs := []rune(s)
	var b []rune
	var asides []Aside
	depth, start := 0, 0
	for i, r := range rs {
		nd := parenDepth(depth, r)
		if depth == 0 && nd == 1 {
			start = i
		}
		if depth == 0 && nd == 0 {
			b = append(b, r)
		}
		if depth == 1 && nd == 0 {
			asides = append(asides, Aside{Head: headBefore(b), Text: string(rs[start+1 : i])})
			b = dropTrailingSpace(b)
		}
		depth = nd
	}
	if depth > 0 {
		b = append(b, rs[start:]...)
	}
	return string(b), asides
}
