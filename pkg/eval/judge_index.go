//ff:func feature=eval type=helper control=selection
//ff:what judgeIndex — 판정의 집계 칸 번호(맞음 0, 틀림 1, 미판정 2)
package eval

// judgeIndex maps a judgement to its [correct, wrong, unjudged] slot.
func judgeIndex(j string) int {
	switch j {
	case Correct:
		return 0
	case Wrong:
		return 1
	default:
		return 2
	}
}
