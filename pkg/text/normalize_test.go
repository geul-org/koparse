//ff:func feature=text type=helper control=sequence
//ff:what TestNormalize — 풀어 쓴 자모(NFD)를 완성형(NFC)으로
package text

import "testing"

func TestNormalize(t *testing.T) {
	if got := Normalize("가"); got != "가" {
		t.Fatalf("got %q", got)
	}
}
