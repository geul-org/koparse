//ff:func feature=split type=helper control=sequence
//ff:what knownVerdict — 사전에 있는 어간의 판정: 후보가 많을수록 조금 낮춘다(1 − decay·ln n, 하한 min)
package split

import "math"

// knownVerdict returns 1 − decay·ln(n) clamped to [lo, 1] for n ≥ 1 candidates.
func knownVerdict(n int, decay, lo float64) float64 {
	v := 1 - decay*math.Log(float64(n))
	return math.Min(1, math.Max(lo, v))
}
