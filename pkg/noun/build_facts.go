//ff:func feature=noun type=builder control=iteration dimension=1
//ff:what buildFacts — 명사 단위와 문서 문맥에서 판정 재료를 만든다
package noun

import "github.com/geul-org/koparse/pkg/lexicon"

// buildFacts collects what the rules need about u.
func buildFacts(u Unit, dc DocContext, lex *lexicon.Lexicon) nounFacts {
	f := nounFacts{property: lex.Property[u.Form], idiom: idiomMatches(lex.Idioms, u)}
	for i, c := range u.Cands {
		f.cands = append(f.cands, candInfo{
			QID:       c.QID,
			Topic:     dc.Topic != "" && c.QID == dc.Topic,
			Statement: dc.Statements[c.QID],
			Proper:    anyIn(u.Infos[i].P31, lex.Proper),
		})
	}
	markGeneral(f.cands)
	markNeighbor(f.cands, u, dc.Neighbors)
	return f
}
