//ff:func feature=lexicon type=rule control=selection
//ff:what Accepts — 앞말 끝소리가 이 접미의 받침 조건에 맞는지(알 수 없으면 맞다고 본다)
package lexicon

import "github.com/geul-org/koparse/pkg/hangul"

// Accepts reports whether the suffix may follow text ending in sound t.
// An unknown ending sound (Latin, symbols) is accepted.
func (s Suffix) Accepts(t hangul.Tail) bool {
	if t == hangul.TailUnknown {
		return true
	}
	switch s.Cond {
	case "C":
		return t != hangul.TailVowel
	case "V":
		return t == hangul.TailVowel
	case "VL":
		return t != hangul.TailConsonant
	case "CX":
		return t == hangul.TailConsonant
	default:
		return true
	}
}
