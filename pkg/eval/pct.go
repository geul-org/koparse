//ff:func feature=eval type=helper control=sequence
//ff:what pct — a/b를 "a/b = x.x%"로(b가 0이면 "-")
package eval

import "fmt"

// pct formats a/b as a percentage.
func pct(a, b int) string {
	if b == 0 {
		return fmt.Sprintf("%d/0 = -", a)
	}
	return fmt.Sprintf("%d/%d = %.1f%%", a, b, 100*float64(a)/float64(b))
}
