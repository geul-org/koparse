//ff:func feature=lexicon type=parser control=sequence
//ff:what Default — 저장소에 내장한 자체 목록(data/)을 읽는다
package lexicon

import (
	"embed"
)

//go:embed data
var dataFS embed.FS

// Default parses the lists embedded from data/.
func Default() (*Lexicon, error) {
	return FromFS(dataFS, "data")
}
