//ff:func feature=eval type=parser control=sequence
//ff:what TestLoadJudgments — 빈 경로는 빈 표, 맞음·틀림만 허용, 읽기 오류는 넘긴다
package eval

import "testing"

func TestLoadJudgments(t *testing.T) {
	if j, err := LoadJudgments(""); err != nil || len(j) != 0 {
		t.Fatal(j, err)
	}
	j, err := LoadJudgments(writeTemp(t, "j.tsv", "1\t물\tQ283\t맞음\t메모\n"))
	if err != nil || j["1\t물\tQ283"] != Correct {
		t.Fatal(j, err)
	}
	if _, err := LoadJudgments(writeTemp(t, "b.tsv", "1\t물\tQ283\t글쎄\n")); err == nil {
		t.Fatal("want value error")
	}
	if _, err := LoadJudgments("/nonexistent"); err == nil {
		t.Fatal("want read error")
	}
}
