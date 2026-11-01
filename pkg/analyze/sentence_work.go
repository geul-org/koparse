//ff:type feature=analyze type=model
//ff:what sentenceWork — 문장 하나의 중간 결과(문장, 분리 결정, 분리된 어절, 명사 단위)
package analyze

import (
	"github.com/geul-org/koparse/pkg/noun"
	"github.com/geul-org/koparse/pkg/split"
	"github.com/geul-org/koparse/pkg/text"
)

// sentenceWork carries one sentence through the pipeline.
type sentenceWork struct {
	sent      text.Sentence
	decisions []split.Decision
	words     []noun.Word
	units     []noun.Unit
}
