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
  "\n  query appTutorialPage($tutorial: String!, $step: String) {\n    page(tutorial: $tutorial, step: $step) {\n      ...GqlTutorialComponent\n      ...GqlNavigation\n    }\n  }\n":
    types.AppTutorialPageDocument,
  "\n  fragment GqlNavigation on Page {\n    prevStep\n    nextStep\n    isTrivial\n  }\n":
    types.GqlNavigationFragmentDoc,
  "\n  fragment GqlSourceCodeColumn on SourceCodeColumn {\n    sourceCode {\n      ...GqlFileTreePane\n\n      openFile {\n        ...GqlOpenFilePane\n      }\n    }\n  }\n":
    types.GqlSourceCodeColumnFragmentDoc,
  "\n  fragment GqlFileTreePane on SourceCode {\n    ...GqlFileTreeHeader\n    ...GqlFileTreeComponent\n    isFoldFileTree\n  }\n":
    types.GqlFileTreePaneFragmentDoc,
  "\n  fragment GqlFileNodeComponent on FileNode {\n    ...GqlFileNodeIcon\n    nodeType\n    name\n    filePath\n    offset\n    isUpdated\n  }\n":
    types.GqlFileNodeComponentFragmentDoc,
  "\n  fragment GqlFileNodeIcon on FileNode {\n    nodeType\n  }\n":
    types.GqlFileNodeIconFragmentDoc,
  "\n  fragment GqlFileTreeHeader on SourceCode {\n    projectDir\n  }\n":
    types.GqlFileTreeHeaderFragmentDoc,
  "\n  fragment GqlFileTreeComponent on SourceCode {\n    fileTree {\n      filePath\n      ...GqlFileNodeComponent\n    }\n  }\n":
    types.GqlFileTreeComponentFragmentDoc,
  "\n  fragment GqlOpenFilePane on OpenFile {\n    ...GqlFileNameTabBar\n    ...GqlSourceCodeEditor\n  }\n":
    types.GqlOpenFilePaneFragmentDoc,
  "\n  fragment GqlSourceCodeEditor on OpenFile {\n    content\n    oldContent\n    language\n    editSequence {\n      id\n      edits {\n        text\n        range {\n          startLineNumber\n          startColumn\n          endLineNumber\n          endColumn\n        }\n      }\n    }\n    tooltip {\n      markdownBody\n      lineNumber\n      timing\n    }\n  }\n":
    types.GqlSourceCodeEditorFragmentDoc,
  "\n  fragment GqlFileNameTabBar on OpenFile {\n    fileName\n  }\n":
    types.GqlFileNameTabBarFragmentDoc,
  "\n  fragment GqlTerminalColumn on TerminalColumn {\n    ...GqlTerminalHeader\n    terminals {\n      ...GqlTerminalContents\n    }\n  }\n":
    types.GqlTerminalColumnFragmentDoc,
  "\n  fragment GqlTerminalContents on Terminal {\n    entries {\n      id\n      ...GqlTerminalEntryComponent\n    }\n    tooltip {\n      ...GqlTerminalTooltip\n    }\n  }\n":
    types.GqlTerminalContentsFragmentDoc,
  "\n  fragment GqlTerminalEntryComponent on TerminalEntry {\n    entryType\n    text\n  }\n":
    types.GqlTerminalEntryComponentFragmentDoc,
  "\n  fragment GqlTerminalHeader on TerminalColumn {\n    terminals {\n      name\n      currentDirectory\n    }\n  }\n":
    types.GqlTerminalHeaderFragmentDoc,
  "\n  fragment GqlTerminalTooltip on TerminalTooltip {\n    markdownBody\n    timing\n  }\n":
    types.GqlTerminalTooltipFragmentDoc,
  "\n  fragment GqlTutorialComponent on Page {\n    ...GqlTutorialHeader\n    ...GqlColumnWrappers\n  }\n":
    types.GqlTutorialComponentFragmentDoc,
  "\n  fragment GqlColumnWrapper on ColumnWrapper {\n    columnName\n    column {\n      # if you forget this, the resulting fragment will have __typename = undefined\n      __typename\n      #\n      # for each column type\n      #\n      ... on TerminalColumn {\n        ...GqlTerminalColumn\n      }\n      ... on SourceCodeColumn {\n        ...GqlSourceCodeColumn\n      }\n    }\n  }\n":
    types.GqlColumnWrapperFragmentDoc,
  "\n  fragment GqlColumnWrappers on Page {\n    focusColumn\n    columns {\n      columnName\n      ...GqlColumnWrapper\n    }\n  }\n":
    types.GqlColumnWrappersFragmentDoc,
  "\n  fragment GqlTutorialHeader on Page {\n    ...GqlColumnTabs\n  }\n":
    types.GqlTutorialHeaderFragmentDoc,
  "\n  fragment GqlColumnTab on ColumnWrapper {\n    columnName\n    columnDisplayName\n    ...GqlColumnTabIcon\n  }\n":
    types.GqlColumnTabFragmentDoc,
  "\n  fragment GqlColumnTabIcon on ColumnWrapper {\n    column {\n      __typename\n    }\n  }\n":
    types.GqlColumnTabIconFragmentDoc,
  "\n  fragment GqlColumnTabs on Page {\n    columns {\n      columnName\n      ...GqlColumnTab\n    }\n    focusColumn\n  }\n":
    types.GqlColumnTabsFragmentDoc,
  "\n  query appTestTerminalPage($step: Int) {\n    _test {\n      appTestTerminalPage(step: $step) {\n        ...GqlTerminalColumn\n      }\n    }\n  }\n":
    types.AppTestTerminalPageDocument,
  "\n  query appTestTutorialColumnsPage {\n    _test {\n      appTestTutorialColumnsPage {\n        ...GqlTutorialComponent\n      }\n    }\n  }\n":
    types.AppTestTutorialColumnsPageDocument,
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
  source: "\n  query appTutorialPage($tutorial: String!, $step: String) {\n    page(tutorial: $tutorial, step: $step) {\n      ...GqlTutorialComponent\n      ...GqlNavigation\n    }\n  }\n",
): (typeof documents)["\n  query appTutorialPage($tutorial: String!, $step: String) {\n    page(tutorial: $tutorial, step: $step) {\n      ...GqlTutorialComponent\n      ...GqlNavigation\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlNavigation on Page {\n    prevStep\n    nextStep\n    isTrivial\n  }\n",
): (typeof documents)["\n  fragment GqlNavigation on Page {\n    prevStep\n    nextStep\n    isTrivial\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlSourceCodeColumn on SourceCodeColumn {\n    sourceCode {\n      ...GqlFileTreePane\n\n      openFile {\n        ...GqlOpenFilePane\n      }\n    }\n  }\n",
): (typeof documents)["\n  fragment GqlSourceCodeColumn on SourceCodeColumn {\n    sourceCode {\n      ...GqlFileTreePane\n\n      openFile {\n        ...GqlOpenFilePane\n      }\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlFileTreePane on SourceCode {\n    ...GqlFileTreeHeader\n    ...GqlFileTreeComponent\n    isFoldFileTree\n  }\n",
): (typeof documents)["\n  fragment GqlFileTreePane on SourceCode {\n    ...GqlFileTreeHeader\n    ...GqlFileTreeComponent\n    isFoldFileTree\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlFileNodeComponent on FileNode {\n    ...GqlFileNodeIcon\n    nodeType\n    name\n    filePath\n    offset\n    isUpdated\n  }\n",
): (typeof documents)["\n  fragment GqlFileNodeComponent on FileNode {\n    ...GqlFileNodeIcon\n    nodeType\n    name\n    filePath\n    offset\n    isUpdated\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlFileNodeIcon on FileNode {\n    nodeType\n  }\n",
): (typeof documents)["\n  fragment GqlFileNodeIcon on FileNode {\n    nodeType\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlFileTreeHeader on SourceCode {\n    projectDir\n  }\n",
): (typeof documents)["\n  fragment GqlFileTreeHeader on SourceCode {\n    projectDir\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlFileTreeComponent on SourceCode {\n    fileTree {\n      filePath\n      ...GqlFileNodeComponent\n    }\n  }\n",
): (typeof documents)["\n  fragment GqlFileTreeComponent on SourceCode {\n    fileTree {\n      filePath\n      ...GqlFileNodeComponent\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlOpenFilePane on OpenFile {\n    ...GqlFileNameTabBar\n    ...GqlSourceCodeEditor\n  }\n",
): (typeof documents)["\n  fragment GqlOpenFilePane on OpenFile {\n    ...GqlFileNameTabBar\n    ...GqlSourceCodeEditor\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlSourceCodeEditor on OpenFile {\n    content\n    oldContent\n    language\n    editSequence {\n      id\n      edits {\n        text\n        range {\n          startLineNumber\n          startColumn\n          endLineNumber\n          endColumn\n        }\n      }\n    }\n    tooltip {\n      markdownBody\n      lineNumber\n      timing\n    }\n  }\n",
): (typeof documents)["\n  fragment GqlSourceCodeEditor on OpenFile {\n    content\n    oldContent\n    language\n    editSequence {\n      id\n      edits {\n        text\n        range {\n          startLineNumber\n          startColumn\n          endLineNumber\n          endColumn\n        }\n      }\n    }\n    tooltip {\n      markdownBody\n      lineNumber\n      timing\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlFileNameTabBar on OpenFile {\n    fileName\n  }\n",
): (typeof documents)["\n  fragment GqlFileNameTabBar on OpenFile {\n    fileName\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlTerminalColumn on TerminalColumn {\n    ...GqlTerminalHeader\n    terminals {\n      ...GqlTerminalContents\n    }\n  }\n",
): (typeof documents)["\n  fragment GqlTerminalColumn on TerminalColumn {\n    ...GqlTerminalHeader\n    terminals {\n      ...GqlTerminalContents\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlTerminalContents on Terminal {\n    entries {\n      id\n      ...GqlTerminalEntryComponent\n    }\n    tooltip {\n      ...GqlTerminalTooltip\n    }\n  }\n",
): (typeof documents)["\n  fragment GqlTerminalContents on Terminal {\n    entries {\n      id\n      ...GqlTerminalEntryComponent\n    }\n    tooltip {\n      ...GqlTerminalTooltip\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlTerminalEntryComponent on TerminalEntry {\n    entryType\n    text\n  }\n",
): (typeof documents)["\n  fragment GqlTerminalEntryComponent on TerminalEntry {\n    entryType\n    text\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlTerminalHeader on TerminalColumn {\n    terminals {\n      name\n      currentDirectory\n    }\n  }\n",
): (typeof documents)["\n  fragment GqlTerminalHeader on TerminalColumn {\n    terminals {\n      name\n      currentDirectory\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlTerminalTooltip on TerminalTooltip {\n    markdownBody\n    timing\n  }\n",
): (typeof documents)["\n  fragment GqlTerminalTooltip on TerminalTooltip {\n    markdownBody\n    timing\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlTutorialComponent on Page {\n    ...GqlTutorialHeader\n    ...GqlColumnWrappers\n  }\n",
): (typeof documents)["\n  fragment GqlTutorialComponent on Page {\n    ...GqlTutorialHeader\n    ...GqlColumnWrappers\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlColumnWrapper on ColumnWrapper {\n    columnName\n    column {\n      # if you forget this, the resulting fragment will have __typename = undefined\n      __typename\n      #\n      # for each column type\n      #\n      ... on TerminalColumn {\n        ...GqlTerminalColumn\n      }\n      ... on SourceCodeColumn {\n        ...GqlSourceCodeColumn\n      }\n    }\n  }\n",
): (typeof documents)["\n  fragment GqlColumnWrapper on ColumnWrapper {\n    columnName\n    column {\n      # if you forget this, the resulting fragment will have __typename = undefined\n      __typename\n      #\n      # for each column type\n      #\n      ... on TerminalColumn {\n        ...GqlTerminalColumn\n      }\n      ... on SourceCodeColumn {\n        ...GqlSourceCodeColumn\n      }\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlColumnWrappers on Page {\n    focusColumn\n    columns {\n      columnName\n      ...GqlColumnWrapper\n    }\n  }\n",
): (typeof documents)["\n  fragment GqlColumnWrappers on Page {\n    focusColumn\n    columns {\n      columnName\n      ...GqlColumnWrapper\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlTutorialHeader on Page {\n    ...GqlColumnTabs\n  }\n",
): (typeof documents)["\n  fragment GqlTutorialHeader on Page {\n    ...GqlColumnTabs\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlColumnTab on ColumnWrapper {\n    columnName\n    columnDisplayName\n    ...GqlColumnTabIcon\n  }\n",
): (typeof documents)["\n  fragment GqlColumnTab on ColumnWrapper {\n    columnName\n    columnDisplayName\n    ...GqlColumnTabIcon\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlColumnTabIcon on ColumnWrapper {\n    column {\n      __typename\n    }\n  }\n",
): (typeof documents)["\n  fragment GqlColumnTabIcon on ColumnWrapper {\n    column {\n      __typename\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlColumnTabs on Page {\n    columns {\n      columnName\n      ...GqlColumnTab\n    }\n    focusColumn\n  }\n",
): (typeof documents)["\n  fragment GqlColumnTabs on Page {\n    columns {\n      columnName\n      ...GqlColumnTab\n    }\n    focusColumn\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  query appTestTerminalPage($step: Int) {\n    _test {\n      appTestTerminalPage(step: $step) {\n        ...GqlTerminalColumn\n      }\n    }\n  }\n",
): (typeof documents)["\n  query appTestTerminalPage($step: Int) {\n    _test {\n      appTestTerminalPage(step: $step) {\n        ...GqlTerminalColumn\n      }\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  query appTestTutorialColumnsPage {\n    _test {\n      appTestTutorialColumnsPage {\n        ...GqlTutorialComponent\n      }\n    }\n  }\n",
): (typeof documents)["\n  query appTestTutorialColumnsPage {\n    _test {\n      appTestTutorialColumnsPage {\n        ...GqlTutorialComponent\n      }\n    }\n  }\n"];

export function graphql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> =
  TDocumentNode extends DocumentNode<infer TType, any> ? TType : never;
