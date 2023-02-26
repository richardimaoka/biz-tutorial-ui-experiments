package model2

func openFile(filePath string, content string) *OpenFile {
	split := filePathPtrSlice(filePath)
	trueValue := true

	file := OpenFile{
		FilePath:      split,
		FileName:      split[len(split)-1],
		Content:       &content,
		IsFullContent: &trueValue,
	}

	return &file
}
