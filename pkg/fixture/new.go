//ff:func feature=fixture type=builder control=sequence
//ff:what New — 표에서 fixture 어댑터를 만든다(nil 표는 빈 표)
package fixture

// New builds a Fixture from its tables; nil tables are treated as empty.
func New(forms map[string]map[string]formEntry, entities map[string]entityRecord, statements map[string]map[string][]string) *Fixture {
	if forms == nil {
		forms = map[string]map[string]formEntry{}
	}
	if entities == nil {
		entities = map[string]entityRecord{}
	}
	if statements == nil {
		statements = map[string]map[string][]string{}
	}
	return &Fixture{forms: forms, entities: entities, statements: statements, misses: map[string]int{}}
}
