package state

import (
	"fmt"
	"sort"
	"strings"

	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/google/uuid"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/edits"
)

type File struct {
	// Path related fields
	fileName    string
	filePath    string
	oldFilePath string
	offset      int

	// Contents fields
	language    string
	contents    string
	oldContents string
	size        int64
	edits       []edits.SingleEditOperation

	// Tooltip fields
	tooltip *SourceCodeTooltip

	// Flags
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

func fileUnChanged(currentFile *object.File, filePath string) (*File, error) {
	if currentFile == nil {
		return nil, fmt.Errorf("failed in FileUnChanged, currentFile is nil")
	}

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

func (f *File) markDeleted() {
	f.oldContents = f.contents
	f.contents = ""

	f.isUpdated = false
	f.isAdded = false
	f.isDeleted = true
	f.isRenamed = false
}

func (f *File) markUpdated(oldContents string, editOps []edits.SingleEditOperation) {
	f.oldContents = oldContents
	f.edits = editOps

	f.isUpdated = true
	f.isAdded = false
	f.isDeleted = false
	f.isRenamed = false
}

func (f *File) markAdded() {
	f.isUpdated = false
	f.isAdded = true
	f.isDeleted = false
	f.isRenamed = false
}

func (f *File) markRenamed(oldFilePath string) {
	f.oldFilePath = oldFilePath

	f.isUpdated = false
	f.isAdded = true
	f.isDeleted = false
	f.isRenamed = false
}

func (f *File) clearEdits() {
	f.edits = nil
}

func toEditSequence(edits []edits.SingleEditOperation) *model.EditSequence {
	if edits == nil || len(edits) == 0 {
		return nil
	}

	var monacoEdits []*model.MonacoEditOperation
	for _, e := range edits {
		editRange := model.MonacoEditRange{
			StartLineNumber: e.Range.StartLineNumber,
			StartColumn:     e.Range.StartColumn,
			EndLineNumber:   e.Range.EndLineNumber,
			EndColumn:       e.Range.EndColumn,
		}
		monacoEdits = append(monacoEdits, &model.MonacoEditOperation{
			Text:  e.Text,
			Range: &editRange,
		})
	}

	id := uuid.NewString()

	return &model.EditSequence{
		ID:    id,
		Edits: monacoEdits,
	}
}

func (s *File) ToGraphQLOpenFile() *model.OpenFile {
	// copy to avoid mutation effects afterwards
	filePath := s.filePath
	fileName := s.fileName
	language := s.language
	contents := s.contents
	oldContents := s.oldContents
	trueValue := true
	size := float64(s.size)
	editSequence := toEditSequence(s.edits)

	var tooltip *model.SourceCodeTooltip
	if s.tooltip != nil {
		tooltip = s.tooltip.ToGraphQL()
	}

	return &model.OpenFile{
		FilePath:      &filePath,
		FileName:      &fileName,
		IsFullContent: &trueValue,
		Content:       &contents,
		OldContent:    &oldContents,
		Language:      &language,
		Size:          &size,
		EditSequence:  editSequence,
		Tooltip:       tooltip,
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
		NodeType:  fileType,
		FilePath:  filePath,
		Name:      &fileName,
		Offset:    &offset,
		IsUpdated: &isUpdated,
		IsDeleted: &isDeleted,
	}
}
