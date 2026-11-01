//ff:func feature=noun type=judge control=iteration dimension=1
//ff:what rankScores — 합친 점수 1위와 2위 후보 번호(2위가 없으면 −1)
package noun

// rankScores returns the indexes of the best and second-best candidates by
// support 1 − Miss (ties keep the earlier candidate).
func rankScores(scores []candScore) (int, int) {
	best, second := -1, -1
	for i, s := range scores {
		switch {
		case best < 0 || s.Miss < scores[best].Miss:
			best, second = i, best
		case second < 0 || s.Miss < scores[second].Miss:
			second = i
		}
	}
	return best, second
}
