//ff:func feature=eval type=helper control=sequence
//ff:what TestJudgeIndex — 맞음 0, 틀림 1, 그 밖 2
package eval

import "testing"

func TestJudgeIndex(t *testing.T) {
	if judgeIndex(Correct) != 0 || judgeIndex(Wrong) != 1 || judgeIndex(Unjudged) != 2 {
		t.Fatal("judgeIndex")
	}
}
