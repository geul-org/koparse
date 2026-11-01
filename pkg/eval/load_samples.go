//ff:func feature=eval type=parser control=sequence
//ff:what LoadSamples — 표본 JSON({"items":[{id,qid,text}]})을 읽는다. 원문은 저장소 밖에 둔다
package eval

import (
	"encoding/json"
	"fmt"
	"os"
)

// LoadSamples reads a sample file of the form {"items": [{"id", "qid", "text"}]}.
// Evaluation texts may carry non-CC0 licenses and are never copied into the repository.
func LoadSamples(path string) ([]Sample, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var doc struct {
		Items []Sample `json:"items"`
	}
	if err := json.Unmarshal(b, &doc); err != nil {
		return nil, fmt.Errorf("samples %s: %w", path, err)
	}
	return doc.Items, nil
}
