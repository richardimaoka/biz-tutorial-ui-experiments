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
