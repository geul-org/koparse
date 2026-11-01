//ff:type feature=eval type=model
//ff:what NounScore — 명사 식별 집계(상태별 수, 판정별 수, 재현, 규칙별 판정, 채택 목록)
package eval

// NounScore aggregates the noun identification metrics.
type NounScore struct {
	Nouns         int               `json:"nouns"`
	Adopted       int               `json:"adopted"`
	Dropped       int               `json:"dropped"`
	Undecided     int               `json:"undecided"`
	DroppedReason map[string]int    `json:"dropped_reason"`
	Judged        map[string]int    `json:"judged"`  // Correct/Wrong/Unjudged → count
	ByRule        map[string][3]int `json:"by_rule"` // rule → [correct, wrong, unjudged]
	GoldGood      int               `json:"gold_good"`
	Recalled      int               `json:"recalled"`
	Rows          []string          `json:"rows"` // key, form, qid, judgement, rules
}
