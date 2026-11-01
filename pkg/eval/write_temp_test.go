//ff:func feature=eval type=helper control=sequence
//ff:what writeTemp — 시험용: 임시 디렉터리에 파일을 쓰고 경로를 돌려준다
package eval

import (
	"os"
	"path/filepath"
	"testing"
)

// writeTemp writes body to a new file in t's temp dir.
func writeTemp(t *testing.T, name, body string) string {
	p := filepath.Join(t.TempDir(), name)
	if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	return p
}
