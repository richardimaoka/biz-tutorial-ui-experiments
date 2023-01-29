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
  "\n  fragment SourceCodeViewer_Fragment on SourceCode {\n    ...FileTreePane_Fragment\n    openFile {\n      ...FileContentPane_Fragment\n    }\n  }\n":
    types.SourceCodeViewer_FragmentFragmentDoc,
  "\n  fragment FileNodeComponent_Fragment on FileNode {\n    ...FileNodeIcon_Fragment\n    name\n    filePath\n    offset\n    isUpdated\n  }\n":
    types.FileNodeComponent_FragmentFragmentDoc,
  "\n  fragment FileNodeIcon_Fragment on FileNode {\n    nodeType\n  }\n":
    types.FileNodeIcon_FragmentFragmentDoc,
  "\n  fragment FileTreeComponent_Fragment on SourceCode {\n    fileTree {\n      filePath\n      ...FileNodeComponent_Fragment\n    }\n  }\n":
    types.FileTreeComponent_FragmentFragmentDoc,
  "\n  fragment FileTreePane_Fragment on SourceCode {\n    ...FileTreeComponent_Fragment\n  }\n":
    types.FileTreePane_FragmentFragmentDoc,
  "\n  fragment FileContentPane_Fragment on OpenFile {\n    ...FileNameTabBar_Fragment\n    ...FileContentViewer_Fragment\n  }\n":
    types.FileContentPane_FragmentFragmentDoc,
  "\n  fragment FileContentViewer_Fragment on OpenFile {\n    content\n    language\n  }\n":
    types.FileContentViewer_FragmentFragmentDoc,
  "\n  fragment FileNameTab_Fragment on OpenFile {\n    fileName\n  }\n":
    types.FileNameTab_FragmentFragmentDoc,
  "\n  fragment FileNameTabBar_Fragment on OpenFile {\n    ...FileNameTab_Fragment\n  }\n":
    types.FileNameTabBar_FragmentFragmentDoc,
  "\n  fragment TerminalCommand_Fragment on TerminalCommand {\n    command\n  }\n":
    types.TerminalCommand_FragmentFragmentDoc,
  "\n  fragment TerminalCommandWriting_Fragment on TerminalCommand {\n    command\n  }\n":
    types.TerminalCommandWriting_FragmentFragmentDoc,
  "\n  fragment TerminalComponent_Fragment on Terminal {\n    currentDirectory\n    nodes {\n      index\n      ...TerminalNodeComponent_Fragment\n    }\n  }\n":
    types.TerminalComponent_FragmentFragmentDoc,
  "\n  fragment TerminalNodeComponent_Fragment on TerminalNode {\n    index\n    content {\n      __typename\n      ... on TerminalCommand {\n        beforeExecution\n        ...TerminalCommand_Fragment\n        ...TerminalCommandWriting_Fragment\n      }\n      ... on TerminalOutput {\n        ...TerminalOutput_Fragment\n      }\n    }\n  }\n":
    types.TerminalNodeComponent_FragmentFragmentDoc,
  "\n  fragment TerminalOutput_Fragment on TerminalOutput {\n    output\n  }\n":
    types.TerminalOutput_FragmentFragmentDoc,
  "\n  query PageQuery($step: Int!) {\n    step(stepNum: $step) {\n      sourceCode {\n        ...SourceCodeViewer_Fragment\n      }\n      terminals {\n        name\n        currentDirectory\n        ...TerminalComponent_Fragment\n      }\n    }\n  }\n":
    types.PageQueryDocument,
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
  source: "\n  fragment SourceCodeViewer_Fragment on SourceCode {\n    ...FileTreePane_Fragment\n    openFile {\n      ...FileContentPane_Fragment\n    }\n  }\n"
): (typeof documents)["\n  fragment SourceCodeViewer_Fragment on SourceCode {\n    ...FileTreePane_Fragment\n    openFile {\n      ...FileContentPane_Fragment\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment FileNodeComponent_Fragment on FileNode {\n    ...FileNodeIcon_Fragment\n    name\n    filePath\n    offset\n    isUpdated\n  }\n"
): (typeof documents)["\n  fragment FileNodeComponent_Fragment on FileNode {\n    ...FileNodeIcon_Fragment\n    name\n    filePath\n    offset\n    isUpdated\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment FileNodeIcon_Fragment on FileNode {\n    nodeType\n  }\n"
): (typeof documents)["\n  fragment FileNodeIcon_Fragment on FileNode {\n    nodeType\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment FileTreeComponent_Fragment on SourceCode {\n    fileTree {\n      filePath\n      ...FileNodeComponent_Fragment\n    }\n  }\n"
): (typeof documents)["\n  fragment FileTreeComponent_Fragment on SourceCode {\n    fileTree {\n      filePath\n      ...FileNodeComponent_Fragment\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment FileTreePane_Fragment on SourceCode {\n    ...FileTreeComponent_Fragment\n  }\n"
): (typeof documents)["\n  fragment FileTreePane_Fragment on SourceCode {\n    ...FileTreeComponent_Fragment\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment FileContentPane_Fragment on OpenFile {\n    ...FileNameTabBar_Fragment\n    ...FileContentViewer_Fragment\n  }\n"
): (typeof documents)["\n  fragment FileContentPane_Fragment on OpenFile {\n    ...FileNameTabBar_Fragment\n    ...FileContentViewer_Fragment\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment FileContentViewer_Fragment on OpenFile {\n    content\n    language\n  }\n"
): (typeof documents)["\n  fragment FileContentViewer_Fragment on OpenFile {\n    content\n    language\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment FileNameTab_Fragment on OpenFile {\n    fileName\n  }\n"
): (typeof documents)["\n  fragment FileNameTab_Fragment on OpenFile {\n    fileName\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment FileNameTabBar_Fragment on OpenFile {\n    ...FileNameTab_Fragment\n  }\n"
): (typeof documents)["\n  fragment FileNameTabBar_Fragment on OpenFile {\n    ...FileNameTab_Fragment\n  }\n"];
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
  source: "\n  fragment TerminalCommandWriting_Fragment on TerminalCommand {\n    command\n  }\n"
): (typeof documents)["\n  fragment TerminalCommandWriting_Fragment on TerminalCommand {\n    command\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment TerminalComponent_Fragment on Terminal {\n    currentDirectory\n    nodes {\n      index\n      ...TerminalNodeComponent_Fragment\n    }\n  }\n"
): (typeof documents)["\n  fragment TerminalComponent_Fragment on Terminal {\n    currentDirectory\n    nodes {\n      index\n      ...TerminalNodeComponent_Fragment\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment TerminalNodeComponent_Fragment on TerminalNode {\n    index\n    content {\n      __typename\n      ... on TerminalCommand {\n        beforeExecution\n        ...TerminalCommand_Fragment\n        ...TerminalCommandWriting_Fragment\n      }\n      ... on TerminalOutput {\n        ...TerminalOutput_Fragment\n      }\n    }\n  }\n"
): (typeof documents)["\n  fragment TerminalNodeComponent_Fragment on TerminalNode {\n    index\n    content {\n      __typename\n      ... on TerminalCommand {\n        beforeExecution\n        ...TerminalCommand_Fragment\n        ...TerminalCommandWriting_Fragment\n      }\n      ... on TerminalOutput {\n        ...TerminalOutput_Fragment\n      }\n    }\n  }\n"];
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
  source: "\n  query PageQuery($step: Int!) {\n    step(stepNum: $step) {\n      sourceCode {\n        ...SourceCodeViewer_Fragment\n      }\n      terminals {\n        name\n        currentDirectory\n        ...TerminalComponent_Fragment\n      }\n    }\n  }\n"
): (typeof documents)["\n  query PageQuery($step: Int!) {\n    step(stepNum: $step) {\n      sourceCode {\n        ...SourceCodeViewer_Fragment\n      }\n      terminals {\n        name\n        currentDirectory\n        ...TerminalComponent_Fragment\n      }\n    }\n  }\n"];

export function graphql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> =
  TDocumentNode extends DocumentNode<infer TType, any> ? TType : never;
