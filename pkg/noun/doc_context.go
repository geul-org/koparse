//ff:type feature=noun type=model
//ff:what DocContext — 문서 문맥(주제 QID, 주제 개체의 진술 값, 이웃 분야 집합)
package noun

// DocContext is the document-level context of the noun rules.
type DocContext struct {
	Topic string
	// Statements holds every statement value of the topic entity.
	Statements map[string]bool
	// Neighbors holds the QIDs and P31·P279 classes of nouns adopted elsewhere in the document.
	Neighbors map[string]bool
}
