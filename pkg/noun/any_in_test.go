//ff:func feature=noun type=helper control=sequence
//ff:what TestAnyIn — 하나라도 집합에 있으면 참
package noun

import "testing"

func TestAnyIn(t *testing.T) {
	set := map[string]bool{"Q5": true}
	if !anyIn([]string{"Q1", "Q5"}, set) || anyIn([]string{"Q1"}, set) || anyIn(nil, set) {
		t.Fatal("anyIn")
	}
}
