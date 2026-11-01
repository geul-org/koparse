//ff:type feature=cli type=model
//ff:what options — 평가 명령 인자(fixture, 설정, 표본, 정답 파일, 손 판정, 상세 출력 디렉터리)
package main

// options are the koparse-eval arguments. Sample texts, gold files and
// outputs live outside the repository (files/koparse/).
type options struct {
	fixture   string
	config    string
	samples   string
	goldSplit string
	goldNoun  string
	judgments string
	out       string
}
