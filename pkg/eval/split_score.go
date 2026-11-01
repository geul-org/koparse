//ff:type feature=eval type=model
//ff:what SplitScore — 분리 정확도 집계(전체, 맞음, 범주별, 미정, 정렬 실패, 틀린 목록)
package eval

// SplitScore aggregates the split accuracy.
type SplitScore struct {
	Total            int               `json:"total"`
	Correct          int               `json:"correct"`
	ByCat            map[string][2]int `json:"by_cat"` // category → [correct, total]
	Undecided        int               `json:"undecided"`
	UndecidedCorrect int               `json:"undecided_correct"`
	Unaligned        int               `json:"unaligned"`
	Mismatches       []string          `json:"mismatches"`
}
