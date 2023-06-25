package internal_test

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/preprocess/processing/internal"
)

func comparisonLetter(less bool) string {
	if less {
		return "<"
	} else {
		return ">"
	}
}
func TestNodeLessFile(t *testing.T) {
	type Entry struct {
		node1 internal.FileTreeNode
		node2 internal.FileTreeNode
		less  bool
	}

	var entries []Entry = []Entry{
		// {
		// 	processing.NewFileProcessorNode("public/images/first.png"),
		// 	processing.NewDirectoryProcessorNode("src"),
		// 	true,
		// },
		// {
		// 	processing.NewFileProcessorNode("src/components/App.tsx"),
		// 	processing.NewFileProcessorNode("src/components/index.tsx"),
		// 	true,
		// },
		// {
		// 	processing.NewDirectoryProcessorNode("src/components"),
		// 	processing.NewFileProcessorNode("src/components/index.tsx"),
		// 	true,
		// },
		// {
		// 	processing.NewDirectoryProcessorNode("src/components"),
		// 	processing.NewFileProcessorNode("package.json"),
		// 	true,
		// },
		// {
		// 	processing.NewDirectoryProcessorNode("src/components"),
		// 	processing.NewFileProcessorNode(".gitignore"),
		// 	true,
		// },
	}

	for _, e := range entries {
		if internal.LessFileNode(e.node1, e.node2) != e.less {
			t.Errorf(
				"%+v %s %+v is expected, but they did not make it",
				e.node1,
				comparisonLetter(e.less),
				e.node2)
		}
	}
}
