//ff:type feature=config type=model
//ff:what Config — 판정 설정(분리·명사의 Qualifier, 공격 세기, 문턱)
package config

// Config holds every tunable number of the TAN judgements. Values start
// hand-set (default.json) and may later be learned from gold data.
type Config struct {
	Split Split `json:"split"`
	Noun  Noun  `json:"noun"`
}
