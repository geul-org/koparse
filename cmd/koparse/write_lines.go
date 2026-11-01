//ff:func feature=cli type=command control=iteration dimension=1
//ff:what writeLines — 문장 출력을 JSON 줄로 쓴다
package main

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/geul-org/koparse/pkg/analyze"
)

// writeLines encodes each sentence as one JSON line; returns the exit code.
func writeLines(enc *json.Encoder, out []analyze.SentenceOut, stderr io.Writer) int {
	for _, s := range out {
		if err := enc.Encode(s); err != nil {
			fmt.Fprintln(stderr, "koparse:", err)
			return 1
		}
	}
	return 0
}
