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

export type TerminalCommand_FragmentFragment = {
  __typename: "TerminalCommand";
  command?: string | null;
} & { " $fragmentName"?: "TerminalCommand_FragmentFragment" };

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
