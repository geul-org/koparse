//ff:type feature=cli type=model
//ff:what inputs — 평가 입력(표본, 분리 정답, 명사 정답, 손 판정)
package main

import "github.com/geul-org/koparse/pkg/eval"

// inputs are the loaded evaluation files.
type inputs struct {
	samples   []eval.Sample
	goldSplit map[string][]eval.GoldSplit
	goldNoun  []eval.GoldNoun
	judgments map[string]string
}
