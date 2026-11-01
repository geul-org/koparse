//ff:func feature=eval type=builder control=iteration dimension=1
//ff:what newGoldIndex — 명사 정답과 추가 손 판정으로 색인을 만든다
package eval

// newGoldIndex indexes the gold nouns and extra judgements.
func newGoldIndex(gold []GoldNoun, extra map[string]string) goldIndex {
	ix := goldIndex{byKey: map[string]nounSets{}, byItem: map[string]nounSets{}, extra: extra}
	for _, g := range gold {
		addSets(ix.byKey, g.Key+"\t"+g.Noun, g)
		addSets(ix.byItem, itemOf(g.Key)+"\t"+g.Noun, g)
	}
	return ix
}
