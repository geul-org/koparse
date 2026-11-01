//ff:func feature=cli type=command control=iteration dimension=1
//ff:what writeDetails — 상세 결과 파일(분리 틀림, 채택 명사 판정, fixture 놓침, 문장 출력 JSON 줄)을 쓴다
package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/geul-org/koparse/pkg/analyze"
	"github.com/geul-org/koparse/pkg/eval"
)

// writeDetails writes split-mismatch.tsv, nouns.tsv, misses.txt and outputs.jsonl into dir.
func writeDetails(dir string, sp eval.SplitScore, ns eval.NounScore, outs map[string]analyze.SentenceOut, misses []string) error {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	files := map[string]string{
		"split-mismatch.tsv": "# item.sent\t정답\t출력\t범주\n" + strings.Join(sp.Mismatches, "\n") + "\n",
		"nouns.tsv":          "# item.sent\t명사\tQID\t판정\t규칙\n" + strings.Join(ns.Rows, "\n") + "\n",
		"misses.txt":         strings.Join(misses, "\n") + "\n",
		"outputs.jsonl":      outputLines(outs),
	}
	for name, body := range files {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(body), 0o644); err != nil {
			return err
		}
	}
	return nil
}
