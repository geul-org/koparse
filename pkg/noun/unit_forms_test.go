//ff:func feature=noun type=helper control=iteration dimension=1
//ff:what unitForms — 시험용: 단위들의 형태 목록
package noun

// unitForms returns the forms of us.
func unitForms(us []Unit) []string {
	var out []string
	for _, u := range us {
		out = append(out, u.Form)
	}
	return out
}
