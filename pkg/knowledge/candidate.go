//ff:type feature=knowledge type=model
//ff:what Candidate — 표면 일치 후보 개체(QID, 일치 종류 label·alias, 원문 표면, 언어)
package knowledge

// Candidate is an entity whose label or alias matched a surface form.
type Candidate struct {
	QID     string `json:"qid"`
	Match   string `json:"match"` // MatchLabel or MatchAlias
	Surface string `json:"surface"`
	Lang    string `json:"lang"`
}
