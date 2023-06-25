package internal

import (
	"strings"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type NodeType string

const (
	FileType      NodeType = "FILE"
	DirectoryType NodeType = "DIRECTORY"
)

type FileTreeNode interface {
	NodeType() NodeType
	FilePath() string
	Offset() int
	Name() string
	ParentDirs() []string
	IsUpdated() bool
	ClearIsUpdated()
	ToGraphQLNode() *model.FileNode
	Clone() FileTreeNode
	Matched(comparedTo FileTreeNode) bool
}

type FileProcessorNode struct {
	filePath   string
	content    string
	isUpdated  bool
	highlights []FileHighlight
}

type DirectoryProcessorNode struct {
	filePath  string
	isUpdated bool
}

func NewFileNode(filePath string, content string, isUpdated bool) *FileProcessorNode {
	return &FileProcessorNode{
		filePath:  filePath,
		content:   content,
		isUpdated: isUpdated,
	}
}

func NewDirNode(filePath string, isUpdated bool) *DirectoryProcessorNode {
	return &DirectoryProcessorNode{
		filePath:  filePath,
		isUpdated: isUpdated,
	}
}

func (n *FileProcessorNode) NodeType() NodeType {
	return FileType
}

func (n *DirectoryProcessorNode) NodeType() NodeType {
	return DirectoryType
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

func (n *FileProcessorNode) ClearIsUpdated() {
	n.isUpdated = false
}

func (n *DirectoryProcessorNode) ClearIsUpdated() {
	n.isUpdated = false
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

func (n *FileProcessorNode) ParentDirs() []string {
	split := strings.Split(n.FilePath(), "/")
	if len(split) == 1 {
		return []string{}
	} else {
		return split[:len(split)-2]
	}
}

func (n *DirectoryProcessorNode) ParentDirs() []string {
	split := strings.Split(n.FilePath(), "/")
	if len(split) == 1 {
		return []string{}
	} else {
		return split[:len(split)-2]
	}
}

func (n *FileProcessorNode) Offset() int {
	split := strings.Split(n.FilePath(), "/")
	return len(split) - 1
}

func (n *DirectoryProcessorNode) Offset() int {
	split := strings.Split(n.FilePath(), "/")
	return len(split) - 1
}

func (n *FileProcessorNode) Name() string {
	split := strings.Split(n.FilePath(), "/")
	return split[len(split)-1]
}

func (n *DirectoryProcessorNode) Name() string {
	split := strings.Split(n.FilePath(), "/")
	return split[len(split)-1]
}

func (n *FileProcessorNode) ToGraphQLNode() *model.FileNode {
	filePath := n.filePath   // copy to avoid mutation effect afterwards
	isUpdated := n.isUpdated // copy to avoid mutation effect afterwards
	nodeType := model.FileNodeTypeFile
	offset := n.Offset()
	name := n.Name()

	return &model.FileNode{
		NodeType:  &nodeType,
		Name:      &name,
		FilePath:  &filePath,
		Offset:    &offset,
		IsUpdated: &isUpdated,
	}
}

func (n *DirectoryProcessorNode) ToGraphQLNode() *model.FileNode {
	filePath := n.filePath   // copy to avoid mutation effect afterwards
	isUpdated := n.isUpdated // copy to avoid mutation effect afterwards
	nodeType := model.FileNodeTypeDirectory
	offset := n.Offset()
	name := n.Name()

	return &model.FileNode{
		NodeType:  &nodeType,
		Name:      &name,
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

func (n *FileProcessorNode) ClearHighlights() {
	n.highlights = nil
}

func (n *FileProcessorNode) Contents() string {
	return n.content
}

func (n *FileProcessorNode) SetHighlights(highlights []FileHighlight) {
	n.highlights = highlights
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

func lessFilePath2(aSplitPath, bSplitPath []string) bool {
	a := aSplitPath[0] //supposedly len(aSplitPath) > 0
	b := bSplitPath[0] //supposedly len(bSplitPath) > 0

	if a == b {
		if len(aSplitPath) == 1 {
			// (e.g.)
			//   aSplitPath = ["src", "components", "books"]
			//   bSplitPath = ["src", "components", "books", "BookView.tsx"]
			// no more path part to compare, then aSplitPath is "less"
			return true
		} else if len(bSplitPath) == 1 {
			// (e.g.)
			//   aSplitPath = ["src", "components", "books", "BookTab.tsx"]
			//   bSplitPath = ["src", "components"]
			return false
		}

		// more path parts to compare in both aSplitPath and bSplitPath
		return lessFilePath2(aSplitPath[1:], bSplitPath[1:])
	} else {
		return a < b
	}
}

func LessFileNode(a, b FileTreeNode) bool {
	// (e.g.) src/components/books/BookView.tsx
	// splitPathA := ["src", "components", "books", "BookView.tsx"]
	aSplitPath := strings.Split(a.FilePath(), "/")

	// (e.g.) src/libs/authentication.ts
	// splitPathB := ["src", "libs", "authentication.ts"]
	bSplitPath := strings.Split(b.FilePath(), "/")

	return lessFilePath2(aSplitPath, bSplitPath)
}
