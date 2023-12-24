package input

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type CsvString string //Whether it's an integer value or a string, forcefully convert to a String in Go

func (v *CsvString) UnmarshalJSON(b []byte) error {
	var intValue int
	if err := json.Unmarshal(b, &intValue); err == nil {
		// if no error, successfully unmarshaled to int, then convert the int to CsvString
		*v = CsvString(strconv.Itoa(intValue))
		return nil
	}

	var stringValue string
	if err := json.Unmarshal(b, &stringValue); err == nil {
		// if no error, successfully unmarshaled to string
		*v = CsvString(stringValue)
		return nil
	} else {
		return fmt.Errorf("unmarshal to CsvString failed, %s", err)
	}
}

type CsvInt int

func (v *CsvInt) UnmarshalJSON(b []byte) error {
	var stringValue string
	if err := json.Unmarshal(b, &stringValue); err == nil {
		if stringValue == "" {
			*v = CsvInt(0)
			return nil
		}

		intValue, err := strconv.Atoi(stringValue)
		if err != nil {
			return fmt.Errorf("unmarshan to CsvInt failed, %s", err)
		}
		*v = CsvInt(intValue)
		return nil
	}

	var intValue int
	err := json.Unmarshal(b, &intValue)
	// if no error, successfully unmarshaled to int
	if err == nil {
		*v = CsvInt(intValue)
		return nil
	}

	return fmt.Errorf("unmarshal to CsvInt failed, %s", err)
}
