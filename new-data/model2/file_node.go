package model2

import "strings"

func (f *FileNode) FilePathString() string {
	var s []string
	for _, v := range f.FilePath {
		s = append(s, *v)
	}
	return strings.Join(s, "/")
}
