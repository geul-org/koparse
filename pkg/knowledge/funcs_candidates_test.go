//ff:func feature=knowledge type=adapter control=sequence
//ff:what TestFuncsCandidates — CandidatesFn으로 넘긴다
package knowledge

import (
	"context"
	"testing"
)

func TestFuncsCandidates(t *testing.T) {
	f := Funcs{CandidatesFn: func(_ context.Context, s, l string) ([]Candidate, error) {
		return []Candidate{{QID: "Q1", Surface: s, Lang: l}}, nil
	}}
	var k Knowledge = f
	got, err := k.Candidates(context.Background(), "김치", "ko")
	if err != nil || len(got) != 1 || got[0].Surface != "김치" || got[0].Lang != "ko" {
		t.Fatal(got, err)
	}
}
