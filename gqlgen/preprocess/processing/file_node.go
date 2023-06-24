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
	// Uppercase exported fields in lowercase unexported struct, as exported fields are necessary for reflection-based testing
	FromLine int
	ToLine   int
}

type fileTreeNode interface {
	NodeType() nodeType
	FilePath() string
	Offset() int
	Name() string
	ParentDirs() []string
	IsUpdated() bool
	ClearIsUpdated()
	ToGraphQLNode() *model.FileNode
	Clone() fileTreeNode
	Matched(comparedTo fileTreeNode) bool
}

type fileProcessorNode struct {
	filePath   string
	content    string
	isUpdated  bool
	highlights []fileHighlight
}

type directoryProcessorNode struct {
	filePath  string
	isUpdated bool
}

func (n *fileProcessorNode) NodeType() nodeType {
	return fileType
}

func (n *directoryProcessorNode) NodeType() nodeType {
	return directoryType
}

func (n *fileProcessorNode) FilePath() string {
	return n.filePath
}

func (n *directoryProcessorNode) FilePath() string {
	return n.filePath
}

func (n *fileProcessorNode) IsUpdated() bool {
	return n.isUpdated
}

func (n *directoryProcessorNode) IsUpdated() bool {
	return n.isUpdated
}

func (n *fileProcessorNode) ClearIsUpdated() {
	n.isUpdated = false
}

func (n *directoryProcessorNode) ClearIsUpdated() {
	n.isUpdated = false
}

func (n *fileProcessorNode) Matched(comparedTo fileTreeNode) bool {
	comparedFile, ok := comparedTo.(*fileProcessorNode)
	if ok {
		return n.filePath == comparedFile.filePath && n.content == comparedFile.content
	}

	return false
}

func (n *directoryProcessorNode) Matched(comparedTo fileTreeNode) bool {
	comparedDir, ok := comparedTo.(*directoryProcessorNode)
	if ok {
		return n.filePath == comparedDir.filePath
	}

	return false
}

func (n *fileProcessorNode) ParentDirs() []string {
	split := strings.Split(n.FilePath(), "/")
	if len(split) == 1 {
		return []string{}
	} else {
		return split[:len(split)-2]
	}
}

func (n *directoryProcessorNode) ParentDirs() []string {
	split := strings.Split(n.FilePath(), "/")
	if len(split) == 1 {
		return []string{}
	} else {
		return split[:len(split)-2]
	}
}

func (n *fileProcessorNode) Offset() int {
	split := strings.Split(n.FilePath(), "/")
	return len(split) - 1
}

func (n *directoryProcessorNode) Offset() int {
	split := strings.Split(n.FilePath(), "/")
	return len(split) - 1
}

func (n *fileProcessorNode) Name() string {
	split := strings.Split(n.FilePath(), "/")
	return split[len(split)-1]
}

func (n *directoryProcessorNode) Name() string {
	split := strings.Split(n.FilePath(), "/")
	return split[len(split)-1]
}

func (n *fileProcessorNode) ToGraphQLNode() *model.FileNode {
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

func (n *directoryProcessorNode) ToGraphQLNode() *model.FileNode {
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
func (n *fileProcessorNode) Clone() fileTreeNode {
	copied := *n // copy to avoid mutation effect afterwards
	return &copied
}

func (n *directoryProcessorNode) Clone() fileTreeNode {
	copied := *n // copy to avoid mutation effect afterwards
	return &copied
}

func (n *fileProcessorNode) ClearHighlights() {
	n.highlights = nil
}

func (n *fileProcessorNode) language() string {
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
	fromLine := h.FromLine // copy to avoid mutation effect afterwards
	toLine := h.ToLine     // copy to avoid mutation effect afterwards

	return &model.FileHighlight{
		FromLine: &fromLine,
		ToLine:   &toLine,
	}
}

func (n *fileProcessorNode) ToGraphQLOpenFile() *model.OpenFile {
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

func LessFileNode(a, b fileTreeNode) bool {
	// (e.g.) src/components/books/BookView.tsx
	// splitPathA := ["src", "components", "books", "BookView.tsx"]
	aSplitPath := strings.Split(a.FilePath(), "/")

	// (e.g.) src/libs/authentication.ts
	// splitPathB := ["src", "libs", "authentication.ts"]
	bSplitPath := strings.Split(b.FilePath(), "/")

	return lessFilePath2(aSplitPath, bSplitPath)
}
