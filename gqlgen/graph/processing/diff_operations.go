package processing

type Diff struct {
	FilesAdded         []FileAdd         `json:"filesAdded"`
	FilesUpdated       []FileUpdate      `json:"filesUpdated"`
	FilesDeleted       []FileDelete      `json:"filesDeleted"`
	DirectoriesAdded   []DirectoryAdd    `json:"directoriesAdded"`
	DirectoriesDeleted []DirectoryDelete `json:"directoriesDeleted"`
}

func (d *Diff) Append(op FileSystemOperation) {
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
