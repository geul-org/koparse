//ff:func feature=eval type=judge control=iteration dimension=1
//ff:what TestGoldIndexJudge — 손 판정(문장·표본) → 문장 정답 → 표본 정답 → 미판정 순서
package eval

import "testing"

func TestGoldIndexJudge(t *testing.T) {
	gold := []GoldNoun{{Key: "1.1", Noun: "물", Good: []string{"Q1"}, Bad: []string{"Q2"}}, {Key: "1.2", Noun: "불", Good: []string{"Q3"}, Bad: []string{"Q4"}}}
	extra := map[string]string{"1.1\t물\tQ9": Wrong, "1\t물\tQ8": Correct}
	ix := newGoldIndex(gold, extra)
	cases := []struct{ key, form, qid, want string }{
		{"1.1", "물", "Q9", Wrong}, {"1.1", "물", "Q8", Correct}, {"1.1", "물", "Q1", Correct}, {"1.1", "물", "Q2", Wrong},
		{"1.5", "불", "Q3", Correct}, {"1.5", "불", "Q4", Wrong}, {"1.5", "불", "Q5", Unjudged},
	}
	for _, c := range cases {
		if got := ix.judge(c.key, c.form, c.qid); got != c.want {
			t.Errorf("%v: %s", c, got)
		}
	}
}
