package csvfield

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
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
			return fmt.Errorf("unmarshal to CsvInt failed, %s", err)
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

type CsvMultiInt struct {
	singularValue int
	multiValues   []int
	isZeroValue   bool
	isMultiValue  bool
}

func (v *CsvMultiInt) Delimiter() string {
	return "\n"
}

func (v *CsvMultiInt) Length() int {
	if v.isZeroValue {
		return 0
	} else if v.isMultiValue {
		return len(v.multiValues)
	} else {
		return 1
	}
}

func (v *CsvMultiInt) UnmarshalJSON(b []byte) error {
	// If it is a string value, suposedly empty string "" or multi `int` values delimited by "\n"
	var stringValue string
	if err := json.Unmarshal(b, &stringValue); err == nil {
		if stringValue == "" {
			v.isZeroValue = true
			return nil
		}

		// Supposedly multi `int` values
		numberStrings := strings.Split(stringValue, v.Delimiter())
		for _, n := range numberStrings {
			intValue, err := strconv.Atoi(n)
			if err != nil {
				return fmt.Errorf("CsvMultiInt failed to unmarshal, `%s` cannot be converted to int, %s", n, err)
			}
			v.multiValues = append(v.multiValues, intValue)
			v.isMultiValue = true
		}
		return nil
	}

	// If Unmarshal to string failed above, then it should be single-value int
	var intValue int
	err := json.Unmarshal(b, &intValue)
	// if no error, successfully unmarshaled to int
	if err != nil {
		return fmt.Errorf("CsvMultiInt failed to unmarshal, %s", err)
	}
	v.singularValue = intValue
	return nil
}

func (v CsvMultiInt) MarshalJSON() ([]byte, error) {
	if v.isZeroValue {
		return nil, nil
	}
	if v.isMultiValue {
		var ss []string
		for _, i := range v.multiValues {
			ss = append(ss, strconv.Itoa(i))
		}
		joined := strings.Join(ss, v.Delimiter())
		return json.Marshal(joined)
	} else {
		return json.Marshal(v.singularValue)
	}
}
