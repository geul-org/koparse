//ff:func feature=config type=parser control=sequence
//ff:what TestDefault — 내장 default.json이 읽히고 설계서 초기값(R1 0.6, R2 1.0, R3 0.9)과 맞는다
package config

import "testing"

func TestDefault(t *testing.T) {
	c := Default()
	if c.Noun.QR1 != 0.6 || c.Noun.QR2 != 1.0 || c.Noun.QR3 != 0.9 || c.Split.QStem != 1.0 || c.Split.ForceBatchim != 1.0 {
		t.Fatalf("%+v", c)
	}
	if _, err := decode(defaultJSON, Config{}); err != nil {
		t.Fatal(err)
	}
}
