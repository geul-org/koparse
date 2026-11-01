//ff:func feature=fixture type=adapter control=sequence
//ff:what TestFixtureCandidates — 라벨 먼저, 겹치는 별칭은 한 번, 표에 없는 표면은 놓침으로 센다
package fixture

import (
	"context"
	"testing"

	"github.com/geul-org/koparse/pkg/knowledge"
)

func TestFixtureCandidates(t *testing.T) {
	f := New(map[string]map[string]formEntry{"ko": {"김치": {Label: []string{"Q1"}, Alias: []string{"Q1", "Q2"}}}}, nil, nil)
	var k knowledge.Knowledge = f
	got, err := k.Candidates(context.Background(), "김치", "ko")
	if err != nil || len(got) != 2 || got[0] != (knowledge.Candidate{QID: "Q1", Match: "label", Surface: "김치", Lang: "ko"}) || got[1].Match != "alias" {
		t.Fatal(got, err)
	}
	if got, _ := f.Candidates(context.Background(), "없음", "ko"); got != nil || f.misses["ko:없음"] != 1 {
		t.Fatal(got, f.misses)
	}
}
