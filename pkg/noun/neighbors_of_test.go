//ff:func feature=noun type=helper control=sequence
//ff:what TestNeighborsOf — 채택된 결과만 이웃이 된다
package noun

import (
	"testing"

	"github.com/geul-org/koparse/pkg/knowledge"
)

func TestNeighborsOf(t *testing.T) {
	us := []Unit{{Cands: []knowledge.Candidate{{QID: "Q1"}}, Infos: []knowledge.EntityInfo{{P31: []string{"Q10"}}}}, {}}
	nb := neighborsOf(us, []Result{{Status: StatusAdopted, QID: "Q1"}, {Status: StatusDropped}})
	if len(nb) != 2 || !nb["Q1"] || !nb["Q10"] {
		t.Fatal(nb)
	}
}
