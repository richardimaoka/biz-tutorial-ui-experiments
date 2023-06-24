package processing

type fileNodeIterator struct {
	node          FileTreeNode
	currentOffset int
}

func newFileNodeIterator(node FileTreeNode) *fileNodeIterator {
	return &fileNodeIterator{
		node:          node,
		currentOffset: 0,
	}
}

// // return the next iterator, and a bool value indicating whether the next exists
// func (i *fileNodeIterator) Node() FileTreeNode {
// 	if i.currentOffset < len(i.node.FilePath())-1 {
// 		split := strings.Split(i.node.FilePath(), "/")
// 		split[:i.currentOffset]
// 		return i.node.Children()[i.currentOffset]
// 	} else {
// 		return nil
// 	}
// }

// // return the next iterator, and a bool value indicating whether the next exists
// func (i *fileNodeIterator) Next() (fileNodeIterator, bool) {
// }
