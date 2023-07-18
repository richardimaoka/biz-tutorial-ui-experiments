package model

import "encoding/json"

func (this BackgroundImageColumn) MarshalJSON() ([]byte, error) {
	extendedOutput := struct {
		TypeName    string  `json:"__typename"`
		Placeholder *string `json:"_placeholder,omitempty"`
		Width       *int    `json:"width,omitempty"`
		Height      *int    `json:"height,omitempty"`
		Path        *string `json:"path,omitempty"`
		URL         *string `json:"url,omitempty"`
		Modal       *Modal  `json:"modal,omitempty"`
	}{
		TypeName:    "BackgroundImageColumn",
		Placeholder: this.Placeholder,
		Width:       this.Width,
		Height:      this.Height,
		Path:        this.Path,
		URL:         this.URL,
		Modal:       this.Modal,
	}

	return json.Marshal(extendedOutput)
}
