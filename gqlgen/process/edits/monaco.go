package edits

type Range struct {
	StartLineNumber int `json:"startLineNumber"`
	StartColumn     int `json:"startColumn"`
	EndLineNumber   int `json:"endLineNumber"`
	EndColumn       int `json:"endColumn"`
}

type SingleEditOperation struct {
	Text  string `json:"text"`
	Range Range  `json:"range"`
}

func a() {}
