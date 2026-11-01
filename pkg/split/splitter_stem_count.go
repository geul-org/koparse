//ff:func feature=split type=helper control=sequence
//ff:what stemCount — 어간 표면의 지식 후보 수(캐시). 어절 단위로 모아 같은 어간은 한 번만 묻는다
package split

import "context"

// stemCount returns how many entities have surface as a Korean label or alias.
func (s *Splitter) stemCount(ctx context.Context, surface string) (int, error) {
	if n, ok := s.count[surface]; ok {
		return n, nil
	}
	cands, err := s.K.Candidates(ctx, surface, "ko")
	if err != nil {
		return 0, err
	}
	s.count[surface] = len(cands)
	return len(cands), nil
}
