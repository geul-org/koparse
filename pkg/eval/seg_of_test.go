//ff:func feature=eval type=helper control=sequence
//ff:what TestSegOf — 접미가 있으면 어간+접미들, 없으면 어간
package eval

import (
	"testing"

	"github.com/geul-org/koparse/pkg/analyze"
)

func TestSegOf(t *testing.T) {
	e := analyze.EojeolOut{Split: []analyze.Part{{Form: "선수"}, {Form: "들"}, {Form: "은"}}}
	if segOf(e) != "선수+들은" || segOf(analyze.EojeolOut{Split: []analyze.Part{{Form: "작다"}}}) != "작다" {
		t.Fatal(segOf(e))
	}
}
