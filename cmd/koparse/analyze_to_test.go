//ff:func feature=cli type=command control=sequence
//ff:what TestAnalyzeTo — 지은 문장을 JSON 줄로(HTML 문자 그대로), 분석 오류는 종료 코드 1
package main

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/geul-org/koparse/pkg/analyze"
	"github.com/geul-org/koparse/pkg/config"
	"github.com/geul-org/koparse/pkg/knowledge"
	"github.com/geul-org/koparse/pkg/lexicon"
)

func TestAnalyzeTo(t *testing.T) {
	a, _ := newAnalyzer("../../testdata/fixture.json", "")
	var out, errb strings.Builder
	if code := analyzeTo(a, "고양이는 <작다>.", "Q146", &out, &errb); code != 0 || !strings.Contains(out.String(), `"qid":"Q146"`) || !strings.Contains(out.String(), "<작다>") {
		t.Fatal(code, out.String(), errb.String())
	}
	lex, _ := lexicon.Default()
	bad := analyze.New(knowledge.Funcs{StatementsFn: func(context.Context, string, []string) (map[string][]string, error) {
		return nil, errors.New("boom")
	}}, lex, config.Default())
	if code := analyzeTo(bad, "가", "Q1", &out, &errb); code != 1 {
		t.Fatal(code)
	}
}
