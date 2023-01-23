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

export type Command = {
  __typename: "Command";
  command?: Maybe<Scalars["String"]>;
};

export type CommandOutput = {
  __typename: "CommandOutput";
  output?: Maybe<Scalars["String"]>;
};

export type File = {
  __typename: "File";
  content?: Maybe<Scalars["String"]>;
  filePath?: Maybe<Array<Maybe<Scalars["String"]>>>;
  isFullContent?: Maybe<Scalars["Boolean"]>;
};

export type Ide = {
  __typename: "IDE";
  focusedFile?: Maybe<File>;
};

export type Query = {
  __typename: "Query";
  step?: Maybe<Step>;
};

export type QueryStepArgs = {
  stepId?: InputMaybe<Scalars["ID"]>;
  tutorialId?: InputMaybe<Scalars["ID"]>;
};

export type Step = {
  __typename: "Step";
  id?: Maybe<Scalars["ID"]>;
  ide?: Maybe<Ide>;
  terminal?: Maybe<Array<Maybe<TerminalElement>>>;
};

export type Terminal = {
  __typename: "Terminal";
  elements?: Maybe<Array<Maybe<TerminalElement>>>;
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
