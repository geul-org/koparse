//ff:func feature=noun type=helper control=iteration dimension=1
//ff:what lookup — 한국어 표면의 후보(캐시). 위키 내부 문서(분류 목록 internal)는 뺀다
package noun

import "context"

// lookup returns the candidates of a Korean surface without wiki-internal
// pages, each with its entity summary.
func (m *Memo) lookup(ctx context.Context, surface string) (lookupResult, error) {
	if r, ok := m.cands[surface]; ok {
		return r, nil
	}
	cands, err := m.K.Candidates(ctx, surface, "ko")
	if err != nil {
		return lookupResult{}, err
	}
	var r lookupResult
	for _, c := range cands {
		e, err := m.Entity(ctx, c.QID)
		if err != nil {
			return lookupResult{}, err
		}
		if anyIn(e.P31, m.Lex.Internal) {
			continue
		}
		r.cands = append(r.cands, c)
		r.infos = append(r.infos, e)
	}
	m.cands[surface] = r
	return r, nil
}
