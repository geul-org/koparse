//ff:func feature=fixture type=parser control=sequence
//ff:what Load — fixture JSON 파일을 읽어 어댑터를 만든다
package fixture

import (
	"fmt"
	"os"
)

// Load reads a fixture JSON file.
func Load(path string) (*Fixture, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	f, err := FromJSON(b)
	if err != nil {
		return nil, fmt.Errorf("fixture %s: %w", path, err)
	}
	return f, nil
}
