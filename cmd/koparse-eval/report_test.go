//ff:func feature=cli type=command control=sequence
//ff:what TestReport — -out 없으면 요약만, 있으면 상세도, 상세 쓰기 실패는 종료 코드 1
package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestReport(t *testing.T) {
	var out strings.Builder
	if code := report(options{}, inputs{}, nil, nil, &out, io.Discard); code != 0 || !strings.Contains(out.String(), "분리 정확도") {
		t.Fatal(code)
	}
	if code := report(options{out: filepath.Join(t.TempDir(), "d")}, inputs{}, nil, nil, io.Discard, io.Discard); code != 0 {
		t.Fatal(code)
	}
	blocker := filepath.Join(t.TempDir(), "file")
	os.WriteFile(blocker, nil, 0o644)
	if code := report(options{out: filepath.Join(blocker, "sub")}, inputs{}, nil, nil, io.Discard, io.Discard); code != 1 {
		t.Fatal(code)
	}
}
