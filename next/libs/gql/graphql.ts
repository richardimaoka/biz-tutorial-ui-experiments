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

export type DirectoryNode = {
  __typename: "DirectoryNode";
  filePath?: Maybe<Array<Maybe<Scalars["String"]>>>;
  isUpdated?: Maybe<Scalars["Boolean"]>;
};

export type FileHighlight = {
  __typename: "FileHighlight";
  fromLine?: Maybe<Scalars["Int"]>;
  toLine?: Maybe<Scalars["Int"]>;
};

export type FileNode = {
  __typename: "FileNode";
  filePath?: Maybe<Array<Maybe<Scalars["String"]>>>;
  isUpdated?: Maybe<Scalars["Boolean"]>;
};

export type FileTreeNode = DirectoryNode | FileNode;

export type OpenFile = {
  __typename: "OpenFile";
  content?: Maybe<Scalars["String"]>;
  filePath?: Maybe<Array<Maybe<Scalars["String"]>>>;
  highlight?: Maybe<Array<Maybe<FileHighlight>>>;
  isFullContent?: Maybe<Scalars["Boolean"]>;
};

export type Query = {
  __typename: "Query";
  step?: Maybe<Step>;
  terminal?: Maybe<Terminal>;
};

export type QueryStepArgs = {
  stepNum?: InputMaybe<Scalars["Int"]>;
};

export type SourceCode = {
  __typename: "SourceCode";
  fileTree?: Maybe<Array<Maybe<FileTreeNode>>>;
  openFile?: Maybe<OpenFile>;
};

export type Step = {
  __typename: "Step";
  nextAction?: Maybe<Scalars["String"]>;
  sourceCode?: Maybe<SourceCode>;
  stepNum?: Maybe<Scalars["Int"]>;
  terminalis?: Maybe<Array<Maybe<Terminal>>>;
};

export type Terminal = {
  __typename: "Terminal";
  currentDirectory?: Maybe<Array<Maybe<Scalars["String"]>>>;
  elements?: Maybe<Array<Maybe<TerminalElement>>>;
  name?: Maybe<Scalars["String"]>;
};

export type TerminalCommand = {
  __typename: "TerminalCommand";
  command?: Maybe<Scalars["String"]>;
};

export type TerminalCommandSet = {
  __typename: "TerminalCommandSet";
  commands?: Maybe<Array<Maybe<TerminalCommand>>>;
};

export type TerminalElement =
  | TerminalCommand
  | TerminalCommandSet
  | TerminalOutput;

export type TerminalOutput = {
  __typename: "TerminalOutput";
  output?: Maybe<Scalars["String"]>;
};

export type TerminalComponent_FragmentFragment = {
  __typename: "Terminal";
  currentDirectory?: Array<string | null> | null;
  elements?: Array<
    | ({ __typename: "TerminalCommand" } & {
        " $fragmentRefs"?: {
          TerminalElementComponent_Fragment_TerminalCommand_Fragment: TerminalElementComponent_Fragment_TerminalCommand_Fragment;
        };
      })
    | ({ __typename: "TerminalCommandSet" } & {
        " $fragmentRefs"?: {
          TerminalElementComponent_Fragment_TerminalCommandSet_Fragment: TerminalElementComponent_Fragment_TerminalCommandSet_Fragment;
        };
      })
    | ({ __typename: "TerminalOutput" } & {
        " $fragmentRefs"?: {
          TerminalElementComponent_Fragment_TerminalOutput_Fragment: TerminalElementComponent_Fragment_TerminalOutput_Fragment;
        };
      })
    | null
  > | null;
} & { " $fragmentName"?: "TerminalComponent_FragmentFragment" };

export type TerminalCommand_FragmentFragment = {
  __typename: "TerminalCommand";
  command?: string | null;
} & { " $fragmentName"?: "TerminalCommand_FragmentFragment" };

type TerminalElementComponent_Fragment_TerminalCommand_Fragment = ({
  __typename: "TerminalCommand";
} & {
  " $fragmentRefs"?: {
    TerminalCommand_FragmentFragment: TerminalCommand_FragmentFragment;
  };
}) & {
  " $fragmentName"?: "TerminalElementComponent_Fragment_TerminalCommand_Fragment";
};

type TerminalElementComponent_Fragment_TerminalCommandSet_Fragment = {
  __typename: "TerminalCommandSet";
} & {
  " $fragmentName"?: "TerminalElementComponent_Fragment_TerminalCommandSet_Fragment";
};

type TerminalElementComponent_Fragment_TerminalOutput_Fragment = ({
  __typename: "TerminalOutput";
} & {
  " $fragmentRefs"?: {
    TerminalOutput_FragmentFragment: TerminalOutput_FragmentFragment;
  };
}) & {
  " $fragmentName"?: "TerminalElementComponent_Fragment_TerminalOutput_Fragment";
};

export type TerminalElementComponent_FragmentFragment =
  | TerminalElementComponent_Fragment_TerminalCommand_Fragment
  | TerminalElementComponent_Fragment_TerminalCommandSet_Fragment
  | TerminalElementComponent_Fragment_TerminalOutput_Fragment;

export type TerminalOutput_FragmentFragment = {
  __typename: "TerminalOutput";
  output?: string | null;
} & { " $fragmentName"?: "TerminalOutput_FragmentFragment" };

export type Home2_QueryQueryVariables = Exact<{ [key: string]: never }>;

export type Home2_QueryQuery = {
  __typename: "Query";
  terminal?: {
    __typename: "Terminal";
    name?: string | null;
    currentDirectory?: Array<string | null> | null;
    elements?: Array<
      | ({ __typename: "TerminalCommand" } & {
          " $fragmentRefs"?: {
            TerminalCommand_FragmentFragment: TerminalCommand_FragmentFragment;
          };
        })
      | { __typename: "TerminalCommandSet" }
      | ({ __typename: "TerminalOutput" } & {
          " $fragmentRefs"?: {
            TerminalOutput_FragmentFragment: TerminalOutput_FragmentFragment;
          };
        })
      | null
    > | null;
  } | null;
};

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
export const TerminalElementComponent_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalElementComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalElement" },
      },
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
                  name: { kind: "Name", value: "TerminalCommand_Fragment" },
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
                  name: { kind: "Name", value: "TerminalOutput_Fragment" },
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
} as unknown as DocumentNode<
  TerminalElementComponent_FragmentFragment,
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
          { kind: "Field", name: { kind: "Name", value: "currentDirectory" } },
          {
            kind: "Field",
            name: { kind: "Name", value: "elements" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: {
                    kind: "Name",
                    value: "TerminalElementComponent_Fragment",
                  },
                },
              ],
            },
          },
        ],
      },
    },
    ...TerminalElementComponent_FragmentFragmentDoc.definitions,
  ],
} as unknown as DocumentNode<TerminalComponent_FragmentFragment, unknown>;
export const Home2_QueryDocument = {
  kind: "Document",
  definitions: [
    {
      kind: "OperationDefinition",
      operation: "query",
      name: { kind: "Name", value: "Home2_Query" },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "terminal" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "name" } },
                {
                  kind: "Field",
                  name: { kind: "Name", value: "currentDirectory" },
                },
                {
                  kind: "Field",
                  name: { kind: "Name", value: "elements" },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "Field",
                        name: { kind: "Name", value: "__typename" },
                      },
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
      },
    },
    ...TerminalCommand_FragmentFragmentDoc.definitions,
    ...TerminalOutput_FragmentFragmentDoc.definitions,
  ],
} as unknown as DocumentNode<Home2_QueryQuery, Home2_QueryQueryVariables>;
