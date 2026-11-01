//ff:func feature=noun type=helper control=sequence
//ff:what TestMarkNeighbor — QID·P31·P279 중 하나라도 이웃과 겹치는 일반 후보가 하나뿐이면 표시, 둘이면 없음, 고유 후보는 제외
package noun

import (
	"testing"

	"github.com/geul-org/koparse/pkg/knowledge"
)

func TestMarkNeighbor(t *testing.T) {
	u := Unit{Infos: []knowledge.EntityInfo{{P31: []string{"Q4022"}}, {P279: []string{"Q9"}}, {}, {P31: []string{"Q4022"}}}}
	nb := map[string]bool{"Q4022": true}
	cs := []candInfo{{QID: "Q1"}, {QID: "Q2"}, {QID: "Q3"}, {QID: "Q4", Proper: true}}
	markNeighbor(cs, u, nb)
	if !cs[0].Neighbor || cs[1].Neighbor || cs[3].Neighbor {
		t.Fatal(cs)
	}
	two := []candInfo{{QID: "Q1"}, {QID: "Q2"}}
	markNeighbor(two, u, map[string]bool{"Q4022": true, "Q9": true})
	if two[0].Neighbor || two[1].Neighbor {
		t.Fatal(two)
	}
	byQID := []candInfo{{QID: "Q3"}}
	markNeighbor(byQID, Unit{Infos: []knowledge.EntityInfo{{}}}, map[string]bool{"Q3": true})
	if !byQID[0].Neighbor {
		t.Fatal(byQID)
	}
}
