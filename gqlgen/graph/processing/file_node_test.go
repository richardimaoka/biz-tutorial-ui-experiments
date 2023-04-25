package processing

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

func TestFileNode_MutateDirectory(t *testing.T) {
	processor := NewSourceCodeProcessor()
	processor.AddDirectory(model.DirectoryAdd{FilePath: "hello/world"})
	result := processor.fileMap["hello/world"].ToGraphQLNode()
	compareAfterMarshal(t, "testdata/source_code/file-node-mutation1.json", result)

	// after SourceCode is materialized to GraphQL object, mutation should have no effect
	processor.fileMap["hello/world"].(*directoryProcessorNode).isUpdated = true
	processor.fileMap["hello/world"].(*directoryProcessorNode).filePath = "something/else"

	compareAfterMarshal(t, "testdata/source_code/file-node-mutation1.json", result)
}

func TestFileNode_MutateFile(t *testing.T) {
	processor := NewSourceCodeProcessor()
	processor.AddFile(model.FileAdd{FilePath: "hello/world"})
	result := processor.fileMap["hello/world"].ToGraphQLNode()
	compareAfterMarshal(t, "testdata/source_code/file-node-mutation2.json", result)

	// after SourceCode is materialized to GraphQL object, mutation should have no effect
	processor.fileMap["hello/world"].(*fileProcessorNode).isUpdated = false
	processor.fileMap["hello/world"].(*fileProcessorNode).filePath = "something/else"
	processor.fileMap["hello/world"].(*fileProcessorNode).content = "contenttttt ttt tt"

	compareAfterMarshal(t, "testdata/source_code/file-node-mutation2.json", result)
}

func TestFileNodeOpenFile_MutateFile(t *testing.T) {
	processor := NewSourceCodeProcessor()
	processor.AddFile(model.FileAdd{FilePath: "hello/world"})
	result := processor.fileMap["hello/world"].(*fileProcessorNode).ToGraphQLOpenFile()
	compareAfterMarshal(t, "testdata/source_code/file-node-mutation-openfile1.json", result)

	// after SourceCode is materialized to GraphQL object, mutation should have no effect
	processor.fileMap["hello/world"].(*fileProcessorNode).isUpdated = false
	processor.fileMap["hello/world"].(*fileProcessorNode).filePath = "something/else"
	processor.fileMap["hello/world"].(*fileProcessorNode).content = "contenttttt ttt tt"

	compareAfterMarshal(t, "testdata/source_code/file-node-mutation-openfile1.json", result)
}
