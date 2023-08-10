package state_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

func TestFileHighlight(t *testing.T) {

	cases := []struct {
		repoUrl       string
		prevCommit    string
		currentCommit string
		filePath      string
		expected      []state.FileHighlight
	}{
		{
			"https://github.com/richardimaoka/file-highlight-test.git",
			"f1df093152800852cf892d4015ff56a56427716b",
			"cf3bc8ae215607bd18d50c72a48868bc4f2b5e49",
			"1.txt",
			[]state.FileHighlight{
				{FromLine: 2, ToLine: 3},
				{FromLine: 5, ToLine: 5},
			},
		},
		{
			"https://github.com/richardimaoka/file-highlight-test.git",
			"cf3bc8ae215607bd18d50c72a48868bc4f2b5e49",
			"692fdc4b925a44517baccd16bfc1c0812d2d0ec4",
			"1.txt",
			[]state.FileHighlight{
				{FromLine: 3, ToLine: 3},
				{FromLine: 5, ToLine: 7},
			},
		},
		{
			"https://github.com/richardimaoka/file-highlight-test.git",
			"692fdc4b925a44517baccd16bfc1c0812d2d0ec4",
			"6586fedbc42864e18a9bbafe4dae9e9b335d9e90",
			"1.txt",
			[]state.FileHighlight{
				{FromLine: 1, ToLine: 3},
			},
		},
		{
			"https://github.com/richardimaoka/file-highlight-test.git",
			"694e06fbd51903f32f6611d21a285f2818f6bb6f",
			"c7ef279eb4afc96b2ff78af8a3370d3cb7595dbb",
			"1.txt",
			[]state.FileHighlight{
				{FromLine: 1, ToLine: 7},
			},
		},
		{
			"https://github.com/richardimaoka/file-highlight-test.git",
			"db1abbb5cb7cd8c429434a57bc2b4ddfd5ea2976",
			"ae11689775ecc97bc10855862ee46bb48da8b1bd",
			"1.txt",
			[]state.FileHighlight{
				{FromLine: 1, ToLine: 2},
				{FromLine: 4, ToLine: 6},
			},
		},
		{
			"https://github.com/richardimaoka/file-highlight-test.git",
			"ae11689775ecc97bc10855862ee46bb48da8b1bd",
			"9766d3f67663d99db4e2d40f2278805d0c660bea",
			"1.txt",
			[]state.FileHighlight{
				{FromLine: 1, ToLine: 2},
				{FromLine: 4, ToLine: 4},
			},
		},
		{
			"https://github.com/richardimaoka/file-highlight-test.git",
			"9766d3f67663d99db4e2d40f2278805d0c660bea",
			"8a0772e35e57f4d1fe2cfc0b4d7f2a19f3d13db1",
			"1.txt",
			[]state.FileHighlight{
				{FromLine: 2, ToLine: 8},
			},
		},
		{
			"https://github.com/richardimaoka/file-highlight-test.git",
			"8a0772e35e57f4d1fe2cfc0b4d7f2a19f3d13db1",
			"4d7c3cdb3dd44e501b84c0ac9b0cbd23a9dbe600",
			"1.txt",
			[]state.FileHighlight{
				{FromLine: 7, ToLine: 10},
			},
		},
		{
			"https://github.com/richardimaoka/file-highlight-test.git",
			"1d143392b7e6553d180bba579fb26ed75c55c100",
			"4b8c45ccadad77d1d6ebdc0cbf8a417fa7e93e24",
			"1.txt",
			[]state.FileHighlight{
				{FromLine: 3, ToLine: 5},
			},
		},
		{
			"https://github.com/richardimaoka/file-highlight-test.git",
			"6ed2af145624da2682e691960ff3f07bbe5dc85e",
			"b3beb34fd96adbc1bde65609ce65e3a9f69b79fc",
			"6.g0",
			[]state.FileHighlight{
				{FromLine: 25, ToLine: 25},
				{FromLine: 67, ToLine: 68},
				{FromLine: 71, ToLine: 72},
			},
		},
		{
			"https://github.com/richardimaoka/file-highlight-test.git",
			"e2be05446a2c01b108ece7d8ec69354bf14f5393",
			"af63b0a71cba4dd14e2b7e5c46edb42236a4d7b3",
			"go.mod",
			[]state.FileHighlight{
				{FromLine: 4, ToLine: 19},
			},
		},

		//
		// sign-in-with google test cases
		//
		{
			"https://github.com/richardimaoka/sign-in-with-google-experiment.git",
			"80853dd7ab95987fef9255cbfa471cc645adefa5",
			"e02fc71d2481ed3a507926a43244e56a2d7c14c2",
			".gitignore",
			[]state.FileHighlight{
				{FromLine: 1, ToLine: 1},
			},
		},
	}

	for _, c := range cases {
		prevCommit, err := gitCommit(c.repoUrl, c.prevCommit)
		if err != nil {
			t.Fatalf("failed in TestFileHighlight to get prev commit, %s", err)
		}

		currentCommit, err := gitCommit(c.repoUrl, c.currentCommit)
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
