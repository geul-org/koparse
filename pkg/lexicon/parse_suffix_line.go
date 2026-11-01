//ff:func feature=lexicon type=parser control=sequence
//ff:what parseSuffixLine — 접미 목록 한 줄(형태, 종류, 자리, 받침 조건, 무게)을 검사해 읽는다
package lexicon

import (
	"fmt"
	"strconv"
)

var validConds = map[string]bool{"*": true, "C": true, "V": true, "VL": true, "CX": true}
var validKinds = map[string]bool{KindParticle: true, KindCopula: true, KindVerbal: true}

// parseSuffixLine parses the five fields of one suffix entry.
func parseSuffixLine(f []string) (Suffix, error) {
	if len(f) != 5 {
		return Suffix{}, fmt.Errorf("want 5 fields, got %d", len(f))
	}
	if !validKinds[f[1]] {
		return Suffix{}, fmt.Errorf("unknown kind %q", f[1])
	}
	slot, err := strconv.Atoi(f[2])
	if err != nil {
		return Suffix{}, fmt.Errorf("slot: %w", err)
	}
	if !validConds[f[3]] {
		return Suffix{}, fmt.Errorf("unknown cond %q", f[3])
	}
	w, err := strconv.ParseFloat(f[4], 64)
	if err != nil || w < 0 || w > 1 {
		return Suffix{}, fmt.Errorf("weight %q not in [0,1]", f[4])
	}
	return Suffix{Form: f[0], Kind: f[1], Slot: slot, Cond: f[3], Weight: w}, nil
}
