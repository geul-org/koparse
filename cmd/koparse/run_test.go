//ff:func feature=cli type=command control=sequence
//ff:what TestRun — 표준입력 분석, 잘못된 인자 2, fixture 없음 1, 없는 입력 파일 1
package main

import (
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	var out, errb strings.Builder
	fx := "../../testdata/fixture.json"
	if code := run([]string{"-fixture", fx, "-topic", "Q146", "-trace"}, strings.NewReader("고양이는 고기를 먹는다."), &out, &errb); code != 0 || !strings.Contains(out.String(), `"ranked"`) {
		t.Fatal(code, errb.String())
	}
	if code := run([]string{"-nope"}, nil, &out, &errb); code != 2 {
		t.Fatal(code)
	}
	if code := run(nil, nil, &out, &errb); code != 1 {
		t.Fatal(code)
	}
	if code := run([]string{"-fixture", fx, "-in", "/nonexistent.txt"}, nil, &out, &errb); code != 1 {
		t.Fatal(code)
	}
}
