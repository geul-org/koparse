//ff:func feature=hangul type=helper control=sequence
//ff:what TestCount — 글자 수는 바이트가 아니라 룬
package hangul

import "testing"

func TestCount(t *testing.T) {
	if Count("김치") != 2 || Count("GPU") != 3 || Count("") != 0 {
		t.Fatal("Count wrong")
	}
}
