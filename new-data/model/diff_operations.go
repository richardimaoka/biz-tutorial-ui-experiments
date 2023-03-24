package model

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
