package model

import "encoding/json"

func (this YouTubeColumn) MarshalJSON() ([]byte, error) {
	extendedOutput := struct {
		TypeName    string        `json:"__typename"`
		Placeholder *string       `json:"_placeholder,omitempty"`
		Youtube     *YouTubeEmbed `json:"youtube,omitempty"`
	}{
		TypeName:    "YouTubeColumn",
		Placeholder: this.Placeholder,
		Youtube:     this.Youtube,
	}

	return json.Marshal(extendedOutput)
}
