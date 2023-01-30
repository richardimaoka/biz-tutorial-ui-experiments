package pkg2

type Input struct {
	Action  ActionNode
	Results []Result
}

type ActionNode struct {
	Content Action
}

type Action interface {
	IsAction()
}

type Command struct {
	Command string
}

func (c *Command) IsAction() {}

type Result struct {
}

// // ActionNode has a non-empty interface field, so custom unmarshal
// // implementation is needed to make json.Unmarshal(bytes, &actionNode) work
// func (t *ActionNode) UnmarshalJSON(b []byte) error {
// 	// First, unmarshal to a general map, to avoid an unmarshal error
// 	// for the non-empty interface field(s)
// 	var unmarshald map[string]interface{}
// 	err := json.Unmarshal(b, &unmarshald)
// 	if err != nil {
// 		return err
// 	}

// 	// Second, re-marshal only the part for the non-empty interface
// 	// bytes, err := json.Marshal(unmarshald["content"])
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	// content, err := terminalElementFromBytes(bytes)
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	// t.Content = content

// 	return nil
// }

// type TerminalCommand struct {
// 	Command *string `json:"command"`
// }
// type TerminalOutput struct {
// 	Output *string `json:"output"`
// }

// func inSwitchh(typeName string, bytes []byte) error {
// 	switch typeName {
// 	case "TerminalCommand":
// 		var cmd TerminalCommand
// 		if err := json.Unmarshal(bytes, &cmd); err != nil {
// 			return err
// 		}
// 		return nil

// 	case "TerminalOutput":
// 		var output TerminalOutput
// 		if err := json.Unmarshal(bytes, &output); err != nil {
// 			return err
// 		}
// 		return nil
// 	default:
// 		return fmt.Errorf("not a valid type")
// 	}
// }

// func switchh(bytes []byte, typeField string, f func(string, []byte) error) error {
// 	var unmarshaled interface{}
// 	if err := json.Unmarshal(bytes, &unmarshaled); err != nil {
// 		return err
// 	}

// 	asserted, ok := unmarshaled.(map[string]interface{}) //type assertion
// 	if !ok {
// 		return nil, fmt.Errorf("perhaps the given JSON is not a JSON 'object', as it is unmarshaled to type = %v", reflect.TypeOf(unmarshaled))
// 	}

// 	typename, ok := asserted["__typename"]
// 	if !ok {
// 		return nil, fmt.Errorf("\"__typename\" does not exist in JSON")
// 	}
// }
