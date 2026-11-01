//ff:type feature=stats type=model
//ff:what Pair — 거리 쌍(다른 토큰, 품사, 부호 있는 거리 구간, 같은 어절, 절 경계 넘음)
package stats

// Pair relates a noun to another noun, particle or predicate of the sentence.
// The exact distance is not stored; it can be recomputed from the text.
type Pair struct {
	Other       string `json:"other"`
	Tag         string `json:"tag"`
	Bucket      string `json:"bucket"`
	SameEojeol  bool   `json:"same_eojeol"`
	CrossClause bool   `json:"cross_clause"`
}
