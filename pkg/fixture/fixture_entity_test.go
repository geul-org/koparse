//ff:func feature=fixture type=adapter control=sequence
//ff:what TestFixtureEntity — 있는 개체는 요약, 없는 개체는 ErrNotFound
package fixture

import (
	"context"
	"errors"
	"testing"

	"github.com/geul-org/koparse/pkg/knowledge"
)

func TestFixtureEntity(t *testing.T) {
	f := New(nil, map[string]entityRecord{"Q1": {Labels: map[string]string{"ko": "김치"}, P31: []string{"Q2"}, P279: []string{"Q3"}}}, nil)
	e, err := f.Entity(context.Background(), "Q1")
	if err != nil || e.QID != "Q1" || e.Labels["ko"] != "김치" || e.P31[0] != "Q2" || e.P279[0] != "Q3" {
		t.Fatal(e, err)
	}
	if _, err := f.Entity(context.Background(), "Q9"); !errors.Is(err, knowledge.ErrNotFound) {
		t.Fatal(err)
	}
}
