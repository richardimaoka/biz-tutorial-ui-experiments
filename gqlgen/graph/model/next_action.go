package model

import (
	"encoding/json"
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
)

func (t *NextAction) UnmarshalJSON(b []byte) error {
	var unmarshald map[string]interface{}
	err := json.Unmarshal(b, &unmarshald)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(unmarshald["content"])
	if err != nil {
		return err
	}

	content, err := nextActionContentFromBytes(bytes)
	if err != nil {
		return err
	}
	t.Content = content

	return nil
}

func nextActionContentFromBytes(bytes []byte) (NextActionContent, error) {
	fromField := "contentType"
	typename, err := internal.ExtractTypeName(bytes, fromField)
	if err != nil {
		return nil, err
	}

	switch typename {
	case "NextActionTerminal":
		var action NextActionTerminal
		if err := json.Unmarshal(bytes, &action); err != nil {
			return nil, err
		}
		return &action, nil

	case "NextActionManual":
		var action NextActionManual
		if err := json.Unmarshal(bytes, &action); err != nil {
			return nil, err
		}

		return &action, nil

	default:
		return nil, fmt.Errorf("\"%s\" = %s is not a valid NextActionContent type", fromField, typename)
	}
}
