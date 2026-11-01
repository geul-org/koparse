//ff:func feature=cli type=command control=sequence
//ff:what analyzeTo — 입력을 문서 하나로 분석해 문장마다 JSON 한 줄을 쓴다
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/geul-org/koparse/pkg/analyze"
)

// analyzeTo analyzes input and writes the JSON lines; returns the exit code.
func analyzeTo(a *analyze.Analyzer, input, topic string, stdout, stderr io.Writer) int {
	out, err := a.Document(context.Background(), input, topic)
	if err != nil {
		fmt.Fprintln(stderr, "koparse:", err)
		return 1
	}
	enc := json.NewEncoder(stdout)
	enc.SetEscapeHTML(false)
	return writeLines(enc, out, stderr)
}
