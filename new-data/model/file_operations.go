package model

import (
	"encoding/json"
	"fmt"
)

type FileSystemOperation interface {
	IsFileSystemOperation()
}

type DirectoryAdd struct {
	FilePath string `json:"filePath"`
}

type DirectoryDelete struct {
	FilePath string `json:"filePath"`
}

type FileAdd struct {
	FilePath      string `json:"filePath"`
	Content       string `json:"content"`
	IsFullContent bool   `json:"isFullContent"`
}

type FileUpdate struct {
	FilePath string `json:"filePath"`
	Content  string `json:"content"`
}

type FileDelete struct {
	FilePath string `json:"filePath"`
}

func (o DirectoryAdd) IsFileSystemOperation()    {}
func (o DirectoryDelete) IsFileSystemOperation() {}
func (o FileAdd) IsFileSystemOperation()         {}
func (o FileUpdate) IsFileSystemOperation()      {}
func (o FileDelete) IsFileSystemOperation()      {}

//marshal DirectoryAdd to json string
func (o DirectoryAdd) MarshalJSON() ([]byte, error) {
	typeName := "DirectoryAdd"
	m := make(map[string]interface{})
	m["operationType"] = &typeName
	m["filePath"] = o.FilePath

	return json.Marshal(m)
}

func (o DirectoryDelete) MarshalJSON() ([]byte, error) {
	typeName := "DirectoryDelete"

	m := make(map[string]interface{})
	m["operationType"] = &typeName
	m["filePath"] = o.FilePath

	return json.Marshal(m)
}

func (o FileAdd) MarshalJSON() ([]byte, error) {
	typeName := "FileAdd"

	m := make(map[string]interface{})
	m["operationType"] = &typeName
	m["filePath"] = o.FilePath
	m["content"] = o.Content
	m["isFullContent"] = o.IsFullContent

	return json.Marshal(m)
}

func (o FileUpdate) MarshalJSON() ([]byte, error) {
	typeName := "FileUpdate"

	m := make(map[string]interface{})
	m["operationType"] = &typeName
	m["filePath"] = o.FilePath
	m["content"] = o.Content

	return json.Marshal(m)
}

func (o FileDelete) MarshalJSON() ([]byte, error) {
	typeName := "FileDelete"

	m := make(map[string]interface{})
	m["operationType"] = &typeName
	m["filePath"] = o.FilePath

	return json.Marshal(m)
}

func unmarshalFileSystemOperation(bytes []byte) (FileSystemOperation, error) {
	typeName, err := extractTypeName(bytes, "operationType")
	if err != nil {
		return nil, fmt.Errorf("unmarshalAction() failed to extract operationType %s", err)
	}

	switch typeName {
	case "FileAdd":
		var op FileAdd
		if err := json.Unmarshal(bytes, &op); err != nil {
			return nil, err
		}
		return op, nil
	case "FileUpdate":
		var op FileUpdate
		if err := json.Unmarshal(bytes, &op); err != nil {
			return nil, err
		}
		return op, nil
	case "FileDelete":
		var op FileDelete
		if err := json.Unmarshal(bytes, &op); err != nil {
			return nil, err
		}
		return op, nil
	case "DirectoryAdd":
		var op DirectoryAdd
		if err := json.Unmarshal(bytes, &op); err != nil {
			return nil, err
		}
		return op, nil
	case "DirectoryDelete":
		var op DirectoryDelete
		if err := json.Unmarshal(bytes, &op); err != nil {
			return nil, err
		}
		return op, nil
	default:
		return nil, fmt.Errorf("unmarshalFileSystemOperation() found invalid operationType = %s", typeName)
	}
}
