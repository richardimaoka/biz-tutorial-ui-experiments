package state_test

import (
	"fmt"
	"testing"

	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/filemode"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

func gitTreeForCommit(repoUrl, commitHashStr, dirPath string) (*object.Tree, error) {
	repo, err := gitOpenOrClone(repoUrl)
	if err != nil {
		return nil, fmt.Errorf("failed in gitFileFromCommit, %s", err)
	}

	commitHash := plumbing.NewHash(commitHashStr)
	if commitHash.String() != commitHashStr {
		return nil, fmt.Errorf("failed in gitFileFromCommit, commit hash = %s is invalid as its re-calculated hash is mismatched = %s", commitHashStr, commitHash.String())
	}

	commit, err := repo.CommitObject(commitHash)
	if err != nil {
		return nil, fmt.Errorf("failed in gitFileFromCommit, cannot get commit = %s, %s", commitHashStr, err)
	}

	rootTree, err := commit.Tree()
	if err != nil {
		return nil, fmt.Errorf("failed in gitFileFromCommit, cannot get tree for commit = %s, %s", commitHashStr, err)

	}

	gitTree, err := rootTree.Tree(dirPath)
	if err != nil {
		return nil, fmt.Errorf("failed in gitFileFromCommit, cannot get tree = %s in commit = %s, %s", dirPath, commitHashStr, err)
	}

	return gitTree, nil
}

func TestSortTreeEntries(t *testing.T) {
	cases := []struct {
		entries  []object.TreeEntry
		expected []object.TreeEntry
	}{
		{
			// files to sort
			[]object.TreeEntry{
				{Mode: filemode.Regular, Hash: plumbing.ZeroHash, Name: "package-lock.json"},
				{Mode: filemode.Regular, Hash: plumbing.ZeroHash, Name: "package.json"},
				{Mode: filemode.Regular, Hash: plumbing.ZeroHash, Name: "next.config.js"},
				{Mode: filemode.Regular, Hash: plumbing.ZeroHash, Name: ".gitignore"},
				{Mode: filemode.Regular, Hash: plumbing.ZeroHash, Name: "README.md"},
				{Mode: filemode.Regular, Hash: plumbing.ZeroHash, Name: ".eslintrc.json"},
				{Mode: filemode.Regular, Hash: plumbing.ZeroHash, Name: "tsconfig.json"},
			},
			[]object.TreeEntry{
				{Mode: filemode.Regular, Hash: plumbing.ZeroHash, Name: ".eslintrc.json"},
				{Mode: filemode.Regular, Hash: plumbing.ZeroHash, Name: ".gitignore"},
				{Mode: filemode.Regular, Hash: plumbing.ZeroHash, Name: "next.config.js"},
				{Mode: filemode.Regular, Hash: plumbing.ZeroHash, Name: "package-lock.json"},
				{Mode: filemode.Regular, Hash: plumbing.ZeroHash, Name: "package.json"},
				{Mode: filemode.Regular, Hash: plumbing.ZeroHash, Name: "README.md"},
				{Mode: filemode.Regular, Hash: plumbing.ZeroHash, Name: "tsconfig.json"},
			},
		},
		{
			// dirs to sort
			[]object.TreeEntry{
				{Mode: filemode.Dir, Hash: plumbing.ZeroHash, Name: "styles"},
				{Mode: filemode.Dir, Hash: plumbing.ZeroHash, Name: "public"},
				{Mode: filemode.Dir, Hash: plumbing.ZeroHash, Name: "pages"},
			},
			[]object.TreeEntry{
				{Mode: filemode.Dir, Hash: plumbing.ZeroHash, Name: "pages"},
				{Mode: filemode.Dir, Hash: plumbing.ZeroHash, Name: "public"},
				{Mode: filemode.Dir, Hash: plumbing.ZeroHash, Name: "styles"},
			},
		},
	}

	for _, c := range cases {
		if len(c.entries) != len(c.expected) {
			t.Fatalf("len(entries) = %d mismatched with len(expected) = %d", len(c.entries), len(c.expected))
		}

		state.SortEntries(c.entries)

		for i, exp := range c.expected {
			if exp.Name != c.entries[i].Name {
				t.Errorf("[%d] expected = %s != result = %s", i, exp.Name, c.entries[i].Name)
			}
		}
	}
}

func TestTreeFilesDirs(t *testing.T) {
	tree, err := gitTreeForCommit(
		"https://github.com/richardimaoka/next-sandbox.git",
		"8adac375628219e020d4b5957ff24f45954cbd3f", // commit = 'npx create-next-app@latest'
		"next",
	)
	if err != nil {
		t.Fatalf("failed in TestTreeDirectories to get git tree, %s", err)
	}

	files, dirs := state.TreeFilesDirs(tree)
	expectedFiles := []string{
		".eslintrc.json",
		".gitignore",
		"README.md",
		"next.config.js",
		"package-lock.json",
		"package.json",
		"tsconfig.json",
	}
	expectedDirs := []string{
		"pages",
		"public",
		"styles",
	}

	if len(files) != len(expectedFiles) {
		t.Fatalf("len(files) = %d mismatched with len(expectedFiles) = %d", len(files), len(expectedFiles))
	}

	for i, f := range files {
		if f.Name != expectedFiles[i] {
			t.Errorf("files[%d] = %s mismatched with expectedFiles[%d] = %s", i, f.Name, i, expectedFiles[i])
		}
	}

	if len(dirs) != len(expectedDirs) {
		t.Fatalf("len(dirs) = %d mismatched with len(expectedDirs) = %d", len(dirs), len(expectedDirs))
	}

	for i, d := range dirs {
		if d.Name != expectedDirs[i] {
			t.Errorf("dirs[%d] = %s mismatched with expectedDirs[%d] = %s", i, d.Name, i, expectedDirs[i])
		}
	}
}

func TestSourceCodePatterns(t *testing.T) {
	repoUrl := "https://github.com/richardimaoka/next-sandbox.git"
	repo, err := gitOpenOrClone(repoUrl)
	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		prevCommit    string
		currentCommit string
		goldenFile    string
	}{
		{
			"55c98498a85f4503e3922586ceeb86ab5100e91f", //cleanup (delete all files)
			"8adac375628219e020d4b5957ff24f45954cbd3f", //npx create-next-app@latest
			"testdata/source_code_golden.json",
		},
	}

	for _, c := range cases {
		sc, err := state.NewSourceCode(repo, c.currentCommit, c.prevCommit)
		if err != nil {
			t.Fatal(err)
		}

		internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFile, sc.ToGraphQLSourceCode())
	}
}
