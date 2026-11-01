//ff:func feature=analyze type=builder control=sequence
//ff:what TestDocument — 지은 두 문장을 주제 고양이로: 고양이는 R2로 채택, 진술·문장 오류는 넘긴다
package analyze

import (
	"context"
	"testing"
)

func TestDocument(t *testing.T) {
	out, err := newTestAnalyzer(t, "").Document(context.Background(), "고양이는 고기를 먹는다. 생쥐는 작다.", "Q146")
	if err != nil || len(out) != 2 || out[0].Nouns[0].QID != "Q146" || len(out[0].Stats) != len(out[0].Nouns) {
		t.Fatalf("%+v %v", out, err)
	}
	if _, err := newTestAnalyzer(t, "statements").Document(context.Background(), "고양이", "Q146"); err == nil {
		t.Fatal("want statements error")
	}
	if _, err := newTestAnalyzer(t, "candidates").Document(context.Background(), "고양이", ""); err == nil {
		t.Fatal("want sentence error")
	}
}
