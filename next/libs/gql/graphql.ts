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

export type BackgroundImageColumn = Column & {
  __typename: "BackgroundImageColumn";
  _placeholder?: Maybe<Scalars["String"]["output"]>;
  height?: Maybe<Scalars["Int"]["output"]>;
  modal?: Maybe<Modal>;
  path?: Maybe<Scalars["String"]["output"]>;
  url?: Maybe<Scalars["String"]["output"]>;
  width?: Maybe<Scalars["Int"]["output"]>;
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
  height?: Maybe<Scalars["Int"]["output"]>;
  path?: Maybe<Scalars["String"]["output"]>;
  width?: Maybe<Scalars["Int"]["output"]>;
};

export type BrowserColumn2 = Column2 & {
  __typename: "BrowserColumn2";
  _placeholder?: Maybe<Scalars["String"]["output"]>;
  browser: Browser;
};

export type Column = {
  _placeholder?: Maybe<Scalars["String"]["output"]>;
};

export type Column2 = {
  _placeholder?: Maybe<Scalars["String"]["output"]>;
};

export type ColumnVerticalPosition = "BOTTOM" | "CENTER" | "TOP";

export type ColumnWrapper = {
  __typename: "ColumnWrapper";
  column?: Maybe<Column>;
  index?: Maybe<Scalars["Int"]["output"]>;
  name?: Maybe<Scalars["String"]["output"]>;
};

export type ColumnWrapper2 = {
  __typename: "ColumnWrapper2";
  column: Column2;
  columnDisplayName?: Maybe<Scalars["String"]["output"]>;
  columnName: Scalars["String"]["output"];
};

export type DevToolsColumn = {
  __typename: "DevToolsColumn";
  _placeholder?: Maybe<Scalars["String"]["output"]>;
  height?: Maybe<Scalars["Int"]["output"]>;
  path?: Maybe<Scalars["String"]["output"]>;
  width?: Maybe<Scalars["Int"]["output"]>;
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

export type ImageCentered = {
  __typename: "ImageCentered";
  height?: Maybe<Scalars["Int"]["output"]>;
  path?: Maybe<Scalars["String"]["output"]>;
  url?: Maybe<Scalars["String"]["output"]>;
  width?: Maybe<Scalars["Int"]["output"]>;
};

export type ImageDescriptionColumn = Column & {
  __typename: "ImageDescriptionColumn";
  _placeholder?: Maybe<Scalars["String"]["output"]>;
  contentsPosition?: Maybe<ColumnVerticalPosition>;
  description?: Maybe<Markdown>;
  image?: Maybe<ImageCentered>;
  order?: Maybe<ImageDescriptionOrder>;
};

export type ImageDescriptionOrder =
  | "DESCRIPTION_THEN_IMAGE"
  | "IMAGE_THEN_DESCRIPTION";

export type Markdown = {
  __typename: "Markdown";
  alignment?: Maybe<MarkdownAlignment>;
  contents?: Maybe<Scalars["String"]["output"]>;
  step?: Maybe<Scalars["String"]["output"]>;
};

export type MarkdownAlignment = "CENTER" | "LEFT";

export type MarkdownColumn = Column & {
  __typename: "MarkdownColumn";
  _placeholder?: Maybe<Scalars["String"]["output"]>;
  contentsPosition?: Maybe<ColumnVerticalPosition>;
  description?: Maybe<Markdown>;
};

export type MarkdownOld = {
  __typename: "MarkdownOld";
  contents?: Maybe<Scalars["String"]["output"]>;
  step?: Maybe<Scalars["String"]["output"]>;
};

export type Modal = {
  __typename: "Modal";
  position?: Maybe<ModalPosition>;
  text?: Maybe<Scalars["String"]["output"]>;
};

export type ModalPosition = "BOTTOM" | "CENTER" | "TOP";

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
  autoNextSeconds?: Maybe<Scalars["Int"]["output"]>;
  columns?: Maybe<Array<Maybe<ColumnWrapper>>>;
  durationSeconds?: Maybe<Scalars["Int"]["output"]>;
  focusColumn?: Maybe<Scalars["String"]["output"]>;
  isTrivialStep?: Maybe<Scalars["Boolean"]["output"]>;
  modal?: Maybe<Modal>;
  nextStep?: Maybe<Scalars["String"]["output"]>;
  prevStep?: Maybe<Scalars["String"]["output"]>;
  step?: Maybe<Scalars["String"]["output"]>;
};

export type Page2 = {
  __typename: "Page2";
  columns?: Maybe<Array<Maybe<ColumnWrapper2>>>;
  focusColumn?: Maybe<Scalars["String"]["output"]>;
  isTrivial?: Maybe<Scalars["Boolean"]["output"]>;
  modal?: Maybe<Modal>;
  nextStep?: Maybe<Scalars["String"]["output"]>;
  prevStep?: Maybe<Scalars["String"]["output"]>;
  step?: Maybe<Scalars["String"]["output"]>;
};

export type Query = {
  __typename: "Query";
  _test?: Maybe<TestObjs>;
  page?: Maybe<Page>;
  page2?: Maybe<Page2>;
};

export type QueryPageArgs = {
  step?: InputMaybe<Scalars["String"]["input"]>;
  tutorial: Scalars["String"]["input"];
};

export type QueryPage2Args = {
  step?: InputMaybe<Scalars["String"]["input"]>;
  tutorial: Scalars["String"]["input"];
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

export type SourceCode2 = {
  __typename: "SourceCode2";
  fileTree?: Maybe<Array<Maybe<FileNode>>>;
  isFoldFileTree?: Maybe<Scalars["Boolean"]["output"]>;
  openFile?: Maybe<OpenFile>;
  projectDir?: Maybe<Scalars["String"]["output"]>;
  step?: Maybe<Scalars["String"]["output"]>;
};

export type SourceCode2OpenFileArgs = {
  filePath?: InputMaybe<Scalars["String"]["input"]>;
};

export type SourceCodeColumn = Column & {
  __typename: "SourceCodeColumn";
  _placeholder?: Maybe<Scalars["String"]["output"]>;
  sourceCode?: Maybe<SourceCode>;
};

export type SourceCodeColumn2 = Column2 & {
  __typename: "SourceCodeColumn2";
  _placeholder?: Maybe<Scalars["String"]["output"]>;
  sourceCode: SourceCode2;
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
  currentDirectory?: Maybe<Scalars["String"]["output"]>;
  name?: Maybe<Scalars["String"]["output"]>;
  nodes?: Maybe<Array<Maybe<TerminalNode>>>;
  step?: Maybe<Scalars["String"]["output"]>;
};

export type Terminal2 = {
  __typename: "Terminal2";
  currentDirectory: Scalars["String"]["output"];
  entries: Array<TerminalEntry>;
  name?: Maybe<Scalars["String"]["output"]>;
  tooltip?: Maybe<TerminalTooltip>;
};

export type TerminalColumn = Column & {
  __typename: "TerminalColumn";
  _placeholder?: Maybe<Scalars["String"]["output"]>;
  terminal?: Maybe<Terminal>;
};

export type TerminalColumn2 = Column2 & {
  __typename: "TerminalColumn2";
  _placeholder?: Maybe<Scalars["String"]["output"]>;
  terminals: Array<Terminal2>;
};

export type TerminalCommand = {
  __typename: "TerminalCommand";
  beforeExecution?: Maybe<Scalars["Boolean"]["output"]>;
  command?: Maybe<Scalars["String"]["output"]>;
  tooltip?: Maybe<Scalars["String"]["output"]>;
};

export type TerminalElement = TerminalCommand | TerminalOutput;

export type TerminalEntry = {
  __typename: "TerminalEntry";
  entryType: TerminalEntryType;
  id: Scalars["ID"]["output"];
  text: Scalars["String"]["output"];
};

export type TerminalEntryType = "COMMAND" | "OUTPUT";

export type TerminalNode = {
  __typename: "TerminalNode";
  content?: Maybe<TerminalElement>;
};

export type TerminalOutput = {
  __typename: "TerminalOutput";
  output?: Maybe<Scalars["String"]["output"]>;
  tooltip?: Maybe<Scalars["String"]["output"]>;
};

export type TerminalTooltip = {
  __typename: "TerminalTooltip";
  markdownBody: Scalars["String"]["output"];
  timing?: Maybe<TerminalTooltipTiming>;
};

export type TerminalTooltipTiming = "END" | "START";

export type TestObjs = {
  __typename: "TestObjs";
  appTestSourcecodeFilecontentPage?: Maybe<OpenFile>;
  appTestTerminalPage?: Maybe<TerminalColumn2>;
  appTestTutorialColumnsPage?: Maybe<Page2>;
  appTestTutorialTutorialPage?: Maybe<Page2>;
};

export type TestObjsAppTestSourcecodeFilecontentPageArgs = {
  step: Scalars["Int"]["input"];
};

export type TestObjsAppTestTerminalPageArgs = {
  step?: InputMaybe<Scalars["Int"]["input"]>;
};

export type YouTubeColumn = Column & {
  __typename: "YouTubeColumn";
  _placeholder?: Maybe<Scalars["String"]["output"]>;
  youtube?: Maybe<YouTubeEmbed>;
};

export type YouTubeEmbed = {
  __typename: "YouTubeEmbed";
  embedUrl?: Maybe<Scalars["String"]["output"]>;
  height?: Maybe<Scalars["Int"]["output"]>;
  width?: Maybe<Scalars["Int"]["output"]>;
};

export type BrowserColumn_FragmentFragment = {
  __typename: "BrowserColumn";
  width?: number | null;
  height?: number | null;
  path?: string | null;
} & { " $fragmentName"?: "BrowserColumn_FragmentFragment" };

export type Navigation_FragmentFragment = {
  __typename: "Page";
  step?: string | null;
  nextStep?: string | null;
  prevStep?: string | null;
  durationSeconds?: number | null;
  isTrivialStep?: boolean | null;
} & { " $fragmentName"?: "Navigation_FragmentFragment" };

export type SourceCodeColumn_FragmentFragment = {
  __typename: "SourceCodeColumn";
  sourceCode?:
    | ({
        __typename: "SourceCode";
        openFile?:
          | ({ __typename: "OpenFile" } & {
              " $fragmentRefs"?: {
                FileContentPane_FragmentFragment: FileContentPane_FragmentFragment;
              };
            })
          | null;
      } & {
        " $fragmentRefs"?: {
          FileTreePane_FragmentFragment: FileTreePane_FragmentFragment;
        };
      })
    | null;
} & { " $fragmentName"?: "SourceCodeColumn_FragmentFragment" };

export type FileNodeComponent_FragmentFragment = ({
  __typename: "FileNode";
  nodeType: FileNodeType;
  name?: string | null;
  filePath: string;
  offset?: number | null;
  isUpdated?: boolean | null;
} & {
  " $fragmentRefs"?: {
    FileNodeIcon_FragmentFragment: FileNodeIcon_FragmentFragment;
  };
}) & { " $fragmentName"?: "FileNodeComponent_FragmentFragment" };

export type FileNodeIcon_FragmentFragment = {
  __typename: "FileNode";
  nodeType: FileNodeType;
} & { " $fragmentName"?: "FileNodeIcon_FragmentFragment" };

export type FileTreeComponent_FragmentFragment = {
  __typename: "SourceCode";
  fileTree?: Array<
    | ({ __typename: "FileNode"; filePath: string } & {
        " $fragmentRefs"?: {
          FileNodeComponent_FragmentFragment: FileNodeComponent_FragmentFragment;
        };
      })
    | null
  > | null;
} & { " $fragmentName"?: "FileTreeComponent_FragmentFragment" };

export type FileTreeHeader_FragmentFragment = {
  __typename: "SourceCode";
  projectDir?: string | null;
} & { " $fragmentName"?: "FileTreeHeader_FragmentFragment" };

export type FileTreePane_FragmentFragment = ({
  __typename: "SourceCode";
  isFoldFileTree?: boolean | null;
} & {
  " $fragmentRefs"?: {
    FileTreeHeader_FragmentFragment: FileTreeHeader_FragmentFragment;
    FileTreeComponent_FragmentFragment: FileTreeComponent_FragmentFragment;
  };
}) & { " $fragmentName"?: "FileTreePane_FragmentFragment" };

export type FileContentPane_FragmentFragment = ({ __typename: "OpenFile" } & {
  " $fragmentRefs"?: {
    FileNameTabBar_FragmentFragment: FileNameTabBar_FragmentFragment;
    FileContentViewer_FragmentFragment: FileContentViewer_FragmentFragment;
  };
}) & { " $fragmentName"?: "FileContentPane_FragmentFragment" };

export type FileContentViewer_FragmentFragment = {
  __typename: "OpenFile";
  content?: string | null;
  language?: string | null;
  highlight?: Array<{
    __typename: "FileHighlight";
    fromLine?: number | null;
    toLine?: number | null;
  } | null> | null;
} & { " $fragmentName"?: "FileContentViewer_FragmentFragment" };

export type FileNameTab_FragmentFragment = {
  __typename: "OpenFile";
  fileName?: string | null;
} & { " $fragmentName"?: "FileNameTab_FragmentFragment" };

export type FileNameTabBar_FragmentFragment = ({ __typename: "OpenFile" } & {
  " $fragmentRefs"?: {
    FileNameTab_FragmentFragment: FileNameTab_FragmentFragment;
  };
}) & { " $fragmentName"?: "FileNameTabBar_FragmentFragment" };

export type GqlSourceCodeColumnFragment = {
  __typename: "SourceCodeColumn2";
  sourceCode: {
    __typename: "SourceCode2";
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
  __typename: "SourceCode2";
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
  __typename: "SourceCode2";
  projectDir?: string | null;
} & { " $fragmentName"?: "GqlFileTreeHeaderFragment" };

export type GqlFileTreeComponentFragment = {
  __typename: "SourceCode2";
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
  __typename: "TerminalColumn2";
  terminals: Array<
    { __typename: "Terminal2" } & {
      " $fragmentRefs"?: {
        GqlTerminalContentsFragment: GqlTerminalContentsFragment;
      };
    }
  >;
} & {
  " $fragmentRefs"?: { GqlTerminalHeaderFragment: GqlTerminalHeaderFragment };
}) & { " $fragmentName"?: "GqlTerminalColumnFragment" };

export type GqlTerminalContentsFragment = {
  __typename: "Terminal2";
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
  __typename: "TerminalColumn2";
  terminals: Array<{
    __typename: "Terminal2";
    name?: string | null;
    currentDirectory: string;
  }>;
} & { " $fragmentName"?: "GqlTerminalHeaderFragment" };

export type GqlTerminalTooltipFragment = {
  __typename: "TerminalTooltip";
  markdownBody: string;
  timing?: TerminalTooltipTiming | null;
} & { " $fragmentName"?: "GqlTerminalTooltipFragment" };

export type GqlTutorialComponentFragment = ({ __typename: "Page2" } & {
  " $fragmentRefs"?: {
    GqlTutorialHeaderFragment: GqlTutorialHeaderFragment;
    GqlColumnWrappersFragment: GqlColumnWrappersFragment;
  };
}) & { " $fragmentName"?: "GqlTutorialComponentFragment" };

export type GqlColumnWrapperFragment = {
  __typename: "ColumnWrapper2";
  columnName: string;
  column:
    | { __typename: "BrowserColumn2" }
    | ({ __typename: "SourceCodeColumn2" } & {
        " $fragmentRefs"?: {
          GqlSourceCodeColumnFragment: GqlSourceCodeColumnFragment;
        };
      })
    | ({ __typename: "TerminalColumn2" } & {
        " $fragmentRefs"?: {
          GqlTerminalColumnFragment: GqlTerminalColumnFragment;
        };
      });
} & { " $fragmentName"?: "GqlColumnWrapperFragment" };

export type GqlColumnWrappersFragment = {
  __typename: "Page2";
  columns?: Array<
    | ({ __typename: "ColumnWrapper2"; columnName: string } & {
        " $fragmentRefs"?: {
          GqlColumnWrapperFragment: GqlColumnWrapperFragment;
        };
      })
    | null
  > | null;
} & { " $fragmentName"?: "GqlColumnWrappersFragment" };

export type GqlTutorialHeaderFragment = ({ __typename: "Page2" } & {
  " $fragmentRefs"?: { GqlColumnTabsFragment: GqlColumnTabsFragment };
}) & { " $fragmentName"?: "GqlTutorialHeaderFragment" };

export type GqlColumnTabFragment = ({
  __typename: "ColumnWrapper2";
  columnName: string;
  columnDisplayName?: string | null;
} & {
  " $fragmentRefs"?: { GqlColumnTabIconFragment: GqlColumnTabIconFragment };
}) & { " $fragmentName"?: "GqlColumnTabFragment" };

export type GqlColumnTabIconFragment = {
  __typename: "ColumnWrapper2";
  column:
    | { __typename: "BrowserColumn2" }
    | { __typename: "SourceCodeColumn2" }
    | { __typename: "TerminalColumn2" };
} & { " $fragmentName"?: "GqlColumnTabIconFragment" };

export type GqlColumnTabsFragment = {
  __typename: "Page2";
  columns?: Array<
    | ({ __typename: "ColumnWrapper2"; columnName: string } & {
        " $fragmentRefs"?: { GqlColumnTabFragment: GqlColumnTabFragment };
      })
    | null
  > | null;
} & { " $fragmentName"?: "GqlColumnTabsFragment" };

export type YouTubeColumn_FragmentFragment = {
  __typename: "YouTubeColumn";
  youtube?:
    | ({ __typename: "YouTubeEmbed" } & {
        " $fragmentRefs"?: {
          YouTube_FragmentFragment: YouTube_FragmentFragment;
        };
      })
    | null;
} & { " $fragmentName"?: "YouTubeColumn_FragmentFragment" };

export type YouTube_FragmentFragment = {
  __typename: "YouTubeEmbed";
  embedUrl?: string | null;
  width?: number | null;
  height?: number | null;
} & { " $fragmentName"?: "YouTube_FragmentFragment" };

export type AppTestTerminalPageQueryVariables = Exact<{
  step?: InputMaybe<Scalars["Int"]["input"]>;
}>;

export type AppTestTerminalPageQuery = {
  __typename: "Query";
  _test?: {
    __typename: "TestObjs";
    appTestTerminalPage?:
      | ({ __typename: "TerminalColumn2" } & {
          " $fragmentRefs"?: {
            GqlTerminalColumnFragment: GqlTerminalColumnFragment;
          };
        })
      | null;
  } | null;
};

export type AppTestTutorialColumnsPageQueryVariables = Exact<{
  [key: string]: never;
}>;

export type AppTestTutorialColumnsPageQuery = {
  __typename: "Query";
  _test?: {
    __typename: "TestObjs";
    appTestTutorialColumnsPage?:
      | ({ __typename: "Page2" } & {
          " $fragmentRefs"?: {
            GqlTutorialComponentFragment: GqlTutorialComponentFragment;
          };
        })
      | null;
  } | null;
};

export const BrowserColumn_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "BrowserColumn_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "BrowserColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "width" } },
          { kind: "Field", name: { kind: "Name", value: "height" } },
          { kind: "Field", name: { kind: "Name", value: "path" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<BrowserColumn_FragmentFragment, unknown>;
export const Navigation_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "Navigation_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "step" } },
          { kind: "Field", name: { kind: "Name", value: "nextStep" } },
          { kind: "Field", name: { kind: "Name", value: "prevStep" } },
          { kind: "Field", name: { kind: "Name", value: "durationSeconds" } },
          { kind: "Field", name: { kind: "Name", value: "isTrivialStep" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<Navigation_FragmentFragment, unknown>;
export const FileTreeHeader_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "FileTreeHeader_Fragment" },
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
} as unknown as DocumentNode<FileTreeHeader_FragmentFragment, unknown>;
export const FileNodeIcon_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "FileNodeIcon_Fragment" },
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
} as unknown as DocumentNode<FileNodeIcon_FragmentFragment, unknown>;
export const FileNodeComponent_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "FileNodeComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "FileNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "FileNodeIcon_Fragment" },
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
      name: { kind: "Name", value: "FileNodeIcon_Fragment" },
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
} as unknown as DocumentNode<FileNodeComponent_FragmentFragment, unknown>;
export const FileTreeComponent_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "FileTreeComponent_Fragment" },
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
                  name: { kind: "Name", value: "FileNodeComponent_Fragment" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "FileNodeIcon_Fragment" },
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
      name: { kind: "Name", value: "FileNodeComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "FileNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "FileNodeIcon_Fragment" },
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
} as unknown as DocumentNode<FileTreeComponent_FragmentFragment, unknown>;
export const FileTreePane_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "FileTreePane_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "FileTreeHeader_Fragment" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "FileTreeComponent_Fragment" },
          },
          { kind: "Field", name: { kind: "Name", value: "isFoldFileTree" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "FileNodeIcon_Fragment" },
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
      name: { kind: "Name", value: "FileNodeComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "FileNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "FileNodeIcon_Fragment" },
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
      name: { kind: "Name", value: "FileTreeHeader_Fragment" },
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
      name: { kind: "Name", value: "FileTreeComponent_Fragment" },
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
                  name: { kind: "Name", value: "FileNodeComponent_Fragment" },
                },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<FileTreePane_FragmentFragment, unknown>;
export const FileNameTab_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "FileNameTab_Fragment" },
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
} as unknown as DocumentNode<FileNameTab_FragmentFragment, unknown>;
export const FileNameTabBar_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "FileNameTabBar_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "FileNameTab_Fragment" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "FileNameTab_Fragment" },
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
} as unknown as DocumentNode<FileNameTabBar_FragmentFragment, unknown>;
export const FileContentViewer_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "FileContentViewer_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "content" } },
          { kind: "Field", name: { kind: "Name", value: "language" } },
          {
            kind: "Field",
            name: { kind: "Name", value: "highlight" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "fromLine" } },
                { kind: "Field", name: { kind: "Name", value: "toLine" } },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<FileContentViewer_FragmentFragment, unknown>;
export const FileContentPane_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "FileContentPane_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "FileNameTabBar_Fragment" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "FileContentViewer_Fragment" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "FileNameTab_Fragment" },
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
      name: { kind: "Name", value: "FileNameTabBar_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "FileNameTab_Fragment" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "FileContentViewer_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "content" } },
          { kind: "Field", name: { kind: "Name", value: "language" } },
          {
            kind: "Field",
            name: { kind: "Name", value: "highlight" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "fromLine" } },
                { kind: "Field", name: { kind: "Name", value: "toLine" } },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<FileContentPane_FragmentFragment, unknown>;
export const SourceCodeColumn_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "SourceCodeColumn_Fragment" },
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
                  name: { kind: "Name", value: "FileTreePane_Fragment" },
                },
                {
                  kind: "Field",
                  name: { kind: "Name", value: "openFile" },
                  arguments: [
                    {
                      kind: "Argument",
                      name: { kind: "Name", value: "filePath" },
                      value: {
                        kind: "Variable",
                        name: { kind: "Name", value: "openFilePath" },
                      },
                    },
                  ],
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "FileContentPane_Fragment",
                        },
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
      name: { kind: "Name", value: "FileTreeHeader_Fragment" },
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
      name: { kind: "Name", value: "FileNodeIcon_Fragment" },
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
      name: { kind: "Name", value: "FileNodeComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "FileNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "FileNodeIcon_Fragment" },
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
      name: { kind: "Name", value: "FileTreeComponent_Fragment" },
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
                  name: { kind: "Name", value: "FileNodeComponent_Fragment" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "FileNameTab_Fragment" },
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
      name: { kind: "Name", value: "FileNameTabBar_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "FileNameTab_Fragment" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "FileContentViewer_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "content" } },
          { kind: "Field", name: { kind: "Name", value: "language" } },
          {
            kind: "Field",
            name: { kind: "Name", value: "highlight" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "fromLine" } },
                { kind: "Field", name: { kind: "Name", value: "toLine" } },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "FileTreePane_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "FileTreeHeader_Fragment" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "FileTreeComponent_Fragment" },
          },
          { kind: "Field", name: { kind: "Name", value: "isFoldFileTree" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "FileContentPane_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "FileNameTabBar_Fragment" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "FileContentViewer_Fragment" },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<SourceCodeColumn_FragmentFragment, unknown>;
export const GqlColumnTabIconFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlColumnTabIcon" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ColumnWrapper2" },
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
        name: { kind: "Name", value: "ColumnWrapper2" },
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
        name: { kind: "Name", value: "ColumnWrapper2" },
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
        name: { kind: "Name", value: "Page2" },
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
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlColumnTabIcon" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ColumnWrapper2" },
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
        name: { kind: "Name", value: "ColumnWrapper2" },
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
export const GqlTutorialHeaderFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTutorialHeader" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page2" },
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
        name: { kind: "Name", value: "ColumnWrapper2" },
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
        name: { kind: "Name", value: "ColumnWrapper2" },
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
        name: { kind: "Name", value: "Page2" },
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
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlTutorialHeaderFragment, unknown>;
export const GqlTerminalHeaderFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalHeader" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalColumn2" },
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
        name: { kind: "Name", value: "Terminal2" },
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
        name: { kind: "Name", value: "TerminalColumn2" },
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
        name: { kind: "Name", value: "TerminalColumn2" },
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
        name: { kind: "Name", value: "Terminal2" },
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
        name: { kind: "Name", value: "SourceCode2" },
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
        name: { kind: "Name", value: "SourceCode2" },
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
        name: { kind: "Name", value: "SourceCode2" },
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
        name: { kind: "Name", value: "SourceCode2" },
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
        name: { kind: "Name", value: "SourceCode2" },
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
        name: { kind: "Name", value: "SourceCodeColumn2" },
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
        name: { kind: "Name", value: "SourceCode2" },
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
        name: { kind: "Name", value: "SourceCode2" },
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
        name: { kind: "Name", value: "SourceCode2" },
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
        name: { kind: "Name", value: "ColumnWrapper2" },
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
                    name: { kind: "Name", value: "TerminalColumn2" },
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
                    name: { kind: "Name", value: "SourceCodeColumn2" },
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
        name: { kind: "Name", value: "TerminalColumn2" },
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
        name: { kind: "Name", value: "Terminal2" },
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
        name: { kind: "Name", value: "SourceCode2" },
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
        name: { kind: "Name", value: "SourceCode2" },
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
        name: { kind: "Name", value: "SourceCode2" },
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
        name: { kind: "Name", value: "TerminalColumn2" },
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
        name: { kind: "Name", value: "SourceCodeColumn2" },
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
        name: { kind: "Name", value: "Page2" },
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
        name: { kind: "Name", value: "TerminalColumn2" },
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
        name: { kind: "Name", value: "Terminal2" },
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
        name: { kind: "Name", value: "TerminalColumn2" },
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
        name: { kind: "Name", value: "SourceCode2" },
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
        name: { kind: "Name", value: "SourceCode2" },
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
        name: { kind: "Name", value: "SourceCode2" },
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
        name: { kind: "Name", value: "SourceCodeColumn2" },
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
        name: { kind: "Name", value: "ColumnWrapper2" },
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
                    name: { kind: "Name", value: "TerminalColumn2" },
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
                    name: { kind: "Name", value: "SourceCodeColumn2" },
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
export const GqlTutorialComponentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTutorialComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page2" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlTutorialHeader" },
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
        name: { kind: "Name", value: "ColumnWrapper2" },
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
        name: { kind: "Name", value: "ColumnWrapper2" },
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
        name: { kind: "Name", value: "Page2" },
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
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalHeader" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalColumn2" },
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
        name: { kind: "Name", value: "Terminal2" },
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
        name: { kind: "Name", value: "TerminalColumn2" },
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
        name: { kind: "Name", value: "SourceCode2" },
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
        name: { kind: "Name", value: "SourceCode2" },
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
        name: { kind: "Name", value: "SourceCode2" },
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
        name: { kind: "Name", value: "SourceCodeColumn2" },
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
        name: { kind: "Name", value: "ColumnWrapper2" },
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
                    name: { kind: "Name", value: "TerminalColumn2" },
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
                    name: { kind: "Name", value: "SourceCodeColumn2" },
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
      name: { kind: "Name", value: "GqlTutorialHeader" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page2" },
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
        name: { kind: "Name", value: "Page2" },
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
                  name: { kind: "Name", value: "GqlColumnWrapper" },
                },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlTutorialComponentFragment, unknown>;
export const YouTube_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "YouTube_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "YouTubeEmbed" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "embedUrl" } },
          { kind: "Field", name: { kind: "Name", value: "width" } },
          { kind: "Field", name: { kind: "Name", value: "height" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<YouTube_FragmentFragment, unknown>;
export const YouTubeColumn_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "YouTubeColumn_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "YouTubeColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "youtube" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "YouTube_Fragment" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "YouTube_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "YouTubeEmbed" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "embedUrl" } },
          { kind: "Field", name: { kind: "Name", value: "width" } },
          { kind: "Field", name: { kind: "Name", value: "height" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<YouTubeColumn_FragmentFragment, unknown>;
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
        name: { kind: "Name", value: "TerminalColumn2" },
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
        name: { kind: "Name", value: "Terminal2" },
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
        name: { kind: "Name", value: "TerminalColumn2" },
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
                        name: { kind: "Name", value: "GqlTutorialComponent" },
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
        name: { kind: "Name", value: "ColumnWrapper2" },
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
        name: { kind: "Name", value: "ColumnWrapper2" },
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
        name: { kind: "Name", value: "Page2" },
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
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTutorialHeader" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page2" },
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
        name: { kind: "Name", value: "TerminalColumn2" },
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
        name: { kind: "Name", value: "Terminal2" },
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
        name: { kind: "Name", value: "TerminalColumn2" },
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
        name: { kind: "Name", value: "SourceCode2" },
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
        name: { kind: "Name", value: "SourceCode2" },
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
        name: { kind: "Name", value: "SourceCode2" },
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
        name: { kind: "Name", value: "SourceCodeColumn2" },
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
        name: { kind: "Name", value: "ColumnWrapper2" },
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
                    name: { kind: "Name", value: "TerminalColumn2" },
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
                    name: { kind: "Name", value: "SourceCodeColumn2" },
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
        name: { kind: "Name", value: "Page2" },
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
      name: { kind: "Name", value: "GqlTutorialComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page2" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlTutorialHeader" },
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
