//ff:type feature=fixture type=model
//ff:what entityRecord — fixture 개체 한 줄(라벨, 설명, P31, P279)
package fixture

// entityRecord is one entity of the fixture file.
type entityRecord struct {
	Labels       map[string]string `json:"labels"`
	Descriptions map[string]string `json:"descriptions"`
	P31          []string          `json:"p31"`
	P279         []string          `json:"p279"`
}
