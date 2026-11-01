//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestLoadInputs — 네 파일을 읽고, 어느 하나라도 없으면 오류
package main

import (
	"io"
	"testing"
)

func TestLoadInputs(t *testing.T) {
	o, _ := parseOptions(writeTestFiles(t), io.Discard)
	in, err := loadInputs(o)
	if err != nil || len(in.samples) != 1 || len(in.goldSplit["1.1"]) != 3 || len(in.goldNoun) != 1 || len(in.judgments) != 1 {
		t.Fatal(in, err)
	}
	for _, set := range []func(*options){
		func(o *options) { o.samples = "/nonexistent" },
		func(o *options) { o.goldSplit = "/nonexistent" },
		func(o *options) { o.goldNoun = "/nonexistent" },
	} {
		bad := o
		set(&bad)
		if _, err := loadInputs(bad); err == nil {
			t.Error("want error")
		}
	}
}
