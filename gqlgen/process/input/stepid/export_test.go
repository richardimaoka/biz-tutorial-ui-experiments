package stepid

import (
	"fmt"
)

//
func PredictableFinder(targetFile string) (*Finder, error) {
	finder, err := NewFinder(targetFile)
	if err != nil {
		return nil, fmt.Errorf("failed to create finder, %s", err)
	}

	// replace the generator to always return empty string
	finder.idGenerator = func() string { return "" }

	return finder, nil
}
