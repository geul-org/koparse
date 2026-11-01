//ff:func feature=stats type=builder control=iteration dimension=1
//ff:what Records — 문장의 모든 명사(채택·버림·미정)에 문맥 통계 레코드를 만든다
package stats

import "github.com/geul-org/koparse/pkg/noun"

// Records returns one record per noun unit of the sentence, in text order.
// results[i] is the judgement of units[i].
func Records(words []noun.Word, units []noun.Unit, results []noun.Result) []Record {
	toks := buildTokens(words, units)
	out := []Record{}
	for p, t := range toks {
		if t.Unit < 0 {
			continue
		}
		josa := josaAfter(toks, p)
		gov, g := governorFor(toks, p, josa)
		r := results[t.Unit]
		out = append(out, Record{Noun: t.Text, QID: r.QID, Status: r.Status, Josa: josa, Governor: gov,
			Negated: g >= 0 && negatedAt(words, toks[g].Eojeol), Pairs: pairsFor(toks, p)})
	}
	return out
}
