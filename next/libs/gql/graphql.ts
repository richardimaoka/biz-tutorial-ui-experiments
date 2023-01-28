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

export type QueryTerminalArgs = {
  step: Scalars["Int"];
};

export type SourceCode = {
  __typename: "SourceCode";
  fileTree?: Maybe<Array<Maybe<FileTreeNode>>>;
  openFile?: Maybe<OpenFile>;
};

export type Step = {
  __typename: "Step";
  file?: Maybe<OpenFile>;
  nextAction?: Maybe<Scalars["String"]>;
  sourceCode?: Maybe<SourceCode>;
  stepNum?: Maybe<Scalars["Int"]>;
  terminalis?: Maybe<Array<Maybe<Terminal>>>;
};

export type StepFileArgs = {
  filePath: Array<InputMaybe<Scalars["String"]>>;
};

export type Terminal = {
  __typename: "Terminal";
  currentDirectory?: Maybe<Array<Maybe<Scalars["String"]>>>;
  name?: Maybe<Scalars["String"]>;
  nodes?: Maybe<Array<Maybe<TerminalNode>>>;
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

export type TerminalNode = {
  __typename: "TerminalNode";
  content?: Maybe<TerminalElement>;
  index?: Maybe<Scalars["Int"]>;
};

export type TerminalOutput = {
  __typename: "TerminalOutput";
  output?: Maybe<Scalars["String"]>;
};

export type TerminalCommand_FragmentFragment = {
  __typename: "TerminalCommand";
  command?: string | null;
} & { " $fragmentName"?: "TerminalCommand_FragmentFragment" };

export type TerminalCommandWriting_FragmentFragment = {
  __typename: "TerminalCommand";
  command?: string | null;
} & { " $fragmentName"?: "TerminalCommandWriting_FragmentFragment" };

export type TerminalComponent_FragmentFragment = {
  __typename: "Terminal";
  currentDirectory?: Array<string | null> | null;
  nodes?: Array<
    | ({ __typename: "TerminalNode"; index?: number | null } & {
        " $fragmentRefs"?: {
          TerminalNodeComponent_FragmentFragment: TerminalNodeComponent_FragmentFragment;
        };
      })
    | null
  > | null;
} & { " $fragmentName"?: "TerminalComponent_FragmentFragment" };

export type TerminalNodeComponent_FragmentFragment = {
  __typename: "TerminalNode";
  index?: number | null;
  content?:
    | ({ __typename: "TerminalCommand" } & {
        " $fragmentRefs"?: {
          TerminalCommand_FragmentFragment: TerminalCommand_FragmentFragment;
          TerminalCommandWriting_FragmentFragment: TerminalCommandWriting_FragmentFragment;
        };
      })
    | { __typename: "TerminalCommandSet" }
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

export type Home2_QueryQueryVariables = Exact<{
  step: Scalars["Int"];
}>;

export type Home2_QueryQuery = {
  __typename: "Query";
  terminal?:
    | ({ __typename: "Terminal" } & {
        " $fragmentRefs"?: {
          TerminalComponent_FragmentFragment: TerminalComponent_FragmentFragment;
        };
      })
    | null;
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
export const TerminalCommandWriting_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalCommandWriting_Fragment" },
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
} as unknown as DocumentNode<TerminalCommandWriting_FragmentFragment, unknown>;
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
          { kind: "Field", name: { kind: "Name", value: "index" } },
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
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "TerminalCommandWriting_Fragment",
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
    ...TerminalCommandWriting_FragmentFragmentDoc.definitions,
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
                { kind: "Field", name: { kind: "Name", value: "index" } },
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
export const Home2_QueryDocument = {
  kind: "Document",
  definitions: [
    {
      kind: "OperationDefinition",
      operation: "query",
      name: { kind: "Name", value: "Home2_Query" },
      variableDefinitions: [
        {
          kind: "VariableDefinition",
          variable: { kind: "Variable", name: { kind: "Name", value: "step" } },
          type: {
            kind: "NonNullType",
            type: { kind: "NamedType", name: { kind: "Name", value: "Int" } },
          },
        },
      ],
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "terminal" },
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
                  name: { kind: "Name", value: "TerminalComponent_Fragment" },
                },
              ],
            },
          },
        ],
      },
    },
    ...TerminalComponent_FragmentFragmentDoc.definitions,
  ],
} as unknown as DocumentNode<Home2_QueryQuery, Home2_QueryQueryVariables>;
