//ff:func feature=text type=parser control=sequence
//ff:what TestAttachField — 가운뎃점으로 나눈 조각들을 어절로 붙인다
package text

import "testing"

func TestAttachField(t *testing.T) {
	out := attachField(nil, "경기도·인천광역시와")
	if len(out) != 2 || out[0].Surface != "경기도" || out[1].Surface != "인천광역시와" {
		t.Fatalf("%+v", out)
	}
}
