//ff:func feature=noun type=helper control=sequence
//ff:what TestParticleFits — * 아무것, - 조사 없음, 형태 일치
package noun

import "testing"

func TestParticleFits(t *testing.T) {
	if !particleFits([]string{"*"}, []string{"을"}) || !particleFits([]string{"-"}, nil) || particleFits([]string{"-"}, []string{"을"}) ||
		!particleFits([]string{"으로", "로"}, []string{"로"}) || particleFits([]string{"을"}, []string{"를"}) {
		t.Fatal("particleFits")
	}
}
