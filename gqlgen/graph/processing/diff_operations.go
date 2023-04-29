package processing

import (
	"encoding/json"
	"fmt"
)

type Diff struct {
	FilesAdded         []FileAdd         `json:"filesAdded"`
	FilesUpdated       []FileUpdate      `json:"filesUpdated"`
	FilesDeleted       []FileDelete      `json:"filesDeleted"`
	DirectoriesAdded   []DirectoryAdd    `json:"directoriesAdded"`
	DirectoriesDeleted []DirectoryDelete `json:"directoriesDeleted"`
}

type DiffEffect interface {
	IsDiffEffect()
}

// TODO: remove this
type GitDiff struct {
	Added   []FileAdd    `json:"added"`
	Updated []FileUpdate `json:"updated"`
	Deleted []FileDelete `json:"deleted"`
}

// TODO: remove this
type DirectoryDiff struct {
	Added   []DirectoryAdd    `json:"added"`
	Deleted []DirectoryDelete `json:"deleted"`
}

func (d GitDiff) IsDiffEffect()       {}
func (d DirectoryDiff) IsDiffEffect() {}

func (d GitDiff) size() int {
	return len(d.Added) + len(d.Updated) + len(d.Deleted)
}

func (d DirectoryDiff) size() int {
	return len(d.Added) + len(d.Deleted)
}

func (d Diff) size() int {
	return len(d.FilesAdded) +
		len(d.FilesDeleted) +
		len(d.FilesAdded) +
		len(d.DirectoriesDeleted) +
		len(d.DirectoriesAdded)
}

func (d GitDiff) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["diffType"] = "GitDiff"
	m["added"] = d.Added
	m["updated"] = d.Updated
	m["deleted"] = d.Deleted

	return json.Marshal(m)
}

func (d DirectoryDiff) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["diffType"] = "DirectoryDiff"
	m["added"] = d.Added
	m["deleted"] = d.Deleted

	return json.Marshal(m)
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
	var filePathCount = make(map[string]int)
	for _, p := range filePathUnion {
		if count, ok := filePathCount[p]; ok {
			filePathCount[p] = count + 1
		} else {
			filePathCount[p] = 1
		}
	}
	var duplicate = make(map[string]int)
	for k, v := range filePathCount {
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
	var filePathCount = make(map[string]int)
	for _, p := range filePathUnion {
		if count, ok := filePathCount[p]; ok {
			filePathCount[p] = count + 1
		} else {
			filePathCount[p] = 1
		}
	}
	var duplicate = make(map[string]int)
	for k, v := range filePathCount {
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

func (d Diff) findDuplicate() Diff {
	var filePathUnion []string
	for _, v := range d.FilesAdded {
		filePathUnion = append(filePathUnion, v.FilePath)
	}
	for _, v := range d.FilesUpdated {
		filePathUnion = append(filePathUnion, v.FilePath)
	}
	for _, v := range d.FilesDeleted {
		filePathUnion = append(filePathUnion, v.FilePath)
	}
	for _, v := range d.DirectoriesAdded {
		filePathUnion = append(filePathUnion, v.FilePath)
	}
	for _, v := range d.DirectoriesDeleted {
		filePathUnion = append(filePathUnion, v.FilePath)
	}

	//find duplicate
	var filePathCount = make(map[string]int)
	for _, p := range filePathUnion {
		if count, ok := filePathCount[p]; ok {
			filePathCount[p] = count + 1
		} else {
			filePathCount[p] = 1
		}
	}
	var duplicate = make(map[string]int)
	for k, v := range filePathCount {
		if v > 1 {
			duplicate[k] = v
		}
	}

	//duplicate by add, update, and delete operation
	diffDuplicate := Diff{}
	for dupePath := range duplicate {
		for _, v := range d.FilesAdded {
			if v.FilePath == dupePath {
				diffDuplicate.FilesAdded = append(diffDuplicate.FilesAdded, v)
			}
		}
		for _, v := range d.FilesUpdated {
			if v.FilePath == dupePath {
				diffDuplicate.FilesUpdated = append(diffDuplicate.FilesUpdated, v)
			}
		}
		for _, v := range d.FilesDeleted {
			if v.FilePath == dupePath {
				diffDuplicate.FilesDeleted = append(diffDuplicate.FilesDeleted, v)
			}
		}
		for _, v := range d.DirectoriesAdded {
			if v.FilePath == dupePath {
				diffDuplicate.DirectoriesAdded = append(diffDuplicate.DirectoriesAdded, v)
			}
		}
		for _, v := range d.DirectoriesDeleted {
			if v.FilePath == dupePath {
				diffDuplicate.DirectoriesDeleted = append(diffDuplicate.DirectoriesDeleted, v)
			}
		}
	}

	return diffDuplicate
}

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

func (d *Diff) append(op FileSystemOperation) {
	switch v := op.(type) {
	case FileAdd:
		d.FilesAdded = append(d.FilesAdded, v)
	case FileUpdate:
		d.FilesUpdated = append(d.FilesUpdated, v)
	case FileDelete:
		d.FilesDeleted = append(d.FilesDeleted, v)
	case DirectoryAdd:
		d.DirectoriesAdded = append(d.DirectoriesAdded, v)
	case DirectoryDelete:
		d.DirectoriesDeleted = append(d.DirectoriesDeleted, v)
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

func InitializeDiff(op FileSystemOperation) *Diff {
	switch v := op.(type) {
	case FileAdd:
		return &Diff{
			FilesAdded: []FileAdd{v},
		}
	case FileUpdate:
		return &Diff{
			FilesUpdated: []FileUpdate{v},
		}
	case FileDelete:
		return &Diff{
			FilesDeleted: []FileDelete{v},
		}
	case DirectoryAdd:
		return &Diff{
			DirectoriesAdded: []DirectoryAdd{v},
		}
	case DirectoryDelete:
		return &Diff{
			DirectoriesDeleted: []DirectoryDelete{v},
		}
	default:
		return nil // this should never happen
	}
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
