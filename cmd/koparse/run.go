//ff:func feature=cli type=command control=sequence
//ff:what run — 인자를 해석해 처리기를 만들고 입력(파일 또는 표준입력)을 읽어 분석을 맡긴다
package main

import (
	"flag"
	"fmt"
	"io"
)

// run parses args, builds the analyzer, reads the input and analyzes it as
// one document. It returns the process exit code.
func run(args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("koparse", flag.ContinueOnError)
	fs.SetOutput(stderr)
	fixturePath := fs.String("fixture", "", "fixture JSON (지식 어댑터; 필수, SILK 어댑터는 M3)")
	topic := fs.String("topic", "", "문서 주제 QID (R2·R3)")
	configPath := fs.String("config", "", "설정 JSON (기본: 내장 default.json)")
	trace := fs.Bool("trace", false, "어절마다 분리 후보 순위를 붙인다")
	in := fs.String("in", "", "입력 텍스트 파일 (기본: 표준입력)")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	a, err := newAnalyzer(*fixturePath, *configPath)
	if err != nil {
		fmt.Fprintln(stderr, "koparse:", err)
		return 1
	}
	a.Trace = *trace
	input, err := readInput(*in, stdin)
	if err != nil {
		fmt.Fprintln(stderr, "koparse:", err)
		return 1
	}
	return analyzeTo(a, input, *topic, stdout, stderr)
}
