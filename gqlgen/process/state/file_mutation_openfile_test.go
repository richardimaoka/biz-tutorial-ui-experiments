package state_test

// state.File is effectively immutable, so no need to test mutation to the state
/*
func TestOpenFileMutation1(t *testing.T) {
	s, err := fileStateFromCommit(
		"https://github.com/richardimaoka/next-sandbox.git",
		"8adac375628219e020d4b5957ff24f45954cbd3f", // commit = 'npx create-next-app@latest'
		"next/package.json",
	)
	if err != nil {
		t.Fatalf("failed in TestOpenFileMutation1, %s", err)
	}

	// once GraphQL model is materialized...
	gqlModel := s.ToGraphQLFileNode()
	goldenFile1 := "testdata/openfile_mutation_golden1-1.json"
	testio.CompareWithGoldenFile(t, *updateFlag, goldenFile1, gqlModel)

	// ... mutation to the state ...
  // state.File is effectively immutable, so no need to test mutation to the state

	// ... has NO effect on the materialized GraphQL model
	testio.CompareAfterMarshal(t, goldenFile1, gqlModel)

	// ... has effect on a RE-materialized GraphQL model
	goldenFile2 := "testdata/openfile_mutation_golden1-2.json"
	testio.CompareWithGoldenFile(t, *updateFlag, goldenFile2, s.ToGraphQLOpenFile())

	// and golden files are indeed different
	internal.FilesMustUnmatch(t, goldenFile1, goldenFile2)
}


func TestOpenFileMutation2(t *testing.T) {
	s, err := fileStateFromCommit(
		"https://github.com/richardimaoka/next-sandbox.git",
		"8adac375628219e020d4b5957ff24f45954cbd3f", // commit = 'npx create-next-app@latest'
		"next/package.json",
	)
	if err != nil {
		t.Fatalf("failed in TestOpenFileMutation2, %s", err)
	}

	// once GraphQL model is materialized...
	gqlModel := s.ToGraphQLOpenFile()
	goldenFile1 := "testdata/openfile_mutation_golden2-1.json"
	testio.CompareWithGoldenFile(t, *updateFlag, goldenFile1, gqlModel)

	// ... mutation to the materialized model ...
	*gqlModel.FilePath = "next/package-mutated.json"
	*gqlModel.FileName = "package-mudated.json"
	*gqlModel.IsFullContent = false
	*gqlModel.Content = "mutated contents - " + *gqlModel.Content
	line100 := 100
	line200 := 200
	highlight := model.FileHighlight{
		FromLine: &line100,
		ToLine:   &line200,
	}
	gqlModel.Highlight = append(gqlModel.Highlight, &highlight)

	// ... has NO effect on a RE-materialized GraphQL model
	testio.CompareAfterMarshal(t, goldenFile1, s.ToGraphQLOpenFile())

	// ... has effect on the materialized GraphQL model
	goldenFile2 := "testdata/openfile_mutation_golden2-2.json"
	testio.CompareWithGoldenFile(t, *updateFlag, goldenFile2, gqlModel)

	// and golden files are indeed different
	internal.FilesMustUnmatch(t, goldenFile1, goldenFile2)
}
*/
