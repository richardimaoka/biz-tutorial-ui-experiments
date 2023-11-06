package input

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/jsonwrap"
)

func TestToBrowserSingle(t *testing.T) {
	cases := []struct {
		inputFile  string
		goldenFile string
	}{
		{"testdata/browser/browser1-1.json", "testdata/browser/browser1-1-golden.json"},
		{"testdata/browser/browser1-2.json", "testdata/browser/browser1-2-golden.json"},
		{"testdata/browser/browser1-3.json", "testdata/browser/browser1-3-golden.json"},
	}

	for _, c := range cases {
		t.Run(c.inputFile, func(t *testing.T) {
			var abst Abstract
			err := jsonwrap.JsonRead(c.inputFile, &abst)
			if err != nil {
				t.Fatal(err)
			}

			result, err := toBrowserSingle(&abst)
			if err != nil {
				t.Fatal(err)
			}

			internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFile, result)
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
			var abst Abstract
			err := jsonwrap.JsonRead(c.inputFile, &abst)
			if err != nil {
				t.Fatal(err)
			}

			_, err = toBrowserNumSeq(&abst)
			if err == nil {
				t.Fatal("expected to fail but succeeded")
			}
		})
	}
}

func TestToBrowserNumSeq(t *testing.T) {
	cases := []struct {
		inputFile  string
		goldenFile string
	}{
		{"testdata/browser/browser2-1.json", "testdata/browser/browser2-1-golden.json"},
		{"testdata/browser/browser2-2.json", "testdata/browser/browser2-2-golden.json"},
	}

	for _, c := range cases {
		t.Run(c.inputFile, func(t *testing.T) {
			var abst Abstract
			err := jsonwrap.JsonRead(c.inputFile, &abst)
			if err != nil {
				t.Fatal(err)
			}

			result, err := toBrowserNumSeq(&abst)
			if err != nil {
				t.Fatal(err)
			}

			internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFile, result)
		})
	}
}

func TestToBrowserNumSeqError(t *testing.T) {
	cases := []struct {
		inputFile string
	}{
		{"testdata/browser/browser2-error1.json"},
	}

	for _, c := range cases {
		t.Run(c.inputFile, func(t *testing.T) {
			var abst Abstract
			err := jsonwrap.JsonRead(c.inputFile, &abst)
			if err != nil {
				t.Fatal(err)
			}

			_, err = toBrowserNumSeq(&abst)
			if err == nil {
				t.Fatal("expected to fail but succeeded")
			}
		})
	}
}

func TestToBrowserSequence(t *testing.T) {
	cases := []struct {
		inputFile  string
		goldenFile string
	}{
		{"testdata/browser/browser3-1.json", "testdata/browser/browser3-1-golden.json"},
		{"testdata/browser/browser3-2.json", "testdata/browser/browser3-2-golden.json"},
		{"testdata/browser/browser3-3.json", "testdata/browser/browser3-3-golden.json"},
	}

	for _, c := range cases {
		t.Run(c.inputFile, func(t *testing.T) {
			var abst Abstract
			err := jsonwrap.JsonRead(c.inputFile, &abst)
			if err != nil {
				t.Fatal(err)
			}

			result, err := toBrowserSequence(&abst)
			if err != nil {
				t.Fatal(err)
			}

			internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFile, result)
		})
	}
}

func TestToBrowserSequenceError(t *testing.T) {
	cases := []struct {
		inputFile string
	}{
		{"testdata/browser/browser3-error1.json"},
	}

	for _, c := range cases {
		t.Run(c.inputFile, func(t *testing.T) {
			var abst Abstract
			err := jsonwrap.JsonRead(c.inputFile, &abst)
			if err != nil {
				t.Fatal(err)
			}

			_, err = toBrowserSequence(&abst)
			if err == nil {
				t.Fatal("expected to fail but succeeded")
			}
		})
	}
}
