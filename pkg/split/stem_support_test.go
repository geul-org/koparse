//ff:func feature=split type=rule control=iteration dimension=1
//ff:what TestStemSupport — 비한글 어간, 아는 어간, 모르는 어간(접미 무게 × unknown_scale)
package split

import (
	"math"
	"testing"

	"github.com/geul-org/koparse/pkg/config"
	"github.com/geul-org/koparse/pkg/lexicon"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestStemSupport(t *testing.T) {
	eul := []lexicon.Suffix{{Form: "을", Weight: 0.8}}
	f := &facts{
		cands:     []Candidate{{Stem: "x"}, {Stem: "km2", Suffixes: eul}, {Stem: "책", Suffixes: eul}, {Stem: "모름", Suffixes: eul}},
		stemCount: []int{0, 0, 1, 0},
		cfg:       config.Split{ForeignStem: 0.9, KnownDecay: 0.1, KnownMin: 0.4, UnknownScale: 0.5},
	}
	ctx := toulmin.NewContext()
	ctx.Set(factsKey, f)
	want := []float64{0, 0.9, 1, 2*0.5*0.8 - 1}
	for i := 1; i < 4; i++ {
		v, ev := stemSupport(ctx, toulmin.Specs{candSpec{Idx: i}})
		if math.Abs(v-want[i]) > 1e-9 || ev != i {
			t.Errorf("cand %d: %v %v", i, v, ev)
		}
	}
}
