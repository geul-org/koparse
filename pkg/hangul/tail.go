//ff:type feature=hangul type=model
//ff:what Tail — 앞말 끝소리 종류(모음, ㄹ 받침, 그 밖의 받침, 알 수 없음)
package hangul

// Tail is the kind of sound that ends the preceding text, used to check
// particle allomorphs (이/가, 을/를, 으로/로 ...).
type Tail int

const (
	// TailUnknown means the ending sound cannot be determined (Latin, symbols).
	TailUnknown Tail = iota
	// TailVowel means no final consonant (받침 없음).
	TailVowel
	// TailRieul means the final consonant is ㄹ.
	TailRieul
	// TailConsonant means any final consonant other than ㄹ.
	TailConsonant
)
