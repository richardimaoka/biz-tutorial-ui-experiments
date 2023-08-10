package state

import (
	"strings"

	"github.com/go-git/go-git/v5/plumbing/format/diff"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type FileHighlight struct {
	// Uppercase exported fields in lowercase unexported struct, as exported fields are necessary for reflection-based testing
	FromLine int
	ToLine   int
}

func fuse(h1, h2 FileHighlight) FileHighlight {
	return FileHighlight{FromLine: h1.FromLine, ToLine: h2.ToLine}
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

func CalcHighlight(patch diff.FilePatch) []FileHighlight {
	var highlights []FileHighlight
	currentLine := 1

	chunks := patch.Chunks()
	for _, chunk := range chunks {
		contents := chunk.Content()

		switch chunk.Type() {
		case diff.Add:
			// possibly numNewLines = 0, if diffs are interleaving within the same line
			numNewLines := strings.Count(contents, "\n")

			if strings.HasPrefix(contents, "\n") && !strings.HasPrefix(contents, "\n\n") {
				currentLine++
				numNewLines--
			}

			if strings.HasSuffix(contents, "\n") {
				highlights = append(highlights,
					FileHighlight{FromLine: currentLine, ToLine: currentLine + numNewLines - 1})
			} else {
				highlights = append(highlights,
					FileHighlight{FromLine: currentLine, ToLine: currentLine + numNewLines})
			}
			currentLine = currentLine + numNewLines
		case diff.Equal:
			// possibly numNewLines = 0, if diffs are interleaving within the same line
			numNewLines := strings.Count(contents, "\n")

			if strings.HasPrefix(contents, "\n") {
				currentLine++
				numNewLines--
			}

			currentLine = currentLine + numNewLines
		case diff.Delete:
			// if numLines == 0 {
			// 	// diff in the middle of a remaining line
			// 	highlights = append(highlights,
			// 		FileHighlight{FromLine: currentLine, ToLine: currentLine + numLines},
			// 	)
			// 	nextLine := currentLine + numLines
			// 	currentLine = nextLine
			// }
		}
	}

	return fuseHighlights(highlights)
}

func (h *FileHighlight) ToGraphQLFileHighlight() *model.FileHighlight {
	fromLine := h.FromLine // copy to avoid mutation effect afterwards
	toLine := h.ToLine     // copy to avoid mutation effect afterwards

	return &model.FileHighlight{
		FromLine: &fromLine,
		ToLine:   &toLine,
	}
}
