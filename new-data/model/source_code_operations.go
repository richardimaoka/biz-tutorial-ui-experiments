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

func readOperationFromBytes(bytes []byte) (FileSystemOperation, error) {
	typeName, err := extractTypeName(bytes, "operationType")
	if err != nil {
		return nil, fmt.Errorf("readActionFromBytes() failed to extract operationType %s", err)
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
		return nil, fmt.Errorf("readOperationFromBytes() found invalid operationType = %s", typeName)
	}
}

type GitDiff struct {
	Added   []FileAdd    `json:"added"`
	Updated []FileUpdate `json:"updated"`
	Deleted []FileDelete `json:"deleted"`
}

type DirectoryDiff struct {
	Added   []DirectoryAdd    `json:"added"`
	Deleted []DirectoryDelete `json:"deleted"`
}

func (d GitDiff) size() int {
	return len(d.Added) + len(d.Updated) + len(d.Deleted)
}

func (d DirectoryDiff) size() int {
	return len(d.Added) + len(d.Deleted)
}

func (d GitDiff) findDuplicate() GitDiff {
	var filePathUnion []string
	for _, v := range d.Added {
		filePathUnion = append(filePathUnion, v.FilePath)
	}
	for _, v := range d.Updated {
		filePathUnion = append(filePathUnion, v.FilePath)
	}
	for _, v := range d.Deleted {
		filePathUnion = append(filePathUnion, v.FilePath)
	}

	//find duplicate
	var found = make(map[string]int)
	for _, p := range filePathUnion {
		if count, ok := found[p]; ok {
			found[p] = count + 1
		} else {
			found[p] = 1
		}
	}
	var duplicate = make(map[string]int)
	for k, v := range found {
		if v > 1 {
			duplicate[k] = v
		}
	}

	//duplicate by add, update, and delete operation
	diffDuplicate := GitDiff{}
	for dupePath := range duplicate {
		for _, v := range d.Added {
			if v.FilePath == dupePath {
				diffDuplicate.Added = append(diffDuplicate.Added, v)
			}
		}
		for _, v := range d.Updated {
			if v.FilePath == dupePath {
				diffDuplicate.Updated = append(diffDuplicate.Updated, v)
			}
		}
		for _, v := range d.Deleted {
			if v.FilePath == dupePath {
				diffDuplicate.Deleted = append(diffDuplicate.Deleted, v)
			}
		}
	}

	return diffDuplicate
}

func (d DirectoryDiff) findDuplicate() DirectoryDiff {
	var filePathUnion []string
	for _, v := range d.Added {
		filePathUnion = append(filePathUnion, v.FilePath)
	}
	for _, v := range d.Deleted {
		filePathUnion = append(filePathUnion, v.FilePath)
	}

	//find duplicate
	var found = make(map[string]int)
	for _, p := range filePathUnion {
		if count, ok := found[p]; ok {
			found[p] = count + 1
		} else {
			found[p] = 1
		}
	}
	var duplicate = make(map[string]int)
	for k, v := range found {
		if v > 1 {
			duplicate[k] = v
		}
	}

	//duplicate by add, update, and delete operation
	diffDuplicate := DirectoryDiff{}
	for dupePath := range duplicate {
		for _, v := range d.Added {
			if v.FilePath == dupePath {
				diffDuplicate.Added = append(diffDuplicate.Added, v)
			}
		}
		for _, v := range d.Deleted {
			if v.FilePath == dupePath {
				diffDuplicate.Deleted = append(diffDuplicate.Deleted, v)
			}
		}
	}

	return diffDuplicate
}
