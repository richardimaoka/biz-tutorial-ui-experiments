package model

import "encoding/json"

func (this DevToolsColumn) MarshalJSON() ([]byte, error) {
	extendedOutput := struct {
		TypeName    string  `json:"__typename"`
		Placeholder *string `json:"_placeholder,omitempty"`
		Width       *int    `json:"width,omitempty"`
		Height      *int    `json:"height,omitempty"`
		Path        *string `json:"path,omitempty"`
	}{
		TypeName:    "DevToolsColumn",
		Placeholder: this.Placeholder,
		Width:       this.Width,
		Height:      this.Height,
		Path:        this.Path,
	}

	return json.Marshal(extendedOutput)
}
