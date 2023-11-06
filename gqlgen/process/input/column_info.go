package input

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/result"

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

func setColumns(step *result.Step, cols UsedColumns) {
	step.Column1 = cols[0]
	step.Column2 = cols[1]
	step.Column3 = cols[2]
	step.Column4 = cols[3]
	step.Column5 = cols[4]
	step.Column6 = cols[5]
	step.Column7 = cols[6]
	step.Column8 = cols[7]
	step.Column9 = cols[8]
	step.Column10 = cols[9]
}
