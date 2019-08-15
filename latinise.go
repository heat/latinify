package latinify

import (
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

//Table unicode mapped table
type Table map[rune]rune

//String convert any unicode souce in a Latin compatible text
func String(source string) (string, error) {
	return StringWithTable(source, CyrillicTable, GreekTable)
}

type removeMark struct{}

func (m removeMark) Contains(r rune) bool {

	return unicode.Is(unicode.Mn, r)
}

type mapping func(rune) rune

type mappedTable struct {
	mapping
	tables []Table
}

func replaceTable(table Table) mapping {

	return func(w rune) rune {
		m, ok := table[w]
		if !ok {
			return w
		}
		return m
	}
}

//StringWithTable convert any unicode source in a LAtin compatible text using Tables conversion
func StringWithTable(source string, tables ...Table) (string, error) {

	nfkd := norm.NFKD

	r := runes.Remove(removeMark{})

	mappers := make([]transform.Transformer, len(tables))

	for idx, table := range tables {
		mappers[idx] = runes.Map(replaceTable(table))
	}

	transformers := []transform.Transformer{
		nfkd,
		r,
	}

	transformers = append(transformers, mappers...)
	// decomposity -> remove marks -> map tables
	t := transform.Chain(transformers...)

	res, _, err := transform.String(t, source)
	return res, err
}
