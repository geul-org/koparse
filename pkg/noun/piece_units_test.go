//ff:func feature=noun type=builder control=sequence
//ff:what TestPieceUnits — 아는 조각만 단위가 되고 마지막 아는 조각이 끝 단위
package noun

import "testing"

func TestPieceUnits(t *testing.T) {
	w := []Word{{Surface: "박명성은", Stem: "박명성"}, {Surface: "보인다"}}
	us := pieceUnits(w, 0, []piece{{form: "박명", known: true}, {form: "성"}})
	if len(us) != 1 || us[0].Form != "박명" || !us[0].Final || !us[0].Piece || us[0].Governor != "보인다" {
		t.Fatal(us)
	}
}
