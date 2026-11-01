//ff:func feature=knowledge type=adapter control=sequence
//ff:what Candidates — CandidatesFn을 부른다
package knowledge

import "context"

// Candidates implements Knowledge.
func (f Funcs) Candidates(ctx context.Context, surface, lang string) ([]Candidate, error) {
	return f.CandidatesFn(ctx, surface, lang)
}
