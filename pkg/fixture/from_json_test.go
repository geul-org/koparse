//ff:func feature=fixture type=parser control=sequence
//ff:what TestFromJSON — JSON에서 표를 읽고, 깨진 JSON은 오류
package fixture

import "testing"

func TestFromJSON(t *testing.T) {
	f, err := FromJSON([]byte(`{"forms":{"ko":{"김치":{"label":["Q1"],"alias":[]}}},"entities":{"Q1":{"labels":{"ko":"김치"},"p31":["Q2"]}},"statements":{"Q1":{"P31":["Q2"]}}}`))
	if err != nil || f.forms["ko"]["김치"].Label[0] != "Q1" || f.entities["Q1"].P31[0] != "Q2" || f.statements["Q1"]["P31"][0] != "Q2" {
		t.Fatal(f, err)
	}
	if _, err := FromJSON([]byte("{")); err == nil {
		t.Fatal("want error")
	}
}
