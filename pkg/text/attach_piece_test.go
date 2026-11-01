//ff:func feature=text type=parser control=sequence
//ff:what TestAttachPiece — 어절은 붙이고, 쉼표만 있는 조각은 앞 어절에 쉼표를 옮기며 앞 어절이 없으면 버린다
package text

import "testing"

func TestAttachPiece(t *testing.T) {
	out := attachPiece(nil, ",")
	out = attachPiece(out, "가")
	out = attachPiece(out, ",")
	out = attachPiece(out, "!")
	if len(out) != 1 || out[0].Surface != "가" || !out[0].Break {
		t.Fatalf("%+v", out)
	}
}
