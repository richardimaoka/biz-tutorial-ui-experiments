package internal_test

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/test_util"
)

func TestGitDiff(t *testing.T) {
	repoUrl := "https://github.com/richardimaoka/biz-tutorial-ui-experiments"
	repo := test_util.GitOpenOrClone(t, repoUrl)

	cases := []struct {
		inputFile  string
		fromCommit string
		toCommit   string
		filePath   string
		// goldenFile string
	}{
		{
			"testdata/commit1.json",
			"8446ae73ef2df52a841b49840b776ecfd11751b4",
			"cc43c06d4ada58059d7defb899a7f191e012555b",
			"next/app/components/sourcecode2/editor/internal/EditorBare.tsx",
		},
	}

	for _, c := range cases {
		t.Run(c.inputFile, func(t *testing.T) {
			patch, err := internal.GetPatch(repo, c.fromCommit, c.toCommit)
			if err != nil {
				t.Fatal(err)
			}

			filePatch := internal.FindFilePatch(patch.FilePatches(), c.filePath)
			internal.CompareWitGoldenFile(t, *updateFlag, c.inputFile, filePatch)
		})
	}
}
