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

export type BrowserColumn = Column & {
  __typename: "BrowserColumn";
  _placeholder?: Maybe<Scalars["String"]["output"]>;
  height?: Maybe<Scalars["Int"]["output"]>;
  path?: Maybe<Scalars["String"]["output"]>;
  width?: Maybe<Scalars["Int"]["output"]>;
};

export type Column = {
  _placeholder?: Maybe<Scalars["String"]["output"]>;
};

export type ColumnVerticalPosition = "BOTTOM" | "CENTER" | "TOP";

export type ColumnWrapper = {
  __typename: "ColumnWrapper";
  column?: Maybe<Column>;
  index?: Maybe<Scalars["Int"]["output"]>;
  name?: Maybe<Scalars["String"]["output"]>;
};

export type DevToolsColumn = Column & {
  __typename: "DevToolsColumn";
  _placeholder?: Maybe<Scalars["String"]["output"]>;
  height?: Maybe<Scalars["Int"]["output"]>;
  path?: Maybe<Scalars["String"]["output"]>;
  width?: Maybe<Scalars["Int"]["output"]>;
};

export type FileHighlight = {
  __typename: "FileHighlight";
  fromLine?: Maybe<Scalars["Int"]["output"]>;
  toLine?: Maybe<Scalars["Int"]["output"]>;
};

export type FileNode = {
  __typename: "FileNode";
  filePath?: Maybe<Scalars["String"]["output"]>;
  isDeleted?: Maybe<Scalars["Boolean"]["output"]>;
  isUpdated?: Maybe<Scalars["Boolean"]["output"]>;
  name?: Maybe<Scalars["String"]["output"]>;
  nodeType?: Maybe<FileNodeType>;
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

export type NextAction = {
  __typename: "NextAction";
  markdown?: Maybe<MarkdownOld>;
  terminalCommand?: Maybe<TerminalCommand>;
  terminalName?: Maybe<Scalars["String"]["output"]>;
};

export type OpenFile = {
  __typename: "OpenFile";
  content?: Maybe<Scalars["String"]["output"]>;
  fileName?: Maybe<Scalars["String"]["output"]>;
  filePath?: Maybe<Scalars["String"]["output"]>;
  highlight?: Maybe<Array<Maybe<FileHighlight>>>;
  isFullContent?: Maybe<Scalars["Boolean"]["output"]>;
  language?: Maybe<Scalars["String"]["output"]>;
  size?: Maybe<Scalars["Float"]["output"]>;
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

export type PageState = {
  __typename: "PageState";
  markdown?: Maybe<MarkdownOld>;
  nextAction?: Maybe<NextAction>;
  nextStep?: Maybe<Scalars["String"]["output"]>;
  prevStep?: Maybe<Scalars["String"]["output"]>;
  sourceCode?: Maybe<SourceCode>;
  step?: Maybe<Scalars["String"]["output"]>;
  terminals?: Maybe<Array<Maybe<Terminal>>>;
};

export type Query = {
  __typename: "Query";
  _test?: Maybe<TestObjs>;
  page?: Maybe<Page>;
};

export type QueryPageArgs = {
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

export type SourceCodeColumn = Column & {
  __typename: "SourceCodeColumn";
  _placeholder?: Maybe<Scalars["String"]["output"]>;
  sourceCode?: Maybe<SourceCode>;
};

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
  name?: Maybe<Scalars["String"]["output"]>;
  nodes: Array<TerminalEntry2>;
  step?: Maybe<Scalars["String"]["output"]>;
  tooltip?: Maybe<TerminalTooltip2>;
};

export type TerminalColumn = Column & {
  __typename: "TerminalColumn";
  _placeholder?: Maybe<Scalars["String"]["output"]>;
  terminal?: Maybe<Terminal>;
};

export type TerminalCommand = {
  __typename: "TerminalCommand";
  beforeExecution?: Maybe<Scalars["Boolean"]["output"]>;
  command?: Maybe<Scalars["String"]["output"]>;
  tooltip?: Maybe<Scalars["String"]["output"]>;
};

export type TerminalCommand2 = TerminalEntry2 & {
  __typename: "TerminalCommand2";
  command: Scalars["String"]["output"];
  id: Scalars["ID"]["output"];
};

export type TerminalElement = TerminalCommand | TerminalOutput;

export type TerminalEntry2 = {
  id: Scalars["ID"]["output"];
};

export type TerminalNode = {
  __typename: "TerminalNode";
  content?: Maybe<TerminalElement>;
};

export type TerminalOutput = {
  __typename: "TerminalOutput";
  output?: Maybe<Scalars["String"]["output"]>;
  tooltip?: Maybe<Scalars["String"]["output"]>;
};

export type TerminalOutput2 = TerminalEntry2 & {
  __typename: "TerminalOutput2";
  id: Scalars["ID"]["output"];
  output: Scalars["String"]["output"];
};

export type TerminalTooltip2 = {
  __typename: "TerminalTooltip2";
  markdownBody: Scalars["String"]["output"];
  timing?: Maybe<TerminalTooltipTiming2>;
};

export type TerminalTooltipTiming2 = "END" | "START";

export type TestObjs = {
  __typename: "TestObjs";
  appTestTerminalTooltipPage?: Maybe<Terminal2>;
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

export type Carousel_FragmentFragment = {
  __typename: "Page";
  focusColumn?: string | null;
  step?: string | null;
  columns?: Array<
    | ({ __typename: "ColumnWrapper"; name?: string | null } & {
        " $fragmentRefs"?: {
          ColumnWrapperComponent_FragmentFragment: ColumnWrapperComponent_FragmentFragment;
        };
      })
    | null
  > | null;
} & { " $fragmentName"?: "Carousel_FragmentFragment" };

export type ColumnHeader_FragmentFragment = ({ __typename: "Page" } & {
  " $fragmentRefs"?: {
    ColumnTabs_FragmentFragment: ColumnTabs_FragmentFragment;
  };
}) & { " $fragmentName"?: "ColumnHeader_FragmentFragment" };

export type ColumnTab_FragmentFragment = {
  __typename: "ColumnWrapper";
  name?: string | null;
  column?:
    | { __typename: "BackgroundImageColumn" }
    | { __typename: "BrowserColumn" }
    | { __typename: "DevToolsColumn" }
    | { __typename: "ImageDescriptionColumn" }
    | { __typename: "MarkdownColumn" }
    | { __typename: "SourceCodeColumn" }
    | { __typename: "TerminalColumn" }
    | { __typename: "YouTubeColumn" }
    | null;
} & { " $fragmentName"?: "ColumnTab_FragmentFragment" };

export type ColumnTabs_FragmentFragment = {
  __typename: "Page";
  columns?: Array<
    | ({ __typename: "ColumnWrapper"; name?: string | null } & {
        " $fragmentRefs"?: {
          ColumnTab_FragmentFragment: ColumnTab_FragmentFragment;
        };
      })
    | null
  > | null;
} & { " $fragmentName"?: "ColumnTabs_FragmentFragment" };

export type ColumnWrapperComponent_FragmentFragment = {
  __typename: "ColumnWrapper";
  name?: string | null;
  column?:
    | { __typename: "BackgroundImageColumn" }
    | ({ __typename: "BrowserColumn" } & {
        " $fragmentRefs"?: {
          BrowserColumn_FragmentFragment: BrowserColumn_FragmentFragment;
        };
      })
    | ({ __typename: "DevToolsColumn" } & {
        " $fragmentRefs"?: {
          DevToolsColumn_FragmentFragment: DevToolsColumn_FragmentFragment;
        };
      })
    | { __typename: "ImageDescriptionColumn" }
    | ({ __typename: "MarkdownColumn" } & {
        " $fragmentRefs"?: {
          MarkdownColumn_FragmentFragment: MarkdownColumn_FragmentFragment;
        };
      })
    | ({ __typename: "SourceCodeColumn" } & {
        " $fragmentRefs"?: {
          SourceCodeColumn_FragmentFragment: SourceCodeColumn_FragmentFragment;
        };
      })
    | ({ __typename: "TerminalColumn" } & {
        " $fragmentRefs"?: {
          TerminalColumn_FragmentFragment: TerminalColumn_FragmentFragment;
        };
      })
    | ({ __typename: "YouTubeColumn" } & {
        " $fragmentRefs"?: {
          YouTubeColumn_FragmentFragment: YouTubeColumn_FragmentFragment;
        };
      })
    | null;
} & { " $fragmentName"?: "ColumnWrapperComponent_FragmentFragment" };

export type VisibleColumn_FragmentFragment = ({
  __typename: "Page";
  columns?: Array<
    | ({ __typename: "ColumnWrapper"; name?: string | null } & {
        " $fragmentRefs"?: {
          ColumnWrapperComponent_FragmentFragment: ColumnWrapperComponent_FragmentFragment;
        };
      })
    | null
  > | null;
  modal?:
    | ({ __typename: "Modal" } & {
        " $fragmentRefs"?: {
          ModalComponentFragmentFragment: ModalComponentFragmentFragment;
        };
      })
    | null;
} & {
  " $fragmentRefs"?: {
    ColumnHeader_FragmentFragment: ColumnHeader_FragmentFragment;
    Carousel_FragmentFragment: Carousel_FragmentFragment;
    Navigation_FragmentFragment: Navigation_FragmentFragment;
  };
}) & { " $fragmentName"?: "VisibleColumn_FragmentFragment" };

export type ModalComponentFragmentFragment = {
  __typename: "Modal";
  text?: string | null;
  position?: ModalPosition | null;
} & { " $fragmentName"?: "ModalComponentFragmentFragment" };

export type DevToolsColumn_FragmentFragment = {
  __typename: "DevToolsColumn";
  width?: number | null;
  height?: number | null;
  path?: string | null;
} & { " $fragmentName"?: "DevToolsColumn_FragmentFragment" };

export type MarkdownColumn_FragmentFragment = {
  __typename: "MarkdownColumn";
  contentsPosition?: ColumnVerticalPosition | null;
  description?:
    | ({ __typename: "Markdown" } & {
        " $fragmentRefs"?: {
          MarkdownFragmentFragment: MarkdownFragmentFragment;
        };
      })
    | null;
} & { " $fragmentName"?: "MarkdownColumn_FragmentFragment" };

export type MarkdownFragmentFragment = {
  __typename: "Markdown";
  contents?: string | null;
  alignment?: MarkdownAlignment | null;
} & { " $fragmentName"?: "MarkdownFragmentFragment" };

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
  nodeType?: FileNodeType | null;
  name?: string | null;
  filePath?: string | null;
  offset?: number | null;
  isUpdated?: boolean | null;
} & {
  " $fragmentRefs"?: {
    FileNodeIcon_FragmentFragment: FileNodeIcon_FragmentFragment;
  };
}) & { " $fragmentName"?: "FileNodeComponent_FragmentFragment" };

export type FileNodeIcon_FragmentFragment = {
  __typename: "FileNode";
  nodeType?: FileNodeType | null;
} & { " $fragmentName"?: "FileNodeIcon_FragmentFragment" };

export type FileTreeComponent_FragmentFragment = {
  __typename: "SourceCode";
  fileTree?: Array<
    | ({ __typename: "FileNode"; filePath?: string | null } & {
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

export type TerminalCurrentDirectory_FragmentFragment = {
  __typename: "Terminal";
  currentDirectory?: string | null;
} & { " $fragmentName"?: "TerminalCurrentDirectory_FragmentFragment" };

export type TerminalColumn_FragmentFragment = {
  __typename: "TerminalColumn";
  terminal?:
    | ({ __typename: "Terminal" } & {
        " $fragmentRefs"?: {
          TerminalComponent_FragmentFragment: TerminalComponent_FragmentFragment;
        };
      })
    | null;
} & { " $fragmentName"?: "TerminalColumn_FragmentFragment" };

export type TerminalCommand_FragmentFragment = ({
  __typename: "TerminalCommand";
  command?: string | null;
  beforeExecution?: boolean | null;
} & {
  " $fragmentRefs"?: {
    TerminalCommandTooltipFragment: TerminalCommandTooltipFragment;
  };
}) & { " $fragmentName"?: "TerminalCommand_FragmentFragment" };

export type TerminalCommandTooltipFragment = {
  __typename: "TerminalCommand";
  tooltip?: string | null;
} & { " $fragmentName"?: "TerminalCommandTooltipFragment" };

export type TerminalComponent_FragmentFragment = ({ __typename: "Terminal" } & {
  " $fragmentRefs"?: {
    TerminalCurrentDirectory_FragmentFragment: TerminalCurrentDirectory_FragmentFragment;
    TerminalContentsComponent_FragmentFragment: TerminalContentsComponent_FragmentFragment;
  };
}) & { " $fragmentName"?: "TerminalComponent_FragmentFragment" };

export type TerminalContentsComponent_FragmentFragment = {
  __typename: "Terminal";
  nodes?: Array<
    | ({ __typename: "TerminalNode" } & {
        " $fragmentRefs"?: {
          TerminalNodeComponent_FragmentFragment: TerminalNodeComponent_FragmentFragment;
        };
      })
    | null
  > | null;
} & { " $fragmentName"?: "TerminalContentsComponent_FragmentFragment" };

export type TerminalNodeComponent_FragmentFragment = {
  __typename: "TerminalNode";
  content?:
    | ({ __typename: "TerminalCommand" } & {
        " $fragmentRefs"?: {
          TerminalCommand_FragmentFragment: TerminalCommand_FragmentFragment;
        };
      })
    | ({ __typename: "TerminalOutput" } & {
        " $fragmentRefs"?: {
          TerminalOutput_FragmentFragment: TerminalOutput_FragmentFragment;
        };
      })
    | null;
} & { " $fragmentName"?: "TerminalNodeComponent_FragmentFragment" };

export type TerminalOutput_FragmentFragment = {
  __typename: "TerminalOutput";
  output?: string | null;
} & { " $fragmentName"?: "TerminalOutput_FragmentFragment" };

export type GqlTerminalComponentFragment = ({ __typename: "Terminal2" } & {
  " $fragmentRefs"?: {
    GqlTerminalHeaderFragment: GqlTerminalHeaderFragment;
    GqlTerminalContentsFragment: GqlTerminalContentsFragment;
  };
}) & { " $fragmentName"?: "GqlTerminalComponentFragment" };

export type GqlCommandComponentFragment = {
  __typename: "TerminalCommand2";
  command: string;
} & { " $fragmentName"?: "GqlCommandComponentFragment" };

export type GqlTerminalContentsFragment = {
  __typename: "Terminal2";
  nodes: Array<
    | ({ __typename: "TerminalCommand2"; id: string } & {
        " $fragmentRefs"?: {
          GqlTerminalEntryComponent_TerminalCommand2_Fragment: GqlTerminalEntryComponent_TerminalCommand2_Fragment;
        };
      })
    | ({ __typename: "TerminalOutput2"; id: string } & {
        " $fragmentRefs"?: {
          GqlTerminalEntryComponent_TerminalOutput2_Fragment: GqlTerminalEntryComponent_TerminalOutput2_Fragment;
        };
      })
  >;
  tooltip?:
    | ({ __typename: "TerminalTooltip2" } & {
        " $fragmentRefs"?: {
          GqlTerminalTooltipFragment: GqlTerminalTooltipFragment;
        };
      })
    | null;
} & { " $fragmentName"?: "GqlTerminalContentsFragment" };

type GqlTerminalEntryComponent_TerminalCommand2_Fragment = ({
  __typename: "TerminalCommand2";
} & {
  " $fragmentRefs"?: {
    GqlCommandComponentFragment: GqlCommandComponentFragment;
  };
}) & {
  " $fragmentName"?: "GqlTerminalEntryComponent_TerminalCommand2_Fragment";
};

type GqlTerminalEntryComponent_TerminalOutput2_Fragment = ({
  __typename: "TerminalOutput2";
} & {
  " $fragmentRefs"?: { GqlOutputComponentFragment: GqlOutputComponentFragment };
}) & {
  " $fragmentName"?: "GqlTerminalEntryComponent_TerminalOutput2_Fragment";
};

export type GqlTerminalEntryComponentFragment =
  | GqlTerminalEntryComponent_TerminalCommand2_Fragment
  | GqlTerminalEntryComponent_TerminalOutput2_Fragment;

export type GqlTerminalHeaderFragment = {
  __typename: "Terminal2";
  currentDirectory: string;
} & { " $fragmentName"?: "GqlTerminalHeaderFragment" };

export type GqlOutputComponentFragment = {
  __typename: "TerminalOutput2";
  output: string;
} & { " $fragmentName"?: "GqlOutputComponentFragment" };

export type GqlTerminalTooltipFragment = {
  __typename: "TerminalTooltip2";
  markdownBody: string;
  timing?: TerminalTooltipTiming2 | null;
} & { " $fragmentName"?: "GqlTerminalTooltipFragment" };

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

export type PageQueryQueryVariables = Exact<{
  tutorial: Scalars["String"]["input"];
  step?: InputMaybe<Scalars["String"]["input"]>;
  openFilePath?: InputMaybe<Scalars["String"]["input"]>;
}>;

export type PageQueryQuery = {
  __typename: "Query";
  page?:
    | ({
        __typename: "Page";
        step?: string | null;
        focusColumn?: string | null;
        autoNextSeconds?: number | null;
      } & {
        " $fragmentRefs"?: {
          VisibleColumn_FragmentFragment: VisibleColumn_FragmentFragment;
        };
      })
    | null;
};

export type AppTestTerminalTooltipPageQueryVariables = Exact<{
  [key: string]: never;
}>;

export type AppTestTerminalTooltipPageQuery = {
  __typename: "Query";
  _test?: {
    __typename: "TestObjs";
    appTestTerminalTooltipPage?:
      | ({ __typename: "Terminal2" } & {
          " $fragmentRefs"?: {
            GqlTerminalComponentFragment: GqlTerminalComponentFragment;
          };
        })
      | null;
  } | null;
};

export const ColumnTab_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "ColumnTab_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ColumnWrapper" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "name" } },
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
} as unknown as DocumentNode<ColumnTab_FragmentFragment, unknown>;
export const ColumnTabs_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "ColumnTabs_Fragment" },
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
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "ColumnTab_Fragment" },
                },
                { kind: "Field", name: { kind: "Name", value: "name" } },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "ColumnTab_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ColumnWrapper" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "name" } },
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
} as unknown as DocumentNode<ColumnTabs_FragmentFragment, unknown>;
export const ColumnHeader_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "ColumnHeader_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "ColumnTabs_Fragment" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "ColumnTab_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ColumnWrapper" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "name" } },
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
      name: { kind: "Name", value: "ColumnTabs_Fragment" },
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
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "ColumnTab_Fragment" },
                },
                { kind: "Field", name: { kind: "Name", value: "name" } },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<ColumnHeader_FragmentFragment, unknown>;
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
export const TerminalCurrentDirectory_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalCurrentDirectory_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "currentDirectory" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<
  TerminalCurrentDirectory_FragmentFragment,
  unknown
>;
export const TerminalCommandTooltipFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalCommandTooltip" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalCommand" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "tooltip" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<TerminalCommandTooltipFragment, unknown>;
export const TerminalCommand_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalCommand_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalCommand" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "command" } },
          { kind: "Field", name: { kind: "Name", value: "beforeExecution" } },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "TerminalCommandTooltip" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalCommandTooltip" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalCommand" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "tooltip" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<TerminalCommand_FragmentFragment, unknown>;
export const TerminalOutput_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalOutput_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalOutput" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "output" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<TerminalOutput_FragmentFragment, unknown>;
export const TerminalNodeComponent_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalNodeComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "content" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "__typename" } },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "TerminalCommand" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "TerminalCommand_Fragment",
                        },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "TerminalOutput" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "TerminalOutput_Fragment",
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
      name: { kind: "Name", value: "TerminalCommandTooltip" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalCommand" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "tooltip" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalCommand_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalCommand" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "command" } },
          { kind: "Field", name: { kind: "Name", value: "beforeExecution" } },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "TerminalCommandTooltip" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalOutput_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalOutput" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "output" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<TerminalNodeComponent_FragmentFragment, unknown>;
export const TerminalContentsComponent_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalContentsComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "nodes" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: {
                    kind: "Name",
                    value: "TerminalNodeComponent_Fragment",
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
      name: { kind: "Name", value: "TerminalCommandTooltip" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalCommand" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "tooltip" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalCommand_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalCommand" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "command" } },
          { kind: "Field", name: { kind: "Name", value: "beforeExecution" } },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "TerminalCommandTooltip" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalOutput_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalOutput" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "output" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalNodeComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "content" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "__typename" } },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "TerminalCommand" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "TerminalCommand_Fragment",
                        },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "TerminalOutput" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "TerminalOutput_Fragment",
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
  ],
} as unknown as DocumentNode<
  TerminalContentsComponent_FragmentFragment,
  unknown
>;
export const TerminalComponent_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "TerminalCurrentDirectory_Fragment" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "TerminalContentsComponent_Fragment" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalCommandTooltip" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalCommand" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "tooltip" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalCommand_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalCommand" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "command" } },
          { kind: "Field", name: { kind: "Name", value: "beforeExecution" } },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "TerminalCommandTooltip" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalOutput_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalOutput" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "output" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalNodeComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "content" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "__typename" } },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "TerminalCommand" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "TerminalCommand_Fragment",
                        },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "TerminalOutput" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "TerminalOutput_Fragment",
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
      name: { kind: "Name", value: "TerminalCurrentDirectory_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "currentDirectory" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalContentsComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "nodes" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: {
                    kind: "Name",
                    value: "TerminalNodeComponent_Fragment",
                  },
                },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<TerminalComponent_FragmentFragment, unknown>;
export const TerminalColumn_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalColumn_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "terminal" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "TerminalComponent_Fragment" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalCurrentDirectory_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "currentDirectory" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalCommandTooltip" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalCommand" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "tooltip" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalCommand_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalCommand" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "command" } },
          { kind: "Field", name: { kind: "Name", value: "beforeExecution" } },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "TerminalCommandTooltip" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalOutput_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalOutput" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "output" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalNodeComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "content" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "__typename" } },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "TerminalCommand" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "TerminalCommand_Fragment",
                        },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "TerminalOutput" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "TerminalOutput_Fragment",
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
      name: { kind: "Name", value: "TerminalContentsComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "nodes" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: {
                    kind: "Name",
                    value: "TerminalNodeComponent_Fragment",
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
      name: { kind: "Name", value: "TerminalComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "TerminalCurrentDirectory_Fragment" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "TerminalContentsComponent_Fragment" },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<TerminalColumn_FragmentFragment, unknown>;
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
export const DevToolsColumn_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "DevToolsColumn_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "DevToolsColumn" },
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
} as unknown as DocumentNode<DevToolsColumn_FragmentFragment, unknown>;
export const MarkdownFragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "MarkdownFragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Markdown" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "contents" } },
          { kind: "Field", name: { kind: "Name", value: "alignment" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<MarkdownFragmentFragment, unknown>;
export const MarkdownColumn_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "MarkdownColumn_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "MarkdownColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "description" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "MarkdownFragment" },
                },
              ],
            },
          },
          { kind: "Field", name: { kind: "Name", value: "contentsPosition" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "MarkdownFragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Markdown" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "contents" } },
          { kind: "Field", name: { kind: "Name", value: "alignment" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<MarkdownColumn_FragmentFragment, unknown>;
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
export const ColumnWrapperComponent_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "ColumnWrapperComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ColumnWrapper" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "name" } },
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
                    name: { kind: "Name", value: "SourceCodeColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "SourceCodeColumn_Fragment",
                        },
                      },
                    ],
                  },
                },
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
                        name: {
                          kind: "Name",
                          value: "TerminalColumn_Fragment",
                        },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "BrowserColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "BrowserColumn_Fragment" },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "DevToolsColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "DevToolsColumn_Fragment",
                        },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "MarkdownColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "MarkdownColumn_Fragment",
                        },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "YouTubeColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "YouTubeColumn_Fragment" },
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
      name: { kind: "Name", value: "TerminalCurrentDirectory_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "currentDirectory" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalCommandTooltip" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalCommand" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "tooltip" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalCommand_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalCommand" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "command" } },
          { kind: "Field", name: { kind: "Name", value: "beforeExecution" } },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "TerminalCommandTooltip" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalOutput_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalOutput" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "output" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalNodeComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "content" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "__typename" } },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "TerminalCommand" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "TerminalCommand_Fragment",
                        },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "TerminalOutput" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "TerminalOutput_Fragment",
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
      name: { kind: "Name", value: "TerminalContentsComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "nodes" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: {
                    kind: "Name",
                    value: "TerminalNodeComponent_Fragment",
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
      name: { kind: "Name", value: "TerminalComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "TerminalCurrentDirectory_Fragment" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "TerminalContentsComponent_Fragment" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "MarkdownFragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Markdown" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "contents" } },
          { kind: "Field", name: { kind: "Name", value: "alignment" } },
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
      name: { kind: "Name", value: "TerminalColumn_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "terminal" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "TerminalComponent_Fragment" },
                },
              ],
            },
          },
        ],
      },
    },
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
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "DevToolsColumn_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "DevToolsColumn" },
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
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "MarkdownColumn_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "MarkdownColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "description" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "MarkdownFragment" },
                },
              ],
            },
          },
          { kind: "Field", name: { kind: "Name", value: "contentsPosition" } },
        ],
      },
    },
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
  ],
} as unknown as DocumentNode<ColumnWrapperComponent_FragmentFragment, unknown>;
export const Carousel_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "Carousel_Fragment" },
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
                {
                  kind: "FragmentSpread",
                  name: {
                    kind: "Name",
                    value: "ColumnWrapperComponent_Fragment",
                  },
                },
                { kind: "Field", name: { kind: "Name", value: "name" } },
              ],
            },
          },
          { kind: "Field", name: { kind: "Name", value: "focusColumn" } },
          { kind: "Field", name: { kind: "Name", value: "step" } },
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
      name: { kind: "Name", value: "TerminalCurrentDirectory_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "currentDirectory" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalCommandTooltip" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalCommand" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "tooltip" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalCommand_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalCommand" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "command" } },
          { kind: "Field", name: { kind: "Name", value: "beforeExecution" } },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "TerminalCommandTooltip" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalOutput_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalOutput" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "output" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalNodeComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "content" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "__typename" } },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "TerminalCommand" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "TerminalCommand_Fragment",
                        },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "TerminalOutput" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "TerminalOutput_Fragment",
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
      name: { kind: "Name", value: "TerminalContentsComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "nodes" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: {
                    kind: "Name",
                    value: "TerminalNodeComponent_Fragment",
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
      name: { kind: "Name", value: "TerminalComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "TerminalCurrentDirectory_Fragment" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "TerminalContentsComponent_Fragment" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalColumn_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "terminal" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "TerminalComponent_Fragment" },
                },
              ],
            },
          },
        ],
      },
    },
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
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "DevToolsColumn_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "DevToolsColumn" },
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
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "MarkdownFragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Markdown" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "contents" } },
          { kind: "Field", name: { kind: "Name", value: "alignment" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "MarkdownColumn_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "MarkdownColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "description" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "MarkdownFragment" },
                },
              ],
            },
          },
          { kind: "Field", name: { kind: "Name", value: "contentsPosition" } },
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
      name: { kind: "Name", value: "ColumnWrapperComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ColumnWrapper" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "name" } },
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
                    name: { kind: "Name", value: "SourceCodeColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "SourceCodeColumn_Fragment",
                        },
                      },
                    ],
                  },
                },
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
                        name: {
                          kind: "Name",
                          value: "TerminalColumn_Fragment",
                        },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "BrowserColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "BrowserColumn_Fragment" },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "DevToolsColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "DevToolsColumn_Fragment",
                        },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "MarkdownColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "MarkdownColumn_Fragment",
                        },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "YouTubeColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "YouTubeColumn_Fragment" },
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
} as unknown as DocumentNode<Carousel_FragmentFragment, unknown>;
export const ModalComponentFragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "ModalComponentFragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Modal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "text" } },
          { kind: "Field", name: { kind: "Name", value: "position" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<ModalComponentFragmentFragment, unknown>;
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
export const VisibleColumn_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "VisibleColumn_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "ColumnHeader_Fragment" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "Carousel_Fragment" },
          },
          {
            kind: "Field",
            name: { kind: "Name", value: "columns" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: {
                    kind: "Name",
                    value: "ColumnWrapperComponent_Fragment",
                  },
                },
                { kind: "Field", name: { kind: "Name", value: "name" } },
              ],
            },
          },
          {
            kind: "Field",
            name: { kind: "Name", value: "modal" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "ModalComponentFragment" },
                },
              ],
            },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "Navigation_Fragment" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "ColumnTab_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ColumnWrapper" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "name" } },
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
      name: { kind: "Name", value: "ColumnTabs_Fragment" },
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
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "ColumnTab_Fragment" },
                },
                { kind: "Field", name: { kind: "Name", value: "name" } },
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
      name: { kind: "Name", value: "TerminalCurrentDirectory_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "currentDirectory" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalCommandTooltip" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalCommand" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "tooltip" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalCommand_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalCommand" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "command" } },
          { kind: "Field", name: { kind: "Name", value: "beforeExecution" } },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "TerminalCommandTooltip" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalOutput_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalOutput" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "output" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalNodeComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "content" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "__typename" } },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "TerminalCommand" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "TerminalCommand_Fragment",
                        },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "TerminalOutput" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "TerminalOutput_Fragment",
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
      name: { kind: "Name", value: "TerminalContentsComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "nodes" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: {
                    kind: "Name",
                    value: "TerminalNodeComponent_Fragment",
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
      name: { kind: "Name", value: "TerminalComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "TerminalCurrentDirectory_Fragment" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "TerminalContentsComponent_Fragment" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalColumn_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "terminal" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "TerminalComponent_Fragment" },
                },
              ],
            },
          },
        ],
      },
    },
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
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "DevToolsColumn_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "DevToolsColumn" },
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
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "MarkdownFragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Markdown" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "contents" } },
          { kind: "Field", name: { kind: "Name", value: "alignment" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "MarkdownColumn_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "MarkdownColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "description" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "MarkdownFragment" },
                },
              ],
            },
          },
          { kind: "Field", name: { kind: "Name", value: "contentsPosition" } },
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
      name: { kind: "Name", value: "ColumnWrapperComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ColumnWrapper" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "name" } },
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
                    name: { kind: "Name", value: "SourceCodeColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "SourceCodeColumn_Fragment",
                        },
                      },
                    ],
                  },
                },
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
                        name: {
                          kind: "Name",
                          value: "TerminalColumn_Fragment",
                        },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "BrowserColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "BrowserColumn_Fragment" },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "DevToolsColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "DevToolsColumn_Fragment",
                        },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "MarkdownColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "MarkdownColumn_Fragment",
                        },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "YouTubeColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "YouTubeColumn_Fragment" },
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
      name: { kind: "Name", value: "ColumnHeader_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "ColumnTabs_Fragment" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "Carousel_Fragment" },
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
                {
                  kind: "FragmentSpread",
                  name: {
                    kind: "Name",
                    value: "ColumnWrapperComponent_Fragment",
                  },
                },
                { kind: "Field", name: { kind: "Name", value: "name" } },
              ],
            },
          },
          { kind: "Field", name: { kind: "Name", value: "focusColumn" } },
          { kind: "Field", name: { kind: "Name", value: "step" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "ModalComponentFragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Modal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "text" } },
          { kind: "Field", name: { kind: "Name", value: "position" } },
        ],
      },
    },
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
} as unknown as DocumentNode<VisibleColumn_FragmentFragment, unknown>;
export const GqlTerminalHeaderFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalHeader" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal2" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "currentDirectory" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlTerminalHeaderFragment, unknown>;
export const GqlCommandComponentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlCommandComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalCommand2" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "command" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlCommandComponentFragment, unknown>;
export const GqlOutputComponentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlOutputComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalOutput2" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "output" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<GqlOutputComponentFragment, unknown>;
export const GqlTerminalEntryComponentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalEntryComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalEntry2" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "InlineFragment",
            typeCondition: {
              kind: "NamedType",
              name: { kind: "Name", value: "TerminalCommand2" },
            },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlCommandComponent" },
                },
              ],
            },
          },
          {
            kind: "InlineFragment",
            typeCondition: {
              kind: "NamedType",
              name: { kind: "Name", value: "TerminalOutput2" },
            },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlOutputComponent" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlCommandComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalCommand2" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "command" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlOutputComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalOutput2" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "output" } },
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
        name: { kind: "Name", value: "TerminalTooltip2" },
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
            name: { kind: "Name", value: "nodes" },
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
      name: { kind: "Name", value: "GqlCommandComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalCommand2" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "command" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlOutputComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalOutput2" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "output" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalEntryComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalEntry2" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "InlineFragment",
            typeCondition: {
              kind: "NamedType",
              name: { kind: "Name", value: "TerminalCommand2" },
            },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlCommandComponent" },
                },
              ],
            },
          },
          {
            kind: "InlineFragment",
            typeCondition: {
              kind: "NamedType",
              name: { kind: "Name", value: "TerminalOutput2" },
            },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlOutputComponent" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalTooltip" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalTooltip2" },
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
export const GqlTerminalComponentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal2" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlTerminalHeader" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlTerminalContents" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlCommandComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalCommand2" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "command" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlOutputComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalOutput2" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "output" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalEntryComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalEntry2" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "InlineFragment",
            typeCondition: {
              kind: "NamedType",
              name: { kind: "Name", value: "TerminalCommand2" },
            },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlCommandComponent" },
                },
              ],
            },
          },
          {
            kind: "InlineFragment",
            typeCondition: {
              kind: "NamedType",
              name: { kind: "Name", value: "TerminalOutput2" },
            },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlOutputComponent" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalTooltip" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalTooltip2" },
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
        name: { kind: "Name", value: "Terminal2" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "currentDirectory" } },
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
            name: { kind: "Name", value: "nodes" },
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
} as unknown as DocumentNode<GqlTerminalComponentFragment, unknown>;
export const PageQueryDocument = {
  kind: "Document",
  definitions: [
    {
      kind: "OperationDefinition",
      operation: "query",
      name: { kind: "Name", value: "PageQuery" },
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
        {
          kind: "VariableDefinition",
          variable: {
            kind: "Variable",
            name: { kind: "Name", value: "openFilePath" },
          },
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
                  name: { kind: "Name", value: "VisibleColumn_Fragment" },
                },
                { kind: "Field", name: { kind: "Name", value: "step" } },
                { kind: "Field", name: { kind: "Name", value: "focusColumn" } },
                {
                  kind: "Field",
                  name: { kind: "Name", value: "autoNextSeconds" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "ColumnTab_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ColumnWrapper" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "name" } },
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
      name: { kind: "Name", value: "ColumnTabs_Fragment" },
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
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "ColumnTab_Fragment" },
                },
                { kind: "Field", name: { kind: "Name", value: "name" } },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "ColumnHeader_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "ColumnTabs_Fragment" },
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
      name: { kind: "Name", value: "TerminalCurrentDirectory_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "currentDirectory" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalCommandTooltip" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalCommand" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "tooltip" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalCommand_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalCommand" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "command" } },
          { kind: "Field", name: { kind: "Name", value: "beforeExecution" } },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "TerminalCommandTooltip" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalOutput_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalOutput" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "output" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalNodeComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "content" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "__typename" } },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "TerminalCommand" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "TerminalCommand_Fragment",
                        },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "TerminalOutput" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "TerminalOutput_Fragment",
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
      name: { kind: "Name", value: "TerminalContentsComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "nodes" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: {
                    kind: "Name",
                    value: "TerminalNodeComponent_Fragment",
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
      name: { kind: "Name", value: "TerminalComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "TerminalCurrentDirectory_Fragment" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "TerminalContentsComponent_Fragment" },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalColumn_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "terminal" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "TerminalComponent_Fragment" },
                },
              ],
            },
          },
        ],
      },
    },
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
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "DevToolsColumn_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "DevToolsColumn" },
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
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "MarkdownFragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Markdown" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "contents" } },
          { kind: "Field", name: { kind: "Name", value: "alignment" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "MarkdownColumn_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "MarkdownColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "description" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "MarkdownFragment" },
                },
              ],
            },
          },
          { kind: "Field", name: { kind: "Name", value: "contentsPosition" } },
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
      name: { kind: "Name", value: "ColumnWrapperComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ColumnWrapper" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "name" } },
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
                    name: { kind: "Name", value: "SourceCodeColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "SourceCodeColumn_Fragment",
                        },
                      },
                    ],
                  },
                },
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
                        name: {
                          kind: "Name",
                          value: "TerminalColumn_Fragment",
                        },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "BrowserColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "BrowserColumn_Fragment" },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "DevToolsColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "DevToolsColumn_Fragment",
                        },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "MarkdownColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "MarkdownColumn_Fragment",
                        },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "YouTubeColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "YouTubeColumn_Fragment" },
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
      name: { kind: "Name", value: "Carousel_Fragment" },
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
                {
                  kind: "FragmentSpread",
                  name: {
                    kind: "Name",
                    value: "ColumnWrapperComponent_Fragment",
                  },
                },
                { kind: "Field", name: { kind: "Name", value: "name" } },
              ],
            },
          },
          { kind: "Field", name: { kind: "Name", value: "focusColumn" } },
          { kind: "Field", name: { kind: "Name", value: "step" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "ModalComponentFragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Modal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "text" } },
          { kind: "Field", name: { kind: "Name", value: "position" } },
        ],
      },
    },
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
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "VisibleColumn_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "ColumnHeader_Fragment" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "Carousel_Fragment" },
          },
          {
            kind: "Field",
            name: { kind: "Name", value: "columns" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: {
                    kind: "Name",
                    value: "ColumnWrapperComponent_Fragment",
                  },
                },
                { kind: "Field", name: { kind: "Name", value: "name" } },
              ],
            },
          },
          {
            kind: "Field",
            name: { kind: "Name", value: "modal" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "ModalComponentFragment" },
                },
              ],
            },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "Navigation_Fragment" },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<PageQueryQuery, PageQueryQueryVariables>;
export const AppTestTerminalTooltipPageDocument = {
  kind: "Document",
  definitions: [
    {
      kind: "OperationDefinition",
      operation: "query",
      name: { kind: "Name", value: "appTestTerminalTooltipPage" },
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
                  name: { kind: "Name", value: "appTestTerminalTooltipPage" },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "GqlTerminalComponent" },
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
        name: { kind: "Name", value: "Terminal2" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "currentDirectory" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlCommandComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalCommand2" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "command" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlOutputComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalOutput2" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "output" } },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalEntryComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalEntry2" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "InlineFragment",
            typeCondition: {
              kind: "NamedType",
              name: { kind: "Name", value: "TerminalCommand2" },
            },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlCommandComponent" },
                },
              ],
            },
          },
          {
            kind: "InlineFragment",
            typeCondition: {
              kind: "NamedType",
              name: { kind: "Name", value: "TerminalOutput2" },
            },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "GqlOutputComponent" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "GqlTerminalTooltip" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalTooltip2" },
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
            name: { kind: "Name", value: "nodes" },
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
      name: { kind: "Name", value: "GqlTerminalComponent" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal2" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlTerminalHeader" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "GqlTerminalContents" },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<
  AppTestTerminalTooltipPageQuery,
  AppTestTerminalTooltipPageQueryVariables
>;
