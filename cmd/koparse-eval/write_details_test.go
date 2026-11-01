//ff:func feature=cli type=command control=sequence
//ff:what TestWriteDetails — 네 파일을 쓰고, 디렉터리를 못 만들거나 파일을 못 쓰면 오류
package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/koparse/pkg/analyze"
	"github.com/geul-org/koparse/pkg/eval"
)

func TestWriteDetails(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "d")
	err := writeDetails(dir, eval.SplitScore{Mismatches: []string{"x"}}, eval.NounScore{Rows: []string{"y"}}, map[string]analyze.SentenceOut{"1.1": {Text: "가"}}, []string{"ko:없음"})
	b, _ := os.ReadFile(filepath.Join(dir, "misses.txt"))
	if err != nil || string(b) != "ko:없음\n" {
		t.Fatal(err, string(b))
	}
	blocker := filepath.Join(t.TempDir(), "file")
	os.WriteFile(blocker, nil, 0o644)
	if err := writeDetails(filepath.Join(blocker, "sub"), eval.SplitScore{}, eval.NounScore{}, nil, nil); err == nil {
		t.Fatal("want mkdir error")
	}
	ro := filepath.Join(t.TempDir(), "ro")
	os.MkdirAll(filepath.Join(ro, "nouns.tsv"), 0o755)
	if err := writeDetails(ro, eval.SplitScore{}, eval.NounScore{}, nil, nil); err == nil {
		t.Fatal("want write error")
	}
}
