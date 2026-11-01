//ff:func feature=cli type=command control=sequence
//ff:what parseOptions — 평가 명령 인자를 해석한다(fixture·samples·gold-split·gold-noun 필수)
package main

import (
	"errors"
	"flag"
	"io"
)

// parseOptions parses args.
func parseOptions(args []string, stderr io.Writer) (options, error) {
	var o options
	fs := flag.NewFlagSet("koparse-eval", flag.ContinueOnError)
	fs.SetOutput(stderr)
	fs.StringVar(&o.fixture, "fixture", "", "fixture JSON")
	fs.StringVar(&o.config, "config", "", "설정 JSON (기본: 내장)")
	fs.StringVar(&o.samples, "samples", "", "표본 JSON ({items:[{id,qid,text}]}, 저장소 밖)")
	fs.StringVar(&o.goldSplit, "gold-split", "", "분리 정답 TSV")
	fs.StringVar(&o.goldNoun, "gold-noun", "", "명사 정답 TSV")
	fs.StringVar(&o.judgments, "judgments", "", "추가 손 판정 TSV (선택)")
	fs.StringVar(&o.out, "out", "", "상세 결과 디렉터리 (선택)")
	if err := fs.Parse(args); err != nil {
		return o, err
	}
	if o.fixture == "" || o.samples == "" || o.goldSplit == "" || o.goldNoun == "" {
		return o, errors.New("-fixture, -samples, -gold-split and -gold-noun are required")
	}
	return o, nil
}
