//ff:func feature=lexicon type=parser control=sequence
//ff:what TestParseSuffixes — 긴 형태부터 정렬, @allow 예외, 잘못된 줄 오류
package lexicon

import "testing"

func TestParseSuffixes(t *testing.T) {
	list, allow, err := parseSuffixes("을\tparticle\t4\tC\t0.9\n에서는\tparticle\t2\t*\t0.9\n@allow\t만\t으로\n")
	if err != nil || len(list) != 2 || list[0].Form != "에서는" || !allow["만\t으로"] {
		t.Fatalf("got %v %v %v", list, allow, err)
	}
	if _, _, err := parseSuffixes("을\tparticle\n"); err == nil {
		t.Fatal("want error for a short line")
	}
}
