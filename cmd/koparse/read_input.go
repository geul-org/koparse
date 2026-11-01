//ff:func feature=cli type=command control=sequence
//ff:what readInput — 파일 경로가 있으면 파일을, 없으면 표준입력을 다 읽는다
package main

import (
	"io"
	"os"
)

// readInput reads the whole input from path, or from stdin when path is "".
func readInput(path string, stdin io.Reader) (string, error) {
	if path != "" {
		b, err := os.ReadFile(path)
		return string(b), err
	}
	b, err := io.ReadAll(stdin)
	return string(b), err
}
