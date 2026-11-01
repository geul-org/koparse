//ff:func feature=eval type=parser control=iteration dimension=1
//ff:what LoadGoldSplit — 분리 정답 TSV(item.sent, 순번, 분리, 범주)를 문장별로 읽는다
package eval

// LoadGoldSplit reads the gold split file keyed by "item.sent".
func LoadGoldSplit(path string) (map[string][]GoldSplit, error) {
	rows, err := readTSV(path, 4)
	if err != nil {
		return nil, err
	}
	out := map[string][]GoldSplit{}
	for _, f := range rows {
		out[f[0]] = append(out[f[0]], GoldSplit{Key: f[0], Seg: f[2], Cat: f[3]})
	}
	return out, nil
}
