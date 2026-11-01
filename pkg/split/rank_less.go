//ff:func feature=split type=judge control=sequence
//ff:what rankLess — 순위: verdict 높은 것, 같으면 통째, 접미 적은 것, 어간 긴 것
package split

import "unicode/utf8"

// rankLess orders a before b: higher verdict, then fewer suffixes (whole
// first), then the longer stem.
func rankLess(a, b Scored) bool {
	if a.Verdict != b.Verdict {
		return a.Verdict > b.Verdict
	}
	if len(a.Candidate.Suffixes) != len(b.Candidate.Suffixes) {
		return len(a.Candidate.Suffixes) < len(b.Candidate.Suffixes)
	}
	return utf8.RuneCountInString(a.Candidate.Stem) > utf8.RuneCountInString(b.Candidate.Stem)
}
