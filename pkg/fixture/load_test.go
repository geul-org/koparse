//ff:func feature=fixture type=parser control=sequence
//ff:what TestLoad — 저장소 fixture를 읽고, 없는 파일·깨진 파일은 오류
package fixture

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	f, err := Load("../../testdata/fixture.json")
	if err != nil || len(f.forms["ko"]) == 0 || len(f.entities) == 0 {
		t.Fatal(err)
	}
	if _, err := Load(filepath.Join(t.TempDir(), "none.json")); err == nil {
		t.Fatal("want error for a missing file")
	}
	bad := filepath.Join(t.TempDir(), "bad.json")
	os.WriteFile(bad, []byte("{"), 0o644)
	if _, err := Load(bad); err == nil {
		t.Fatal("want error for broken JSON")
	}
}
