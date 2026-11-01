//ff:func feature=noun type=helper control=sequence
//ff:what TestSpaceVariants — 세 음절 이상 한글만 공백 변형, 짧거나 섞인 어간은 그대로
package noun

import (
	"reflect"
	"testing"
)

func TestSpaceVariants(t *testing.T) {
	if got := spaceVariants("발효식품"); !reflect.DeepEqual(got, []string{"발효식품", "발 효식품", "발효 식품", "발효식 품"}) {
		t.Fatal(got)
	}
	if len(spaceVariants("김치")) != 1 || len(spaceVariants("GPU칩셋")) != 1 {
		t.Fatal("short or mixed stems must not vary")
	}
}
