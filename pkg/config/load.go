//ff:func feature=config type=parser control=sequence
//ff:what Load — 설정 파일을 기본값 위에 덮어 읽는다(파일에 없는 값은 기본값)
package config

import (
	"fmt"
	"os"
)

// Load reads a JSON configuration file over the defaults.
func Load(path string) (Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return Default(), err
	}
	c, err := decode(b, Default())
	if err != nil {
		return c, fmt.Errorf("config %s: %w", path, err)
	}
	return c, nil
}
