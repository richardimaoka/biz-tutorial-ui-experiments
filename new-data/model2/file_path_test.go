package model2

import "testing"

func comparisonLetter(less bool) string {
	if less {
		return "<"
	} else {
		return ">"
	}
}

func TestFilePathLessFile(t *testing.T) {
	type Entry struct {
		path1 string
		path2 string
		less  bool
	}

	var entries []Entry = []Entry{
		{
			path1: "aaa",
			path2: "aaab",
			less:  true,
		},
		{
			path1: "aaab",
			path2: "aaa",
			less:  false,
		},
		{
			path1: "aaa/bbb.txt",
			path2: "bbb/aaa.txt",
			less:  true},
		{
			path1: "aaa",
			path2: "aaa/abc.txt",
			less:  true},
		{
			path1: "aaa",
			path2: "aa/child.txt",
			less:  false},
		{
			path1: "aba",
			path2: "aaa.txt",
			less:  false},
		{
			path1: "aaaa/abcde/a.txt",
			path2: "aaaa/abcde",
			less:  false,
		},
	}

	for _, e := range entries {
		if lessFilePath2(e.path1, e.path2) != e.less {
			t.Errorf(
				"%s %s %s is expected, but they did not make it",
				e.path1,
				comparisonLetter(e.less),
				e.path2)
		}
	}
}
