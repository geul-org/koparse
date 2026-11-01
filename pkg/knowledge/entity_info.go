//ff:type feature=knowledge type=model
//ff:what EntityInfo — 개체 요약(언어별 라벨·설명, P31·P279 일부)
package knowledge

// EntityInfo is a summary of one entity.
type EntityInfo struct {
	QID          string            `json:"qid"`
	Labels       map[string]string `json:"labels,omitempty"`
	Descriptions map[string]string `json:"descriptions,omitempty"`
	P31          []string          `json:"p31,omitempty"`
	P279         []string          `json:"p279,omitempty"`
}
