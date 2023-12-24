package input

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type CsvString string

func (v *CsvString) UnmarshalJSON(b []byte) error {

	var intValue int
	err1 := json.Unmarshal(b, &intValue)
	// if no error, successfully unmarshaled to int
	if err1 == nil {
		*v = CsvString(strconv.Itoa(intValue))
		return nil
	}

	var stringValue string
	err2 := json.Unmarshal(b, &stringValue)
	if err2 == nil {
		*v = CsvString(stringValue)
		return nil
	}

	return fmt.Errorf("unmarshal to CsvString failed, %s, %s", err1, err2)
}
