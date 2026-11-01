//ff:func feature=eval type=builder control=sequence
//ff:what TestRunSamples — 표본·문장 번호 열쇠로 모으고 분석 오류는 표본 번호와 함께 넘긴다
package eval

import (
	"context"
	"testing"
)

func TestRunSamples(t *testing.T) {
	s := []Sample{{ID: 7, QID: "Q146", Text: "고양이는 작다. 고기를 먹는다."}}
	outs, err := RunSamples(context.Background(), newTestAnalyzer(t, false), s)
	if err != nil || len(outs) != 2 || outs["7.2"].Text != "고기를 먹는다." {
		t.Fatal(outs, err)
	}
	if _, err := RunSamples(context.Background(), newTestAnalyzer(t, true), s); err == nil {
		t.Fatal("want error")
	}
}
