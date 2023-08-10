package process_test

/*
func TestReadStepEntries(t *testing.T) {
	cases := []struct {
		dirPath string
	}{
		{"testdata/basic"},
	}

	for _, c := range cases {
		effects, err := process.ReadStepEntries(c.dirPath)
		if err != nil {
			t.Fatalf("ReadStepEntries failed to read file, %s", err)
		}

		internal.CompareWitGoldenFile(t, *updateFlag, c.dirPath+"/golden/step_entries_golden.json", effects)
	}
}

func TestToGraphQLPages(t *testing.T) {
	cases := []struct {
		dirPath string
	}{
		// {"testdata/basic"},
		{"testdata/sign-in-with-google"},
	}

	for _, c := range cases {
		effects, err := process.ReadStepEntries(c.dirPath)
		if err != nil {
			t.Fatalf("ReadStepEntries failed to read file, %s", err)
		}

		pages := effects.ToGraphQLPages()
		for i, p := range pages {
			goldenFile := fmt.Sprintf("%s/golden/pages_golden%03d.json", c.dirPath, i)
			internal.CompareWitGoldenFile(t, *updateFlag, goldenFile, p)
		}
	}
}
*/
