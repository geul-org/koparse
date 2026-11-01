//ff:func feature=config type=parser control=sequence
//ff:what decode — JSON 설정을 base 위에 덮어 읽는다
package config

import "encoding/json"

// decode unmarshals b over base; fields absent from b keep base's values.
func decode(b []byte, base Config) (Config, error) {
	c := base
	err := json.Unmarshal(b, &c)
	return c, err
}
