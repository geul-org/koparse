//ff:func feature=noun type=helper control=sequence
//ff:what TestChainForms — 형태들만
package noun

import (
	"reflect"
	"testing"

	"github.com/geul-org/koparse/pkg/lexicon"
)

func TestChainForms(t *testing.T) {
	if got := chainForms([]lexicon.Suffix{{Form: "들"}, {Form: "은"}}); !reflect.DeepEqual(got, []string{"들", "은"}) {
		t.Fatal(got)
	}
}
