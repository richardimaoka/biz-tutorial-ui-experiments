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

	return fuseHighlights(highlights)
}

func fuseHighlights(highlights []fileHighlight) []fileHighlight {
	var updated []fileHighlight
	var last *fileHighlight
	for i := 1; i < len(highlights); i++ {
		if last == nil {
			last = &highlights[i-1]
		}

		current := highlights[i]
		if last.ToLine == current.FromLine || last.ToLine+1 == current.FromLine {
			fused := fuse(*last, current)
			last = &fused
		} else {
			updated = append(updated, *last)
			last = nil
		}
	}

	if last != nil {
		updated = append(updated, *last)
	} else {
		updated = append(updated, highlights[len(highlights)-1])
	}

	return updated
}

func fuse(h1, h2 fileHighlight) fileHighlight {
	return fileHighlight{FromLine: h1.FromLine, ToLine: h2.ToLine}
}
