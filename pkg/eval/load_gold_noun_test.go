//ff:func feature=eval type=parser control=sequence
//ff:what TestLoadGoldNoun — 맞는·틀린 QID 목록을 읽고 읽기 오류는 넘긴다
package eval

import "testing"

func TestLoadGoldNoun(t *testing.T) {
	g, err := LoadGoldNoun(writeTemp(t, "n.tsv", "1.1\t고양이\tQ146\t\t채택\t메모\n1.1\t전통\t\tQ1,Q2\t버림\t\n"))
	if err != nil || len(g) != 2 || g[0].Good[0] != "Q146" || g[0].Bad != nil || len(g[1].Bad) != 2 {
		t.Fatal(g, err)
	}
	if _, err := LoadGoldNoun("/nonexistent"); err == nil {
		t.Fatal("want error")
	}
}
