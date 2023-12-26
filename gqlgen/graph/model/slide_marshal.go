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

func (this ImageSlide) MarshalJSON() ([]byte, error) {
	extendedOutput := struct {
		TypeName    string  `json:"__typename"`
		Placeholder *string `json:"_placeholder"`
		Title       string  `json:"title"`
		Image       *Image  `json:"image"`
	}{
		TypeName:    "ImageSlide",
		Placeholder: this.Placeholder,
		Image:       this.Image,
	}

	return json.Marshal(extendedOutput)
}

func (this MarkdownSlide) MarshalJSON() ([]byte, error) {
	extendedOutput := struct {
		TypeName     string  `json:"__typename"`
		Placeholder  *string `json:"_placeholder"`
		MarkdownBody string  `json:"markdownBody"`
	}{
		TypeName:     "MarkdownSlide",
		Placeholder:  this.Placeholder,
		MarkdownBody: this.MarkdownBody,
	}

	return json.Marshal(extendedOutput)
}

func (this SectionTitleSlide) MarshalJSON() ([]byte, error) {
	extendedOutput := struct {
		TypeName    string  `json:"__typename"`
		Placeholder *string `json:"_placeholder"`
		Title       string  `json:"title"`
		SectionNum  int     `json:"sectionNum"`
	}{
		TypeName:    "SectionTitleSlide",
		Placeholder: this.Placeholder,
		Title:       this.Title,
		SectionNum:  this.SectionNum,
	}

	return json.Marshal(extendedOutput)
}
