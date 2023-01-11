package pkg

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type Action interface {
	IsAction()
}

type Command struct {
	TypeName string `json:"__typename"`
	Command  string
}

func (c *Command) IsAction() {}

func GetActionFromFile(filename string) (Action, error) {
	errorPreceding := "Error in getActionFromFile for filename = " + filename

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPreceding, err)
	}

	action, err := GetActionFromBytes(bytes)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPreceding, err)
	}

	return action, nil
}

func GetActionFromBytes(bytes []byte) (Action, error) {
	var unmarshaled interface{}
	if err := json.Unmarshal(bytes, &unmarshaled); err != nil {
		return nil, err
	}

	asserted, ok := unmarshaled.(map[string]interface{}) //type assertion
	if !ok {
		return nil, fmt.Errorf("perhaps the given JSON is not a JSON 'object'")
	}

	typename, ok := asserted["__typename"]
	if !ok {
		return nil, fmt.Errorf("\"__typename\" does not exist in JSON")
	}

	switch t := typename.(type) {
	case string:
		switch t {
		case "Command":
			var command Command
			if err := json.Unmarshal(bytes, &command); err != nil {
				return nil, err
			}

			return &command, nil
		default:
			return nil, fmt.Errorf("\"__typename\" = %s is not a valid action type", t)
		}
	default:
		return nil, fmt.Errorf("\"__typename\" = %v is in wrong type %v", t, reflect.TypeOf(t))
	}
}
