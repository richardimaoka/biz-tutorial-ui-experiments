/* eslint-disable */
import { TypedDocumentNode as DocumentNode } from "@graphql-typed-document-node/core";
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = {
  [K in keyof T]: T[K];
};
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]?: Maybe<T[SubKey]>;
};
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]: Maybe<T[SubKey]>;
};
export type MakeEmpty<
  T extends { [key: string]: unknown },
  K extends keyof T,
> = { [_ in K]?: never };
export type Incremental<T> =
  | T
  | {
      [P in keyof T]?: P extends " $fragmentName" | "__typename" ? T[P] : never;
    };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string };
  String: { input: string; output: string };
  Boolean: { input: boolean; output: boolean };
  Int: { input: number; output: number };
  Float: { input: number; output: number };
};

export type Browser = {
  __typename: "Browser";
  height: Scalars["Int"]["output"];
  path: Scalars["String"]["output"];
  width: Scalars["Int"]["output"];
};

export type BrowserColumn = Column & {
  __typename: "BrowserColumn";
  _placeholder?: Maybe<Scalars["String"]["output"]>;
  browser: Browser;
};

export type Column = {
  _placeholder?: Maybe<Scalars["String"]["output"]>;
};

export type ColumnWrapper = {
  __typename: "ColumnWrapper";
  column: Column;
  columnDisplayName?: Maybe<Scalars["String"]["output"]>;
  columnName: Scalars["String"]["output"];
};

export type EditSequence = {
  __typename: "EditSequence";
  edits?: Maybe<Array<MonacoEditOperation>>;
  id: Scalars["ID"]["output"];
};

export type FileHighlight = {
  __typename: "FileHighlight";
  fromLine?: Maybe<Scalars["Int"]["output"]>;
  toLine?: Maybe<Scalars["Int"]["output"]>;
};

export type FileNode = {
  __typename: "FileNode";
  filePath: Scalars["String"]["output"];
  isDeleted?: Maybe<Scalars["Boolean"]["output"]>;
  isUpdated?: Maybe<Scalars["Boolean"]["output"]>;
  name?: Maybe<Scalars["String"]["output"]>;
  nodeType: FileNodeType;
  offset?: Maybe<Scalars["Int"]["output"]>;
};

export type FileNodeType = "DIRECTORY" | "FILE";

export type Image = {
  __typename: "Image";
  caption?: Maybe<Scalars["String"]["output"]>;
  height: Scalars["Int"]["output"];
  src: Scalars["String"]["output"];
  width: Scalars["Int"]["output"];
};

export type ImageSlide = Slide & {
  __typename: "ImageSlide";
  _placeholder?: Maybe<Scalars["String"]["output"]>;
  image: Image;
};

export type MarkdownSlide = Slide & {
  __typename: "MarkdownSlide";
  _placeholder?: Maybe<Scalars["String"]["output"]>;
  markdownBody: Scalars["String"]["output"];
};

export type Modal = {
  __typename: "Modal";
  markdownBody?: Maybe<Scalars["String"]["output"]>;
};

export type MonacoEditOperation = {
  __typename: "MonacoEditOperation";
  range: MonacoEditRange;
  text: Scalars["String"]["output"];
};

export type MonacoEditRange = {
  __typename: "MonacoEditRange";
  endColumn: Scalars["Int"]["output"];
  endLineNumber: Scalars["Int"]["output"];
  startColumn: Scalars["Int"]["output"];
  startLineNumber: Scalars["Int"]["output"];
};

export type OpenFile = {
  __typename: "OpenFile";
  content?: Maybe<Scalars["String"]["output"]>;
  editSequence?: Maybe<EditSequence>;
  fileName?: Maybe<Scalars["String"]["output"]>;
  filePath?: Maybe<Scalars["String"]["output"]>;
  highlight?: Maybe<Array<Maybe<FileHighlight>>>;
  isFullContent?: Maybe<Scalars["Boolean"]["output"]>;
  language?: Maybe<Scalars["String"]["output"]>;
  oldContent?: Maybe<Scalars["String"]["output"]>;
  size?: Maybe<Scalars["Float"]["output"]>;
  tooltip?: Maybe<SourceCodeTooltip>;
};

export type Page = {
  __typename: "Page";
  columns?: Maybe<Array<Maybe<ColumnWrapper>>>;
  focusColumn?: Maybe<Scalars["String"]["output"]>;
  isTrivial?: Maybe<Scalars["Boolean"]["output"]>;
  modal?: Maybe<Modal>;
  mode?: Maybe<PageMode>;
  nextStep?: Maybe<Scalars["String"]["output"]>;
  prevStep?: Maybe<Scalars["String"]["output"]>;
  slide?: Maybe<SlideWrapper>;
  step?: Maybe<Scalars["String"]["output"]>;
};

export type PageMode = "HANDSON" | "SLIDESHOW";

export type Query = {
  __typename: "Query";
  _test?: Maybe<TestObjs>;
  page?: Maybe<Page>;
};

export type QueryPageArgs = {
  step?: InputMaybe<Scalars["String"]["input"]>;
  tutorial: Scalars["String"]["input"];
};

export type SectionTitleSlide = Slide & {
  __typename: "SectionTitleSlide";
  _placeholder?: Maybe<Scalars["String"]["output"]>;
  sectionNum: Scalars["Int"]["output"];
  title: Scalars["String"]["output"];
};

export type Slide = {
  _placeholder?: Maybe<Scalars["String"]["output"]>;
};

export type SlideWrapper = {
  __typename: "SlideWrapper";
  slide: Slide;
};

export type SourceCode = {
  __typename: "SourceCode";
  fileTree?: Maybe<Array<Maybe<FileNode>>>;
  isFoldFileTree?: Maybe<Scalars["Boolean"]["output"]>;
  openFile?: Maybe<OpenFile>;
  projectDir?: Maybe<Scalars["String"]["output"]>;
  step?: Maybe<Scalars["String"]["output"]>;
};

export type SourceCodeOpenFileArgs = {
  filePath?: InputMaybe<Scalars["String"]["input"]>;
};

export type SourceCodeColumn = Column & {
  __typename: "SourceCodeColumn";
  _placeholder?: Maybe<Scalars["String"]["output"]>;
  sourceCode: SourceCode;
};

export type SourceCodeTooltip = {
  __typename: "SourceCodeTooltip";
  lineNumber: Scalars["Int"]["output"];
  markdownBody: Scalars["String"]["output"];
  timing?: Maybe<SourceCodeTooltipTiming>;
};

export type SourceCodeTooltipTiming = "END" | "START";

export type Terminal = {
  __typename: "Terminal";
  currentDirectory: Scalars["String"]["output"];
  entries: Array<TerminalEntry>;
  name?: Maybe<Scalars["String"]["output"]>;
  tooltip?: Maybe<TerminalTooltip>;
};

export type TerminalColumn = Column & {
  __typename: "TerminalColumn";
  _placeholder?: Maybe<Scalars["String"]["output"]>;
  terminals: Array<Terminal>;
};

export type TerminalEntry = {
  __typename: "TerminalEntry";
  entryType: TerminalEntryType;
  id: Scalars["ID"]["output"];
  text: Scalars["String"]["output"];
};

export type TerminalEntryType = "COMMAND" | "OUTPUT";

export type TerminalTooltip = {
  __typename: "TerminalTooltip";
  markdownBody: Scalars["String"]["output"];
  timing?: Maybe<TerminalTooltipTiming>;
};

export type TerminalTooltipTiming = "END" | "START";

export type TestObjs = {
  __typename: "TestObjs";
  appTestSourcecodeFilecontentPage?: Maybe<OpenFile>;
  appTestTerminalPage?: Maybe<TerminalColumn>;
  appTestTutorialColumnsPage?: Maybe<Page>;
  appTestTutorialTutorialPage?: Maybe<Page>;
};

export type TestObjsAppTestSourcecodeFilecontentPageArgs = {
  step: Scalars["Int"]["input"];
};

export type TestObjsAppTestTerminalPageArgs = {
  step?: InputMaybe<Scalars["Int"]["input"]>;
};

export type TutorialTitleSlide = Slide & {
  __typename: "TutorialTitleSlide";
  _placeholder?: Maybe<Scalars["String"]["output"]>;
  images?: Maybe<Array<Image>>;
  title: Scalars["String"]["output"];
};

export type AppTutorialPageQueryVariables = Exact<{
  tutorial: Scalars["String"]["input"];
  step?: InputMaybe<Scalars["String"]["input"]>;
}>;

export type AppTutorialPageQuery = {
  __typename: "Query";
  page?:
    | ({ __typename: "Page"; mode?: PageMode | null } & {
        " $fragmentRefs"?: {
          GqlHandsonComponentFragment: GqlHandsonComponentFragment;
          GqlNavigationFragment: GqlNavigationFragment;
          GqlSlideshowComponentFragment: GqlSlideshowComponentFragment;
        };
      })
    | null;
};

export type GqlHandsonComponentFragment = ({ __typename: "Page" } & {
  " $fragmentRefs"?: {
    GqlHandsonHeaderFragment: GqlHandsonHeaderFragment;
    GqlColumnWrappersFragment: GqlColumnWrappersFragment;
  };
}) & { " $fragmentName"?: "GqlHandsonComponentFragment" };

export type GqlColumnWrapperFragment = {
  __typename: "ColumnWrapper";
  columnName: string;
  column:
    | { __typename: "BrowserColumn" }
    | ({ __typename: "SourceCodeColumn" } & {
        " $fragmentRefs"?: {
          GqlSourceCodeColumnFragment: GqlSourceCodeColumnFragment;
        };
      })
    | ({ __typename: "TerminalColumn" } & {
        " $fragmentRefs"?: {
          GqlTerminalColumnFragment: GqlTerminalColumnFragment;
        };
      });
} & { " $fragmentName"?: "GqlColumnWrapperFragment" };

export type GqlColumnWrappersFragment = {
  __typename: "Page";
  focusColumn?: string | null;
  columns?: Array<
    | ({ __typename: "ColumnWrapper"; columnName: string } & {
        " $fragmentRefs"?: {
          GqlColumnWrapperFragment: GqlColumnWrapperFragment;
        };
      })
    | null
  > | null;
} & { " $fragmentName"?: "GqlColumnWrappersFragment" };

export type GqlHandsonHeaderFragment = ({ __typename: "Page" } & {
  " $fragmentRefs"?: { GqlColumnTabsFragment: GqlColumnTabsFragment };
}) & { " $fragmentName"?: "GqlHandsonHeaderFragment" };

export type GqlColumnTabFragment = ({
  __typename: "ColumnWrapper";
  columnName: string;
  columnDisplayName?: string | null;
} & {
  " $fragmentRefs"?: { GqlColumnTabIconFragment: GqlColumnTabIconFragment };
}) & { " $fragmentName"?: "GqlColumnTabFragment" };

export type GqlColumnTabIconFragment = {
  __typename: "ColumnWrapper";
  column:
    | { __typename: "BrowserColumn" }
    | { __typename: "SourceCodeColumn" }
    | { __typename: "TerminalColumn" };
} & { " $fragmentName"?: "GqlColumnTabIconFragment" };

export type GqlColumnTabsFragment = {
  __typename: "Page";
  focusColumn?: string | null;
  columns?: Array<
    | ({ __typename: "ColumnWrapper"; columnName: string } & {
        " $fragmentRefs"?: { GqlColumnTabFragment: GqlColumnTabFragment };
      })
    | null
  > | null;
} & { " $fragmentName"?: "GqlColumnTabsFragment" };

export type GqlImageSlideFragment = {
  __typename: "ImageSlide";
  image: {
    __typename: "Image";
    src: string;
    width: number;
    height: number;
    caption?: string | null;
  };
} & { " $fragmentName"?: "GqlImageSlideFragment" };

export type GqlNavigationFragment = {
  __typename: "Page";
  prevStep?: string | null;
  nextStep?: string | null;
  isTrivial?: boolean | null;
} & { " $fragmentName"?: "GqlNavigationFragment" };

export type GqlSectionTitleSlideFragment = {
  __typename: "SectionTitleSlide";
  title: string;
  sectionNum: number;
} & { " $fragmentName"?: "GqlSectionTitleSlideFragment" };

export type GqlSlideWrapperFragment = {
  __typename: "SlideWrapper";
  slide:
    | ({ __typename: "ImageSlide" } & {
        " $fragmentRefs"?: { GqlImageSlideFragment: GqlImageSlideFragment };
      })
    | { __typename: "MarkdownSlide"; markdownBody: string }
    | ({ __typename: "SectionTitleSlide" } & {
        " $fragmentRefs"?: {
          GqlSectionTitleSlideFragment: GqlSectionTitleSlideFragment;
        };
      })
    | ({ __typename: "TutorialTitleSlide" } & {
        " $fragmentRefs"?: {
          GqlTutorialTitleSlideFragment: GqlTutorialTitleSlideFragment;
        };
      });
} & { " $fragmentName"?: "GqlSlideWrapperFragment" };

export type GqlSlideshowComponentFragment = {
  __typename: "Page";
  slide?:
    | ({ __typename: "SlideWrapper" } & {
        " $fragmentRefs"?: { GqlSlideWrapperFragment: GqlSlideWrapperFragment };
      })
    | null;
} & { " $fragmentName"?: "GqlSlideshowComponentFragment" };

export type GqlSourceCodeColumnFragment = {
  __typename: "SourceCodeColumn";
  sourceCode: {
    __typename: "SourceCode";
    openFile?:
      | ({ __typename: "OpenFile" } & {
          " $fragmentRefs"?: {
            GqlOpenFilePaneFragment: GqlOpenFilePaneFragment;
          };
        })
      | null;
  } & {
    " $fragmentRefs"?: { GqlFileTreePaneFragment: GqlFileTreePaneFragment };
  };
} & { " $fragmentName"?: "GqlSourceCodeColumnFragment" };

export type GqlFileTreePaneFragment = ({
  __typename: "SourceCode";
  isFoldFileTree?: boolean | null;
} & {
  " $fragmentRefs"?: {
    GqlFileTreeHeaderFragment: GqlFileTreeHeaderFragment;
    GqlFileTreeComponentFragment: GqlFileTreeComponentFragment;
  };
}) & { " $fragmentName"?: "GqlFileTreePaneFragment" };

export type GqlFileNodeComponentFragment = ({
  __typename: "FileNode";
  nodeType: FileNodeType;
  name?: string | null;
  filePath: string;
  offset?: number | null;
  isUpdated?: boolean | null;
} & {
  " $fragmentRefs"?: { GqlFileNodeIconFragment: GqlFileNodeIconFragment };
}) & { " $fragmentName"?: "GqlFileNodeComponentFragment" };

export type GqlFileNodeIconFragment = {
  __typename: "FileNode";
  nodeType: FileNodeType;
} & { " $fragmentName"?: "GqlFileNodeIconFragment" };

export type GqlFileTreeHeaderFragment = {
  __typename: "SourceCode";
  projectDir?: string | null;
} & { " $fragmentName"?: "GqlFileTreeHeaderFragment" };

export type GqlFileTreeComponentFragment = {
  __typename: "SourceCode";
  fileTree?: Array<
    | ({ __typename: "FileNode"; filePath: string } & {
        " $fragmentRefs"?: {
          GqlFileNodeComponentFragment: GqlFileNodeComponentFragment;
        };
      })
    | null
  > | null;
} & { " $fragmentName"?: "GqlFileTreeComponentFragment" };

export type GqlOpenFilePaneFragment = ({ __typename: "OpenFile" } & {
  " $fragmentRefs"?: {
    GqlFileNameTabBarFragment: GqlFileNameTabBarFragment;
    GqlSourceCodeEditorFragment: GqlSourceCodeEditorFragment;
  };
}) & { " $fragmentName"?: "GqlOpenFilePaneFragment" };

export type GqlSourceCodeEditorFragment = {
  __typename: "OpenFile";
  content?: string | null;
  oldContent?: string | null;
  language?: string | null;
  editSequence?: {
    __typename: "EditSequence";
    id: string;
    edits?: Array<{
      __typename: "MonacoEditOperation";
      text: string;
      range: {
        __typename: "MonacoEditRange";
        startLineNumber: number;
        startColumn: number;
        endLineNumber: number;
        endColumn: number;
      };
    }> | null;
  } | null;
  tooltip?: {
    __typename: "SourceCodeTooltip";
    markdownBody: string;
    lineNumber: number;
    timing?: SourceCodeTooltipTiming | null;
  } | null;
} & { " $fragmentName"?: "GqlSourceCodeEditorFragment" };

export type GqlFileNameTabBarFragment = {
  __typename: "OpenFile";
  fileName?: string | null;
} & { " $fragmentName"?: "GqlFileNameTabBarFragment" };

export type GqlTerminalColumnFragment = ({
  __typename: "TerminalColumn";
  terminals: Array<
    { __typename: "Terminal" } & {
      " $fragmentRefs"?: {
        GqlTerminalContentsFragment: GqlTerminalContentsFragment;
      };
    }
  >;
} & {
  " $fragmentRefs"?: { GqlTerminalHeaderFragment: GqlTerminalHeaderFragment };
}) & { " $fragmentName"?: "GqlTerminalColumnFragment" };

export type GqlTerminalContentsFragment = {
  __typename: "Terminal";
  entries: Array<
    { __typename: "TerminalEntry"; id: string } & {
      " $fragmentRefs"?: {
        GqlTerminalEntryComponentFragment: GqlTerminalEntryComponentFragment;
      };
    }
  >;
  tooltip?:
    | ({ __typename: "TerminalTooltip" } & {
        " $fragmentRefs"?: {
          GqlTerminalTooltipFragment: GqlTerminalTooltipFragment;
        };
      })
    | null;
} & { " $fragmentName"?: "GqlTerminalContentsFragment" };

export type GqlTerminalEntryComponentFragment = {
  __typename: "TerminalEntry";
  entryType: TerminalEntryType;
  text: string;
} & { " $fragmentName"?: "GqlTerminalEntryComponentFragment" };

export type GqlTerminalHeaderFragment = {
  __typename: "TerminalColumn";
  terminals: Array<{
    __typename: "Terminal";
    name?: string | null;
    currentDirectory: string;
  }>;
} & { " $fragmentName"?: "GqlTerminalHeaderFragment" };

export type GqlTerminalTooltipFragment = {
  __typename: "TerminalTooltip";
  markdownBody: string;
  timing?: TerminalTooltipTiming | null;
} & { " $fragmentName"?: "GqlTerminalTooltipFragment" };

export type GqlTutorialTitleSlideFragment = {
  __typename: "TutorialTitleSlide";
  title: string;
  images?: Array<{
    __typename: "Image";
    src: string;
    width: number;
    height: number;
    caption?: string | null;
  }> | null;
} & { " $fragmentName"?: "GqlTutorialTitleSlideFragment" };

export type AppTestTutorialColumnsPageQueryVariables = Exact<{
  [key: string]: never;
}>;

export type AppTestTutorialColumnsPageQuery = {
  __typename: "Query";
  _test?: {
    __typename: "TestObjs";
    appTestTutorialColumnsPage?:
      | ({ __typename: "Page" } & {
          " $fragmentRefs"?: {
            GqlHandsonComponentFragment: GqlHandsonComponentFragment;
          };
        })
      | null;
  } | null;
};

export type AppTestTerminalPageQueryVariables = Exact<{
  step?: InputMaybe<Scalars["Int"]["input"]>;
}>;

export type AppTestTerminalPageQuery = {
  __typename: "Query";
  _test?: {
    __typename: "TestObjs";
    appTestTerminalPage?:
      | ({ __typename: "TerminalColumn" } & {
          " $fragmentRefs"?: {
            GqlTerminalColumnFragment: GqlTerminalColumnFragment;
          };
        })
      | null;
  } | null;
};

export const GqlColumnTabIconFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlColumnTabIcon" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ColumnWrapper" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "column" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "__typename" } },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlColumnTabIconFragment, unknown>;
export const GqlColumnTabFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlColumnTab" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ColumnWrapper" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "columnName" } },
          { kind: "Field", name: { kind: "Name", value: "columnDisplayName" } },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlColumnTabIcon" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlColumnTabIcon" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ColumnWrapper" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "column" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "__typename" } },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlColumnTabFragment, unknown>;
export const GqlColumnTabsFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlColumnTabs" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "columns" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "columnName" } },
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlColumnTab" },
                },
              ],
            },
          },
          { kind: "Field", name: { kind: "Name", value: "focusColumn" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlColumnTabIcon" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ColumnWrapper" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "column" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "__typename" } },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlColumnTab" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ColumnWrapper" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "columnName" } },
          { kind: "Field", name: { kind: "Name", value: "columnDisplayName" } },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlColumnTabIcon" },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlColumnTabsFragment, unknown>;
export const GqlHandsonHeaderFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlHandsonHeader" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlColumnTabs" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlColumnTabIcon" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ColumnWrapper" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "column" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "__typename" } },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlColumnTab" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ColumnWrapper" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "columnName" } },
          { kind: "Field", name: { kind: "Name", value: "columnDisplayName" } },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlColumnTabIcon" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlColumnTabs" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "columns" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "columnName" } },
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlColumnTab" },
                },
              ],
            },
          },
          { kind: "Field", name: { kind: "Name", value: "focusColumn" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlHandsonHeaderFragment, unknown>;
export const GqlTerminalHeaderFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalHeader" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "terminals" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "name" } },
                {
                  kind: "Field",
                  name: { kind: "Name", value: "currentDirectory" },
                },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlTerminalHeaderFragment, unknown>;
export const GqlTerminalEntryComponentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalEntryComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalEntry" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "entryType" } },
          { kind: "Field", name: { kind: "Name", value: "text" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlTerminalEntryComponentFragment, unknown>;
export const GqlTerminalTooltipFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalTooltip" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalTooltip" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "markdownBody" } },
          { kind: "Field", name: { kind: "Name", value: "timing" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlTerminalTooltipFragment, unknown>;
export const GqlTerminalContentsFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalContents" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "entries" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "id" } },
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlTerminalEntryComponent" },
                },
              ],
            },
          },
          {
            kind: "Field",
            name: { kind: "Name", value: "tooltip" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlTerminalTooltip" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalEntryComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalEntry" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "entryType" } },
          { kind: "Field", name: { kind: "Name", value: "text" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalTooltip" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalTooltip" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "markdownBody" } },
          { kind: "Field", name: { kind: "Name", value: "timing" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlTerminalContentsFragment, unknown>;
export const GqlTerminalColumnFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalColumn" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlTerminalHeader" },
          },
          {
            kind: "Field",
            name: { kind: "Name", value: "terminals" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlTerminalContents" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalEntryComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalEntry" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "entryType" } },
          { kind: "Field", name: { kind: "Name", value: "text" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalTooltip" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalTooltip" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "markdownBody" } },
          { kind: "Field", name: { kind: "Name", value: "timing" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalHeader" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "terminals" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "name" } },
                {
                  kind: "Field",
                  name: { kind: "Name", value: "currentDirectory" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalContents" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "entries" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "id" } },
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlTerminalEntryComponent" },
                },
              ],
            },
          },
          {
            kind: "Field",
            name: { kind: "Name", value: "tooltip" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlTerminalTooltip" },
                },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlTerminalColumnFragment, unknown>;
export const GqlFileTreeHeaderFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileTreeHeader" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "projectDir" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlFileTreeHeaderFragment, unknown>;
export const GqlFileNodeIconFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileNodeIcon" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "FileNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "nodeType" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlFileNodeIconFragment, unknown>;
export const GqlFileNodeComponentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileNodeComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "FileNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlFileNodeIcon" },
          },
          { kind: "Field", name: { kind: "Name", value: "nodeType" } },
          { kind: "Field", name: { kind: "Name", value: "name" } },
          { kind: "Field", name: { kind: "Name", value: "filePath" } },
          { kind: "Field", name: { kind: "Name", value: "offset" } },
          { kind: "Field", name: { kind: "Name", value: "isUpdated" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileNodeIcon" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "FileNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "nodeType" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlFileNodeComponentFragment, unknown>;
export const GqlFileTreeComponentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileTreeComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "fileTree" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "filePath" } },
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlFileNodeComponent" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileNodeIcon" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "FileNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "nodeType" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileNodeComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "FileNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlFileNodeIcon" },
          },
          { kind: "Field", name: { kind: "Name", value: "nodeType" } },
          { kind: "Field", name: { kind: "Name", value: "name" } },
          { kind: "Field", name: { kind: "Name", value: "filePath" } },
          { kind: "Field", name: { kind: "Name", value: "offset" } },
          { kind: "Field", name: { kind: "Name", value: "isUpdated" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlFileTreeComponentFragment, unknown>;
export const GqlFileTreePaneFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileTreePane" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlFileTreeHeader" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlFileTreeComponent" },
          },
          { kind: "Field", name: { kind: "Name", value: "isFoldFileTree" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileNodeIcon" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "FileNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "nodeType" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileNodeComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "FileNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlFileNodeIcon" },
          },
          { kind: "Field", name: { kind: "Name", value: "nodeType" } },
          { kind: "Field", name: { kind: "Name", value: "name" } },
          { kind: "Field", name: { kind: "Name", value: "filePath" } },
          { kind: "Field", name: { kind: "Name", value: "offset" } },
          { kind: "Field", name: { kind: "Name", value: "isUpdated" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileTreeHeader" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "projectDir" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileTreeComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "fileTree" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "filePath" } },
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlFileNodeComponent" },
                },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlFileTreePaneFragment, unknown>;
export const GqlFileNameTabBarFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileNameTabBar" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "fileName" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlFileNameTabBarFragment, unknown>;
export const GqlSourceCodeEditorFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlSourceCodeEditor" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "content" } },
          { kind: "Field", name: { kind: "Name", value: "oldContent" } },
          { kind: "Field", name: { kind: "Name", value: "language" } },
          {
            kind: "Field",
            name: { kind: "Name", value: "editSequence" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "id" } },
                {
                  kind: "Field",
                  name: { kind: "Name", value: "edits" },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      { kind: "Field", name: { kind: "Name", value: "text" } },
                      {
                        kind: "Field",
                        name: { kind: "Name", value: "range" },
                        selectionSet: {
                          kind: "SelectionSet",
                          selections: [
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "startLineNumber" },
                            },
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "startColumn" },
                            },
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "endLineNumber" },
                            },
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "endColumn" },
                            },
                          ],
                        },
                      },
                    ],
                  },
                },
              ],
            },
          },
          {
            kind: "Field",
            name: { kind: "Name", value: "tooltip" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "Field",
                  name: { kind: "Name", value: "markdownBody" },
                },
                { kind: "Field", name: { kind: "Name", value: "lineNumber" } },
                { kind: "Field", name: { kind: "Name", value: "timing" } },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlSourceCodeEditorFragment, unknown>;
export const GqlOpenFilePaneFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlOpenFilePane" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlFileNameTabBar" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlSourceCodeEditor" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileNameTabBar" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "fileName" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlSourceCodeEditor" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "content" } },
          { kind: "Field", name: { kind: "Name", value: "oldContent" } },
          { kind: "Field", name: { kind: "Name", value: "language" } },
          {
            kind: "Field",
            name: { kind: "Name", value: "editSequence" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "id" } },
                {
                  kind: "Field",
                  name: { kind: "Name", value: "edits" },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      { kind: "Field", name: { kind: "Name", value: "text" } },
                      {
                        kind: "Field",
                        name: { kind: "Name", value: "range" },
                        selectionSet: {
                          kind: "SelectionSet",
                          selections: [
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "startLineNumber" },
                            },
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "startColumn" },
                            },
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "endLineNumber" },
                            },
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "endColumn" },
                            },
                          ],
                        },
                      },
                    ],
                  },
                },
              ],
            },
          },
          {
            kind: "Field",
            name: { kind: "Name", value: "tooltip" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "Field",
                  name: { kind: "Name", value: "markdownBody" },
                },
                { kind: "Field", name: { kind: "Name", value: "lineNumber" } },
                { kind: "Field", name: { kind: "Name", value: "timing" } },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlOpenFilePaneFragment, unknown>;
export const GqlSourceCodeColumnFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlSourceCodeColumn" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCodeColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "sourceCode" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlFileTreePane" },
                },
                {
                  kind: "Field",
                  name: { kind: "Name", value: "openFile" },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "GqlOpenFilePane" },
                      },
                    ],
                  },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileTreeHeader" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "projectDir" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileNodeIcon" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "FileNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "nodeType" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileNodeComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "FileNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlFileNodeIcon" },
          },
          { kind: "Field", name: { kind: "Name", value: "nodeType" } },
          { kind: "Field", name: { kind: "Name", value: "name" } },
          { kind: "Field", name: { kind: "Name", value: "filePath" } },
          { kind: "Field", name: { kind: "Name", value: "offset" } },
          { kind: "Field", name: { kind: "Name", value: "isUpdated" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileTreeComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "fileTree" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "filePath" } },
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlFileNodeComponent" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileNameTabBar" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "fileName" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlSourceCodeEditor" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "content" } },
          { kind: "Field", name: { kind: "Name", value: "oldContent" } },
          { kind: "Field", name: { kind: "Name", value: "language" } },
          {
            kind: "Field",
            name: { kind: "Name", value: "editSequence" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "id" } },
                {
                  kind: "Field",
                  name: { kind: "Name", value: "edits" },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      { kind: "Field", name: { kind: "Name", value: "text" } },
                      {
                        kind: "Field",
                        name: { kind: "Name", value: "range" },
                        selectionSet: {
                          kind: "SelectionSet",
                          selections: [
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "startLineNumber" },
                            },
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "startColumn" },
                            },
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "endLineNumber" },
                            },
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "endColumn" },
                            },
                          ],
                        },
                      },
                    ],
                  },
                },
              ],
            },
          },
          {
            kind: "Field",
            name: { kind: "Name", value: "tooltip" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "Field",
                  name: { kind: "Name", value: "markdownBody" },
                },
                { kind: "Field", name: { kind: "Name", value: "lineNumber" } },
                { kind: "Field", name: { kind: "Name", value: "timing" } },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileTreePane" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlFileTreeHeader" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlFileTreeComponent" },
          },
          { kind: "Field", name: { kind: "Name", value: "isFoldFileTree" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlOpenFilePane" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlFileNameTabBar" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlSourceCodeEditor" },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlSourceCodeColumnFragment, unknown>;
export const GqlColumnWrapperFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlColumnWrapper" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ColumnWrapper" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "columnName" } },
          {
            kind: "Field",
            name: { kind: "Name", value: "column" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "__typename" } },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "TerminalColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "GqlTerminalColumn" },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "SourceCodeColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "GqlSourceCodeColumn" },
                      },
                    ],
                  },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalHeader" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "terminals" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "name" } },
                {
                  kind: "Field",
                  name: { kind: "Name", value: "currentDirectory" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalEntryComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalEntry" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "entryType" } },
          { kind: "Field", name: { kind: "Name", value: "text" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalTooltip" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalTooltip" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "markdownBody" } },
          { kind: "Field", name: { kind: "Name", value: "timing" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalContents" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "entries" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "id" } },
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlTerminalEntryComponent" },
                },
              ],
            },
          },
          {
            kind: "Field",
            name: { kind: "Name", value: "tooltip" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlTerminalTooltip" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileTreeHeader" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "projectDir" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileNodeIcon" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "FileNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "nodeType" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileNodeComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "FileNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlFileNodeIcon" },
          },
          { kind: "Field", name: { kind: "Name", value: "nodeType" } },
          { kind: "Field", name: { kind: "Name", value: "name" } },
          { kind: "Field", name: { kind: "Name", value: "filePath" } },
          { kind: "Field", name: { kind: "Name", value: "offset" } },
          { kind: "Field", name: { kind: "Name", value: "isUpdated" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileTreeComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "fileTree" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "filePath" } },
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlFileNodeComponent" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileTreePane" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlFileTreeHeader" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlFileTreeComponent" },
          },
          { kind: "Field", name: { kind: "Name", value: "isFoldFileTree" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileNameTabBar" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "fileName" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlSourceCodeEditor" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "content" } },
          { kind: "Field", name: { kind: "Name", value: "oldContent" } },
          { kind: "Field", name: { kind: "Name", value: "language" } },
          {
            kind: "Field",
            name: { kind: "Name", value: "editSequence" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "id" } },
                {
                  kind: "Field",
                  name: { kind: "Name", value: "edits" },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      { kind: "Field", name: { kind: "Name", value: "text" } },
                      {
                        kind: "Field",
                        name: { kind: "Name", value: "range" },
                        selectionSet: {
                          kind: "SelectionSet",
                          selections: [
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "startLineNumber" },
                            },
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "startColumn" },
                            },
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "endLineNumber" },
                            },
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "endColumn" },
                            },
                          ],
                        },
                      },
                    ],
                  },
                },
              ],
            },
          },
          {
            kind: "Field",
            name: { kind: "Name", value: "tooltip" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "Field",
                  name: { kind: "Name", value: "markdownBody" },
                },
                { kind: "Field", name: { kind: "Name", value: "lineNumber" } },
                { kind: "Field", name: { kind: "Name", value: "timing" } },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlOpenFilePane" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlFileNameTabBar" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlSourceCodeEditor" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalColumn" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlTerminalHeader" },
          },
          {
            kind: "Field",
            name: { kind: "Name", value: "terminals" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlTerminalContents" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlSourceCodeColumn" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCodeColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "sourceCode" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlFileTreePane" },
                },
                {
                  kind: "Field",
                  name: { kind: "Name", value: "openFile" },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "GqlOpenFilePane" },
                      },
                    ],
                  },
                },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlColumnWrapperFragment, unknown>;
export const GqlColumnWrappersFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlColumnWrappers" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "focusColumn" } },
          {
            kind: "Field",
            name: { kind: "Name", value: "columns" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "columnName" } },
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlColumnWrapper" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalHeader" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "terminals" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "name" } },
                {
                  kind: "Field",
                  name: { kind: "Name", value: "currentDirectory" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalEntryComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalEntry" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "entryType" } },
          { kind: "Field", name: { kind: "Name", value: "text" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalTooltip" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalTooltip" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "markdownBody" } },
          { kind: "Field", name: { kind: "Name", value: "timing" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalContents" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "entries" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "id" } },
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlTerminalEntryComponent" },
                },
              ],
            },
          },
          {
            kind: "Field",
            name: { kind: "Name", value: "tooltip" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlTerminalTooltip" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalColumn" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlTerminalHeader" },
          },
          {
            kind: "Field",
            name: { kind: "Name", value: "terminals" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlTerminalContents" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileTreeHeader" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "projectDir" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileNodeIcon" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "FileNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "nodeType" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileNodeComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "FileNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlFileNodeIcon" },
          },
          { kind: "Field", name: { kind: "Name", value: "nodeType" } },
          { kind: "Field", name: { kind: "Name", value: "name" } },
          { kind: "Field", name: { kind: "Name", value: "filePath" } },
          { kind: "Field", name: { kind: "Name", value: "offset" } },
          { kind: "Field", name: { kind: "Name", value: "isUpdated" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileTreeComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "fileTree" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "filePath" } },
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlFileNodeComponent" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileTreePane" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlFileTreeHeader" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlFileTreeComponent" },
          },
          { kind: "Field", name: { kind: "Name", value: "isFoldFileTree" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileNameTabBar" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "fileName" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlSourceCodeEditor" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "content" } },
          { kind: "Field", name: { kind: "Name", value: "oldContent" } },
          { kind: "Field", name: { kind: "Name", value: "language" } },
          {
            kind: "Field",
            name: { kind: "Name", value: "editSequence" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "id" } },
                {
                  kind: "Field",
                  name: { kind: "Name", value: "edits" },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      { kind: "Field", name: { kind: "Name", value: "text" } },
                      {
                        kind: "Field",
                        name: { kind: "Name", value: "range" },
                        selectionSet: {
                          kind: "SelectionSet",
                          selections: [
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "startLineNumber" },
                            },
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "startColumn" },
                            },
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "endLineNumber" },
                            },
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "endColumn" },
                            },
                          ],
                        },
                      },
                    ],
                  },
                },
              ],
            },
          },
          {
            kind: "Field",
            name: { kind: "Name", value: "tooltip" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "Field",
                  name: { kind: "Name", value: "markdownBody" },
                },
                { kind: "Field", name: { kind: "Name", value: "lineNumber" } },
                { kind: "Field", name: { kind: "Name", value: "timing" } },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlOpenFilePane" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlFileNameTabBar" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlSourceCodeEditor" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlSourceCodeColumn" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCodeColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "sourceCode" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlFileTreePane" },
                },
                {
                  kind: "Field",
                  name: { kind: "Name", value: "openFile" },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "GqlOpenFilePane" },
                      },
                    ],
                  },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlColumnWrapper" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ColumnWrapper" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "columnName" } },
          {
            kind: "Field",
            name: { kind: "Name", value: "column" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "__typename" } },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "TerminalColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "GqlTerminalColumn" },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "SourceCodeColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "GqlSourceCodeColumn" },
                      },
                    ],
                  },
                },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlColumnWrappersFragment, unknown>;
export const GqlHandsonComponentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlHandsonComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlHandsonHeader" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlColumnWrappers" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlColumnTabIcon" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ColumnWrapper" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "column" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "__typename" } },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlColumnTab" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ColumnWrapper" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "columnName" } },
          { kind: "Field", name: { kind: "Name", value: "columnDisplayName" } },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlColumnTabIcon" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlColumnTabs" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "columns" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "columnName" } },
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlColumnTab" },
                },
              ],
            },
          },
          { kind: "Field", name: { kind: "Name", value: "focusColumn" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalHeader" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "terminals" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "name" } },
                {
                  kind: "Field",
                  name: { kind: "Name", value: "currentDirectory" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalEntryComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalEntry" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "entryType" } },
          { kind: "Field", name: { kind: "Name", value: "text" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalTooltip" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalTooltip" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "markdownBody" } },
          { kind: "Field", name: { kind: "Name", value: "timing" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalContents" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "entries" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "id" } },
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlTerminalEntryComponent" },
                },
              ],
            },
          },
          {
            kind: "Field",
            name: { kind: "Name", value: "tooltip" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlTerminalTooltip" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalColumn" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlTerminalHeader" },
          },
          {
            kind: "Field",
            name: { kind: "Name", value: "terminals" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlTerminalContents" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileTreeHeader" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "projectDir" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileNodeIcon" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "FileNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "nodeType" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileNodeComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "FileNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlFileNodeIcon" },
          },
          { kind: "Field", name: { kind: "Name", value: "nodeType" } },
          { kind: "Field", name: { kind: "Name", value: "name" } },
          { kind: "Field", name: { kind: "Name", value: "filePath" } },
          { kind: "Field", name: { kind: "Name", value: "offset" } },
          { kind: "Field", name: { kind: "Name", value: "isUpdated" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileTreeComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "fileTree" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "filePath" } },
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlFileNodeComponent" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileTreePane" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlFileTreeHeader" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlFileTreeComponent" },
          },
          { kind: "Field", name: { kind: "Name", value: "isFoldFileTree" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileNameTabBar" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "fileName" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlSourceCodeEditor" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "content" } },
          { kind: "Field", name: { kind: "Name", value: "oldContent" } },
          { kind: "Field", name: { kind: "Name", value: "language" } },
          {
            kind: "Field",
            name: { kind: "Name", value: "editSequence" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "id" } },
                {
                  kind: "Field",
                  name: { kind: "Name", value: "edits" },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      { kind: "Field", name: { kind: "Name", value: "text" } },
                      {
                        kind: "Field",
                        name: { kind: "Name", value: "range" },
                        selectionSet: {
                          kind: "SelectionSet",
                          selections: [
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "startLineNumber" },
                            },
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "startColumn" },
                            },
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "endLineNumber" },
                            },
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "endColumn" },
                            },
                          ],
                        },
                      },
                    ],
                  },
                },
              ],
            },
          },
          {
            kind: "Field",
            name: { kind: "Name", value: "tooltip" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "Field",
                  name: { kind: "Name", value: "markdownBody" },
                },
                { kind: "Field", name: { kind: "Name", value: "lineNumber" } },
                { kind: "Field", name: { kind: "Name", value: "timing" } },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlOpenFilePane" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlFileNameTabBar" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlSourceCodeEditor" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlSourceCodeColumn" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCodeColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "sourceCode" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlFileTreePane" },
                },
                {
                  kind: "Field",
                  name: { kind: "Name", value: "openFile" },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "GqlOpenFilePane" },
                      },
                    ],
                  },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlColumnWrapper" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ColumnWrapper" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "columnName" } },
          {
            kind: "Field",
            name: { kind: "Name", value: "column" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "__typename" } },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "TerminalColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "GqlTerminalColumn" },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "SourceCodeColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "GqlSourceCodeColumn" },
                      },
                    ],
                  },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlHandsonHeader" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlColumnTabs" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlColumnWrappers" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "focusColumn" } },
          {
            kind: "Field",
            name: { kind: "Name", value: "columns" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "columnName" } },
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlColumnWrapper" },
                },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlHandsonComponentFragment, unknown>;
export const GqlNavigationFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlNavigation" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "prevStep" } },
          { kind: "Field", name: { kind: "Name", value: "nextStep" } },
          { kind: "Field", name: { kind: "Name", value: "isTrivial" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlNavigationFragment, unknown>;
export const GqlTutorialTitleSlideFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTutorialTitleSlide" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TutorialTitleSlide" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "title" } },
          {
            kind: "Field",
            name: { kind: "Name", value: "images" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "src" } },
                { kind: "Field", name: { kind: "Name", value: "width" } },
                { kind: "Field", name: { kind: "Name", value: "height" } },
                { kind: "Field", name: { kind: "Name", value: "caption" } },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlTutorialTitleSlideFragment, unknown>;
export const GqlSectionTitleSlideFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlSectionTitleSlide" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SectionTitleSlide" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "title" } },
          { kind: "Field", name: { kind: "Name", value: "sectionNum" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlSectionTitleSlideFragment, unknown>;
export const GqlImageSlideFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlImageSlide" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ImageSlide" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "image" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "src" } },
                { kind: "Field", name: { kind: "Name", value: "width" } },
                { kind: "Field", name: { kind: "Name", value: "height" } },
                { kind: "Field", name: { kind: "Name", value: "caption" } },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlImageSlideFragment, unknown>;
export const GqlSlideWrapperFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlSlideWrapper" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SlideWrapper" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "slide" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "__typename" } },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "TutorialTitleSlide" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "GqlTutorialTitleSlide" },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "MarkdownSlide" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "Field",
                        name: { kind: "Name", value: "markdownBody" },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "SectionTitleSlide" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "GqlSectionTitleSlide" },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "ImageSlide" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "GqlImageSlide" },
                      },
                    ],
                  },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTutorialTitleSlide" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TutorialTitleSlide" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "title" } },
          {
            kind: "Field",
            name: { kind: "Name", value: "images" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "src" } },
                { kind: "Field", name: { kind: "Name", value: "width" } },
                { kind: "Field", name: { kind: "Name", value: "height" } },
                { kind: "Field", name: { kind: "Name", value: "caption" } },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlSectionTitleSlide" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SectionTitleSlide" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "title" } },
          { kind: "Field", name: { kind: "Name", value: "sectionNum" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlImageSlide" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ImageSlide" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "image" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "src" } },
                { kind: "Field", name: { kind: "Name", value: "width" } },
                { kind: "Field", name: { kind: "Name", value: "height" } },
                { kind: "Field", name: { kind: "Name", value: "caption" } },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlSlideWrapperFragment, unknown>;
export const GqlSlideshowComponentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlSlideshowComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "slide" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlSlideWrapper" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTutorialTitleSlide" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TutorialTitleSlide" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "title" } },
          {
            kind: "Field",
            name: { kind: "Name", value: "images" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "src" } },
                { kind: "Field", name: { kind: "Name", value: "width" } },
                { kind: "Field", name: { kind: "Name", value: "height" } },
                { kind: "Field", name: { kind: "Name", value: "caption" } },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlSectionTitleSlide" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SectionTitleSlide" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "title" } },
          { kind: "Field", name: { kind: "Name", value: "sectionNum" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlImageSlide" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ImageSlide" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "image" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "src" } },
                { kind: "Field", name: { kind: "Name", value: "width" } },
                { kind: "Field", name: { kind: "Name", value: "height" } },
                { kind: "Field", name: { kind: "Name", value: "caption" } },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlSlideWrapper" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SlideWrapper" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "slide" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "__typename" } },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "TutorialTitleSlide" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "GqlTutorialTitleSlide" },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "MarkdownSlide" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "Field",
                        name: { kind: "Name", value: "markdownBody" },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "SectionTitleSlide" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "GqlSectionTitleSlide" },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "ImageSlide" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "GqlImageSlide" },
                      },
                    ],
                  },
                },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlSlideshowComponentFragment, unknown>;
export const AppTutorialPageDocument = {
  kind: "Document",
  definitions: [
    {
      kind: "OperationDefinition",
      operation: "query",
      name: { kind: "Name", value: "appTutorialPage" },
      variableDefinitions: [
        {
          kind: "VariableDefinition",
          variable: {
            kind: "Variable",
            name: { kind: "Name", value: "tutorial" },
          },
          type: {
            kind: "NonNullType",
            type: {
              kind: "NamedType",
              name: { kind: "Name", value: "String" },
            },
          },
        },
        {
          kind: "VariableDefinition",
          variable: { kind: "Variable", name: { kind: "Name", value: "step" } },
          type: { kind: "NamedType", name: { kind: "Name", value: "String" } },
        },
      ],
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "page" },
            arguments: [
              {
                kind: "Argument",
                name: { kind: "Name", value: "tutorial" },
                value: {
                  kind: "Variable",
                  name: { kind: "Name", value: "tutorial" },
                },
              },
              {
                kind: "Argument",
                name: { kind: "Name", value: "step" },
                value: {
                  kind: "Variable",
                  name: { kind: "Name", value: "step" },
                },
              },
            ],
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlHandsonComponent" },
                },
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlNavigation" },
                },
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlSlideshowComponent" },
                },
                { kind: "Field", name: { kind: "Name", value: "mode" } },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlColumnTabIcon" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ColumnWrapper" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "column" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "__typename" } },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlColumnTab" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ColumnWrapper" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "columnName" } },
          { kind: "Field", name: { kind: "Name", value: "columnDisplayName" } },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlColumnTabIcon" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlColumnTabs" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "columns" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "columnName" } },
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlColumnTab" },
                },
              ],
            },
          },
          { kind: "Field", name: { kind: "Name", value: "focusColumn" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlHandsonHeader" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlColumnTabs" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalHeader" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "terminals" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "name" } },
                {
                  kind: "Field",
                  name: { kind: "Name", value: "currentDirectory" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalEntryComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalEntry" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "entryType" } },
          { kind: "Field", name: { kind: "Name", value: "text" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalTooltip" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalTooltip" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "markdownBody" } },
          { kind: "Field", name: { kind: "Name", value: "timing" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalContents" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "entries" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "id" } },
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlTerminalEntryComponent" },
                },
              ],
            },
          },
          {
            kind: "Field",
            name: { kind: "Name", value: "tooltip" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlTerminalTooltip" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalColumn" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlTerminalHeader" },
          },
          {
            kind: "Field",
            name: { kind: "Name", value: "terminals" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlTerminalContents" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileTreeHeader" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "projectDir" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileNodeIcon" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "FileNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "nodeType" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileNodeComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "FileNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlFileNodeIcon" },
          },
          { kind: "Field", name: { kind: "Name", value: "nodeType" } },
          { kind: "Field", name: { kind: "Name", value: "name" } },
          { kind: "Field", name: { kind: "Name", value: "filePath" } },
          { kind: "Field", name: { kind: "Name", value: "offset" } },
          { kind: "Field", name: { kind: "Name", value: "isUpdated" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileTreeComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "fileTree" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "filePath" } },
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlFileNodeComponent" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileTreePane" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlFileTreeHeader" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlFileTreeComponent" },
          },
          { kind: "Field", name: { kind: "Name", value: "isFoldFileTree" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileNameTabBar" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "fileName" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlSourceCodeEditor" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "content" } },
          { kind: "Field", name: { kind: "Name", value: "oldContent" } },
          { kind: "Field", name: { kind: "Name", value: "language" } },
          {
            kind: "Field",
            name: { kind: "Name", value: "editSequence" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "id" } },
                {
                  kind: "Field",
                  name: { kind: "Name", value: "edits" },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      { kind: "Field", name: { kind: "Name", value: "text" } },
                      {
                        kind: "Field",
                        name: { kind: "Name", value: "range" },
                        selectionSet: {
                          kind: "SelectionSet",
                          selections: [
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "startLineNumber" },
                            },
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "startColumn" },
                            },
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "endLineNumber" },
                            },
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "endColumn" },
                            },
                          ],
                        },
                      },
                    ],
                  },
                },
              ],
            },
          },
          {
            kind: "Field",
            name: { kind: "Name", value: "tooltip" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "Field",
                  name: { kind: "Name", value: "markdownBody" },
                },
                { kind: "Field", name: { kind: "Name", value: "lineNumber" } },
                { kind: "Field", name: { kind: "Name", value: "timing" } },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlOpenFilePane" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlFileNameTabBar" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlSourceCodeEditor" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlSourceCodeColumn" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCodeColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "sourceCode" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlFileTreePane" },
                },
                {
                  kind: "Field",
                  name: { kind: "Name", value: "openFile" },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "GqlOpenFilePane" },
                      },
                    ],
                  },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlColumnWrapper" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ColumnWrapper" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "columnName" } },
          {
            kind: "Field",
            name: { kind: "Name", value: "column" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "__typename" } },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "TerminalColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "GqlTerminalColumn" },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "SourceCodeColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "GqlSourceCodeColumn" },
                      },
                    ],
                  },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlColumnWrappers" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "focusColumn" } },
          {
            kind: "Field",
            name: { kind: "Name", value: "columns" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "columnName" } },
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlColumnWrapper" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTutorialTitleSlide" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TutorialTitleSlide" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "title" } },
          {
            kind: "Field",
            name: { kind: "Name", value: "images" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "src" } },
                { kind: "Field", name: { kind: "Name", value: "width" } },
                { kind: "Field", name: { kind: "Name", value: "height" } },
                { kind: "Field", name: { kind: "Name", value: "caption" } },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlSectionTitleSlide" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SectionTitleSlide" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "title" } },
          { kind: "Field", name: { kind: "Name", value: "sectionNum" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlImageSlide" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ImageSlide" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "image" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "src" } },
                { kind: "Field", name: { kind: "Name", value: "width" } },
                { kind: "Field", name: { kind: "Name", value: "height" } },
                { kind: "Field", name: { kind: "Name", value: "caption" } },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlSlideWrapper" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SlideWrapper" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "slide" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "__typename" } },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "TutorialTitleSlide" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "GqlTutorialTitleSlide" },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "MarkdownSlide" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "Field",
                        name: { kind: "Name", value: "markdownBody" },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "SectionTitleSlide" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "GqlSectionTitleSlide" },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "ImageSlide" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "GqlImageSlide" },
                      },
                    ],
                  },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlHandsonComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlHandsonHeader" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlColumnWrappers" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlNavigation" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "prevStep" } },
          { kind: "Field", name: { kind: "Name", value: "nextStep" } },
          { kind: "Field", name: { kind: "Name", value: "isTrivial" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlSlideshowComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "slide" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlSlideWrapper" },
                },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<
  AppTutorialPageQuery,
  AppTutorialPageQueryVariables
>;
export const AppTestTutorialColumnsPageDocument = {
  kind: "Document",
  definitions: [
    {
      kind: "OperationDefinition",
      operation: "query",
      name: { kind: "Name", value: "appTestTutorialColumnsPage" },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "_test" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "Field",
                  name: { kind: "Name", value: "appTestTutorialColumnsPage" },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "GqlHandsonComponent" },
                      },
                    ],
                  },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlColumnTabIcon" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ColumnWrapper" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "column" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "__typename" } },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlColumnTab" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ColumnWrapper" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "columnName" } },
          { kind: "Field", name: { kind: "Name", value: "columnDisplayName" } },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlColumnTabIcon" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlColumnTabs" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "columns" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "columnName" } },
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlColumnTab" },
                },
              ],
            },
          },
          { kind: "Field", name: { kind: "Name", value: "focusColumn" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlHandsonHeader" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlColumnTabs" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalHeader" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "terminals" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "name" } },
                {
                  kind: "Field",
                  name: { kind: "Name", value: "currentDirectory" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalEntryComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalEntry" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "entryType" } },
          { kind: "Field", name: { kind: "Name", value: "text" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalTooltip" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalTooltip" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "markdownBody" } },
          { kind: "Field", name: { kind: "Name", value: "timing" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalContents" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "entries" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "id" } },
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlTerminalEntryComponent" },
                },
              ],
            },
          },
          {
            kind: "Field",
            name: { kind: "Name", value: "tooltip" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlTerminalTooltip" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalColumn" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlTerminalHeader" },
          },
          {
            kind: "Field",
            name: { kind: "Name", value: "terminals" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlTerminalContents" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileTreeHeader" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "projectDir" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileNodeIcon" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "FileNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "nodeType" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileNodeComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "FileNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlFileNodeIcon" },
          },
          { kind: "Field", name: { kind: "Name", value: "nodeType" } },
          { kind: "Field", name: { kind: "Name", value: "name" } },
          { kind: "Field", name: { kind: "Name", value: "filePath" } },
          { kind: "Field", name: { kind: "Name", value: "offset" } },
          { kind: "Field", name: { kind: "Name", value: "isUpdated" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileTreeComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "fileTree" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "filePath" } },
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlFileNodeComponent" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileTreePane" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlFileTreeHeader" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlFileTreeComponent" },
          },
          { kind: "Field", name: { kind: "Name", value: "isFoldFileTree" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlFileNameTabBar" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "fileName" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlSourceCodeEditor" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "content" } },
          { kind: "Field", name: { kind: "Name", value: "oldContent" } },
          { kind: "Field", name: { kind: "Name", value: "language" } },
          {
            kind: "Field",
            name: { kind: "Name", value: "editSequence" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "id" } },
                {
                  kind: "Field",
                  name: { kind: "Name", value: "edits" },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      { kind: "Field", name: { kind: "Name", value: "text" } },
                      {
                        kind: "Field",
                        name: { kind: "Name", value: "range" },
                        selectionSet: {
                          kind: "SelectionSet",
                          selections: [
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "startLineNumber" },
                            },
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "startColumn" },
                            },
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "endLineNumber" },
                            },
                            {
                              kind: "Field",
                              name: { kind: "Name", value: "endColumn" },
                            },
                          ],
                        },
                      },
                    ],
                  },
                },
              ],
            },
          },
          {
            kind: "Field",
            name: { kind: "Name", value: "tooltip" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "Field",
                  name: { kind: "Name", value: "markdownBody" },
                },
                { kind: "Field", name: { kind: "Name", value: "lineNumber" } },
                { kind: "Field", name: { kind: "Name", value: "timing" } },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlOpenFilePane" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlFileNameTabBar" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlSourceCodeEditor" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlSourceCodeColumn" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCodeColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "sourceCode" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlFileTreePane" },
                },
                {
                  kind: "Field",
                  name: { kind: "Name", value: "openFile" },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "GqlOpenFilePane" },
                      },
                    ],
                  },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlColumnWrapper" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ColumnWrapper" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "columnName" } },
          {
            kind: "Field",
            name: { kind: "Name", value: "column" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "__typename" } },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "TerminalColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "GqlTerminalColumn" },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "SourceCodeColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "GqlSourceCodeColumn" },
                      },
                    ],
                  },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlColumnWrappers" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "focusColumn" } },
          {
            kind: "Field",
            name: { kind: "Name", value: "columns" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "columnName" } },
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlColumnWrapper" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlHandsonComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlHandsonHeader" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlColumnWrappers" },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<
  AppTestTutorialColumnsPageQuery,
  AppTestTutorialColumnsPageQueryVariables
>;
export const AppTestTerminalPageDocument = {
  kind: "Document",
  definitions: [
    {
      kind: "OperationDefinition",
      operation: "query",
      name: { kind: "Name", value: "appTestTerminalPage" },
      variableDefinitions: [
        {
          kind: "VariableDefinition",
          variable: { kind: "Variable", name: { kind: "Name", value: "step" } },
          type: { kind: "NamedType", name: { kind: "Name", value: "Int" } },
        },
      ],
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "_test" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "Field",
                  name: { kind: "Name", value: "appTestTerminalPage" },
                  arguments: [
                    {
                      kind: "Argument",
                      name: { kind: "Name", value: "step" },
                      value: {
                        kind: "Variable",
                        name: { kind: "Name", value: "step" },
                      },
                    },
                  ],
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "GqlTerminalColumn" },
                      },
                    ],
                  },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalHeader" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "terminals" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "name" } },
                {
                  kind: "Field",
                  name: { kind: "Name", value: "currentDirectory" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalEntryComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalEntry" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "entryType" } },
          { kind: "Field", name: { kind: "Name", value: "text" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalTooltip" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalTooltip" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "markdownBody" } },
          { kind: "Field", name: { kind: "Name", value: "timing" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalContents" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "entries" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "id" } },
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlTerminalEntryComponent" },
                },
              ],
            },
          },
          {
            kind: "Field",
            name: { kind: "Name", value: "tooltip" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlTerminalTooltip" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalColumn" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlTerminalHeader" },
          },
          {
            kind: "Field",
            name: { kind: "Name", value: "terminals" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlTerminalContents" },
                },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<
  AppTestTerminalPageQuery,
  AppTestTerminalPageQueryVariables
>;
