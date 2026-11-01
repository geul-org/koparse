//ff:func feature=fixture type=adapter control=sequence
//ff:what TestFixtureStatements — 속성 목록 nil은 전부, 목록이면 있는 속성만
package fixture

import (
	"context"
	"testing"
)

func TestFixtureStatements(t *testing.T) {
	f := New(nil, nil, map[string]map[string][]string{"Q1": {"P31": {"Q2"}, "P279": {"Q3"}}})
	all, _ := f.Statements(context.Background(), "Q1", nil)
	some, err := f.Statements(context.Background(), "Q1", []string{"P279", "P17"})
	if err != nil || len(all) != 2 || len(some) != 1 || some["P279"][0] != "Q3" {
		t.Fatal(all, some, err)
	}
}
