//ff:func feature=knowledge type=adapter control=sequence
//ff:what TestFuncsStatements — StatementsFn으로 넘긴다
package knowledge

import (
	"context"
	"testing"
)

func TestFuncsStatements(t *testing.T) {
	f := Funcs{StatementsFn: func(_ context.Context, q string, p []string) (map[string][]string, error) {
		return map[string][]string{p[0]: {q}}, nil
	}}
	if m, err := f.Statements(context.Background(), "Q3", []string{"P31"}); err != nil || m["P31"][0] != "Q3" {
		t.Fatal(m, err)
	}
}
