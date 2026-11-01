//ff:func feature=noun type=helper control=iteration dimension=1
//ff:what runEnd — a에서 시작하는 맨 어절 묶음의 끝 번호(접미가 붙거나 쉼표가 오면 끝)
package noun

// runEnd returns the last index of the run starting at a: it extends over
// bare eojeols not followed by a clause break.
func runEnd(words []Word, a int) int {
	b := a
	for b < len(words)-1 && len(words[b].Chain) == 0 && !words[b].Break {
		b++
	}
	return b
}
