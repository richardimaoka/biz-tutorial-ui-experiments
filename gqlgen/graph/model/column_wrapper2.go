package model

import (
	"encoding/json"
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
)

func (t *ColumnWrapper2) UnmarshalJSON(b []byte) error {
	/**
	 * Read JSON into JsonObj(i.e. map), and validate fields
	 */
	var obj internal.JsonObj
	err := json.Unmarshal(b, &obj)
	if err != nil {
		return fmt.Errorf("failed in ColumnWrapper2 UnmarshalJSON() while unmarshaling to Go data, %w", err)
	}

	// See if necessary fields exist
	name, ok := obj["columnName"]
	if !ok {
		return fmt.Errorf(`failed in ColumnWrapper2 UnmarshalJSON(), "columnName" field does not exist`)
	}
	nameStr, ok := name.(string)
	if !ok {
		return fmt.Errorf(`failed in ColumnWrapper2 UnmarshalJSON(), "name" field = %v is not string`, name)
	}

	// See if the column field exists
	columnObj, ok := obj["column"]
	if !ok {
		return fmt.Errorf(`failed in ColumnWrapper2 UnmarshalJSON(), "column" field does not exist`)
	}

	/**
	 * Marshal `column` and unmarshal it back
	 */
	bytes, err := json.Marshal(columnObj)
	if err != nil {
		return fmt.Errorf("failed in ColumnWrapper2 UnmarshalJSON() while marshaling the column object, %w", err)
	}
	column, err := columnFromBytes2(bytes)
	if err != nil {
		return fmt.Errorf("failed in ColumnWrapper2 UnmarshalJSON() while unmarshaling the column object, %w", err)
	}

	/**
	 * If all successful, use the column and the name
	 */
	t.Column = column
	t.ColumnName = nameStr

	return nil
}

func columnFromBytes2(bytes []byte) (Column2, error) {
	fromField := "__typename"
	typename, err := internal.ExtractTypeName(bytes, fromField)
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
	default:
		return nil, fmt.Errorf("\"%s\" = %s is not a valid Column type. If it should be valid, define it in column_wrapper.go", fromField, typename)
	}
}
