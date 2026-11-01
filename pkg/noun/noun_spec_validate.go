//ff:func feature=noun type=model control=sequence
//ff:what Validate — 후보 번호가 0 이상인지
package noun

import "fmt"

// Validate implements toulmin.Spec.
func (s nounSpec) Validate() error {
	if s.Cand < 0 {
		return fmt.Errorf("noun: negative candidate index %d", s.Cand)
	}
	return nil
}
