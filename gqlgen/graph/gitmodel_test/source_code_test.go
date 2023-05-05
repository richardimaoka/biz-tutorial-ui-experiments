package gitmodel_test

import (
	"flag"
	"testing"
)

var update = flag.Bool("update", false, "update golden files")

func TestSourceCodeFromGit(t *testing.T) {
	t.Errorf("update = %t", *update)
}
