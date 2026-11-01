//ff:func feature=noun type=helper control=iteration dimension=1
//ff:what anyIn — 목록의 값 하나라도 집합에 있는지
package noun

// anyIn reports whether any value of vs is in set.
func anyIn(vs []string, set map[string]bool) bool {
	for _, v := range vs {
		if set[v] {
			return true
		}
	}
	return false
}
