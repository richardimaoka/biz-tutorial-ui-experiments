package input

type UsedColumns = [10]string
type CurrentColumn = string

type ColumnInfo struct {
	Columns UsedColumns
	Current CurrentColumn
}

// similar to append() for slice
func appendIfNotExists(columns UsedColumns, colName string) UsedColumns {
	for _, col := range columns {
		if col == colName {
			// if already exists, do nothing
			return columns
		}
	}

	// here we didn't find the column, so append it
	for i, col := range columns {
		if col == "" {
			// columns is copied as an argument, so we can modify it without affecting the caller
			columns[i] = colName
			break
		}
	}

	return columns
}
