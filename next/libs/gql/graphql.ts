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

export type NextAction = {
  __typename: "NextAction";
  content?: Maybe<NextActionContent>;
};

export type NextActionContent = NextActionManual | NextActionTerminal;

export type NextActionManual = {
  __typename: "NextActionManual";
  comment?: Maybe<Scalars["String"]>;
};

export type NextActionTerminal = {
  __typename: "NextActionTerminal";
  command?: Maybe<Scalars["String"]>;
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

export type PageState = {
  __typename: "PageState";
  nextAction?: Maybe<NextAction>;
  nextStep?: Maybe<Scalars["String"]>;
  prevStep?: Maybe<Scalars["String"]>;
  sourceCode?: Maybe<SourceCode>;
  step?: Maybe<Scalars["String"]>;
  terminals?: Maybe<Array<Maybe<Terminal>>>;
};

export type Query = {
  __typename: "Query";
  pageState?: Maybe<PageState>;
};

export type QueryPageStateArgs = {
  step?: InputMaybe<Scalars["String"]>;
};

export type SourceCode = {
  __typename: "SourceCode";
  defaultOpenFile?: Maybe<OpenFile>;
  fileTree?: Maybe<Array<Maybe<FileNode>>>;
  openFile?: Maybe<OpenFile>;
  step?: Maybe<Scalars["String"]>;
};

export type SourceCodeOpenFileArgs = {
  filePath?: InputMaybe<Scalars["String"]>;
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
  step?: InputMaybe<Scalars["String"]>;
  openFilePath?: InputMaybe<Scalars["String"]>;
}>;

export type IndexSsrPageQuery = {
  __typename: "Query";
  pageState?: {
    __typename: "PageState";
    nextStep?: string | null;
    prevStep?: string | null;
    step?: string | null;
    sourceCode?:
      | ({ __typename: "SourceCode" } & {
          " $fragmentRefs"?: {
            SourceCodeViewer_FragmentFragment: SourceCodeViewer_FragmentFragment;
          };
        })
      | null;
    terminals?: Array<
      | ({
          __typename: "Terminal";
          name?: string | null;
          currentDirectory?: string | null;
        } & {
          " $fragmentRefs"?: {
            TerminalComponent_FragmentFragment: TerminalComponent_FragmentFragment;
          };
        })
      | null
    > | null;
  } | null;
};

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
            name: { kind: "Name", value: "pageState" },
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
                { kind: "Field", name: { kind: "Name", value: "nextStep" } },
                { kind: "Field", name: { kind: "Name", value: "prevStep" } },
                { kind: "Field", name: { kind: "Name", value: "step" } },
                {
                  kind: "Field",
                  name: { kind: "Name", value: "sourceCode" },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "SourceCodeViewer_Fragment",
                        },
                      },
                    ],
                  },
                },
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
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "TerminalComponent_Fragment",
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
    ...SourceCodeViewer_FragmentFragmentDoc.definitions,
    ...TerminalComponent_FragmentFragmentDoc.definitions,
  ],
} as unknown as DocumentNode<IndexSsrPageQuery, IndexSsrPageQueryVariables>;
