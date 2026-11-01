//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what writeTestFiles — 시험용: 지은 표본·정답 파일을 임시 디렉터리에 쓰고 인자 목록을 돌려준다
package main

import (
	"os"
	"path/filepath"
	"testing"
)

// writeTestFiles writes a made-up sample and gold files and returns the
// koparse-eval arguments pointing at them.
func writeTestFiles(t *testing.T) []string {
	dir := t.TempDir()
	files := map[string]string{
		"s.json":   `{"items":[{"id":1,"qid":"Q146","text":"고양이는 고기를 먹는다."}]}`,
		"gs.tsv":   "1.1\t0\t고양이+는\tnoun+particle\n1.1\t1\t고기+를\tnoun+particle\n1.1\t2\t먹는다\twhole-other\n",
		"gn.tsv":   "1.1\t고양이\tQ146\t\t채택\t\n",
		"j.tsv":    "1\t고기\tQ10990\t맞음\t\n",
		"cfg.json": `{"noun":{"adopt":0.5}}`,
	}
	for name, body := range files {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	p := func(n string) string { return filepath.Join(dir, n) }
	return []string{"-fixture", "../../testdata/fixture.json", "-samples", p("s.json"), "-gold-split", p("gs.tsv"),
		"-gold-noun", p("gn.tsv"), "-judgments", p("j.tsv"), "-config", p("cfg.json"), "-out", p("out")}
}
