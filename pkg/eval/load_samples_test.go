//ff:func feature=eval type=parser control=sequence
//ff:what TestLoadSamples — items를 읽고, 없는 파일·깨진 JSON은 오류
package eval

import "testing"

func TestLoadSamples(t *testing.T) {
	s, err := LoadSamples(writeTemp(t, "s.json", `{"items":[{"id":1,"qid":"Q146","text":"고양이는 작다."}]}`))
	if err != nil || len(s) != 1 || s[0].QID != "Q146" {
		t.Fatal(s, err)
	}
	if _, err := LoadSamples(writeTemp(t, "b.json", "{")); err == nil {
		t.Fatal("want JSON error")
	}
	if _, err := LoadSamples("/nonexistent/s.json"); err == nil {
		t.Fatal("want read error")
	}
}
