//ff:func feature=cli type=command control=sequence
//ff:what newAnalyzer — fixture·설정 경로로 처리기를 만든다(fixture 필수)
package main

import (
	"errors"

	"github.com/geul-org/koparse/pkg/analyze"
	"github.com/geul-org/koparse/pkg/config"
	"github.com/geul-org/koparse/pkg/fixture"
	"github.com/geul-org/koparse/pkg/lexicon"
)

// newAnalyzer loads the fixture adapter, the built-in lists and the config.
func newAnalyzer(fixturePath, configPath string) (*analyze.Analyzer, error) {
	if fixturePath == "" {
		return nil, errors.New("-fixture is required (the SILK adapter comes with M3)")
	}
	fx, err := fixture.Load(fixturePath)
	if err != nil {
		return nil, err
	}
	lex, _ := lexicon.Default() // embedded lists, checked by lexicon.TestDefault
	cfg := config.Default()
	if configPath != "" {
		if cfg, err = config.Load(configPath); err != nil {
			return nil, err
		}
	}
	return analyze.New(fx, lex, cfg), nil
}
