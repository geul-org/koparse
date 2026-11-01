package stats

// endingSyllables are syllables that commonly end a predicate eojeol.
var endingSyllables = map[rune]bool{
	'다': true, '며': true, '고': true, '는': true, '던': true, '서': true, '여': true, '어': true,
	'아': true, '게': true, '면': true, '록': true, '자': true, '지': true, '해': true, '라': true,
	'니': true, '까': true, '요': true, '져': true, '워': true, '와': true, '나': true, '야': true,
}

// connectiveSyllables end a clause (연결어미): 고, 며, 지만, 서, 면, 자, 여, 아/어, 라, 니, 데, 나.
var connectiveSyllables = map[rune]bool{
	'고': true, '며': true, '만': true, '서': true, '면': true, '자': true, '여': true, '어': true,
	'아': true, '라': true, '니': true, '데': true, '나': true, '해': true, '져': true,
}

// negationMarks mark a negated predicate.
var negationMarks = []string{"않", "못", "아니", "없"}
