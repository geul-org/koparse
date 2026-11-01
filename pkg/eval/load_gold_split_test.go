//ff:func feature=eval type=parser control=sequence
//ff:what TestLoadGoldSplit — 문장 열쇠별로 모으고 읽기 오류는 넘긴다
package eval

import "testing"

func TestLoadGoldSplit(t *testing.T) {
	g, err := LoadGoldSplit(writeTemp(t, "g.tsv", "1.1\t0\t고양이+는\tnoun+particle\n1.1\t1\t작다\twhole-other\n"))
	if err != nil || len(g["1.1"]) != 2 || g["1.1"][0].Seg != "고양이+는" || g["1.1"][1].Cat != "whole-other" {
		t.Fatal(g, err)
	}
	if _, err := LoadGoldSplit("/nonexistent"); err == nil {
		t.Fatal("want error")
	}
}
