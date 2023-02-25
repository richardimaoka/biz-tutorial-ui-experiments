package model2

import (
	"testing"
)

func TestNewSourceCode(t *testing.T) {
	sc := newSourceCode()
	compareAfterMarshal(t, "testdata/new-source-code.json", sc)
}
