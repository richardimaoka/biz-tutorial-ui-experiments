/* eslint-disable */
import * as types from "./graphql";
import { TypedDocumentNode as DocumentNode } from "@graphql-typed-document-node/core";

/**
 * Map of all GraphQL operations in the project.
 *
 * This map has several performance disadvantages:
 * 1. It is not tree-shakeable, so it will include all operations in the project.
 * 2. It is not minifiable, so the string of a GraphQL query will be multiple times inside the bundle.
 * 3. It does not support dead code elimination, so it will add unused operations.
 *
 * Therefore it is highly recommended to use the babel-plugin for production.
 */
const documents = {
  "\n  fragment TerminalComponent_Fragment on Terminal {\n    currentDirectory\n    elements {\n      ...TerminalElementComponent_Fragment\n    }\n  }\n":
    types.TerminalComponent_FragmentFragmentDoc,
  "\n  fragment TerminalCommand_Fragment on TerminalCommand {\n    command\n  }\n":
    types.TerminalCommand_FragmentFragmentDoc,
  "\n  fragment TerminalElementComponent_Fragment on TerminalElement {\n    __typename\n    ... on TerminalCommand {\n      ...TerminalCommand_Fragment\n    }\n    ... on TerminalOutput {\n      ...TerminalOutput_Fragment\n    }\n  }\n":
    types.TerminalElementComponent_FragmentFragmentDoc,
  "\n  fragment TerminalOutput_Fragment on TerminalOutput {\n    output\n  }\n":
    types.TerminalOutput_FragmentFragmentDoc,
  "\n  query Home2_Query {\n    terminal {\n      name\n      currentDirectory\n      elements {\n        __typename\n        ... on TerminalCommand {\n          ...TerminalCommand_Fragment\n        }\n        ... on TerminalOutput {\n          ...TerminalOutput_Fragment\n        }\n      }\n    }\n  }\n":
    types.Home2_QueryDocument,
};

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 *
 *
 * @example
 * ```ts
 * const query = gql(`query GetUser($id: ID!) { user(id: $id) { name } }`);
 * ```
 *
 * The query argument is unknown!
 * Please regenerate the types.
 **/
export function graphql(source: string): unknown;

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment TerminalComponent_Fragment on Terminal {\n    currentDirectory\n    elements {\n      ...TerminalElementComponent_Fragment\n    }\n  }\n"
): (typeof documents)["\n  fragment TerminalComponent_Fragment on Terminal {\n    currentDirectory\n    elements {\n      ...TerminalElementComponent_Fragment\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment TerminalCommand_Fragment on TerminalCommand {\n    command\n  }\n"
): (typeof documents)["\n  fragment TerminalCommand_Fragment on TerminalCommand {\n    command\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment TerminalElementComponent_Fragment on TerminalElement {\n    __typename\n    ... on TerminalCommand {\n      ...TerminalCommand_Fragment\n    }\n    ... on TerminalOutput {\n      ...TerminalOutput_Fragment\n    }\n  }\n"
): (typeof documents)["\n  fragment TerminalElementComponent_Fragment on TerminalElement {\n    __typename\n    ... on TerminalCommand {\n      ...TerminalCommand_Fragment\n    }\n    ... on TerminalOutput {\n      ...TerminalOutput_Fragment\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment TerminalOutput_Fragment on TerminalOutput {\n    output\n  }\n"
): (typeof documents)["\n  fragment TerminalOutput_Fragment on TerminalOutput {\n    output\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  query Home2_Query {\n    terminal {\n      name\n      currentDirectory\n      elements {\n        __typename\n        ... on TerminalCommand {\n          ...TerminalCommand_Fragment\n        }\n        ... on TerminalOutput {\n          ...TerminalOutput_Fragment\n        }\n      }\n    }\n  }\n"
): (typeof documents)["\n  query Home2_Query {\n    terminal {\n      name\n      currentDirectory\n      elements {\n        __typename\n        ... on TerminalCommand {\n          ...TerminalCommand_Fragment\n        }\n        ... on TerminalOutput {\n          ...TerminalOutput_Fragment\n        }\n      }\n    }\n  }\n"];

export function graphql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> =
  TDocumentNode extends DocumentNode<infer TType, any> ? TType : never;
