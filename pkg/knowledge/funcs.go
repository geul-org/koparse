//ff:type feature=knowledge type=adapter
//ff:what Funcs — 함수 세 개로 만드는 지식 어댑터(다른 어댑터를 감싸거나 시험에서 오류를 넣을 때)
package knowledge

import "context"

// Funcs adapts three functions to Knowledge, e.g. to wrap another adapter.
type Funcs struct {
	CandidatesFn func(ctx context.Context, surface, lang string) ([]Candidate, error)
	EntityFn     func(ctx context.Context, qid string) (EntityInfo, error)
	StatementsFn func(ctx context.Context, qid string, props []string) (map[string][]string, error)
}
