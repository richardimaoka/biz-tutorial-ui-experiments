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
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

export type BackgroundImageColumn = Column & {
  __typename: "BackgroundImageColumn";
  _placeholder?: Maybe<Scalars["String"]>;
  height?: Maybe<Scalars["Int"]>;
  modal?: Maybe<Modal>;
  path?: Maybe<Scalars["String"]>;
  url?: Maybe<Scalars["String"]>;
  width?: Maybe<Scalars["Int"]>;
};

export type Column = {
  _placeholder?: Maybe<Scalars["String"]>;
};

export type ColumnWrapper = {
  __typename: "ColumnWrapper";
  column?: Maybe<Column>;
  index?: Maybe<Scalars["Int"]>;
};

export type FileHighlight = {
  __typename: "FileHighlight";
  fromLine?: Maybe<Scalars["Int"]>;
  toLine?: Maybe<Scalars["Int"]>;
};

export type FileNode = {
  __typename: "FileNode";
  filePath?: Maybe<Scalars["String"]>;
  isUpdated?: Maybe<Scalars["Boolean"]>;
  name?: Maybe<Scalars["String"]>;
  nodeType?: Maybe<FileNodeType>;
  offset?: Maybe<Scalars["Int"]>;
};

export type FileNodeType = "DIRECTORY" | "FILE";

export type ImageCentered = {
  __typename: "ImageCentered";
  height?: Maybe<Scalars["Int"]>;
  path?: Maybe<Scalars["String"]>;
  url?: Maybe<Scalars["String"]>;
  width?: Maybe<Scalars["Int"]>;
};

export type ImageDescriptionColumn = Column & {
  __typename: "ImageDescriptionColumn";
  _placeholder?: Maybe<Scalars["String"]>;
  description?: Maybe<Markdown>;
  image?: Maybe<ImageCentered>;
  order?: Maybe<ImageDescriptionOrder>;
  position?: Maybe<ImageDescriptionColumnPosition>;
};

export type ImageDescriptionColumnPosition = "BOTTOM" | "CENTER" | "TOP";

export type ImageDescriptionOrder =
  | "DESCRIPTION_THEN_IMAGE"
  | "IMAGE_THEN_DESCRIPTION";

export type Markdown = {
  __typename: "Markdown";
  alignment?: Maybe<MarkdownAlignment>;
  contents?: Maybe<Scalars["String"]>;
  step?: Maybe<Scalars["String"]>;
};

export type MarkdownAlignment = "CENTER" | "LEFT";

export type MarkdownColumn = Column & {
  __typename: "MarkdownColumn";
  _placeholder?: Maybe<Scalars["String"]>;
  description?: Maybe<Markdown>;
  position?: Maybe<MarkdownColumnPosition>;
};

export type MarkdownColumnPosition = "BOTTOM" | "CENTER" | "TOP";

export type MarkdownOld = {
  __typename: "MarkdownOld";
  contents?: Maybe<Scalars["String"]>;
  step?: Maybe<Scalars["String"]>;
};

export type Modal = {
  __typename: "Modal";
  position?: Maybe<ModalPosition>;
  text?: Maybe<Scalars["String"]>;
};

export type ModalPosition = "BOTTOM" | "CENTER" | "TOP";

export type NextAction = {
  __typename: "NextAction";
  markdown?: Maybe<MarkdownOld>;
  terminalCommand?: Maybe<TerminalCommand>;
  terminalName?: Maybe<Scalars["String"]>;
};

export type OpenFile = {
  __typename: "OpenFile";
  content?: Maybe<Scalars["String"]>;
  fileName?: Maybe<Scalars["String"]>;
  filePath?: Maybe<Scalars["String"]>;
  highlight?: Maybe<Array<Maybe<FileHighlight>>>;
  isFullContent?: Maybe<Scalars["Boolean"]>;
  language?: Maybe<Scalars["String"]>;
};

export type Page = {
  __typename: "Page";
  columns?: Maybe<Array<Maybe<ColumnWrapper>>>;
  nextStep?: Maybe<Scalars["String"]>;
  prevStep?: Maybe<Scalars["String"]>;
  step?: Maybe<Scalars["String"]>;
};

export type PageState = {
  __typename: "PageState";
  markdown?: Maybe<MarkdownOld>;
  nextAction?: Maybe<NextAction>;
  nextStep?: Maybe<Scalars["String"]>;
  prevStep?: Maybe<Scalars["String"]>;
  sourceCode?: Maybe<SourceCode>;
  step?: Maybe<Scalars["String"]>;
  terminals?: Maybe<Array<Maybe<Terminal>>>;
};

export type Query = {
  __typename: "Query";
  page?: Maybe<Page>;
  pageState?: Maybe<PageState>;
};

export type QueryPageArgs = {
  step?: InputMaybe<Scalars["String"]>;
  tutorial: Scalars["String"];
};

export type QueryPageStateArgs = {
  step?: InputMaybe<Scalars["String"]>;
};

export type SourceCode = {
  __typename: "SourceCode";
  fileTree?: Maybe<Array<Maybe<FileNode>>>;
  openFile?: Maybe<OpenFile>;
  step?: Maybe<Scalars["String"]>;
};

export type SourceCodeOpenFileArgs = {
  filePath?: InputMaybe<Scalars["String"]>;
};

export type SourceCodeColumn = Column & {
  __typename: "SourceCodeColumn";
  _placeholder?: Maybe<Scalars["String"]>;
  sourceCode?: Maybe<SourceCode>;
};

export type Terminal = {
  __typename: "Terminal";
  currentDirectory?: Maybe<Scalars["String"]>;
  name?: Maybe<Scalars["String"]>;
  nodes?: Maybe<Array<Maybe<TerminalNode>>>;
  step?: Maybe<Scalars["String"]>;
};

export type TerminalCommand = {
  __typename: "TerminalCommand";
  beforeExecution?: Maybe<Scalars["Boolean"]>;
  command?: Maybe<Scalars["String"]>;
};

export type TerminalElement = TerminalCommand | TerminalOutput;

export type TerminalNode = {
  __typename: "TerminalNode";
  content?: Maybe<TerminalElement>;
};

export type TerminalOutput = {
  __typename: "TerminalOutput";
  output?: Maybe<Scalars["String"]>;
};

export type ColumnWrapperFragmentFragment = {
  __typename: "ColumnWrapper";
  column?:
    | { __typename: "BackgroundImageColumn" }
    | ({
        __typename: "ImageDescriptionColumn";
        position?: ImageDescriptionColumnPosition | null;
      } & {
        " $fragmentRefs"?: {
          ImageDescriptionColumnFragmentFragment: ImageDescriptionColumnFragmentFragment;
        };
      })
    | { __typename: "MarkdownColumn" }
    | { __typename: "SourceCodeColumn" }
    | null;
} & { " $fragmentName"?: "ColumnWrapperFragmentFragment" };

export type ImageDescriptionColumnFragmentFragment = {
  __typename: "ImageDescriptionColumn";
  order?: ImageDescriptionOrder | null;
  description?:
    | ({ __typename: "Markdown" } & {
        " $fragmentRefs"?: {
          MarkdownFragmentFragment: MarkdownFragmentFragment;
        };
      })
    | null;
  image?:
    | ({ __typename: "ImageCentered" } & {
        " $fragmentRefs"?: {
          ImageCenteredFragmentFragment: ImageCenteredFragmentFragment;
        };
      })
    | null;
} & { " $fragmentName"?: "ImageDescriptionColumnFragmentFragment" };

export type ImageCenteredFragmentFragment = {
  __typename: "ImageCentered";
  width?: number | null;
  height?: number | null;
  path?: string | null;
} & { " $fragmentName"?: "ImageCenteredFragmentFragment" };

export type MarkdownFragmentFragment = {
  __typename: "Markdown";
  contents?: string | null;
  alignment?: MarkdownAlignment | null;
} & { " $fragmentName"?: "MarkdownFragmentFragment" };

export type SourceCodeViewer_FragmentFragment = ({
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
}) & { " $fragmentName"?: "SourceCodeViewer_FragmentFragment" };

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

export type FileTreePane_FragmentFragment = ({ __typename: "SourceCode" } & {
  " $fragmentRefs"?: {
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

export type TerminalCommand_FragmentFragment = {
  __typename: "TerminalCommand";
  command?: string | null;
  beforeExecution?: boolean | null;
} & { " $fragmentName"?: "TerminalCommand_FragmentFragment" };

export type TerminalComponent_FragmentFragment = {
  __typename: "Terminal";
  currentDirectory?: string | null;
  nodes?: Array<
    | ({ __typename: "TerminalNode" } & {
        " $fragmentRefs"?: {
          TerminalNodeComponent_FragmentFragment: TerminalNodeComponent_FragmentFragment;
        };
      })
    | null
  > | null;
} & { " $fragmentName"?: "TerminalComponent_FragmentFragment" };

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

export type IndexSsrPageQueryVariables = Exact<{
  tutorial: Scalars["String"];
  step?: InputMaybe<Scalars["String"]>;
}>;

export type IndexSsrPageQuery = {
  __typename: "Query";
  page?: {
    __typename: "Page";
    step?: string | null;
    nextStep?: string | null;
    prevStep?: string | null;
    columns?: Array<
      | ({ __typename: "ColumnWrapper" } & {
          " $fragmentRefs"?: {
            ColumnWrapperFragmentFragment: ColumnWrapperFragmentFragment;
          };
        })
      | null
    > | null;
  } | null;
};

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
export const ImageCenteredFragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "ImageCenteredFragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ImageCentered" },
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
} as unknown as DocumentNode<ImageCenteredFragmentFragment, unknown>;
export const ImageDescriptionColumnFragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "ImageDescriptionColumnFragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ImageDescriptionColumn" },
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
          {
            kind: "Field",
            name: { kind: "Name", value: "image" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "ImageCenteredFragment" },
                },
              ],
            },
          },
          { kind: "Field", name: { kind: "Name", value: "order" } },
        ],
      },
    },
    ...MarkdownFragmentFragmentDoc.definitions,
    ...ImageCenteredFragmentFragmentDoc.definitions,
  ],
} as unknown as DocumentNode<ImageDescriptionColumnFragmentFragment, unknown>;
export const ColumnWrapperFragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "ColumnWrapperFragment" },
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
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "ImageDescriptionColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "ImageDescriptionColumnFragment",
                        },
                      },
                      {
                        kind: "Field",
                        name: { kind: "Name", value: "position" },
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
    ...ImageDescriptionColumnFragmentFragmentDoc.definitions,
  ],
} as unknown as DocumentNode<ColumnWrapperFragmentFragment, unknown>;
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
    ...FileNodeIcon_FragmentFragmentDoc.definitions,
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
    ...FileNodeComponent_FragmentFragmentDoc.definitions,
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
            name: { kind: "Name", value: "FileTreeComponent_Fragment" },
          },
        ],
      },
    },
    ...FileTreeComponent_FragmentFragmentDoc.definitions,
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
    ...FileNameTab_FragmentFragmentDoc.definitions,
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
    ...FileNameTabBar_FragmentFragmentDoc.definitions,
    ...FileContentViewer_FragmentFragmentDoc.definitions,
  ],
} as unknown as DocumentNode<FileContentPane_FragmentFragment, unknown>;
export const SourceCodeViewer_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "SourceCodeViewer_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCode" },
      },
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
                  name: { kind: "Name", value: "FileContentPane_Fragment" },
                },
              ],
            },
          },
        ],
      },
    },
    ...FileTreePane_FragmentFragmentDoc.definitions,
    ...FileContentPane_FragmentFragmentDoc.definitions,
  ],
} as unknown as DocumentNode<SourceCodeViewer_FragmentFragment, unknown>;
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
    ...TerminalCommand_FragmentFragmentDoc.definitions,
    ...TerminalOutput_FragmentFragmentDoc.definitions,
  ],
} as unknown as DocumentNode<TerminalNodeComponent_FragmentFragment, unknown>;
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
          { kind: "Field", name: { kind: "Name", value: "currentDirectory" } },
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
    ...TerminalNodeComponent_FragmentFragmentDoc.definitions,
  ],
} as unknown as DocumentNode<TerminalComponent_FragmentFragment, unknown>;
export const IndexSsrPageDocument = {
  kind: "Document",
  definitions: [
    {
      kind: "OperationDefinition",
      operation: "query",
      name: { kind: "Name", value: "IndexSsrPage" },
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
                { kind: "Field", name: { kind: "Name", value: "__typename" } },
                { kind: "Field", name: { kind: "Name", value: "step" } },
                { kind: "Field", name: { kind: "Name", value: "nextStep" } },
                { kind: "Field", name: { kind: "Name", value: "prevStep" } },
                {
                  kind: "Field",
                  name: { kind: "Name", value: "columns" },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "ColumnWrapperFragment" },
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
    ...ColumnWrapperFragmentFragmentDoc.definitions,
  ],
} as unknown as DocumentNode<IndexSsrPageQuery, IndexSsrPageQueryVariables>;
