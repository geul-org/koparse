//ff:func feature=cli type=command control=sequence
//ff:what evaluate — 표본을 분석하고 점수·요약·상세 결과를 낸다
package main

import (
	"context"
	"fmt"
	"io"

	"github.com/geul-org/koparse/pkg/analyze"
	"github.com/geul-org/koparse/pkg/eval"
)

// evaluate runs the samples through a and reports; misses lists the
// adapter's unanswered surfaces after the run.
func evaluate(a *analyze.Analyzer, o options, in inputs, misses func() []string, stdout, stderr io.Writer) int {
	outs, err := eval.RunSamples(context.Background(), a, in.samples)
	if err != nil {
		fmt.Fprintln(stderr, "koparse-eval:", err)
		return 1
	}
	return report(o, in, outs, misses(), stdout, stderr)
}
