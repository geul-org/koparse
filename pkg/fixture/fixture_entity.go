//ff:func feature=fixture type=adapter control=sequence
//ff:what Entity — 개체 요약을 돌려준다(없으면 knowledge.ErrNotFound)
package fixture

import (
	"context"
	"fmt"

	"github.com/geul-org/koparse/pkg/knowledge"
)

// Entity returns the summary of qid.
func (f *Fixture) Entity(_ context.Context, qid string) (knowledge.EntityInfo, error) {
	r, ok := f.entities[qid]
	if !ok {
		return knowledge.EntityInfo{}, fmt.Errorf("%s: %w", qid, knowledge.ErrNotFound)
	}
	return knowledge.EntityInfo{QID: qid, Labels: r.Labels, Descriptions: r.Descriptions, P31: r.P31, P279: r.P279}, nil
}
