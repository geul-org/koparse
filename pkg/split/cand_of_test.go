//ff:func feature=split type=helper control=sequence
//ff:what TestCandOf — Spec의 후보 번호
package split

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestCandOf(t *testing.T) {
	if candOf(toulmin.Specs{candSpec{Idx: 3}}) != 3 {
		t.Fatal("candOf")
	}
}
