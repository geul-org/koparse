//ff:func feature=eval type=judge control=selection
//ff:what judge — 채택된 (문장, 명사, QID)를 판정: 손 판정 → 같은 문장 정답 → 같은 표본 정답 → 미판정
package eval

// judge returns Correct, Wrong or Unjudged for an adopted noun.
func (ix goldIndex) judge(key, form, qid string) string {
	item := itemOf(key)
	k, i := ix.byKey[key+"\t"+form], ix.byItem[item+"\t"+form]
	switch {
	case ix.extra[key+"\t"+form+"\t"+qid] != "":
		return ix.extra[key+"\t"+form+"\t"+qid]
	case ix.extra[item+"\t"+form+"\t"+qid] != "":
		return ix.extra[item+"\t"+form+"\t"+qid]
	case k.good[qid]:
		return Correct
	case k.bad[qid]:
		return Wrong
	case i.good[qid]:
		return Correct
	case i.bad[qid]:
		return Wrong
	default:
		return Unjudged
	}
}
