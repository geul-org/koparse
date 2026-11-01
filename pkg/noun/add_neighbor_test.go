//ff:func feature=noun type=helper control=sequence
//ff:what TestAddNeighbor — 채택 QID와 그 P31·P279만 넣는다
package noun

import (
	"testing"

	"github.com/geul-org/koparse/pkg/knowledge"
)

func TestAddNeighbor(t *testing.T) {
	u := Unit{Cands: []knowledge.Candidate{{QID: "Q1"}, {QID: "Q2"}}, Infos: []knowledge.EntityInfo{{P31: []string{"Q10"}}, {P31: []string{"Q20"}, P279: []string{"Q21"}}}}
	nb := map[string]bool{}
	addNeighbor(nb, u, "Q2")
	if len(nb) != 3 || !nb["Q2"] || !nb["Q20"] || !nb["Q21"] || nb["Q10"] {
		t.Fatal(nb)
	}
}
