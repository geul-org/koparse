//ff:type feature=fixture type=model
//ff:what fileDoc — fixture JSON 파일 구조(forms[lang][surface], entities, statements)
package fixture

// fileDoc is the fixture JSON file.
type fileDoc struct {
	Source     string                          `json:"source"`
	Forms      map[string]map[string]formEntry `json:"forms"`
	Entities   map[string]entityRecord         `json:"entities"`
	Statements map[string]map[string][]string  `json:"statements"`
}
