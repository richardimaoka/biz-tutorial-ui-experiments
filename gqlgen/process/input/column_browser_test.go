package input

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/jsonwrap"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/testio"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

func TestToBrowserSingle(t *testing.T) {
	cases := []struct {
		inputFile  string
		goldenFile string
	}{
		{"testdata/browser/browser1-1.json", "testdata/browser/browser1-1-golden.json"},
		{"testdata/browser/browser1-2.json", "testdata/browser/browser1-2-golden.json"},
	}

	for _, c := range cases {
		t.Run(c.inputFile, func(t *testing.T) {
			var fromRow Row
			err := jsonwrap.Read(c.inputFile, &fromRow)
			if err != nil {
				t.Fatal(err)
			}

			result, err := toBrowserSingleRow(&fromRow)
			if err != nil {
				t.Fatal(err)
			}

			testio.CompareWithGoldenFile(t, *updateFlag, c.goldenFile, result)
		})
	}
}

func TestToBrowserSingle2(t *testing.T) {
	cases := []struct {
		inputFile  string
		goldenFile string
	}{
		{"testdata/browser/browser1-1.json", "testdata/browser/browser1-1-step-golden.json"},
		{"testdata/browser/browser1-2.json", "testdata/browser/browser1-2-step-golden.json"},
	}

	for _, c := range cases {
		t.Run(c.inputFile, func(t *testing.T) {
			var fromRow Row
			err := jsonwrap.Read(c.inputFile, &fromRow)
			if err != nil {
				t.Fatal(err)
			}

			currentColumn := state.NoColumnType
			finder := PredictableFinder(t, "testdata/browser/empty.json")

			steps, err := toBrowserSteps(&fromRow, finder, currentColumn)
			if err != nil {
				t.Error(err)
			}

			testio.CompareWithGoldenFile(t, *updateFlag, c.goldenFile, steps)
		})
	}
}

func TestToBrowserSingleError(t *testing.T) {
	cases := []struct {
		inputFile string
	}{
		{"testdata/browser/browser1-error1.json"},
	}

	for _, c := range cases {
		t.Run(c.inputFile, func(t *testing.T) {
			var fromRow Row
			err := jsonwrap.Read(c.inputFile, &fromRow)
			if err != nil {
				t.Fatal(err)
			}

			_, err = toBrowserSingleRow(&fromRow)
			if err == nil {
				t.Fatal("expected to fail but succeeded")
			}
		})
	}
}

// func TestToBrowserNumSeq(t *testing.T) {
// 	cases := []struct {
// 		inputFile  string
// 		goldenFile string
// 	}{
// 		{"testdata/browser/browser2-1.json", "testdata/browser/browser2-1-golden.json"},
// 		{"testdata/browser/browser2-2.json", "testdata/browser/browser2-2-golden.json"},
// 	}

// 	for _, c := range cases {
// 		t.Run(c.inputFile, func(t *testing.T) {
// 			var fromRow Row
// 			err := jsonwrap.Read(c.inputFile, &fromRow)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			result, err := toBrowserNumSeqRow(&fromRow)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			testio.CompareWithGoldenFile(t, *updateFlag, c.goldenFile, result)
// 		})
// 	}
// }

// func TestToBrowserNumSeqError(t *testing.T) {
// 	cases := []struct {
// 		inputFile string
// 	}{
// 		{"testdata/browser/browser2-error1.json"},
// 	}

// 	for _, c := range cases {
// 		t.Run(c.inputFile, func(t *testing.T) {
// 			var fromRow Row
// 			err := jsonwrap.Read(c.inputFile, &fromRow)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			_, err = toBrowserNumSeqRow(&fromRow)
// 			if err == nil {
// 				t.Fatal("expected to fail but succeeded")
// 			}
// 		})
// 	}
// }

// func TestToBrowserSequence(t *testing.T) {
// 	cases := []struct {
// 		inputFile  string
// 		goldenFile string
// 	}{
// 		{"testdata/browser/browser3-1.json", "testdata/browser/browser3-1-golden.json"},
// 		{"testdata/browser/browser3-2.json", "testdata/browser/browser3-2-golden.json"},
// 		{"testdata/browser/browser3-3.json", "testdata/browser/browser3-3-golden.json"},
// 	}

// 	for _, c := range cases {
// 		t.Run(c.inputFile, func(t *testing.T) {
// 			var fromRow Row
// 			err := jsonwrap.Read(c.inputFile, &fromRow)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			result, err := toBrowserSequenceRow(&fromRow)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			testio.CompareWithGoldenFile(t, *updateFlag, c.goldenFile, result)
// 		})
// 	}
// }

// func TestToBrowserSequenceError(t *testing.T) {
// 	cases := []struct {
// 		inputFile string
// 	}{
// 		{"testdata/browser/browser3-error1.json"},
// 	}

// 	for _, c := range cases {
// 		t.Run(c.inputFile, func(t *testing.T) {
// 			var fromRow Row
// 			err := jsonwrap.Read(c.inputFile, &fromRow)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			_, err = toBrowserSequenceRow(&fromRow)
// 			if err == nil {
// 				t.Fatal("expected to fail but succeeded")
// 			}
// 		})
// 	}
// }
