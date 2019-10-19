package latinify

import (
	"regexp"
	"strings"
)

type stringTransform func(string) string

func apply(funcs ...stringTransform) func(string) string {

	return func(src string) (res string) {
		res = src
		for _, f := range funcs {
			res = f(res)
		}
		return
	}
}

// Slugify transform a source folder in a no spacede ascii string
func Slugify(source string) (string, error) {
	latin, err := String(source)
	if err != nil {
		return "", err
	}

	spacers := regexp.MustCompile(`\s+`)
	notAWord := regexp.MustCompile(`[^a-zA-Z0-9\-_]+`)
	transformers := apply(
		strings.ToLower,
		strings.TrimSpace,
		func(src string) string {
			return spacers.ReplaceAllString(src, "-")
		},
		func(src string) string {
			return notAWord.ReplaceAllString(src, "")
		},
	)

	res := transformers(latin)
	return res, nil
}
