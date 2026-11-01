//ff:func feature=noun type=helper control=sequence
//ff:what Entity — 개체 요약(캐시). 어댑터에 없는 개체는 빈 요약
package noun

import (
	"context"
	"errors"

	"github.com/geul-org/koparse/pkg/knowledge"
)

// Entity returns the cached summary of qid; an entity unknown to the adapter
// yields an empty summary.
func (m *Memo) Entity(ctx context.Context, qid string) (knowledge.EntityInfo, error) {
	if e, ok := m.ents[qid]; ok {
		return e, nil
	}
	e, err := m.K.Entity(ctx, qid)
	if errors.Is(err, knowledge.ErrNotFound) {
		e, err = knowledge.EntityInfo{QID: qid}, nil
	}
	if err != nil {
		return e, err
	}
	m.ents[qid] = e
	return e, nil
}
