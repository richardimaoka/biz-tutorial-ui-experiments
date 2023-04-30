package model

import "encoding/json"

func (a NextActionManual) MarshalJSON() ([]byte, error) {
	extended := struct {
		ContentType string  `json:"contentType"`
		Comment     *string `json:"comment"`
	}{
		"NextActionManual",
		a.Comment,
	}

	return json.Marshal(extended)
}
