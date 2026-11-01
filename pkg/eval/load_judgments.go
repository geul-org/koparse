//ff:func feature=eval type=parser control=iteration dimension=1
//ff:what LoadJudgments — 추가 손 판정 TSV(열쇠 item.sent 또는 item, 명사, QID, 맞음|틀림)를 읽는다(경로가 비면 빈 표)
package eval

import "fmt"

// LoadJudgments reads extra manual judgements keyed "key\tnoun\tqid", where
// key is "item.sent" or "item".
func LoadJudgments(path string) (map[string]string, error) {
	out := map[string]string{}
	if path == "" {
		return out, nil
	}
	rows, err := readTSV(path, 4)
	if err != nil {
		return nil, err
	}
	for _, f := range rows {
		if f[3] != Correct && f[3] != Wrong {
			return nil, fmt.Errorf("%s: judgement %q is not %s or %s", path, f[3], Correct, Wrong)
		}
		out[f[0]+"\t"+f[1]+"\t"+f[2]] = f[3]
	}
	return out, nil
}
