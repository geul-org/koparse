//ff:func feature=analyze type=builder control=sequence
//ff:what TestSentenceUnits — 명사 단위를 찾고, 분리·개체 조회 오류는 넘긴다
package analyze

import (
	"context"
	"testing"

	"github.com/geul-org/koparse/pkg/text"
)

func TestSentenceUnits(t *testing.T) {
	s := text.Parse("고양이는 고기를 먹는다.")[0]
	w, err := newTestAnalyzer(t, "").sentenceUnits(context.Background(), s)
	if err != nil || len(w.units) != 2 || w.units[0].Form != "고양이" {
		t.Fatal(w.units, err)
	}
	if _, err := newTestAnalyzer(t, "candidates").sentenceUnits(context.Background(), s); err == nil {
		t.Fatal("want split error")
	}
	if _, err := newTestAnalyzer(t, "entity").sentenceUnits(context.Background(), s); err == nil {
		t.Fatal("want unit error")
	}
}
