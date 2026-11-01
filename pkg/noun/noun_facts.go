//ff:type feature=noun type=model
//ff:what nounFacts — 명사 하나의 판정 재료(후보들, 성질·틀 명사 여부, 관용 틀 여부)
package noun

// nounFactsKey is the toulmin context key holding *nounFacts.
const nounFactsKey = "noun.facts"

// nounFacts is what the noun rules read from the toulmin context.
type nounFacts struct {
	cands    []candInfo
	property bool
	idiom    bool
}
