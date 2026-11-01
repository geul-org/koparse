//ff:func feature=cli type=command control=sequence
//ff:what TestOutputLines — 열쇠 순 JSON 줄
package main

import (
	"strings"
	"testing"

	"github.com/geul-org/koparse/pkg/analyze"
)

func TestOutputLines(t *testing.T) {
	got := outputLines(map[string]analyze.SentenceOut{"1.2": {Text: "나"}, "1.1": {Text: "가"}})
	lines := strings.Split(strings.TrimSpace(got), "\n")
	if len(lines) != 2 || !strings.Contains(lines[0], `"key":"1.1"`) {
		t.Fatal(got)
	}
}
