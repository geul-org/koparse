//ff:func feature=cli type=command control=sequence
//ff:what TestWriteLines — 문장마다 한 줄, 쓰기 실패는 종료 코드 1
package main

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/geul-org/koparse/pkg/analyze"
)

func TestWriteLines(t *testing.T) {
	var out, errb strings.Builder
	if code := writeLines(json.NewEncoder(&out), []analyze.SentenceOut{{Text: "가"}, {Text: "나"}}, &errb); code != 0 || strings.Count(out.String(), "\n") != 2 {
		t.Fatal(code, out.String())
	}
	if code := writeLines(json.NewEncoder(failWriter{}), []analyze.SentenceOut{{Text: "가"}}, &errb); code != 1 || errb.Len() == 0 {
		t.Fatal(code)
	}
}
