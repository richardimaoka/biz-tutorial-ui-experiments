package csvfield

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Bool bool //Whether it's an integer value or a string, forcefully convert to a String in Go

func (v *Bool) UnmarshalJSON(b []byte) error {
	var boolValue bool
	if err := json.Unmarshal(b, &boolValue); err == nil {
		// if no error, successfully unmarshaled to bool
		*v = Bool(boolValue)
		return nil
	}

	var stringValue string
	if err := json.Unmarshal(b, &stringValue); err != nil {
		return fmt.Errorf("unmarshal to csvfied.Bool failed, %s", err)
	}

	switch strings.ToLower(stringValue) {
	case "":
		*v = false // zero value
		return nil
	case "true":
		*v = true
		return nil
	case "false":
		*v = true
		return nil
	default:
		return fmt.Errorf("unmarshal to csvfied.Bool failed, '%s' is invalid", stringValue)
	}
}

type String string //Whether it's an integer value or a string, forcefully convert to a String in Go

func (v *String) UnmarshalJSON(b []byte) error {
	var intValue int
	if err := json.Unmarshal(b, &intValue); err == nil {
		// if no error, successfully unmarshaled to int, then convert the int to CsvString
		*v = String(strconv.Itoa(intValue))
		return nil
	}

	var stringValue string
	if err := json.Unmarshal(b, &stringValue); err == nil {
		// if no error, successfully unmarshaled to string
		*v = String(stringValue)
		return nil
	} else {
		return fmt.Errorf("unmarshal to csvfield.String failed, %s", err)
	}
}

type Int int

func (v *Int) UnmarshalJSON(b []byte) error {
	var stringValue string
	if err := json.Unmarshal(b, &stringValue); err == nil {
		if stringValue == "" {
			*v = Int(0)
			return nil
		}

		intValue, err := strconv.Atoi(stringValue)
		if err != nil {
			return fmt.Errorf("unmarshal to csvfield.Int failed, %s", err)
		}
		*v = Int(intValue)
		return nil
	}

	var intValue int
	err := json.Unmarshal(b, &intValue)
	// if no error, successfully unmarshaled to int
	if err == nil {
		*v = Int(intValue)
		return nil
	}

	return fmt.Errorf("unmarshal to csvfield.Int failed, %s", err)
}

type MultiInt struct {
	singularValue int
	multiValues   []int
	isZeroValue   bool
	isMultiValue  bool
}

func (v *MultiInt) Delimiter() string {
	return "\n"
}

func (v *MultiInt) Length() int {
	if v.isZeroValue {
		return 0
	} else if v.isMultiValue {
		return len(v.multiValues)
	} else {
		return 1
	}
}

func (v *MultiInt) Get(index int) (int, error) {
	if v.isZeroValue {
		return 0, fmt.Errorf("trying to get [%d] of zero-value MultiInt", index)
	} else if v.isMultiValue {
		return v.multiValues[index], nil
	} else {
		return 0, fmt.Errorf("trying to get [%d] of single-value MultiInt", index)
	}
}

func (v *MultiInt) GetSingleValue() (int, error) {
	if v.isZeroValue {
		return 0, fmt.Errorf("trying to get single value of zero-value MultiInt")
	} else if v.isMultiValue {
		return 0, fmt.Errorf("trying to get single value of multi-value MultiInt")
	} else {
		return v.singularValue, nil
	}
}

func (v *MultiInt) UnmarshalJSON(b []byte) error {
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
				return fmt.Errorf("csvfield.MultiInt failed to unmarshal, `%s` cannot be converted to int, %s", n, err)
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
		return fmt.Errorf("csvfield.MultiInt failed to unmarshal, %s", err)
	}
	v.singularValue = intValue
	return nil
}

func (v MultiInt) MarshalJSON() ([]byte, error) {
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
