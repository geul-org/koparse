//ff:func feature=stats type=helper control=selection
//ff:what Bucket — 부호 있는 토큰 거리를 구간으로(±1, ±2, ±3~4, ±5~8, ±9~16, 같은 문장)
package stats

// Bucket maps a signed token distance to its interval label.
func Bucket(d int) string {
	s := "+"
	if d < 0 {
		s, d = "-", -d
	}
	switch {
	case d <= 1:
		return s + "1"
	case d == 2:
		return s + "2"
	case d <= 4:
		return s + "3~4"
	case d <= 8:
		return s + "5~8"
	case d <= 16:
		return s + "9~16"
	default:
		return "같은 문장"
	}
}
