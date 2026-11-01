//ff:func feature=cli type=command control=sequence
//ff:what TestEvaluate — 지은 표본을 평가해 요약을 쓰고, 분석 오류는 종료 코드 1
package main

import (
	"context"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/geul-org/koparse/pkg/analyze"
	"github.com/geul-org/koparse/pkg/config"
	"github.com/geul-org/koparse/pkg/fixture"
	"github.com/geul-org/koparse/pkg/knowledge"
	"github.com/geul-org/koparse/pkg/lexicon"
)

func TestEvaluate(t *testing.T) {
	o, _ := parseOptions(writeTestFiles(t), io.Discard)
	o.out = ""
	in, _ := loadInputs(o)
	fx, _ := fixture.Load(o.fixture)
	lex, _ := lexicon.Default()
	var out strings.Builder
	if code := evaluate(analyze.New(fx, lex, config.Default()), o, in, fx.Misses, &out, io.Discard); code != 0 || !strings.Contains(out.String(), "분리 정확도: 3/3") {
		t.Fatal(code, out.String())
	}
	bad := analyze.New(knowledge.Funcs{StatementsFn: func(context.Context, string, []string) (map[string][]string, error) {
		return nil, errors.New("boom")
	}}, lex, config.Default())
	if code := evaluate(bad, o, in, fx.Misses, io.Discard, io.Discard); code != 1 {
		t.Fatal(code)
	}
}
