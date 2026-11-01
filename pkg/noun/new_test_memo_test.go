//ff:func feature=noun type=helper control=sequence
//ff:what newTestMemo — 시험용 조회 캐시: 작은 지은 표(물은 음절 후보와 일반 후보, 속은 음절뿐, 위키는 내부 문서), errOn 표면과 QERR 개체는 오류
package noun

import (
	"context"
	"errors"
	"testing"

	"github.com/geul-org/koparse/pkg/fixture"
	"github.com/geul-org/koparse/pkg/knowledge"
	"github.com/geul-org/koparse/pkg/lexicon"
)

const testFixtureJSON = `{"forms":{"ko":{
 "김치":{"label":["Q1"],"alias":[]}, "발효":{"label":["Q2"],"alias":[]}, "식품":{"label":["Q3"],"alias":[]},
 "발효 식품":{"label":["Q4"],"alias":[]}, "한강":{"label":["Q5","Q6"],"alias":[]}, "연안":{"label":["Q7"],"alias":[]},
 "물":{"label":["Q8","Q9"],"alias":[]}, "속":{"label":["Q10"],"alias":[]}, "땅":{"label":["Q11"],"alias":[]},
 "위키":{"label":["Q12"],"alias":[]}, "고장":{"label":["QERR"],"alias":[]}, "자리":{"label":["Q13"],"alias":[]},
 "강 하구":{"label":["Q14"],"alias":[]}}},
 "entities":{"Q1":{"labels":{"ko":"김치"},"p31":["Q100"]},"Q2":{"p31":[]},"Q3":{"p31":[]},"Q4":{"p31":[]},
 "Q5":{"p31":["Q4022"],"p279":["Q200"]},"Q6":{"p31":["Q5"]},"Q7":{"p31":[]},"Q8":{"p31":["Q55809450"]},"Q9":{"p31":[]},
 "Q10":{"p31":["Q55809450"]},"Q11":{"p31":[]},"Q12":{"p31":["Q4167410"]},"Q13":{"p31":[]},"Q14":{"p31":[]}},
 "statements":{"Q1":{"P17":["Q5"]}}}`

// newTestMemo returns a Memo over the test table; looking up errOn fails.
func newTestMemo(t *testing.T, errOn string) *Memo {
	fx, err := fixture.FromJSON([]byte(testFixtureJSON))
	if err != nil {
		t.Fatal(err)
	}
	lex, _ := lexicon.Default()
	k := knowledge.Funcs{
		CandidatesFn: func(ctx context.Context, s, l string) ([]knowledge.Candidate, error) {
			if s == errOn {
				return nil, errors.New("boom")
			}
			return fx.Candidates(ctx, s, l)
		},
		EntityFn: func(ctx context.Context, q string) (knowledge.EntityInfo, error) {
			if q == "QERR" {
				return knowledge.EntityInfo{}, errors.New("boom")
			}
			return fx.Entity(ctx, q)
		},
		StatementsFn: fx.Statements,
	}
	return NewMemo(k, lex)
}
