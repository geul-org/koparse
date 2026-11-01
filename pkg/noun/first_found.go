//ff:func feature=noun type=helper control=iteration dimension=1
//ff:what firstFound — 형태들을 차례로 조회해 후보가 있는 첫 형태와 결과(없으면 빈 형태)
package noun

import "context"

// firstFound returns the first form with candidates, or "" when none has.
func (m *Memo) firstFound(ctx context.Context, forms []string) (string, lookupResult, error) {
	for _, f := range forms {
		r, err := m.lookup(ctx, f)
		if err != nil {
			return "", r, err
		}
		if len(r.cands) > 0 {
			return f, r, nil
		}
	}
	return "", lookupResult{}, nil
}
