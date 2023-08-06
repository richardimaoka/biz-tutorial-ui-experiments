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
 * Therefore it is highly recommended to use the babel or swc plugin for production.
 */
const documents = {
  "\n  fragment ColumnHeader_Fragment on Page {\n    ...ColumnTabs_Fragment\n  }\n":
    types.ColumnHeader_FragmentFragmentDoc,
  "\n  fragment ColumnTab_Fragment on ColumnWrapper {\n    name\n    column {\n      __typename\n    }\n  }\n":
    types.ColumnTab_FragmentFragmentDoc,
  "\n  fragment ColumnTabs_Fragment on Page {\n    columns {\n      ...ColumnTab_Fragment\n      name\n    }\n  }\n":
    types.ColumnTabs_FragmentFragmentDoc,
  "\n  fragment ColumnWrapperComponent_Fragment on ColumnWrapper {\n    name\n    column {\n      __typename\n      ... on SourceCodeColumn {\n        ...SourceCodeColumn_Fragment\n      }\n\n      ... on TerminalColumn {\n        ...TerminalColumn_Fragment\n      }\n    }\n  }\n":
    types.ColumnWrapperComponent_FragmentFragmentDoc,
  "\n  fragment VisibleColumn_Fragment on Page {\n    ...ColumnHeader_Fragment\n    columns {\n      ...ColumnWrapperComponent_Fragment\n      name\n    }\n  }\n":
    types.VisibleColumn_FragmentFragmentDoc,
  "\n  fragment SourceCodeColumn_Fragment on SourceCodeColumn {\n    sourceCode {\n      ...FileTreePane_Fragment\n      openFile(filePath: $openFilePath) {\n        ...FileContentPane_Fragment\n      }\n    }\n  }\n":
    types.SourceCodeColumn_FragmentFragmentDoc,
  "\n  fragment FileNodeComponent_Fragment on FileNode {\n    ...FileNodeIcon_Fragment\n    nodeType\n    name\n    filePath\n    offset\n    isUpdated\n  }\n":
    types.FileNodeComponent_FragmentFragmentDoc,
  "\n  fragment FileNodeIcon_Fragment on FileNode {\n    nodeType\n  }\n":
    types.FileNodeIcon_FragmentFragmentDoc,
  "\n  fragment FileTreeComponent_Fragment on SourceCode {\n    fileTree {\n      filePath\n      ...FileNodeComponent_Fragment\n    }\n  }\n":
    types.FileTreeComponent_FragmentFragmentDoc,
  "\n  fragment FileTreeHeader_Fragment on SourceCode {\n    projectDir\n  }\n":
    types.FileTreeHeader_FragmentFragmentDoc,
  "\n  fragment FileTreePane_Fragment on SourceCode {\n    ...FileTreeHeader_Fragment\n    ...FileTreeComponent_Fragment\n  }\n":
    types.FileTreePane_FragmentFragmentDoc,
  "\n  fragment FileContentPane_Fragment on OpenFile {\n    ...FileNameTabBar_Fragment\n    ...FileContentViewer_Fragment\n  }\n":
    types.FileContentPane_FragmentFragmentDoc,
  "\n  fragment FileContentViewer_Fragment on OpenFile {\n    content\n    language\n    highlight {\n      fromLine\n      toLine\n    }\n  }\n":
    types.FileContentViewer_FragmentFragmentDoc,
  "\n  fragment FileNameTab_Fragment on OpenFile {\n    fileName\n  }\n":
    types.FileNameTab_FragmentFragmentDoc,
  "\n  fragment FileNameTabBar_Fragment on OpenFile {\n    ...FileNameTab_Fragment\n  }\n":
    types.FileNameTabBar_FragmentFragmentDoc,
  "\n  fragment TerminalCurrentDirectory_Fragment on Terminal {\n    currentDirectory\n  }\n":
    types.TerminalCurrentDirectory_FragmentFragmentDoc,
  "\n  fragment TerminalColumn_Fragment on TerminalColumn {\n    terminal {\n      ...TerminalComponent_Fragment\n    }\n  }\n":
    types.TerminalColumn_FragmentFragmentDoc,
  "\n  fragment TerminalCommand_Fragment on TerminalCommand {\n    command\n    beforeExecution\n  }\n":
    types.TerminalCommand_FragmentFragmentDoc,
  "\n  fragment TerminalComponent_Fragment on Terminal {\n    ...TerminalCurrentDirectory_Fragment\n    ...TerminalContentsComponent_Fragment\n  }\n":
    types.TerminalComponent_FragmentFragmentDoc,
  "\n  fragment TerminalContentsComponent_Fragment on Terminal {\n    nodes {\n      ...TerminalNodeComponent_Fragment\n    }\n  }\n":
    types.TerminalContentsComponent_FragmentFragmentDoc,
  "\n  fragment TerminalNodeComponent_Fragment on TerminalNode {\n    content {\n      __typename\n      ... on TerminalCommand {\n        ...TerminalCommand_Fragment\n      }\n      ... on TerminalOutput {\n        ...TerminalOutput_Fragment\n      }\n    }\n  }\n":
    types.TerminalNodeComponent_FragmentFragmentDoc,
  "\n  fragment TerminalOutput_Fragment on TerminalOutput {\n    output\n  }\n":
    types.TerminalOutput_FragmentFragmentDoc,
  "\n  query PageQuery($tutorial: String!, $step: String, $openFilePath: String) {\n    page(tutorial: $tutorial, step: $step) {\n      ...VisibleColumn_Fragment\n    }\n  }\n":
    types.PageQueryDocument,
};

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 *
 *
 * @example
 * ```ts
 * const query = graphql(`query GetUser($id: ID!) { user(id: $id) { name } }`);
 * ```
 *
 * The query argument is unknown!
 * Please regenerate the types.
 */
export function graphql(source: string): unknown;

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment ColumnHeader_Fragment on Page {\n    ...ColumnTabs_Fragment\n  }\n",
): (typeof documents)["\n  fragment ColumnHeader_Fragment on Page {\n    ...ColumnTabs_Fragment\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment ColumnTab_Fragment on ColumnWrapper {\n    name\n    column {\n      __typename\n    }\n  }\n",
): (typeof documents)["\n  fragment ColumnTab_Fragment on ColumnWrapper {\n    name\n    column {\n      __typename\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment ColumnTabs_Fragment on Page {\n    columns {\n      ...ColumnTab_Fragment\n      name\n    }\n  }\n",
): (typeof documents)["\n  fragment ColumnTabs_Fragment on Page {\n    columns {\n      ...ColumnTab_Fragment\n      name\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment ColumnWrapperComponent_Fragment on ColumnWrapper {\n    name\n    column {\n      __typename\n      ... on SourceCodeColumn {\n        ...SourceCodeColumn_Fragment\n      }\n\n      ... on TerminalColumn {\n        ...TerminalColumn_Fragment\n      }\n    }\n  }\n",
): (typeof documents)["\n  fragment ColumnWrapperComponent_Fragment on ColumnWrapper {\n    name\n    column {\n      __typename\n      ... on SourceCodeColumn {\n        ...SourceCodeColumn_Fragment\n      }\n\n      ... on TerminalColumn {\n        ...TerminalColumn_Fragment\n      }\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment VisibleColumn_Fragment on Page {\n    ...ColumnHeader_Fragment\n    columns {\n      ...ColumnWrapperComponent_Fragment\n      name\n    }\n  }\n",
): (typeof documents)["\n  fragment VisibleColumn_Fragment on Page {\n    ...ColumnHeader_Fragment\n    columns {\n      ...ColumnWrapperComponent_Fragment\n      name\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment SourceCodeColumn_Fragment on SourceCodeColumn {\n    sourceCode {\n      ...FileTreePane_Fragment\n      openFile(filePath: $openFilePath) {\n        ...FileContentPane_Fragment\n      }\n    }\n  }\n",
): (typeof documents)["\n  fragment SourceCodeColumn_Fragment on SourceCodeColumn {\n    sourceCode {\n      ...FileTreePane_Fragment\n      openFile(filePath: $openFilePath) {\n        ...FileContentPane_Fragment\n      }\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment FileNodeComponent_Fragment on FileNode {\n    ...FileNodeIcon_Fragment\n    nodeType\n    name\n    filePath\n    offset\n    isUpdated\n  }\n",
): (typeof documents)["\n  fragment FileNodeComponent_Fragment on FileNode {\n    ...FileNodeIcon_Fragment\n    nodeType\n    name\n    filePath\n    offset\n    isUpdated\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment FileNodeIcon_Fragment on FileNode {\n    nodeType\n  }\n",
): (typeof documents)["\n  fragment FileNodeIcon_Fragment on FileNode {\n    nodeType\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment FileTreeComponent_Fragment on SourceCode {\n    fileTree {\n      filePath\n      ...FileNodeComponent_Fragment\n    }\n  }\n",
): (typeof documents)["\n  fragment FileTreeComponent_Fragment on SourceCode {\n    fileTree {\n      filePath\n      ...FileNodeComponent_Fragment\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment FileTreeHeader_Fragment on SourceCode {\n    projectDir\n  }\n",
): (typeof documents)["\n  fragment FileTreeHeader_Fragment on SourceCode {\n    projectDir\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment FileTreePane_Fragment on SourceCode {\n    ...FileTreeHeader_Fragment\n    ...FileTreeComponent_Fragment\n  }\n",
): (typeof documents)["\n  fragment FileTreePane_Fragment on SourceCode {\n    ...FileTreeHeader_Fragment\n    ...FileTreeComponent_Fragment\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment FileContentPane_Fragment on OpenFile {\n    ...FileNameTabBar_Fragment\n    ...FileContentViewer_Fragment\n  }\n",
): (typeof documents)["\n  fragment FileContentPane_Fragment on OpenFile {\n    ...FileNameTabBar_Fragment\n    ...FileContentViewer_Fragment\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment FileContentViewer_Fragment on OpenFile {\n    content\n    language\n    highlight {\n      fromLine\n      toLine\n    }\n  }\n",
): (typeof documents)["\n  fragment FileContentViewer_Fragment on OpenFile {\n    content\n    language\n    highlight {\n      fromLine\n      toLine\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment FileNameTab_Fragment on OpenFile {\n    fileName\n  }\n",
): (typeof documents)["\n  fragment FileNameTab_Fragment on OpenFile {\n    fileName\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment FileNameTabBar_Fragment on OpenFile {\n    ...FileNameTab_Fragment\n  }\n",
): (typeof documents)["\n  fragment FileNameTabBar_Fragment on OpenFile {\n    ...FileNameTab_Fragment\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment TerminalCurrentDirectory_Fragment on Terminal {\n    currentDirectory\n  }\n",
): (typeof documents)["\n  fragment TerminalCurrentDirectory_Fragment on Terminal {\n    currentDirectory\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment TerminalColumn_Fragment on TerminalColumn {\n    terminal {\n      ...TerminalComponent_Fragment\n    }\n  }\n",
): (typeof documents)["\n  fragment TerminalColumn_Fragment on TerminalColumn {\n    terminal {\n      ...TerminalComponent_Fragment\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment TerminalCommand_Fragment on TerminalCommand {\n    command\n    beforeExecution\n  }\n",
): (typeof documents)["\n  fragment TerminalCommand_Fragment on TerminalCommand {\n    command\n    beforeExecution\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment TerminalComponent_Fragment on Terminal {\n    ...TerminalCurrentDirectory_Fragment\n    ...TerminalContentsComponent_Fragment\n  }\n",
): (typeof documents)["\n  fragment TerminalComponent_Fragment on Terminal {\n    ...TerminalCurrentDirectory_Fragment\n    ...TerminalContentsComponent_Fragment\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment TerminalContentsComponent_Fragment on Terminal {\n    nodes {\n      ...TerminalNodeComponent_Fragment\n    }\n  }\n",
): (typeof documents)["\n  fragment TerminalContentsComponent_Fragment on Terminal {\n    nodes {\n      ...TerminalNodeComponent_Fragment\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment TerminalNodeComponent_Fragment on TerminalNode {\n    content {\n      __typename\n      ... on TerminalCommand {\n        ...TerminalCommand_Fragment\n      }\n      ... on TerminalOutput {\n        ...TerminalOutput_Fragment\n      }\n    }\n  }\n",
): (typeof documents)["\n  fragment TerminalNodeComponent_Fragment on TerminalNode {\n    content {\n      __typename\n      ... on TerminalCommand {\n        ...TerminalCommand_Fragment\n      }\n      ... on TerminalOutput {\n        ...TerminalOutput_Fragment\n      }\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment TerminalOutput_Fragment on TerminalOutput {\n    output\n  }\n",
): (typeof documents)["\n  fragment TerminalOutput_Fragment on TerminalOutput {\n    output\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  query PageQuery($tutorial: String!, $step: String, $openFilePath: String) {\n    page(tutorial: $tutorial, step: $step) {\n      ...VisibleColumn_Fragment\n    }\n  }\n",
): (typeof documents)["\n  query PageQuery($tutorial: String!, $step: String, $openFilePath: String) {\n    page(tutorial: $tutorial, step: $step) {\n      ...VisibleColumn_Fragment\n    }\n  }\n"];

export function graphql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> =
  TDocumentNode extends DocumentNode<infer TType, any> ? TType : never;
