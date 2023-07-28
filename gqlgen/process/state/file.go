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
	}
}

func FileUnChanged(currentFile *object.File, currentDir string) (*File, error) {
	if currentFile == nil {
		return nil, fmt.Errorf("failed in FileUnChanged, currentFile is nil")
	}

	filePath := filePathInDir(currentDir, currentFile.Name)

	// read contents here, to avoid error upon GraphQL materialization
	currentContents, err := currentFile.Contents()
	if err != nil {
		return nil, fmt.Errorf("failed in FileUnChanged for file = %s, cannot get current file contents, %s", filePath, err)
	}

	file := intrinsicFile(currentContents, filePath, currentFile.Size)
	// no need to update flags, as unchanged file has all flags false

	return file, nil
}

func FileAdded(currentFile *object.File, currentDir string) (*File, error) {
	if currentFile == nil {
		return nil, fmt.Errorf("failed in FileAdded, currentFile is nil")
	}

	filePath := filePathInDir(currentDir, currentFile.Name)

	// read contents here, to avoid error upon GraphQL materialization
	currentContents, err := currentFile.Contents()
	if err != nil {
		return nil, fmt.Errorf("failed in FileAdded for file = %s, cannot get current file contents, %s", filePath, err)
	}

	file := intrinsicFile(currentContents, filePath, currentFile.Size)
	// update necessary flags only, as default flags are false
	file.isAdded = true

	return file, nil
}

func FileUpdated(prevFile, currentFile *object.File, currentDir string) (*File, error) {
	if prevFile == nil {
		return nil, fmt.Errorf("failed in FileUpdated, prevFile is nil")
	}
	if currentFile == nil {
		return nil, fmt.Errorf("failed in FileUpdated, currentFile is nil")
	}

	filePath := filePathInDir(currentDir, currentFile.Name)

	// read contents here, to avoid error upon GraphQL materialization
	currentContents, err := currentFile.Contents()
	if err != nil {
		return nil, fmt.Errorf("failed in FileUpdated for file = %s, cannot get current file contents, %s", filePath, err)
	}

	file := intrinsicFile(currentContents, filePath, currentFile.Size)
	// update necessary flags only, as default flags are false
	file.isUpdated = true

	return file, nil
}

func FileDeleted(deletedFile diff.File) (*File, error) {
	if deletedFile == nil {
		return nil, fmt.Errorf("failed in FileDeleted, deletedFile is nil")
	}

	var filePath = deletedFile.Path()

	file := intrinsicFile("", filePath, 0)
	// update necessary flags only, as default flags are false
	file.isDeleted = true

	return file, nil
}

func (f *File) ToFileAdded() (*File, error) {
	// copy to avoid mutation effects afterwards
	file := *f
	// update necessary flags only, as default flags are false
	file.isAdded = true
	file.isUpdated = true
	return &file, nil
}

func (f *File) ToFileUpdated(from diff.File) (*File, error) {
	if from == nil {
		return nil, fmt.Errorf("failed in ToFileUpdated, 'from' File is nil")
	}

	// copy to avoid mutation effects afterwards
	file := *f
	// update necessary flags only, as default flags are false
	file.isUpdated = true
	return &file, nil
}

func NewFile(prevFile *object.File, currentFile *object.File, currentDir string) (*File, error) {
	if currentFile == nil && prevFile == nil {
		return nil, fmt.Errorf("failed in NewFile, currentFile and prevFile are both nil")
	}

	// read contents here, to avoid error upon GraphQL materialization
	var currentContents string
	var err error
	if currentFile != nil {
		currentContents, err = currentFile.Contents()
		if err != nil {
			return nil, fmt.Errorf("failed in ToGraphQLFileNode for file = %s, cannot get current file contents, %s", currentFile.Name, err)
		}
	}

	var filePath string
	if currentFile != nil {
		if currentDir != "" {
			filePath = currentDir + "/" + currentFile.Name
		} else {
			filePath = currentFile.Name
		}
	} else {
		filePath = prevFile.Name
	}

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

	isAdded := currentFile != nil && prevFile == nil
	isDeleted := currentFile == nil && prevFile != nil
	isUpdated := currentFile != nil && prevFile != nil && prevFile.Hash != currentFile.Hash
	isRenamed := currentFile != nil && prevFile != nil && prevFile.Name != currentFile.Name

	return &File{
		contents:  currentContents,
		filePath:  filePath,
		fileName:  fileName,
		language:  language,
		offset:    offset,
		isAdded:   isAdded,
		isUpdated: isUpdated,
		isDeleted: isDeleted,
		isRenamed: isRenamed,
	}, nil
}

func (s *File) ToGraphQLOpenFile() *model.OpenFile {
	// copy to avoid mutation effects afterwards
	filePath := s.filePath
	fileName := s.fileName
	language := s.language
	contents := s.contents // TODO: should we not copy this as contents can be huge?
	trueValue := true
	size := float64(s.size)

	return &model.OpenFile{
		FilePath:      &filePath,
		FileName:      &fileName,
		IsFullContent: &trueValue,
		Content:       &contents,
		Language:      &language,
		Size:          &size,
	}
}

func (s *File) ToGraphQLFileNode() *model.FileNode {
	//copy to avoid mutation effects afterwards
	fileType := model.FileNodeTypeFile
	filePath := s.filePath
	fileName := s.fileName
	offset := s.offset
	isUpdated := s.isUpdated || s.isAdded || s.isRenamed

	return &model.FileNode{
		NodeType:  &fileType,
		FilePath:  &filePath,
		Name:      &fileName,
		Offset:    &offset,
		IsUpdated: &isUpdated,
	}
}
