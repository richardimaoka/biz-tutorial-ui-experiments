package state

import (
	"fmt"
	"sort"
	"strings"

	"github.com/go-git/go-git/v5/plumbing/format/diff"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type File struct {
	// intrinsic fields
	filePath string
	fileName string
	offset   int
	language string
	contents string
	size     int64

	highlights []FileHighlight

	// flags
	isUpdated bool
	isAdded   bool
	isDeleted bool
	isRenamed bool
}

type Files []*File

func (files Files) sortSelf() {
	sort.Slice(files, func(i, j int) bool {
		return strings.ToLower(files[i].fileName) < strings.ToLower(files[j].fileName)
	})
}

func intrinsicFile(contents string, filePath string, size int64) *File {
	split := strings.Split(filePath, "/")
	fileName := split[len(split)-1]

	offset := len(split) - 1

	dotSplit := strings.Split(fileName, ".")
	var suffix string
	if len(dotSplit) > 1 {
		//e.g. fileName = 'some.interesting.name.json', suffix = 'json'
		suffix = dotSplit[len(dotSplit)-1]
	}
	language := fileLanguage(suffix)

	return &File{
		// intrinsic fields
		filePath: filePath,
		fileName: fileName,
		offset:   offset,
		language: language,
		contents: contents,
		size:     size,

		// flags are all false by default = Go's zero value for bool
		// isUpdated bool
		// ...
	}
}

func FileUnChanged(currentFile *object.File, currentDir string) (*File, error) {
	if currentFile == nil {
		return nil, fmt.Errorf("failed in FileUnChanged, currentFile is nil")
	}

	filePath := filePathInDir(currentDir, currentFile.Name)

	isBinary, err := currentFile.IsBinary()
	if err != nil {
		return nil, fmt.Errorf("failed in FileUnChanged for file = %s, cannot get current file binary flag, %s", filePath, err)
	}

	// read contents here, to avoid error upon GraphQL materialization
	var currentContents string
	if isBinary {
		currentContents = "Binary file not shown."
	} else {
		var err error
		currentContents, err = currentFile.Contents()
		if err != nil {
			return nil, fmt.Errorf("failed in FileUnChanged for file = %s, cannot get current file contents, %s", filePath, err)
		}
	}

	file := intrinsicFile(currentContents, filePath, currentFile.Size)
	// no need to update flags, as unchanged file has all flags false

	return file, nil
}

func FileDeleted(filePath string) *File {
	file := intrinsicFile("", filePath, 0)
	// update necessary flags only, as default flags are false
	file.isDeleted = true

	return file
}

// to keep File immutable, return a new File
func (f *File) ToFileAdded() *File {
	// copy to avoid mutation effects afterwards
	file := *f
	// update necessary flags only, as default flags are false
	file.isAdded = true
	file.isUpdated = true
	return &file
}

// to keep File immutable, return a new File
func (f *File) ToFileUpdated(patch diff.FilePatch) *File {
	// copy to avoid mutation effects afterwards
	file := *f
	// update necessary flags only, as default flags are false
	file.isUpdated = true
	file.highlights = CalcHighlight(patch)

	return &file
}

func (s *File) ToGraphQLOpenFile() *model.OpenFile {
	// copy to avoid mutation effects afterwards
	filePath := s.filePath
	fileName := s.fileName
	language := s.language
	contents := s.contents // TODO: should we not copy this as contents can be huge?
	trueValue := true
	size := float64(s.size)

	var highlights []*model.FileHighlight
	for _, h := range s.highlights {
		highlights = append(highlights, h.ToGraphQLFileHighlight())
	}

	return &model.OpenFile{
		FilePath:      &filePath,
		FileName:      &fileName,
		IsFullContent: &trueValue,
		Content:       &contents,
		Language:      &language,
		Size:          &size,
		Highlight:     highlights,
	}
}

func (s *File) ToGraphQLFileNode() *model.FileNode {
	//copy to avoid mutation effects afterwards
	fileType := model.FileNodeTypeFile
	filePath := s.filePath
	fileName := s.fileName
	offset := s.offset
	isUpdated := s.isUpdated || s.isAdded || s.isRenamed
	isDeleted := s.isDeleted

	return &model.FileNode{
		NodeType:  &fileType,
		FilePath:  &filePath,
		Name:      &fileName,
		Offset:    &offset,
		IsUpdated: &isUpdated,
		IsDeleted: &isDeleted,
	}
}
