//ff:func feature=analyze type=helper control=iteration dimension=1
//ff:what statementSet — 주제 개체의 모든 진술 값을 집합으로(주제가 없으면 빈 집합)
package analyze

import "context"

// statementSet returns every statement value of the topic entity.
func (a *Analyzer) statementSet(ctx context.Context, topic string) (map[string]bool, error) {
	set := map[string]bool{}
	if topic == "" {
		return set, nil
	}
	st, err := a.K.Statements(ctx, topic, nil)
	if err != nil {
		return nil, err
	}
	for _, vs := range st {
		for _, v := range vs {
			set[v] = true
		}
	}
	return set, nil
}
