package processing

import (
	"strings"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type nodeType string

const (
	fileType      nodeType = "FILE"
	directoryType nodeType = "DIRECTORY"
)

type fileHighlight struct {
	fromLine int
	toLine   int
}

type FileTreeNode interface {
	NodeType() nodeType
	FilePath() string
	IsUpdated() bool
	SetIsUpdated(isUpdated bool)
	ToGraphQLNode() *model.FileNode
	Clone() FileTreeNode
	Matched(comparedTo FileTreeNode) bool
}

type FileProcessorNode struct {
	filePath   string
	content    string
	isUpdated  bool
	highlights []fileHighlight
}

type DirectoryProcessorNode struct {
	filePath  string
	isUpdated bool
}

func (n *FileProcessorNode) NodeType() nodeType {
	return fileType
}

func (n *DirectoryProcessorNode) NodeType() nodeType {
	return directoryType
}

func (n *FileProcessorNode) FilePath() string {
	return n.filePath
}

func (n *DirectoryProcessorNode) FilePath() string {
	return n.filePath
}

func (n *FileProcessorNode) IsUpdated() bool {
	return n.isUpdated
}

func (n *DirectoryProcessorNode) IsUpdated() bool {
	return n.isUpdated
}

func (n *FileProcessorNode) SetIsUpdated(isUpdated bool) {
	n.isUpdated = isUpdated
}

func (n *DirectoryProcessorNode) SetIsUpdated(isUpdated bool) {
	n.isUpdated = isUpdated
}

func (n *FileProcessorNode) Matched(comparedTo FileTreeNode) bool {
	comparedFile, ok := comparedTo.(*FileProcessorNode)
	if ok {
		return n.filePath == comparedFile.filePath && n.content == comparedFile.content
	}

	return false
}

func (n *DirectoryProcessorNode) Matched(comparedTo FileTreeNode) bool {
	comparedDir, ok := comparedTo.(*DirectoryProcessorNode)
	if ok {
		return n.filePath == comparedDir.filePath
	}

	return false
}

func (n *FileProcessorNode) ToGraphQLNode() *model.FileNode {
	filePath := n.filePath   // copy to avoid mutation effect afterwards
	isUpdated := n.isUpdated // copy to avoid mutation effect afterwards
	nodeType := model.FileNodeTypeFile
	split := strings.Split(filePath, "/")
	offset := len(split) - 1

	return &model.FileNode{
		NodeType:  &nodeType,
		Name:      &split[len(split)-1],
		FilePath:  &filePath,
		Offset:    &offset,
		IsUpdated: &isUpdated,
	}
}

func (n *DirectoryProcessorNode) ToGraphQLNode() *model.FileNode {
	filePath := n.filePath   // copy to avoid mutation effect afterwards
	isUpdated := n.isUpdated // copy to avoid mutation effect afterwards
	nodeType := model.FileNodeTypeDirectory
	split := strings.Split(filePath, "/")
	offset := len(split) - 1

	return &model.FileNode{
		NodeType:  &nodeType,
		Name:      &split[len(split)-1],
		FilePath:  &filePath,
		Offset:    &offset,
		IsUpdated: &isUpdated,
	}
}

// TODO: can this be deleted?? check and delete if possible
func (n *FileProcessorNode) Clone() FileTreeNode {
	copied := *n // copy to avoid mutation effect afterwards
	return &copied
}

func (n *DirectoryProcessorNode) Clone() FileTreeNode {
	copied := *n // copy to avoid mutation effect afterwards
	return &copied
}

func (n *FileProcessorNode) language() string {
	split := strings.Split(n.filePath, ".")

	if len(split) == 1 {
		return ""
	}

	suffix := split[len(split)-1]
	switch suffix {
	case "go":
		return "go"
	case "js":
		return "javascript"
	case "ts":
		return "typescript"
	case "html":
		return "html"
	case "css":
		return "css"
	case "md":
		return "markdown"
	case "json":
		return "json"
	case "yaml":
		return "yaml"
	case "yml":
		return "yaml"
	case "graphql":
		return "graphql"
	case "gql":
		return "graphql"
	case "sql":
		return "sql"
	case "py":
		return "python"
	case "java":
		return "java"
	case "kt":
		return "kotlin"
	case "swift":
		return "swift"
	case "rb":
		return "ruby"
	case "php":
		return "php"
	case "c":
		return "c"
	case "cpp":
		return "cpp"
	case "h":
		return "cpp"
	case "hpp":
		return "cpp"
	case "cs":
		return "csharp"
	case "scala":
		return "scala"
	case "rs":
		return "rust"
	case "sh":
		return "shell"
	case "bash":
		return "shell"
	case "zsh":
		return "shell"
	case "ps1":
		return "powershell"
	case "psm1":
		return "powershell"
	case "bat":
		return "batch"
	case "gradle":
		return "groovy"
	case "xml":
		return "xml"
	case "vue":
		return "vue"
	default:
		return ""
	}
}

func (h *fileHighlight) ToGraphQLFileHighlight() *model.FileHighlight {
	fromLine := h.fromLine // copy to avoid mutation effect afterwards
	toLine := h.toLine     // copy to avoid mutation effect afterwards

	return &model.FileHighlight{
		FromLine: &fromLine,
		ToLine:   &toLine,
	}
}

func (n *FileProcessorNode) ToGraphQLOpenFile() *model.OpenFile {
	filePath := n.filePath // copy to avoid mutation effect afterwards
	content := n.content   // copy to avoid mutation effect afterwards
	isFullContent := true
	split := strings.Split(filePath, "/")

	l := n.language()
	var language *string
	if l != "" {
		language = &l
	} else {
		language = nil
	}

	var highlights []*model.FileHighlight
	for _, h := range n.highlights {
		highlight := h.ToGraphQLFileHighlight()
		highlights = append(highlights, highlight)
	}

	return &model.OpenFile{
		FilePath:      &filePath,
		FileName:      &split[len(split)-1],
		Content:       &content,
		IsFullContent: &isFullContent,
		Language:      language,
		Highlight:     highlights,
	}
}
