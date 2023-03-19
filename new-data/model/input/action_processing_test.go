package input

import "testing"

func TestActionProcessing(t *testing.T) {
	SplitActionListFile("testdata/action_list.json", "testdata", "input")
	t.Fatal("errrrrrrr actionnnnn")
}
