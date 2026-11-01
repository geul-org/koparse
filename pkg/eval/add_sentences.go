//ff:func feature=eval type=builder control=iteration dimension=1
//ff:what addSentences — 표본 하나의 문장 출력을 item.sent 열쇠로 넣는다
package eval

import (
	"fmt"

	"github.com/geul-org/koparse/pkg/analyze"
)

// addSentences stores sents under "id.1", "id.2", ….
func addSentences(outs map[string]analyze.SentenceOut, id int, sents []analyze.SentenceOut) {
	for i, s := range sents {
		outs[fmt.Sprintf("%d.%d", id, i+1)] = s
	}
}
