package state

type Step struct {
	/**
	 * Step meta fields
	 */
	FromRowFields // Fields to make the step searchable for re-generation

	ColumnFields

	IntrinsicFields

	ModalFields

	AnimationFields

	/**
	 * Fields for each slide type
	 */
	TutorialTitleFields

	SectionTitleFields

	MarkdownFields

	/**
	 * Fields for each column type
	 */
	SourceFields

	TerminalFields

	BrowserFields

	BrowserDevToolsFields

	YoutubeFields
}
