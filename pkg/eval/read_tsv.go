//ff:func feature=eval type=parser control=iteration dimension=1
//ff:what readTSV — 탭 구분 파일을 줄별 필드로 읽는다(빈 줄·# 줄 제외, 최소 필드 수 검사)
package eval

import (
	"fmt"
	"os"
	"strings"
)

// readTSV reads a tab-separated file, skipping blank and "#" lines; every
// line must have at least minFields fields.
func readTSV(path string, minFields int) ([][]string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var out [][]string
	for n, line := range strings.Split(string(b), "\n") {
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}
		f := strings.Split(line, "\t")
		if len(f) < minFields {
			return nil, fmt.Errorf("%s:%d: want %d fields, got %d", path, n+1, minFields, len(f))
		}
		out = append(out, f)
	}
	return out, nil
}
