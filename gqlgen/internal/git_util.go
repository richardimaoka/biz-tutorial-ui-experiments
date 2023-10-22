package internal

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/filemode"
	"github.com/go-git/go-git/v5/plumbing/format/diff"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type Operation int

const (
	// Equal item represents an equals diff.
	Equal Operation = iota
	// Add item represents an insert diff.
	Add
	// Delete item represents a delete diff.
	Delete
)

// Chunk represents a portion of a file transformation into another.
type Chunk struct {
	// Content contains the portion of the file.
	Content string
	// Type contains the Operation to do with this Chunk.
	Type string
}

// FileWithinPatch contains all the file metadata necessary to print some patch formats.
type FileWithinPatch struct {
	// Hash returns the File Hash.
	Hash string
	// Mode returns the FileMode.
	Mode string
	// Path returns the complete Path to the file, including the filename.
	Path string
}

// FilePatch represents the necessary steps to transform one file into another.
type FilePatch struct {
	// IsBinary returns true if this patch is representing a binary file.
	IsBinary bool

	// From and to Files in FilePatch, with all the necessary metadata
	// about them. If the patch creates a new file, "from" will be nil.
	// If the patch deletes a file, "to" will be nil.
	FromFile *FileWithinPatch
	ToFile   *FileWithinPatch

	Type string

	// Chunks represent a slice of ordered changes to transform "from" File into
	// "to" File. If the file is a binary one, Chunks will be empty.
	Chunks []Chunk
}

func toChunk(chunk diff.Chunk) *Chunk {
	var chunkType string
	switch chunk.Type() {
	case diff.Equal:
		chunkType = "Equal"
	case diff.Add:
		chunkType = "Add"
	case diff.Delete:
		chunkType = "Delete"
	}

	return &Chunk{
		Content: chunk.Content(),
		Type:    chunkType,
	}
}

func toFileInPatch(file diff.File) *FileWithinPatch {
	var fileMode string
	switch file.Mode() {
	case filemode.Empty:
		fileMode = "Empty"
	case filemode.Dir:
		fileMode = "Dir"
	case filemode.Regular:
		fileMode = "Regular"
	case filemode.Deprecated:
		fileMode = "Deprecated"
	case filemode.Executable:
		fileMode = "Executable"
	case filemode.Symlink:
		fileMode = "Symlink"
	case filemode.Submodule:
		fileMode = "Submodule"
	}

	return &FileWithinPatch{
		Hash: file.Hash().String(),
		Mode: fileMode,
		Path: file.Path(),
	}
}

func validateCommitHash(hashStr string) (plumbing.Hash, error) {
	commitHash := plumbing.NewHash(hashStr)
	if commitHash.String() != hashStr {
		return plumbing.ZeroHash, fmt.Errorf("commit hash = %s mismatched with re-calculated hash = %s", hashStr, commitHash.String())
	}

	return commitHash, nil
}

func errroMessage(prefix, leadingMessage string, underlyingError error) error {
	return fmt.Errorf("%s - %s, %s", prefix, leadingMessage, underlyingError)
}

// Get git commit object from hash string
// bit easier than go-get's equivalent, as this function works with string, not plumbing.Hash
func GetCommit(repo *git.Repository, hashStr string) (*object.Commit, error) {
	funcName := "internal.GetCommit"

	commitHash, err := validateCommitHash(hashStr)
	if err != nil {
		return nil, errroMessage(funcName, "validation error", err)
	}

	commit, err := repo.CommitObject(commitHash)
	if err != nil {
		return nil, errroMessage(funcName, fmt.Sprintf("cannot get commit for %s", hashStr), err)
	}

	return commit, nil
}

// Get git patch object from hash strings
// bit easier than go-get's equivalent, as this function works with string, not plumbing.Hash
func GetPatch(repo *git.Repository, fromCommitHash, toCommitHash string) (*object.Patch, error) {
	funcName := "internal.GetPatch"

	fromCommit, err := GetCommit(repo, fromCommitHash)
	if err != nil {
		return nil, errroMessage(funcName, fmt.Sprintf("cannot get commit for %s", fromCommitHash), err)
	}

	toCommit, err := GetCommit(repo, toCommitHash)
	if err != nil {
		return nil, errroMessage(funcName, fmt.Sprintf("cannot get commit for %s", toCommitHash), err)
	}

	patch, err := fromCommit.Patch(toCommit)
	if err != nil {
		return nil, errroMessage(funcName, fmt.Sprintf("cannot get patch from = %s to = %s", fromCommitHash, toCommitHash), err)
	}

	return patch, nil
}

func FindFilePatch(filePatches []diff.FilePatch, fileFullPath string) *FilePatch {
	for _, patch := range filePatches {
		// Files returns the from and to Files, with all the necessary metadata
		// about them. If the patch creates a new file, "from" will be nil.
		// If the patch deletes a file, "to" will be nil.
		from, to := patch.Files()

		var patchType string
		if from == nil {
			patchType = "Add"
		} else if to == nil {
			patchType = "Delete"
		} else if from != nil && to != nil {
			patchType = "Update"
		}

		// Even with file rename, there must be only one file/patch matching givne `fileFullPath`
		// (i.e.) within a git commit, there can't be both rename-from and rename-to with the same file name
		if from.Path() == fileFullPath || to.Path() == fileFullPath {
			var chunks []Chunk
			for _, c := range patch.Chunks() {
				chunks = append(chunks, *toChunk(c))
			}

			return &FilePatch{
				Type:     patchType,
				IsBinary: patch.IsBinary(),
				FromFile: toFileInPatch(from),
				ToFile:   toFileInPatch(to),
				Chunks:   chunks,
			}
		}
	}

	return nil
}
