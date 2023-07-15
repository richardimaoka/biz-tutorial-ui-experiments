package state_test

import (
	"fmt"
	"testing"

	"github.com/go-git/go-git/v5/storage/memory"
)

var storage *memory.Storage

func TestMain(m *testing.M) {
	storage = memory.NewStorage()

	fmt.Println("before all tests")
	m.Run()
}
