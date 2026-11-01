//ff:type feature=eval type=model
//ff:what GoldNoun — 명사 정답 한 줄(문장 열쇠, 명사, 맞는 QID, 틀린 QID)
package eval

// GoldNoun is the gold judgement of one noun occurrence.
type GoldNoun struct {
	Key  string
	Noun string
	Good []string
	Bad  []string
}
