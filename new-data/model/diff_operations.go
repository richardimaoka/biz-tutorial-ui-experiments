package model

import (
	"encoding/json"
	"fmt"
)

type DiffEffect interface {
	IsDiffEffect()
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

func (d GitDiff) IsDiffEffect()       {}
func (d DirectoryDiff) IsDiffEffect() {}

func (d GitDiff) append(op FileSystemOperation) (GitDiff, error) {
	switch v := op.(type) {
	case FileAdd:
		d.Added = append(d.Added, v)
		return d, nil
	case FileUpdate:
		d.Updated = append(d.Updated, v)
		return d, nil
	case FileDelete:
		d.Deleted = append(d.Deleted, v)
		return d, nil
	default:
		return GitDiff{}, fmt.Errorf("GitDiff.append() received invalid operation type = %T", op)
	}
}

func (d DirectoryDiff) append(op FileSystemOperation) (DirectoryDiff, error) {
	switch v := op.(type) {
	case DirectoryAdd:
		d.Added = append(d.Added, v)
		return d, nil
	case DirectoryDelete:
		d.Deleted = append(d.Deleted, v)
		return d, nil
	default:
		return DirectoryDiff{}, fmt.Errorf("DirectoryDiff.append() received invalid operation type = %T", op)
	}
}

func InitializeDiffEffect(op FileSystemOperation) DiffEffect {
	switch v := op.(type) {
	case FileAdd:
		return GitDiff{
			Added: []FileAdd{v},
		}
	case FileUpdate:
		return GitDiff{
			Updated: []FileUpdate{v},
		}
	case FileDelete:
		return GitDiff{
			Deleted: []FileDelete{v},
		}
	case DirectoryAdd:
		return DirectoryDiff{
			Added: []DirectoryAdd{v},
		}
	case DirectoryDelete:
		return DirectoryDiff{
			Deleted: []DirectoryDelete{v},
		}
	default:
		return nil
	}
}

func AppendDiffEffect(d DiffEffect, op FileSystemOperation) (DiffEffect, error) {
	if d == nil {
		return InitializeDiffEffect(op), nil
	} else {
		switch v := d.(type) {
		case GitDiff:
			return v.append(op)
		case DirectoryDiff:
			return v.append(op)
		default:
			return nil, fmt.Errorf("AppendDiffEffect() received invalid DiffEffect type = %T", d)
		}
	}
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

func unmarshalDiffEffect(jsonBytes []byte) (DiffEffect, error) {
	diffType, err := extractTypeName(jsonBytes, "diffType")
	if err != nil {
		return nil, fmt.Errorf("unmarshalDiffEffect() failed to extract diffType: %s", err)
	}

	switch diffType {
	case "GitDiff":
		var diffEffect GitDiff
		err := json.Unmarshal(jsonBytes, &diffEffect)
		return diffEffect, err
	case "DirectoryDiff":
		var diffEffect DirectoryDiff
		err := json.Unmarshal(jsonBytes, &diffEffect)
		return diffEffect, err
	default:
		return nil, fmt.Errorf("unmarshalDiffEffect() found invalid  diffType = %s", diffType)
	}
}
