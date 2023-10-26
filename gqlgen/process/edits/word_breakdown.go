package edits

import (
	"strings"
	"unicode/utf8"
)

const (
	BREAKDOWN_TO_CHARACTERS = "each-character"
	BREAKDOWN_TO_WORDS      = "word-by-word"
	BREAKDOWN_TO_WHOLE_LINE = "whole-line"
)

func breakdownLineToAdd(lineToAdd string) []string {
	cond := condition(lineToAdd)
	switch cond {
	case BREAKDOWN_TO_CHARACTERS:
		return breakdownToCharacters(lineToAdd)
	case BREAKDOWN_TO_WORDS:
		return breakdownToWords(lineToAdd)
	case BREAKDOWN_TO_WHOLE_LINE:
		return breakdownToWholeLine(lineToAdd)
	default:
		return breakdownToWholeLine(lineToAdd)
	}
}

func breakdownToCharacters(toAdd string) []string {
	if toAdd == "" {
		return []string{}
	}

	var additions []string

	// From プログラミング言語Go chap. 3, p. 78
	// since a UTF-8 character varies in its size, we need to use rune (via utf8.DecodeRuneInString)
	// to extract each character. (e.g.) Source code comments can have multi-byte characters.
	for i := 0; i < len(toAdd); {
		r, size := utf8.DecodeRuneInString(toAdd[i:])
		additions = append(additions, string(r))
		i += size
	}

	return additions
}

func breakdownToWords(toAdd string) []string {
	if toAdd == "" {
		return []string{}
	}

	return strings.SplitAfter(toAdd, " ")
}

func breakdownToWholeLine(toAdd string) []string {
	if toAdd == "" {
		return []string{}
	}

	return []string{toAdd}
}

func condition(toAdd string) string {
	length := len(toAdd)
	if length < 10 {
		return BREAKDOWN_TO_CHARACTERS
	} else if length < 100 {
		return BREAKDOWN_TO_WORDS
	} else {
		return BREAKDOWN_TO_WHOLE_LINE
	}
}
