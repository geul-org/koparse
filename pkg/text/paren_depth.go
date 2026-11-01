//ff:func feature=text type=helper control=selection
//ff:what parenDepth — 괄호((, [) 깊이를 한 글자만큼 갱신한다(음수는 0)
package text

// parenDepth returns the bracket depth after rune r.
func parenDepth(d int, r rune) int {
	switch r {
	case '(', '[':
		return d + 1
	case ')', ']':
		return max(d-1, 0)
	default:
		return d
	}
}
