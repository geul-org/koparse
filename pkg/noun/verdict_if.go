//ff:func feature=noun type=helper control=sequence
//ff:what verdictIf — 조건이 참이면 +1(적용), 아니면 −1(적용 안 됨)
package noun

// verdictIf returns +1 when cond holds, −1 (does not apply) otherwise.
func verdictIf(cond bool) float64 {
	if cond {
		return 1
	}
	return -1
}
