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
