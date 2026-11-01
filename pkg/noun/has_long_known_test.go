//ff:func feature=noun type=helper control=sequence
//ff:what TestHasLongKnown — 두 음절 이상 아는 조각이 있어야 참
package noun

import "testing"

func TestHasLongKnown(t *testing.T) {
	if !hasLongKnown([]piece{{form: "땅", known: true}, {form: "김치", known: true}}) ||
		hasLongKnown([]piece{{form: "땅", known: true}, {form: "물", known: true}}) ||
		hasLongKnown([]piece{{form: "김치"}}) {
		t.Fatal("hasLongKnown")
	}
}
