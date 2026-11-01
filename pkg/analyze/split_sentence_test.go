//ff:func feature=analyze type=builder control=sequence
//ff:what TestSplitSentence — 지은 문장 "고양이는 고기를 먹는다": 고양이+는, 고기+를, 어댑터 오류는 넘긴다
package analyze

import (
	"context"
	"testing"

	"github.com/geul-org/koparse/pkg/text"
)

func TestSplitSentence(t *testing.T) {
	s := text.Parse("고양이는 고기를 먹는다.")[0]
	w, err := newTestAnalyzer(t, "").splitSentence(context.Background(), s)
	if err != nil || len(w.words) != 3 || w.words[0].Stem != "고양이" || w.words[1].Stem != "고기" || len(w.decisions) != 3 {
		t.Fatal(w, err)
	}
	if _, err := newTestAnalyzer(t, "candidates").splitSentence(context.Background(), s); err == nil {
		t.Fatal("want error")
	}
}
