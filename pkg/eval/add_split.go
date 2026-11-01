//ff:func feature=eval type=judge control=sequence
//ff:what addSplit — 어절 하나의 비교 결과를 집계에 더한다
package eval

// addSplit records one compared eojeol.
func addSplit(s *SplitScore, key string, g GoldSplit, got string, undecided bool) {
	ok := got == g.Seg
	c := s.ByCat[g.Cat]
	c[1]++
	s.Total++
	if undecided {
		s.Undecided++
	}
	if ok {
		c[0]++
		s.Correct++
	}
	if ok && undecided {
		s.UndecidedCorrect++
	}
	if !ok {
		s.Mismatches = append(s.Mismatches, key+"\t"+g.Seg+"\t"+got+"\t"+g.Cat)
	}
	s.ByCat[g.Cat] = c
}
