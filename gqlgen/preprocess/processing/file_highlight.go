package processing

import (
	"strings"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func CalcHighlight(oldText, newText string) []FileHighlight {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(oldText, newText, true)

	var highlights []FileHighlight
	currentLine := 1
	for _, diff := range diffs {
		// possibly numLines = 0, if diffs are interleaving within the same line
		numLines := strings.Count(diff.Text, "\n")

		switch diff.Type {
		case diffmatchpatch.DiffInsert:
			if strings.HasPrefix(diff.Text, "\n") {
				currentLine++
				numLines--
			}

			highlights = append(highlights,
				FileHighlight{FromLine: currentLine, ToLine: currentLine + numLines},
			)
			nextLine := currentLine + numLines
			currentLine = nextLine
		case diffmatchpatch.DiffEqual:
			nextLine := currentLine + numLines
			currentLine = nextLine
		case diffmatchpatch.DiffDelete:
			if numLines == 0 {
				// diff in the middle of a remaining line
				highlights = append(highlights,
					FileHighlight{FromLine: currentLine, ToLine: currentLine + numLines},
				)
				nextLine := currentLine + numLines
				currentLine = nextLine
			}
		}
	}

	return fuseHighlights(highlights)
}

func fuseHighlights(highlights []FileHighlight) []FileHighlight {
	if len(highlights) == 0 {
		return highlights
	}

	var updated []FileHighlight
	var last *FileHighlight
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

func fuse(h1, h2 FileHighlight) FileHighlight {
	return FileHighlight{FromLine: h1.FromLine, ToLine: h2.ToLine}
}
