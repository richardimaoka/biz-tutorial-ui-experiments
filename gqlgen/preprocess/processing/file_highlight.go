package processing

import (
	"strings"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func calcHighlight(oldText, newText string) []fileHighlight {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(oldText, newText, true)

	var highlights []fileHighlight
	var h *fileHighlight
	currentLine := 1
	for _, diff := range diffs {
		// possibly numLines = 0, if diffs are interleaving within the same line
		numLines := strings.Count(diff.Text, "\n")
		nextLine := currentLine + numLines

		if diff.Type == diffmatchpatch.DiffInsert {
			if h != nil {
				h.ToLine = currentLine
			} else {
				h = &fileHighlight{FromLine: currentLine, ToLine: currentLine + numLines}
			}
		} else if h != nil {
			if diff.Text != "\n" && nextLine > h.ToLine {
				highlights = append(highlights, *h)
				h = nil
			}
		}

		currentLine = nextLine
	}

	if h != nil {
		highlights = append(highlights, *h)
	}

	return highlights
}
