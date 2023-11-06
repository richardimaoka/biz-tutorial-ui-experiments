package test_util

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/jsonwrap"
)

func JsonRead(t *testing.T, filePath string, v interface{}) {
	err := jsonwrap.JsonRead(filePath, v)
	if err != nil {
		t.Fatalf("failed to read file %s, %s", filePath, err)
		return
	}
}
