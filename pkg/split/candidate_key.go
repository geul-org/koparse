//ff:func feature=split type=helper control=iteration dimension=1
//ff:what Key — 후보를 가르는 열쇠(어간|형태:종류…)
package split

import "strings"

// Key identifies the candidate by stem and the form and kind of each suffix.
func (c Candidate) Key() string {
	var b strings.Builder
	b.WriteString(c.Stem)
	for _, s := range c.Suffixes {
		b.WriteString("|" + s.Form + ":" + s.Kind)
	}
	return b.String()
}
