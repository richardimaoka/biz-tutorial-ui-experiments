package internal

import "strings"

type FileNodeIterator struct {
	node          FileTreeNode
	currentOffset int
}

func (n *FileProcessorNode) IteratorFromRoot() *FileNodeIterator {
	return &FileNodeIterator{
		node:          n,
		currentOffset: 0,
	}
}

func (n *DirectoryProcessorNode) IteratorFromRoot() *FileNodeIterator {
	return &FileNodeIterator{
		node:          n,
		currentOffset: 0,
	}
}

// return the next iterator, and a bool value indicating whether the next exists
func (i *FileNodeIterator) CurrentNode() FileTreeNode {
	split := strings.Split(i.node.FilePath(), "/")
	if i.currentOffset == len(split)-1 {
		// leaf node
		return i.node
	}

	numElements := i.currentOffset + 1
	filePath := strings.Join(split[:numElements], "/")
	node := NewDirNode(filePath, false)

	return node
}

// return the next iterator, and a bool value indicating whether the next exists
func (i *FileNodeIterator) Next() (*FileNodeIterator, bool) {
	split := strings.Split(i.node.FilePath(), "/")

	if i.currentOffset == len(split)-1 {
		//currentOffset is at the leaf node
		return nil, false
	}

	return &FileNodeIterator{node: i.node, currentOffset: i.currentOffset + 1}, true
}

func lessFilePath2(aSplitPath, bSplitPath []string) bool {
	a := aSplitPath[0] //supposedly len(aSplitPath) > 0
	b := bSplitPath[0] //supposedly len(bSplitPath) > 0

	if a == b {
		if len(aSplitPath) == 1 {
			// (e.g.)
			//   aSplitPath = ["src", "components", "shelf"]
			//   bSplitPath = ["src", "components", "books", "BookView.tsx"]
			// no more path part to compare, then bSplitPath is "less"
			return true
		} else if len(bSplitPath) == 1 {
			// (e.g.)
			//   aSplitPath = ["src", "components", "books", "BookTab.tsx"]
			//   bSplitPath = ["src", "components"]
			return false
		}

		// more path parts to compare in both aSplitPath and bSplitPath
		return lessFilePath2(aSplitPath[1:], bSplitPath[1:])
	} else {
		return a < b
	}
}

func lessFileNodeIter(aIter, bIter *FileNodeIterator) bool {
	// nodes at current iterator's offset
	aNode := aIter.CurrentNode()
	bNode := bIter.CurrentNode()

	if aNode.FilePath() == bNode.FilePath() {
		// (e.g.)
		//   a = src/components/books/BookView.tsx
		//   b = src/components
		//   current offset = 1, (i.e.) "components"
		aNext, aNextExists := aIter.Next()
		bNext, bNextExists := bIter.Next()

		if !aNextExists {
			// supposedly bNextExists == true
			// so, aNode is a parent directory containing bNode
			return false //
		} else if !bNextExists {
			// supposedly aNextExists == true
			// so, bNode is a parent directory containing aNode
			return true
		}

		// supposedly aNextExists == true,  and bNextExists == true
		//           (aNextExists == false, and bNextExists == false should not happen, which means comparing identical (i.e.) duplicated files or dirs)
		return lessFileNodeIter(aNext, bNext)
	} else {
		// aNode.FilePath() != bNode.FilePath()
		//
		// if current iterator's offset has different file paths for a and b:
		// (e.g.)
		//   a = src/components/books/BookView.tsx
		//   b = package.json
		//   current offset = 0, (i.e.) "src" vs. "package.json"

		if aNode.NodeType() == DirectoryType {
			if bNode.NodeType() == DirectoryType {
				return aNode.FilePath() < bNode.FilePath()
			} else {
				// supposedly bNode.NodeType() == FileType
				// directory is always less than file
				return true // a < b
			}
		} else {
			// supposedly aNode.NodeType == FileType

			if bNode.NodeType() == DirectoryType {
				// directory is always less than file
				return false // a > b
			} else {
				// supposedly bNode.NodeType() == FileType
				return aNode.FilePath() < bNode.FilePath()
			}
		}
	}

}

func LessFileNode(a, b FileTreeNode) bool {
	aIter := a.IteratorFromRoot()
	bIter := b.IteratorFromRoot()
	return lessFileNodeIter(aIter, bIter)
}
