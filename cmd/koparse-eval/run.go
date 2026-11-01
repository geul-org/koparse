//ff:func feature=cli type=command control=sequence
//ff:what run — 인자 해석, 평가 파일·fixture·설정을 읽어 처리기를 만들고 평가를 맡긴다
package main

import (
	"fmt"
	"io"

	"github.com/geul-org/koparse/pkg/analyze"
	"github.com/geul-org/koparse/pkg/config"
	"github.com/geul-org/koparse/pkg/fixture"
	"github.com/geul-org/koparse/pkg/lexicon"
)

// run evaluates and returns the exit code.
func run(args []string, stdout, stderr io.Writer) int {
	o, err := parseOptions(args, stderr)
	if err != nil {
		fmt.Fprintln(stderr, "koparse-eval:", err)
		return 2
	}
	in, err := loadInputs(o)
	if err != nil {
		fmt.Fprintln(stderr, "koparse-eval:", err)
		return 1
	}
	fx, err := fixture.Load(o.fixture)
	if err != nil {
		fmt.Fprintln(stderr, "koparse-eval:", err)
		return 1
	}
	lex, _ := lexicon.Default() // embedded lists, checked by lexicon.TestDefault
	cfg := config.Default()
	if o.config != "" {
		cfg, err = config.Load(o.config)
	}
	if err != nil {
		fmt.Fprintln(stderr, "koparse-eval:", err)
		return 1
	}
	return evaluate(analyze.New(fx, lex, cfg), o, in, fx.Misses, stdout, stderr)
}
