package edits

type Range struct {
	StartLineNumber int `json:"startLineNumber"`
	StartColumn     int `json:"startColumn"`
	EndLineNumber   int `json:"endLineNumber"`
	EndColumn       int `json:"endColumn"`
}

type SingleEditOperatoin struct {
	Range Range `json:"range"`
}

func a() {}
