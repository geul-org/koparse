//ff:type feature=noun type=model
//ff:what piece — 복합어 조각(형태, 아는 조각인지 아니면 접사로 본 한 음절인지)
package noun

// piece is one piece of a compound stem; unknown pieces are edge affixes.
type piece struct {
	form  string
	known bool
	r     lookupResult
}
