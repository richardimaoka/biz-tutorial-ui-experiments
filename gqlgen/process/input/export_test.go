package input

import "testing"

func PredictableFinder(t *testing.T, targetFile string) *StepIdFinder {
	finder, err := NewFinder(targetFile)
	if err != nil {
		t.Fatalf("failed to create finder, %s", err)
	}

	// replace the generator to always return empty string
	finder.idGenerator = func() string { return "" }

	return finder
}
