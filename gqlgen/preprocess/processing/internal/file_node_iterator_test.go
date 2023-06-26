package internal

import "testing"

func TestFileIterator(t *testing.T) {
	f := NewFileNode("src/components/App.tsx", "", false)
	i := f.IteratorFromRoot()
	if i.node != f {
		t.Fatalf("i.node and f should match")
	}
	if i.CurrentNode().FilePath() != "src" && i.CurrentNode().NodeType() == DirectoryType {
		t.Fatalf("i.node = %+v should %+v", i.CurrentNode(), NewDirNode("src", false))
	}
}

func TestNodeLessFile(t *testing.T) {
	comparisonLetter := func(less bool) string {
		if less {
			return "<"
		} else {
			return ">"
		}
	}

	type Entry struct {
		node1 FileTreeNode
		node2 FileTreeNode
		less  bool
	}

	var entries []Entry = []Entry{
		{
			NewDirNode("public", false),
			NewFileNode("public/favicon.ico", "", false),
			true,
		},
		{
			NewFileNode("public/favicon.ico", "", false),
			NewDirNode("public", false),
			false,
		},
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
		if LessFileNode(e.node1, e.node2) != e.less {
			t.Errorf(
				"%+v %s %+v is expected, but they did not make it",
				e.node1,
				comparisonLetter(e.less),
				e.node2)
		}
	}
}
