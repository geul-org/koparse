//ff:func feature=noun type=helper control=sequence
//ff:what betterState — 조각이 적은 쪽, 같으면 접사가 적은 쪽이 낫다
package noun

// betterState reports whether a beats b: fewer pieces, then fewer affixes.
func betterState(a, b segState) bool {
	if !a.ok {
		return false
	}
	if !b.ok || a.pieces != b.pieces {
		return !b.ok || a.pieces < b.pieces
	}
	return a.affixes < b.affixes
}
