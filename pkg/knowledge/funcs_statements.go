//ff:func feature=knowledge type=adapter control=sequence
//ff:what Statements — StatementsFn을 부른다
package knowledge

import "context"

// Statements implements Knowledge.
func (f Funcs) Statements(ctx context.Context, qid string, props []string) (map[string][]string, error) {
	return f.StatementsFn(ctx, qid, props)
}
