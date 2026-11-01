//ff:func feature=lexicon type=parser control=sequence
//ff:what FromFS — 디렉터리(suffixes.tsv, property_nouns.txt, idioms.tsv, classes.tsv)에서 목록을 읽는다
package lexicon

import (
	"fmt"
	"io/fs"
	"path"
)

// FromFS reads the four list files from dir in fsys.
func FromFS(fsys fs.FS, dir string) (*Lexicon, error) {
	read := func(name string) (string, error) {
		b, err := fs.ReadFile(fsys, path.Join(dir, name))
		return string(b), err
	}
	suf, err := read("suffixes.tsv")
	if err != nil {
		return nil, err
	}
	prop, err := read("property_nouns.txt")
	if err != nil {
		return nil, err
	}
	idi, err := read("idioms.tsv")
	if err != nil {
		return nil, err
	}
	cls, err := read("classes.tsv")
	if err != nil {
		return nil, err
	}
	lex := &Lexicon{Property: parseWordList(prop)}
	if lex.Suffixes, lex.Allow, err = parseSuffixes(suf); err != nil {
		return nil, fmt.Errorf("suffixes.tsv: %w", err)
	}
	if lex.Idioms, err = parseIdioms(idi); err != nil {
		return nil, fmt.Errorf("idioms.tsv: %w", err)
	}
	if lex.Proper, lex.Internal, err = parseClasses(cls); err != nil {
		return nil, fmt.Errorf("classes.tsv: %w", err)
	}
	return lex, nil
}
