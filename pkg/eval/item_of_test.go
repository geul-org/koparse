//ff:func feature=eval type=helper control=sequence
//ff:what TestItemOf — 점 앞부분, 점이 없으면 통째
package eval

import "testing"

func TestItemOf(t *testing.T) {
	if itemOf("10.3") != "10" || itemOf("7") != "7" {
		t.Fatal("itemOf")
	}
}
