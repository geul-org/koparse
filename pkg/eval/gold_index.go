//ff:type feature=eval type=model
//ff:what goldIndex — 명사 정답 색인(문장 열쇠·명사별, 표본·명사별)과 추가 손 판정
package eval

// goldIndex looks up gold judgements by "item.sent\tnoun" and "item\tnoun".
type goldIndex struct {
	byKey  map[string]nounSets
	byItem map[string]nounSets
	extra  map[string]string
}
