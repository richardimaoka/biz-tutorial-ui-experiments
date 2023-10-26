package internal

import (
	"github.com/go-git/go-git/v5/plumbing/filemode"
	"github.com/go-git/go-git/v5/plumbing/format/diff"
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

func toChunk(chunk diff.Chunk) Chunk {
	var chunkType string
	switch chunk.Type() {
	case diff.Equal:
		chunkType = "Equal"
	case diff.Add:
		chunkType = "Add"
	case diff.Delete:
		chunkType = "Delete"
	}

	return Chunk{
		Content: chunk.Content(),
		Type:    chunkType,
	}
}

func toFileInPatch(file diff.File) *FileWithinPatch {
	if file == nil {
		return nil
	}

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

func ToFilePatch(diffFilePatch diff.FilePatch, patchType string) FilePatch {
	var chunks []Chunk
	for _, c := range diffFilePatch.Chunks() {
		chunks = append(chunks, toChunk(c))
	}

	from, to := diffFilePatch.Files()
	return FilePatch{
		Type:     patchType,
		IsBinary: diffFilePatch.IsBinary(),
		FromFile: toFileInPatch(from),
		ToFile:   toFileInPatch(to),
		Chunks:   chunks,
	}
}
