package state2

type BrowserStepType string

const (
	BrowserOpen BrowserStepType = "open"
	BrowserMove BrowserStepType = "move"
)

type BrowserFields struct {
	BrowserStepType  BrowserStepType
	BrowserImagePath string `json:"browserImagePath"`
}
