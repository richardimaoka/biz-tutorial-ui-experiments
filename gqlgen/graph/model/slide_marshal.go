package model

import "encoding/json"

func (this TutorialTitleSlide) MarshalJSON() ([]byte, error) {
	extendedOutput := struct {
		TypeName    string   `json:"__typename"`
		Placeholder *string  `json:"_placeholder"`
		Title       string   `json:"title"`
		Images      []*Image `json:"images"`
	}{
		TypeName:    "TutorialTitleSlide",
		Placeholder: this.Placeholder,
		Title:       this.Title,
		Images:      this.Images,
	}

	return json.Marshal(extendedOutput)
}

// func (this TutorialTitleSlide) MarshalJSON() ([]byte, error) {
// 	extendedOutput := struct {
// 		TypeName    string   `json:"__typename"`
// 		Placeholder *string  `json:"_placeholder"`
// 		Title       string   `json:"title"`
// 		Images      []*Image `json:"images"`
// 	}{
// 		TypeName:    "TutorialTitleSlide",
// 		Placeholder: this.Placeholder,
// 		Title:       this.Title,
// 		Images:      this.Images,
// 	}

// 	return json.Marshal(extendedOutput)
// }

// func (this TutorialTitleSlide) MarshalJSON() ([]byte, error) {
// 	extendedOutput := struct {
// 		TypeName    string   `json:"__typename"`
// 		Placeholder *string  `json:"_placeholder"`
// 		Title       string   `json:"title"`
// 		Images      []*Image `json:"images"`
// 	}{
// 		TypeName:    "TutorialTitleSlide",
// 		Placeholder: this.Placeholder,
// 		Title:       this.Title,
// 		Images:      this.Images,
// 	}

// 	return json.Marshal(extendedOutput)
// }

// func (this TutorialTitleSlide) MarshalJSON() ([]byte, error) {
// 	extendedOutput := struct {
// 		TypeName    string   `json:"__typename"`
// 		Placeholder *string  `json:"_placeholder"`
// 		Title       string   `json:"title"`
// 		Images      []*Image `json:"images"`
// 	}{
// 		TypeName:    "TutorialTitleSlide",
// 		Placeholder: this.Placeholder,
// 		Title:       this.Title,
// 		Images:      this.Images,
// 	}

// 	return json.Marshal(extendedOutput)
// }
