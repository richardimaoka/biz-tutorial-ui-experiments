package model

import (
	"encoding/json"
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/jsonwrap"
)

func (t *ColumnWrapper) UnmarshalJSON(b []byte) error {
	/**
	 * Read JSON into internal struct
	 */
	internals := struct {
		Column            interface{} `json:"column"`
		ColumnName        string      `json:"columnName"`
		ColumnDisplayName *string     `json:"columnDisplayName"`
	}{}

	err := json.Unmarshal(b, &internals)
	if err != nil {
		return fmt.Errorf("failed in ColumnWrapper UnmarshalJSON() while unmarshaling to Go data, %w", err)
	}
	t.ColumnName = internals.ColumnName
	t.ColumnDisplayName = internals.ColumnDisplayName

	/**
	 * Marshal `column` and unmarshal it back
	 */
	bytes, err := json.Marshal(internals.Column)
	if err != nil {
		return fmt.Errorf("failed in ColumnWrapper UnmarshalJSON() while marshaling the column object, %w", err)
	}
	column, err := columnFromBytes(bytes)
	if err != nil {
		return fmt.Errorf("failed in ColumnWrapper UnmarshalJSON() while unmarshaling the column object, %w", err)
	}
	t.Column = column

	return nil
}

func columnFromBytes(bytes []byte) (Column, error) {
	fromField := "__typename"
	typename, err := jsonwrap.ExtractTypeName(bytes, fromField)
	if err != nil {
		return nil, err
	}

	switch typename {
	case "TerminalColumn2":
		var col TerminalColumn2
		if err := json.Unmarshal(bytes, &col); err != nil {
			return nil, err
		}
		return &col, nil

	case "SourceCodeColumn2":
		var col SourceCodeColumn2
		if err := json.Unmarshal(bytes, &col); err != nil {
			return nil, err
		}
		return &col, nil

	case "BrowserColumn2":
		var col BrowserColumn2
		if err := json.Unmarshal(bytes, &col); err != nil {
			return nil, err
		}
		return &col, nil

	default:
		return nil, fmt.Errorf("\"%s\" = %s is not a valid Column type. If it should be valid, define it in column_wrapper.go", fromField, typename)
	}
}
