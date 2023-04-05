package model

func openFile(filePath string, content string) *OpenFile {
	split := filePathPtrSlice(filePath)
	trueValue := true

	file := OpenFile{
		FilePath:      &filePath,
		FileName:      split[len(split)-1],
		Content:       &content,
		IsFullContent: &trueValue,
	}

	return &file
}
