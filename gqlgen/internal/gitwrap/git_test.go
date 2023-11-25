package gitwrap_test

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/gitwrap"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/testio"
)

func TestGitDiff(t *testing.T) {
	repoUrl := "https://github.com/richardimaoka/biz-tutorial-ui-experiments"
	repo := testio.GitOpenOrClone(t, repoUrl)

	cases := []struct {
		inputFile  string
		fromCommit string
		toCommit   string
		filePath   string
		// goldenFile string
	}{
		{
			"testdata/commit1.json",
			"2f551fc2d64dc17b590388dd04c3774869044eb8",
			"8446ae73ef2df52a841b49840b776ecfd11751b4",
			"next/app/components/sourcecode2/editor/internal/EditorBare.tsx",
		},
	}

	for _, c := range cases {
		t.Run(c.inputFile, func(t *testing.T) {
			patch, err := gitwrap.GetPatch(repo, c.fromCommit, c.toCommit)
			if err != nil {
				t.Fatal(err)
			}

			diffFilePatch := gitwrap.FindFilePatch(patch, c.filePath)
			if diffFilePatch == nil {
				t.Fatal("diffFilePatch is nil")
			}
			result := gitwrap.ToFilePatch(diffFilePatch)
			testio.CompareWithGoldenFile(t, *updateFlag, c.inputFile, result)
		})
	}
}
