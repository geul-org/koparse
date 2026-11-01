//ff:func feature=config type=parser control=sequence
//ff:what TestDecode — 파일에 없는 값은 base 값을 지키고, 깨진 JSON은 오류
package config

import "testing"

func TestDecode(t *testing.T) {
	base := Config{Noun: Noun{QR1: 0.6, QR2: 1}}
	c, err := decode([]byte(`{"noun":{"q_r1":0.3}}`), base)
	if err != nil || c.Noun.QR1 != 0.3 || c.Noun.QR2 != 1 {
		t.Fatal(c, err)
	}
	if _, err := decode([]byte("{"), base); err == nil {
		t.Fatal("want error")
	}
}
