//ff:type feature=noun type=model
//ff:what lookupResult — 표면 하나의 조회 결과(후보와 개체 요약을 나란히)
package noun

import "github.com/geul-org/koparse/pkg/knowledge"

// lookupResult is the filtered candidates of one surface with their summaries.
type lookupResult struct {
	cands []knowledge.Candidate
	infos []knowledge.EntityInfo
}
