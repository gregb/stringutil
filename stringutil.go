package stringutil

import "unicode"

const Digits = "01234567890"
const Alpha = "ABCDEFGHIJKLMNOPQSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const AlphaNum = "ABCDEFGHIJKLMNOPQSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
const AlphaNum_ = "ABCDEFGHIJKLMNOPQSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"
const Punctuation = `~!@#$%^&*()_+-={}|[]\;':"<>?,./`
const NonAlphaNum = `~!@#$%^&*()_+-={}|[]\;':"<>?,./ `

type runeSet map[rune]struct{}

var runeSetCache = make(map[string]runeSet)

func Keep(source, keep string) string {
	return filter(source, keep, true)
}

func Strip(source, strip string) string {
	return filter(source, strip, false)
}

func filter(source, targetRunes string, present bool) string {
	runes, ok := runeSetCache[targetRunes]

	if !ok {
		runes = make(map[rune]struct{})
		for _, r := range targetRunes {
			runes[r] = *new(struct{})
		}

		runeSetCache[targetRunes] = runes
	}

	filtered := make([]rune, 0, len(source))

	for _, r := range source {
		_, ok := runes[r]

		if present == ok {
			filtered = append(filtered, r)
		}
	}

	return string(filtered)
}

func CamelCaseToUnderscored(camelCase string) string {

	sizeEstimate := len(camelCase) + len(camelCase)/8

	underscored := make([]rune, 0, sizeEstimate)

	for index, char := range camelCase {
		if unicode.IsUpper(char) {

			if index > 0 {
				underscored = append(underscored, '_')
			}

			underscored = append(underscored, unicode.ToLower(char))
		} else {
			underscored = append(underscored, char)
		}
	}

	return string(underscored)
}
