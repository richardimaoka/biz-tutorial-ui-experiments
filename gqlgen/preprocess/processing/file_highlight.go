package processing

import (
	"strings"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func calcHighlight(oldText, newText string) []fileHighlight {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(oldText, newText, true)

	var highlights []fileHighlight
	currentLine := 1
	for _, diff := range diffs {
		switch diff.Type {
		case diffmatchpatch.DiffInsert:
			fromLine := currentLine
			currentLine += strings.Count(diff.Text, "\n")
			toLine := currentLine
			highlights = append(highlights, fileHighlight{FromLine: fromLine, ToLine: toLine})
		default:
			currentLine += strings.Count(diff.Text, "\n")
		}
	}

	return highlights
}
