package state

import (
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type File struct {
	prevFile        *object.File
	currentFile     *object.File
	currentContents string
	prevContents    string
	filePath        string
	offset          int
	fileName        string
	language        string
	isUpdated       bool
	isAdded         bool
	isDeleted       bool
	isRenamed       bool
}

func (f *File) FilePath() string {
	return f.filePath
}

func FileUnChanged(currentFile *object.File, currentDir string) (*File, error) {
	if currentFile == nil {
		return nil, fmt.Errorf("failed in NewFileUnChanged, currentFile is nil")
	}

	var filePath string
	if currentDir != "" {
		filePath = currentDir + "/" + currentFile.Name
	} else {
		filePath = currentFile.Name
	}

	// read contents here, to avoid error upon GraphQL materialization
	currentContents, err := currentFile.Contents()
	if err != nil {
		return nil, fmt.Errorf("failed in NewFileUnChanged for file = %s, cannot get current file contents, %s", filePath, err)
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

	return &File{
		prevFile:        currentFile,
		currentFile:     currentFile,
		currentContents: currentContents,
		prevContents:    currentContents,
		filePath:        filePath,
		fileName:        fileName,
		language:        language,
		offset:          offset,
		isAdded:         false,
		isUpdated:       false,
		isDeleted:       false,
		isRenamed:       false,
	}, nil
}

func NewFileDeleted(prevFile *object.File, prevDir string) (*File, error) {
	if prevFile == nil {
		return nil, fmt.Errorf("failed in NewFileUnChanged, prevFile is nil")
	}

	var filePath string
	if prevDir != "" {
		filePath = prevDir + "/" + prevFile.Name
	} else {
		filePath = prevFile.Name
	}

	// read contents here, to avoid error upon GraphQL materialization
	prevContents, err := prevFile.Contents()
	if err != nil {
		return nil, fmt.Errorf("failed in NewFileUnChanged for file = %s, cannot get current file contents, %s", filePath, err)
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

	return &File{
		prevFile:        prevFile,
		currentFile:     nil,
		currentContents: "",
		prevContents:    prevContents,
		filePath:        filePath,
		fileName:        fileName,
		language:        language,
		offset:          offset,
		isAdded:         false,
		isUpdated:       false,
		isDeleted:       true,
		isRenamed:       false,
	}, nil
}

func NewFile(prevFile *object.File, currentFile *object.File, currentDir string) (*File, error) {
	if currentFile == nil && prevFile == nil {
		return nil, fmt.Errorf("failed in NewFile, currentFile and prevFile are both nil")
	}

	// read contents here, to avoid error upon GraphQL materialization
	var currentContents, prevContents string
	var err error
	if currentFile != nil {
		currentContents, err = currentFile.Contents()
		if err != nil {
			return nil, fmt.Errorf("failed in ToGraphQLFileNode for file = %s, cannot get current file contents, %s", currentFile.Name, err)
		}
	}
	if prevFile != nil {
		prevContents, err = prevFile.Contents()
		if err != nil {
			return nil, fmt.Errorf("failed in ToGraphQLFileNode for file = %s, cannot get previous file contents, %s", prevFile.Name, err)
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
		prevFile:        prevFile,
		currentFile:     currentFile,
		currentContents: currentContents,
		prevContents:    prevContents,
		filePath:        filePath,
		fileName:        fileName,
		language:        language,
		offset:          offset,
		isAdded:         isAdded,
		isUpdated:       isUpdated,
		isDeleted:       isDeleted,
		isRenamed:       isRenamed,
	}, nil
}

func (s *File) ToGraphQLOpenFile() *model.OpenFile {
	// copy to avoid mutation effects afterwards
	filePath := s.filePath
	fileName := s.fileName
	language := s.language
	contents := s.currentContents // TODO: should we not copy this as contents can be huge?
	trueValue := true

	return &model.OpenFile{
		FilePath:      &filePath,
		FileName:      &fileName,
		IsFullContent: &trueValue,
		Content:       &contents,
		Language:      &language,
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
