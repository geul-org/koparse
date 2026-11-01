//ff:func feature=cli type=command control=sequence
//ff:what loadInputs — 표본·정답·손 판정 파일을 읽는다
package main

import "github.com/geul-org/koparse/pkg/eval"

// loadInputs reads every evaluation file named in o.
func loadInputs(o options) (inputs, error) {
	var in inputs
	var err error
	if in.samples, err = eval.LoadSamples(o.samples); err != nil {
		return in, err
	}
	if in.goldSplit, err = eval.LoadGoldSplit(o.goldSplit); err != nil {
		return in, err
	}
	if in.goldNoun, err = eval.LoadGoldNoun(o.goldNoun); err != nil {
		return in, err
	}
	in.judgments, err = eval.LoadJudgments(o.judgments)
	return in, err
}
