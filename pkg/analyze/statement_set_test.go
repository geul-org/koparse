//ff:func feature=analyze type=helper control=sequence
//ff:what TestStatementSet — 주제 없으면 빈 집합, 주제의 진술 값 모음, 오류는 넘긴다
package analyze

import (
	"context"
	"testing"
)

func TestStatementSet(t *testing.T) {
	a := newTestAnalyzer(t, "")
	if s, _ := a.statementSet(context.Background(), ""); len(s) != 0 {
		t.Fatal(s)
	}
	s, err := a.statementSet(context.Background(), "Q146")
	if err != nil || len(s) == 0 {
		t.Fatal(s, err)
	}
	if _, err := newTestAnalyzer(t, "statements").statementSet(context.Background(), "Q146"); err == nil {
		t.Fatal("want error")
	}
}
