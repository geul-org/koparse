//ff:func feature=cli type=command control=sequence
//ff:what TestParseOptions — 필수 인자, 모르는 인자, 빠진 필수 인자
package main

import (
	"io"
	"testing"
)

func TestParseOptions(t *testing.T) {
	o, err := parseOptions(writeTestFiles(t), io.Discard)
	if err != nil || o.fixture == "" || o.out == "" || o.judgments == "" {
		t.Fatal(o, err)
	}
	if _, err := parseOptions([]string{"-nope"}, io.Discard); err == nil {
		t.Fatal("want flag error")
	}
	if _, err := parseOptions([]string{"-fixture", "x"}, io.Discard); err == nil {
		t.Fatal("want required error")
	}
}
