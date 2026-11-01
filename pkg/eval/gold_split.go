//ff:type feature=eval type=model
//ff:what GoldSplit — 분리 정답 한 어절(문장 열쇠 item.sent, 순번, 정답 분리 어간+접미, 범주)
package eval

// GoldSplit is the gold split of one eojeol.
type GoldSplit struct {
	Key string
	Seg string // "stem+suffixes", or the eojeol itself when unsplit
	Cat string
}
