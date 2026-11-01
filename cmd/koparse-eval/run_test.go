//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestRun — 전체 평가 성공, 인자 오류 2, 파일·fixture·설정 오류 1
package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	args := writeTestFiles(t)
	var out strings.Builder
	if code := run(args, &out, io.Discard); code != 0 || !strings.Contains(out.String(), "명사:") {
		t.Fatal(code, out.String())
	}
	if _, err := os.Stat(filepath.Join(args[len(args)-1], "nouns.tsv")); err != nil {
		t.Fatal(err)
	}
	if run([]string{"-nope"}, io.Discard, io.Discard) != 2 {
		t.Fatal("want 2")
	}
	with := func(flag, value string) []string {
		a := append([]string{}, args...)
		for i := range a {
			if a[i] == flag {
				a[i+1] = value
			}
		}
		return a
	}
	for _, a := range [][]string{with("-samples", "/nonexistent"), with("-fixture", "/nonexistent"), with("-config", "/nonexistent")} {
		if run(a, io.Discard, io.Discard) != 1 {
			t.Errorf("want 1 for %v", a)
		}
	}
}
