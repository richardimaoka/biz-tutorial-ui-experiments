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
		// 	processing.NewFileNode("public/images/first.png"),
		// 	processing.NewDirNode("src"),
		// 	true,
		// },
		// {
		// 	processing.NewFileNode("src/components/App.tsx"),
		// 	processing.NewFileNode("src/components/index.tsx"),
		// 	true,
		// },
		// {
		// 	processing.NewDirNode("src/components"),
		// 	processing.NewFileNode("src/components/index.tsx"),
		// 	true,
		// },
		// {
		// 	processing.NewDirNode("src/components"),
		// 	processing.NewFileNode("package.json"),
		// 	true,
		// },
		// {
		// 	processing.NewDirNode("src/components"),
		// 	processing.NewFileNode(".gitignore"),
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
