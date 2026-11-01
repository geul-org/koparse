//ff:func feature=cli type=command control=sequence
//ff:what TestNewAnalyzer — fixture 없음·읽기 실패·설정 실패는 오류, 설정 파일을 반영한다
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewAnalyzer(t *testing.T) {
	fx := "../../testdata/fixture.json"
	if _, err := newAnalyzer("", ""); err == nil {
		t.Fatal("want missing-fixture error")
	}
	if _, err := newAnalyzer("/nonexistent.json", ""); err == nil {
		t.Fatal("want load error")
	}
	if _, err := newAnalyzer(fx, "/nonexistent.json"); err == nil {
		t.Fatal("want config error")
	}
	cfg := filepath.Join(t.TempDir(), "c.json")
	os.WriteFile(cfg, []byte(`{"noun":{"adopt":0.9}}`), 0o644)
	a, err := newAnalyzer(fx, cfg)
	if err != nil || a.Cfg.Noun.Adopt != 0.9 {
		t.Fatal(a, err)
	}
	if a, err := newAnalyzer(fx, ""); err != nil || a.Cfg.Noun.Adopt != 0.5 {
		t.Fatal(a, err)
	}
}
