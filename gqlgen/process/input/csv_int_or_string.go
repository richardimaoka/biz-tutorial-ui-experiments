package input

import (
	"encoding/json"
	"fmt"
)

type IntOrStringFlag string

const (
	StringFlag IntOrStringFlag = "string"
	IntFlag    IntOrStringFlag = "int"
)

type IntOrString struct {
	intValue    int
	stringValue string
	flag        IntOrStringFlag
}

func (v *IntOrString) UnmarshalJSON(b []byte) error {
	err1 := json.Unmarshal(b, &v.intValue)
	// if no error, successfully unmarshaled to int
	if err1 == nil {
		v.flag = IntFlag
		return nil
	}
	v.intValue = 0 // clear with the zero value for int

	err2 := json.Unmarshal(b, &v.stringValue)
	if err2 == nil {
		v.flag = IntFlag
		return nil
	}

	return fmt.Errorf("unmarshal to int/string both failed, %s, %s", err1, err2)
}

func (v *IntOrString) GetInt() (int, error) {
	if v.flag == IntFlag {
		return v.intValue, nil
	} else if v.flag == StringFlag {
		return 0, fmt.Errorf("IntOrString = '%s' is not int but string", v.stringValue)
	} else {
		return 0, fmt.Errorf("IntOrString internal is somehow messed up")
	}
}

func (v *IntOrString) GetString() (string, error) {
	if v.flag == IntFlag {
		return "", fmt.Errorf("IntOrString = '%d' is not string but int", v.intValue)
	} else if v.flag == StringFlag {
		return v.stringValue, nil
	} else {
		return "", fmt.Errorf("IntOrString internal is somehow messed up")
	}
}
