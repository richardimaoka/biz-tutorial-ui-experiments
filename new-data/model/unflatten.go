package model

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type MapType = map[string]interface{}

func Unflatten(bytes []byte, m *MapType) error {
	var unmarshaled map[string]interface{}
	err := json.Unmarshal(bytes, &unmarshaled)
	if err != nil {
		return err
	}

	*m = make(MapType)
	return unflattenMap(unmarshaled, *m)
}

func unflattenMap(mOriginal MapType, mTarget MapType) error {
	for k, v := range mOriginal {
		err := set(mTarget, k, k, v)
		if err != nil {
			return err
		}
	}

	return nil
}

func set(m MapType, originalKey, key string, v interface{}) error {
	//pre-condition check
	existingValue, already := m[key]
	if already {
		return fmt.Errorf("key = %s (%s) alrady has value = %v, so failed to set value = %v", originalKey, key, existingValue, v)
	}

	// Pattern 1: if key doens't have '.', then simply set value and return
	var dotIndex = strings.IndexRune(key, '.')
	if dotIndex == -1 {
		if v == nil {
			m[key] = nil
			return nil
		} else {
			t := reflect.TypeOf(v)
			switch t.Kind() {
			case reflect.Slice:
				return fmt.Errorf("key = %s (%s) has value = %v of type = %v, but slice value can't be handled", originalKey, key, v, t)
			case reflect.Array:
				return fmt.Errorf("key = %s (%s) has value = %v of type = %v, but array value can't be handled", originalKey, key, v, t)
			case reflect.Map:
				return fmt.Errorf("key = %s (%s) has value = %v of type = %v, but map value can't be handled", originalKey, key, v, t)
			default:
				m[key] = v
				return nil
			}
		}
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

		return set(mm, originalKey, childKey, v)

	} else {
		// Pattern 2.2: parent key not used before
		mm := make(map[string]interface{})
		m[parentKey] = mm
		return set(mm, originalKey, childKey, v)
	}
}
