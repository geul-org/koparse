//ff:func feature=analyze type=builder control=sequence
//ff:what TestAssembleAll — 평평한 결과를 문장별 단위 수대로 나눈다
package analyze

import (
	"testing"

	"github.com/geul-org/koparse/pkg/noun"
)

func TestAssembleAll(t *testing.T) {
	a := newTestAnalyzer(t, "")
	works := []sentenceWork{{units: []noun.Unit{{Form: "가"}}}, {units: []noun.Unit{{Form: "나"}, {Form: "다"}}}}
	out := a.assembleAll(works, []noun.Result{{Form: "가"}, {Form: "나"}, {Form: "다"}})
	if len(out) != 2 || len(out[0].Nouns) != 1 || len(out[1].Nouns) != 2 || out[1].Nouns[1].Form != "다" {
		t.Fatalf("%+v", out)
	}
}
