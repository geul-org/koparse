//ff:type feature=fixture type=adapter
//ff:what Fixture — 메모리 표 지식 어댑터(언어별 표면→라벨·별칭 QID, 개체 요약, 진술). 조회하지 못한 표면을 센다
package fixture

import "sync"

// Fixture is an in-memory Knowledge adapter loaded from a JSON file.
// Surfaces that are not in the table answer "no candidates" and are counted
// as misses, so an evaluation can report how much of its lookups the table covered.
type Fixture struct {
	forms      map[string]map[string]formEntry
	entities   map[string]entityRecord
	statements map[string]map[string][]string
	mu         sync.Mutex
	misses     map[string]int
}
