//ff:func feature=cli type=command control=sequence
//ff:what report — 점수를 매겨 요약을 표준출력에 쓰고 -out이 있으면 상세 파일을 쓴다
package main

import (
	"fmt"
	"io"

	"github.com/geul-org/koparse/pkg/analyze"
	"github.com/geul-org/koparse/pkg/eval"
)

// report scores outs, prints the summary and writes details when o.out is set.
func report(o options, in inputs, outs map[string]analyze.SentenceOut, misses []string, stdout, stderr io.Writer) int {
	sp := eval.ScoreSplit(outs, in.goldSplit)
	ns := eval.ScoreNouns(outs, in.goldNoun, in.judgments)
	eval.WriteReport(stdout, sp, ns, len(misses))
	if o.out == "" {
		return 0
	}
	if err := writeDetails(o.out, sp, ns, outs, misses); err != nil {
		fmt.Fprintln(stderr, "koparse-eval:", err)
		return 1
	}
	return 0
}
