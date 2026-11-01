//ff:func feature=fixture type=parser control=sequence
//ff:what FromJSON — fixture JSON 바이트에서 어댑터를 만든다
package fixture

import "encoding/json"

// FromJSON builds a Fixture from the JSON file format.
func FromJSON(b []byte) (*Fixture, error) {
	var doc fileDoc
	if err := json.Unmarshal(b, &doc); err != nil {
		return nil, err
	}
	return New(doc.Forms, doc.Entities, doc.Statements), nil
}
