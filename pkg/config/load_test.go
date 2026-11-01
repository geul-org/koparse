//ff:func feature=config type=parser control=sequence
//ff:what TestLoad — 파일을 기본값 위에 덮고, 없는 파일·깨진 파일은 오류
package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	dir := t.TempDir()
	good := filepath.Join(dir, "c.json")
	os.WriteFile(good, []byte(`{"split":{"margin":0.2}}`), 0o644)
	c, err := Load(good)
	if err != nil || c.Split.Margin != 0.2 || c.Noun.QR2 != Default().Noun.QR2 {
		t.Fatal(c, err)
	}
	if _, err := Load(filepath.Join(dir, "none.json")); err == nil {
		t.Fatal("want error for a missing file")
	}
	bad := filepath.Join(dir, "bad.json")
	os.WriteFile(bad, []byte("{"), 0o644)
	if _, err := Load(bad); err == nil {
		t.Fatal("want error for broken JSON")
	}
}
