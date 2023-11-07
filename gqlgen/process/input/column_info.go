package input

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/result"

type UsedColumns [10]result.ColumnType
type CurrentColumn = result.ColumnType

type ColumnInfo struct {
	AllUsed UsedColumns
	Focus   CurrentColumn
}

// similar to append() for slice
func appendIfNotExists(columns UsedColumns, colName result.ColumnType) UsedColumns {
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

func setColumns(f *result.ColumnFields, cols UsedColumns) {
	f.Column1 = cols[0]
	f.Column2 = cols[1]
	f.Column3 = cols[2]
	f.Column4 = cols[3]
	f.Column5 = cols[4]
	f.Column6 = cols[5]
	f.Column7 = cols[6]
	f.Column8 = cols[7]
	f.Column9 = cols[8]
	f.Column10 = cols[9]
}
