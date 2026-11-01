//ff:func feature=eval type=parser control=iteration dimension=1
//ff:what LoadGoldNoun — 명사 정답 TSV(item.sent, 명사, 맞는 QID, 틀린 QID, …)를 읽는다
package eval

// LoadGoldNoun reads the gold noun file.
func LoadGoldNoun(path string) ([]GoldNoun, error) {
	rows, err := readTSV(path, 4)
	if err != nil {
		return nil, err
	}
	out := make([]GoldNoun, 0, len(rows))
	for _, f := range rows {
		out = append(out, GoldNoun{Key: f[0], Noun: f[1], Good: splitList(f[2]), Bad: splitList(f[3])})
	}
	return out, nil
}
