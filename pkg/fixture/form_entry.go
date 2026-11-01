//ff:type feature=fixture type=model
//ff:what formEntry — fixture 표면 한 줄(라벨 일치 QID, 별칭 일치 QID)
package fixture

// formEntry lists the entities whose label or alias equals a surface.
type formEntry struct {
	Label []string `json:"label"`
	Alias []string `json:"alias"`
}
