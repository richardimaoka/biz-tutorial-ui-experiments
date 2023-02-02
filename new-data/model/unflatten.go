package model

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func unflatten(bytes []byte) (map[string]interface{}, error) {
	var unmarshaled map[string]interface{}
	err := json.Unmarshal(bytes, &unmarshaled)
	if err != nil {
		panic(err)
	}

	return unflattenMap(unmarshaled)
}

func set(m map[string]interface{}, originalKey, key string, v interface{}) error {
	//pre-condition check
	existingValue, already := m[key]
	if already {
		return fmt.Errorf("key = %s (%s) alrady has value = %v, so failed to set value = %v", originalKey, key, existingValue, v)
	}
	if v == nil {
		m[key] = nil
		return nil
	}
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Slice:
		return fmt.Errorf("key = %s (%s) has value = %v of type = %v, but slice value can't be handled", originalKey, key, v, t)
	case reflect.Array:
		return fmt.Errorf("key = %s (%s) has value = %v of type = %v, but array value can't be handled", originalKey, key, v, t)
	case reflect.Map:
		return fmt.Errorf("key = %s (%s) has value = %v of type = %v, but map value can't be handled", originalKey, key, v, t)
	}

	// Pattern 1: if key doens't have '.', then simply set value and return
	var dotIndex = strings.IndexRune(key, '.')
	if dotIndex == -1 {
		m[key] = v
		return nil
	}

	// Pattern 2: if key has '.', then make children
	parentKey := key[0:dotIndex]
	childKey := key[dotIndex+1:]

	_, already = m[parentKey]
	if already {
		// Pattern 2.1: parent key already used
		mm, ok := m[parentKey].(map[string]interface{})
		if !ok {
			return fmt.Errorf("key = %s already has a non-map value at %s, which is supposedly dupe, so failed to set the value = %v", originalKey, key, v)
		}
		set(mm, originalKey, childKey, v)

		return nil
	} else {
		// Pattern 2.2: parent key not used before
		mm := make(map[string]interface{})
		err := set(mm, originalKey, childKey, v)
		if err != nil {
			return err
		}
		m[parentKey] = mm

		return nil
	}
}

func unflattenMap(mo map[string]interface{}) (map[string]interface{}, error) {
	m := make(map[string]interface{})

	for k, v := range mo {
		err := set(m, k, k, v)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}
