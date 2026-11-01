//ff:func feature=cli type=command control=sequence
//ff:what TestReadInput — 파일 경로면 파일, 아니면 표준입력, 없는 파일은 오류
package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestReadInput(t *testing.T) {
	p := filepath.Join(t.TempDir(), "in.txt")
	os.WriteFile(p, []byte("파일"), 0o644)
	if s, err := readInput(p, nil); err != nil || s != "파일" {
		t.Fatal(s, err)
	}
	if s, err := readInput("", strings.NewReader("표준")); err != nil || s != "표준" {
		t.Fatal(s, err)
	}
	if _, err := readInput("/nonexistent.txt", nil); err == nil {
		t.Fatal("want error")
	}
}
