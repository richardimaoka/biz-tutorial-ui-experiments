package state

// it is very verbose so separated into a dedicated file
func fileLanguage(extension string) string {
	switch extension {
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
