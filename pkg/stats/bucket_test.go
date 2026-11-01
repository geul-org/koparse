//ff:func feature=stats type=helper control=iteration dimension=1
//ff:what TestBucket — 부호 있는 구간 경계(±1, ±2, ±3~4, ±5~8, ±9~16, 같은 문장)
package stats

import "testing"

func TestBucket(t *testing.T) {
	cases := map[int]string{1: "+1", -1: "-1", 2: "+2", -3: "-3~4", 4: "+3~4", 5: "+5~8", -8: "-5~8", 9: "+9~16", 16: "+9~16", 17: "같은 문장", -40: "같은 문장"}
	for d, want := range cases {
		if got := Bucket(d); got != want {
			t.Errorf("Bucket(%d) = %q, want %q", d, got, want)
		}
	}
}
