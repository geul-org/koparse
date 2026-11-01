//ff:func feature=eval type=builder control=sequence
//ff:what TestAddSentences — id.1부터 열쇠를 붙인다
package eval

import (
	"testing"

	"github.com/geul-org/koparse/pkg/analyze"
)

func TestAddSentences(t *testing.T) {
	outs := map[string]analyze.SentenceOut{}
	addSentences(outs, 3, []analyze.SentenceOut{{Text: "가"}, {Text: "나"}})
	if outs["3.1"].Text != "가" || outs["3.2"].Text != "나" {
		t.Fatal(outs)
	}
}
