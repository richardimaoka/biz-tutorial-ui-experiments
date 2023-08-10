package state_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

func TestFileHighlight(t *testing.T) {
	repoUrl := "https://github.com/richardimaoka/file-highlight-test.git"

	cases := []struct {
		prevCommit    string
		currentCommit string
		filePath      string
		expected      []state.FileHighlight
	}{{
		"f1df093152800852cf892d4015ff56a56427716b",
		"cf3bc8ae215607bd18d50c72a48868bc4f2b5e49",
		"1.txt",
		[]state.FileHighlight{
			{FromLine: 2, ToLine: 3},
			{FromLine: 5, ToLine: 5},
		},

		// "cf3bc8ae215607bd18d50c72a48868bc4f2b5e49",
		// "692fdc4b925a44517baccd16bfc1c0812d2d0ec4",

		// "692fdc4b925a44517baccd16bfc1c0812d2d0ec4",
		// "6586fedbc42864e18a9bbafe4dae9e9b335d9e90",

		// "fb22fd0aabd68156911a16a653095281b0ba11c4",
		// "694e06fbd51903f32f6611d21a285f2818f6bb6f",

		// "694e06fbd51903f32f6611d21a285f2818f6bb6f",
		// "c7ef279eb4afc96b2ff78af8a3370d3cb7595dbb",

		// "c7ef279eb4afc96b2ff78af8a3370d3cb7595dbb",
		// "db1abbb5cb7cd8c429434a57bc2b4ddfd5ea2976",

		// "db1abbb5cb7cd8c429434a57bc2b4ddfd5ea2976",
		// "ae11689775ecc97bc10855862ee46bb48da8b1bd",

		// "ae11689775ecc97bc10855862ee46bb48da8b1bd",
		// "9766d3f67663d99db4e2d40f2278805d0c660bea",

		// "9766d3f67663d99db4e2d40f2278805d0c660bea",
		// "8a0772e35e57f4d1fe2cfc0b4d7f2a19f3d13db1",

		// "8a0772e35e57f4d1fe2cfc0b4d7f2a19f3d13db1",
		// "4d7c3cdb3dd44e501b84c0ac9b0cbd23a9dbe600",

		// "4d7c3cdb3dd44e501b84c0ac9b0cbd23a9dbe600",
		// "1d143392b7e6553d180bba579fb26ed75c55c100",

		// "1d143392b7e6553d180bba579fb26ed75c55c100",
		// "4b8c45ccadad77d1d6ebdc0cbf8a417fa7e93e24",

		// "4b8c45ccadad77d1d6ebdc0cbf8a417fa7e93e24",
		// "6ed2af145624da2682e691960ff3f07bbe5dc85e",

		// "6ed2af145624da2682e691960ff3f07bbe5dc85e",
		// "b3beb34fd96adbc1bde65609ce65e3a9f69b79fc",

		// "b3beb34fd96adbc1bde65609ce65e3a9f69b79fc",
		// "e2be05446a2c01b108ece7d8ec69354bf14f5393",

		// "e2be05446a2c01b108ece7d8ec69354bf14f5393",
		// "af63b0a71cba4dd14e2b7e5c46edb42236a4d7b3",
		// }, {
		// 	before: "testdata/file_highlight_golden1-2.txt", after: "testdata/file_highlight_golden1-3.txt", expected: []state.FileHighlight{
		// 		{FromLine: 3, ToLine: 3},
		// 		{FromLine: 5, ToLine: 7},
		// 	},
		// }, {
		// 	before: "testdata/file_highlight_golden1-3.txt", after: "testdata/file_highlight_golden1-4.txt", expected: []state.FileHighlight{
		// 		{FromLine: 1, ToLine: 3},
		// 	},
		// }, {
		// 	before: "testdata/file_highlight_golden2-1.txt", after: "testdata/file_highlight_golden2-2.txt", expected: []state.FileHighlight{
		// 		{FromLine: 1, ToLine: 4},
		// 		{FromLine: 6, ToLine: 7},
		// 	},
		// }, {
		// 	before: "testdata/file_highlight_golden3-1.txt", after: "testdata/file_highlight_golden3-2.txt", expected: []state.FileHighlight{
		// 		{FromLine: 1, ToLine: 2},
		// 		{FromLine: 4, ToLine: 6},
		// 	},
		// }, {
		// 	before: "testdata/file_highlight_golden4-1.txt", after: "testdata/file_highlight_golden4-2.txt", expected: []state.FileHighlight{
		// 		{FromLine: 2, ToLine: 8},
		// 	},
		// }, {
		// 	before: "testdata/file_highlight_golden4-2.txt", after: "testdata/file_highlight_golden4-3.txt", expected: []state.FileHighlight{
		// 		{FromLine: 7, ToLine: 10},
		// 	},
		// }, {
		// 	before: "testdata/file_highlight_golden5-1.txt", after: "testdata/file_highlight_golden5-2.txt", expected: []state.FileHighlight{
		// 		{FromLine: 3, ToLine: 5},
		// 	},
		// }, {
		// 	before: "testdata/file_highlight_golden6-1.go", after: "testdata/file_highlight_golden6-2.go", expected: []state.FileHighlight{
		// 		{FromLine: 25, ToLine: 25},
		// 		{FromLine: 67, ToLine: 68},
		// 		{FromLine: 71, ToLine: 72},
		// 	},
		// }, {
		// 	before: "testdata/file_highlight_golden7-1.go.mod", after: "testdata/file_highlight_golden7-2.go.mod", expected: []state.FileHighlight{
		// 		{FromLine: 4, ToLine: 20},
		// 	},
	}}

	for _, c := range cases {
		prevCommit, err := gitCommit(repoUrl, c.prevCommit)
		if err != nil {
			t.Fatalf("failed in TestFileHighlight to get prev commit, %s", err)
		}

		currentCommit, err := gitCommit(repoUrl, c.currentCommit)
		if err != nil {
			t.Fatalf("failed in TestFileHighlight to get current commit, %s", err)
		}

		compared := false
		patch, _ := prevCommit.Patch(currentCommit)
		for _, p := range patch.FilePatches() {
			_, to := p.Files()
			if to.Path() == c.filePath {
				results := state.CalcHighlight(p)
				if diff := cmp.Diff(c.expected, results); diff != "" {
					t.Fatalf("mismatch (-expected +result):\n%s", diff)
				}
				compared = true
			}
		}

		if !compared {
			t.Fatalf("failed in TestFileHighlight, no patch found for %s in prev = %s and current  = %s", c.filePath, c.prevCommit, c.currentCommit)
		}
	}
}
