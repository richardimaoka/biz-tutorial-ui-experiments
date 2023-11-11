package model

import (
	"encoding/json"
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/jsonwrap"
)

func (t *ColumnWrapper) UnmarshalJSON(b []byte) error {
	var obj jsonwrap.JsonObj
	err := json.Unmarshal(b, &obj)
	if err != nil {
		return fmt.Errorf("failed in ColumnWrapper UnmarshalJSON() while unmarshaling to Go data, %w", err)
	}

	columnObj, ok := obj["column"]
	if !ok {
		return fmt.Errorf(`failed in ColumnWrapper UnmarshalJSON(), "column" field does not exist`)
	}

	name, ok := obj["name"]
	if !ok {
		name = ""
		// return fmt.Errorf(`failed in ColumnWrapper UnmarshalJSON(), "name" field does not exist`)
	}

	nameStr, ok := name.(string)
	if !ok {
		return fmt.Errorf(`failed in ColumnWrapper UnmarshalJSON(), "name" field = %v is not string`, name)
	}

	bytes, err := json.Marshal(columnObj)
	if err != nil {
		return fmt.Errorf("failed in ColumnWrapper UnmarshalJSON() while marshaling the column object, %w", err)
	}

	column, err := columnFromBytes_Old(bytes)
	if err != nil {
		return fmt.Errorf("failed in ColumnWrapper UnmarshalJSON() while unmarshaling the column object, %w", err)
	}

	t.Column = column
	t.Name = &nameStr

	return nil
}

func columnFromBytes_Old(bytes []byte) (Column, error) {
	fromField := "__typename"
	typename, err := jsonwrap.ExtractTypeName(bytes, fromField)
	if err != nil {
		return nil, err
	}

	switch typename {

	case "SourceCodeColumn":
		var col SourceCodeColumn
		if err := json.Unmarshal(bytes, &col); err != nil {
			return nil, err
		}
		return &col, nil

	default:
		return nil, fmt.Errorf("\"%s\" = %s is not a valid Column type. If it should be valid, define it in column_wrapper.go", fromField, typename)
	}
}
