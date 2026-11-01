//ff:type feature=noun type=model
//ff:what hit — warrant의 근거(후보 번호, 규칙 이름). toulmin 결과를 후보로 되돌린다
package noun

// hit is the evidence returned by a noun warrant.
type hit struct {
	Cand int
	Rule string
}
