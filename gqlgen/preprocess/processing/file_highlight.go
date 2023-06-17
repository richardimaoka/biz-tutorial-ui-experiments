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
		// possibly numLines = 0, if diffs are interleaving within the same line
		numLines := strings.Count(diff.Text, "\n")

		switch diff.Type {
		case diffmatchpatch.DiffInsert:
			h := fileHighlight{FromLine: currentLine, ToLine: currentLine + numLines}
			highlights = append(highlights, h)
		default:
			//do nothing
		}

		nextLine := currentLine + numLines
		currentLine = nextLine
	}

	fused := fuseHighliths(highlights)
	return fused
}

func fuseHighliths(highlights []fileHighlight) []fileHighlight {
	return highlights
	// var fused []fileHighlight
	// var lastHighlight *fileHighlight
	// for _, h := range highlights {
	// }
}
