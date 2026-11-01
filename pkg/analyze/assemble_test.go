//ff:func feature=analyze type=builder control=sequence
//ff:what TestAssemble — 명사 없어도 빈 목록, Trace면 후보 순위를 붙인다
package analyze

import (
	"context"
	"testing"

	"github.com/geul-org/koparse/pkg/text"
)

func TestAssemble(t *testing.T) {
	a := newTestAnalyzer(t, "")
	w, _ := a.splitSentence(context.Background(), text.Parse("먹는다.")[0])
	out := a.assemble(w, nil)
	if out.Nouns == nil || len(out.Eojeol) != 1 || out.Eojeol[0].Ranked != nil || out.Text != "먹는다." {
		t.Fatalf("%+v", out)
	}
	a.Trace = true
	if out := a.assemble(w, nil); len(out.Eojeol[0].Ranked) == 0 {
		t.Fatal("trace missing")
	}
}
