//ff:type feature=eval type=model
//ff:what nounSets — 명사 하나의 맞는·틀린 QID 집합
package eval

// nounSets holds the good and bad QIDs of a noun.
type nounSets struct {
	good map[string]bool
	bad  map[string]bool
}
