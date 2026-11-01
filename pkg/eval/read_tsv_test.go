//ff:func feature=eval type=parser control=sequence
//ff:what TestReadTSV — 빈 줄·# 줄을 건너뛰고, 필드가 모자라거나 파일이 없으면 오류
package eval

import "testing"

func TestReadTSV(t *testing.T) {
	rows, err := readTSV(writeTemp(t, "a.tsv", "# h\n\na\tb\n"), 2)
	if err != nil || len(rows) != 1 || rows[0][1] != "b" {
		t.Fatal(rows, err)
	}
	if _, err := readTSV(writeTemp(t, "b.tsv", "a\n"), 2); err == nil {
		t.Fatal("want field error")
	}
	if _, err := readTSV("/nonexistent/a.tsv", 1); err == nil {
		t.Fatal("want read error")
	}
}
