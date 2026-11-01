//ff:func feature=split type=model control=sequence
//ff:what Validate — 후보 번호가 0 이상인지
package split

import "fmt"

// Validate implements toulmin.Spec.
func (s candSpec) Validate() error {
	if s.Idx < 0 {
		return fmt.Errorf("split: negative candidate index %d", s.Idx)
	}
	return nil
}
