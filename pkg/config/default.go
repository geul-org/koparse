//ff:func feature=config type=parser control=sequence
//ff:what Default — 내장 기본 설정(default.json)
package config

import (
	_ "embed"
)

//go:embed default.json
var defaultJSON []byte

// Default returns the built-in configuration (default.json, checked by tests).
func Default() Config {
	c, _ := decode(defaultJSON, Config{})
	return c
}
