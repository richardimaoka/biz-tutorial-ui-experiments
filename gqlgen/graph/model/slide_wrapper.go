package model

import (
	"encoding/json"
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/jsonwrap"
)

func (t *SlideWrapper) UnmarshalJSON(b []byte) error {
	/**
	 * Read JSON into internal struct
	 */
	internals := struct {
		Slide interface{} `json:"slide"`
	}{}

	err := json.Unmarshal(b, &internals)
	if err != nil {
		return fmt.Errorf("failed in SlideWrapper UnmarshalJSON() while unmarshaling to Go data, %w", err)
	}

	/**
	 * Marshal `slide` and unmarshal it back
	 */
	bytes, err := json.Marshal(internals.Slide)
	if err != nil {
		return fmt.Errorf("failed in SlideWrapper UnmarshalJSON() while marshaling the slide object, %w", err)
	}
	slide, err := slideFromBytes(bytes)
	if err != nil {
		return fmt.Errorf("failed in SlideWrapper UnmarshalJSON() while unmarshaling the slide object, %w", err)
	}
	t.Slide = slide

	return nil
}

func slideFromBytes(bytes []byte) (Slide, error) {
	fromField := "__typename"
	typename, err := jsonwrap.ExtractTypeName(bytes, fromField)
	if err != nil {
		return nil, err
	}

	switch typename {
	case "TutorialTitleSlide":
		var slide TutorialTitleSlide
		if err := json.Unmarshal(bytes, &slide); err != nil {
			return nil, err
		}
		return &slide, nil

	case "MarkdownSlide":
		var slide MarkdownSlide
		if err := json.Unmarshal(bytes, &slide); err != nil {
			return nil, err
		}
		return &slide, nil

	default:
		return nil, fmt.Errorf("\"%s\" = %s is not a valid Slide type. If it should be valid, define it in slide_wrapper.go", fromField, typename)
	}
}
