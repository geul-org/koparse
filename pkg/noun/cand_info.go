//ff:type feature=noun type=model
//ff:what candInfo — QID 후보 하나의 판정 재료(주제·진술 일치, 고유 분류, 일반 후보 유일, 이웃 분야)
package noun

// candInfo is what the noun rules know about one QID candidate.
type candInfo struct {
	QID       string
	Topic     bool
	Statement bool
	Proper    bool
	General   bool // the only candidate left after removing proper classes (R4a)
	Neighbor  bool // the only general candidate sharing a class with adopted nouns (R4b)
}
