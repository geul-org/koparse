//ff:type feature=knowledge type=model
//ff:what Knowledge — 지식 어댑터 인터페이스(표면→후보, 개체 요약, 개체 진술). 판정 규칙은 이것만 본다
package knowledge

import "context"

// Knowledge is the knowledge adapter. Judgement rules see only this
// interface; whether it is backed by SILK over the Wikidata GEUL archive or by
// a test fixture is hidden.
type Knowledge interface {
	// Candidates returns entities whose label or alias in lang equals surface exactly.
	Candidates(ctx context.Context, surface string, lang string) ([]Candidate, error)
	// Entity returns a summary of one entity: labels, descriptions (ko·en), some P31·P279.
	Entity(ctx context.Context, qid string) (EntityInfo, error)
	// Statements returns the entity's statement values for props (nil props = all).
	Statements(ctx context.Context, qid string, props []string) (map[string][]string, error)
}
