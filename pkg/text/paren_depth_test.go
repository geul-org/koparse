//ff:func feature=text type=helper control=sequence
//ff:what TestParenDepth — 여는 괄호는 +1, 닫는 괄호는 −1(0 밑으로 안 감), 나머지는 그대로
package text

import "testing"

func TestParenDepth(t *testing.T) {
	if parenDepth(0, '(') != 1 || parenDepth(1, '[') != 2 || parenDepth(2, ')') != 1 || parenDepth(0, ']') != 0 || parenDepth(3, '가') != 3 {
		t.Fatal("parenDepth wrong")
	}
}
