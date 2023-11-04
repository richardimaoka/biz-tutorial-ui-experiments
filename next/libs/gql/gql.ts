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
  "\n  fragment BrowserColumn_Fragment on BrowserColumn {\n    width\n    height\n    path\n  }\n":
    types.BrowserColumn_FragmentFragmentDoc,
  "\n  fragment Carousel_Fragment on Page {\n    columns {\n      ...ColumnWrapperComponent_Fragment\n      name\n    }\n    focusColumn\n    step\n  }\n":
    types.Carousel_FragmentFragmentDoc,
  "\n  fragment ColumnHeader_Fragment on Page {\n    ...ColumnTabs_Fragment\n  }\n":
    types.ColumnHeader_FragmentFragmentDoc,
  "\n  fragment ColumnTab_Fragment on ColumnWrapper {\n    name\n    column {\n      __typename\n    }\n  }\n":
    types.ColumnTab_FragmentFragmentDoc,
  "\n  fragment ColumnTabs_Fragment on Page {\n    columns {\n      ...ColumnTab_Fragment\n      name\n    }\n  }\n":
    types.ColumnTabs_FragmentFragmentDoc,
  "\n  fragment ColumnWrapperComponent_Fragment on ColumnWrapper {\n    name\n    column {\n      __typename\n      ... on SourceCodeColumn {\n        ...SourceCodeColumn_Fragment\n      }\n\n      ... on TerminalColumn {\n        ...TerminalColumn_Fragment\n      }\n\n      ... on BrowserColumn {\n        ...BrowserColumn_Fragment\n      }\n\n      ... on DevToolsColumn {\n        ...DevToolsColumn_Fragment\n      }\n\n      ... on MarkdownColumn {\n        ...MarkdownColumn_Fragment\n      }\n\n      ... on YouTubeColumn {\n        ...YouTubeColumn_Fragment\n      }\n    }\n  }\n":
    types.ColumnWrapperComponent_FragmentFragmentDoc,
  "\n  fragment VisibleColumn_Fragment on Page {\n    ...ColumnHeader_Fragment\n    ...Carousel_Fragment\n    columns {\n      ...ColumnWrapperComponent_Fragment\n      name\n    }\n    modal {\n      ...ModalComponentFragment\n    }\n    ...Navigation_Fragment\n  }\n":
    types.VisibleColumn_FragmentFragmentDoc,
  "\n  fragment ModalComponentFragment on Modal {\n    text\n    position\n  }\n":
    types.ModalComponentFragmentFragmentDoc,
  "\n  fragment DevToolsColumn_Fragment on DevToolsColumn {\n    width\n    height\n    path\n  }\n":
    types.DevToolsColumn_FragmentFragmentDoc,
  "\n  fragment MarkdownColumn_Fragment on MarkdownColumn {\n    description {\n      ...MarkdownFragment\n    }\n    contentsPosition\n  }\n":
    types.MarkdownColumn_FragmentFragmentDoc,
  "\n  fragment MarkdownFragment on Markdown {\n    contents\n    alignment\n  }\n":
    types.MarkdownFragmentFragmentDoc,
  "\n  fragment Navigation_Fragment on Page {\n    step\n    nextStep\n    prevStep\n    durationSeconds\n    isTrivialStep\n  }\n":
    types.Navigation_FragmentFragmentDoc,
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
  "\n  fragment FileTreePane_Fragment on SourceCode {\n    ...FileTreeHeader_Fragment\n    ...FileTreeComponent_Fragment\n    isFoldFileTree\n  }\n":
    types.FileTreePane_FragmentFragmentDoc,
  "\n  fragment FileContentPane_Fragment on OpenFile {\n    ...FileNameTabBar_Fragment\n    ...FileContentViewer_Fragment\n  }\n":
    types.FileContentPane_FragmentFragmentDoc,
  "\n  fragment FileContentViewer_Fragment on OpenFile {\n    content\n    language\n    highlight {\n      fromLine\n      toLine\n    }\n  }\n":
    types.FileContentViewer_FragmentFragmentDoc,
  "\n  fragment FileNameTab_Fragment on OpenFile {\n    fileName\n  }\n":
    types.FileNameTab_FragmentFragmentDoc,
  "\n  fragment FileNameTabBar_Fragment on OpenFile {\n    ...FileNameTab_Fragment\n  }\n":
    types.FileNameTabBar_FragmentFragmentDoc,
  "\n  fragment GqlFileContentPane on OpenFile {\n    ...GqlFileNameTabBar\n    ...GqlSourceCodeEditor\n  }\n":
    types.GqlFileContentPaneFragmentDoc,
  "\n  fragment GqlSourceCodeEditor on OpenFile {\n    content\n    language\n  }\n":
    types.GqlSourceCodeEditorFragmentDoc,
  "\n  fragment GqlFileTreePane on SourceCode {\n    ...GqlFileTreeHeader\n    ...GqlFileTreeComponent\n    isFoldFileTree\n  }\n":
    types.GqlFileTreePaneFragmentDoc,
  "\n  fragment GqlFileNodeComponent on FileNode {\n    ...GqlFileNodeIcon\n    nodeType\n    name\n    filePath\n    offset\n    isUpdated\n  }\n":
    types.GqlFileNodeComponentFragmentDoc,
  "\n  fragment GqlFileNodeIcon on FileNode {\n    nodeType\n  }\n":
    types.GqlFileNodeIconFragmentDoc,
  "\n  fragment GqlFileTreeHeader on SourceCode {\n    projectDir\n  }\n":
    types.GqlFileTreeHeaderFragmentDoc,
  "\n  fragment GqlFileNameTabBar on OpenFile {\n    fileName\n  }\n":
    types.GqlFileNameTabBarFragmentDoc,
  "\n  fragment GqlFileTreeComponent on SourceCode {\n    fileTree {\n      filePath\n      ...GqlFileNodeComponent\n    }\n  }\n":
    types.GqlFileTreeComponentFragmentDoc,
  "\n  fragment TerminalCurrentDirectory_Fragment on Terminal {\n    currentDirectory\n  }\n":
    types.TerminalCurrentDirectory_FragmentFragmentDoc,
  "\n  fragment TerminalColumn_Fragment on TerminalColumn {\n    terminal {\n      ...TerminalComponent_Fragment\n    }\n  }\n":
    types.TerminalColumn_FragmentFragmentDoc,
  "\n  fragment TerminalCommand_Fragment on TerminalCommand {\n    command\n    beforeExecution\n    ...TerminalCommandTooltip\n  }\n":
    types.TerminalCommand_FragmentFragmentDoc,
  "\n  fragment TerminalCommandTooltip on TerminalCommand {\n    tooltip\n  }\n":
    types.TerminalCommandTooltipFragmentDoc,
  "\n  fragment TerminalComponent_Fragment on Terminal {\n    ...TerminalCurrentDirectory_Fragment\n    ...TerminalContentsComponent_Fragment\n  }\n":
    types.TerminalComponent_FragmentFragmentDoc,
  "\n  fragment TerminalContentsComponent_Fragment on Terminal {\n    nodes {\n      ...TerminalNodeComponent_Fragment\n    }\n  }\n":
    types.TerminalContentsComponent_FragmentFragmentDoc,
  "\n  fragment TerminalNodeComponent_Fragment on TerminalNode {\n    content {\n      __typename\n      ... on TerminalCommand {\n        ...TerminalCommand_Fragment\n      }\n      ... on TerminalOutput {\n        ...TerminalOutput_Fragment\n      }\n    }\n  }\n":
    types.TerminalNodeComponent_FragmentFragmentDoc,
  "\n  fragment TerminalOutput_Fragment on TerminalOutput {\n    output\n  }\n":
    types.TerminalOutput_FragmentFragmentDoc,
  "\n  fragment GqlTerminalColumn on TerminalColumn2 {\n    ...GqlTerminalHeader\n    terminals {\n      ...GqlTerminalContents\n    }\n  }\n":
    types.GqlTerminalColumnFragmentDoc,
  "\n  fragment GqlTerminalContents on Terminal2 {\n    entries {\n      id\n      ...GqlTerminalEntryComponent\n    }\n    tooltip {\n      ...GqlTerminalTooltip\n    }\n  }\n":
    types.GqlTerminalContentsFragmentDoc,
  "\n  fragment GqlTerminalEntryComponent on TerminalEntry {\n    entryType\n    text\n  }\n":
    types.GqlTerminalEntryComponentFragmentDoc,
  "\n  fragment GqlTerminalHeader on TerminalColumn2 {\n    terminals {\n      name\n      currentDirectory\n    }\n  }\n":
    types.GqlTerminalHeaderFragmentDoc,
  "\n  fragment GqlTerminalTooltip on TerminalTooltip2 {\n    markdownBody\n    timing\n  }\n":
    types.GqlTerminalTooltipFragmentDoc,
  "\n  fragment GqlTutorialComponent on Page2 {\n    ...GqlTutorialHeader\n    ...GqlColumnWrappers\n  }\n":
    types.GqlTutorialComponentFragmentDoc,
  "\n  fragment GqlColumnWrapper on ColumnWrapper2 {\n    columnName\n    column {\n      # if you forget this, the resulting fragment will have __typename = undefined\n      __typename\n      #\n      # for each column type\n      #\n      ... on TerminalColumn2 {\n        ...GqlTerminalColumn\n      }\n    }\n  }\n":
    types.GqlColumnWrapperFragmentDoc,
  "\n  fragment GqlColumnWrappers on Page2 {\n    columns {\n      columnName\n      ...GqlColumnWrapper\n    }\n  }\n":
    types.GqlColumnWrappersFragmentDoc,
  "\n  fragment GqlTutorialHeader on Page2 {\n    ...GqlColumnTabs\n  }\n":
    types.GqlTutorialHeaderFragmentDoc,
  "\n  fragment GqlColumnTab on ColumnWrapper2 {\n    columnName\n    ...GqlColumnTabIcon\n  }\n":
    types.GqlColumnTabFragmentDoc,
  "\n  fragment GqlColumnTabIcon on ColumnWrapper2 {\n    column {\n      __typename\n    }\n  }\n":
    types.GqlColumnTabIconFragmentDoc,
  "\n  fragment GqlColumnTabs on Page2 {\n    columns {\n      columnName\n      ...GqlColumnTab\n    }\n  }\n":
    types.GqlColumnTabsFragmentDoc,
  "\n  fragment YouTubeColumn_Fragment on YouTubeColumn {\n    youtube {\n      ...YouTube_Fragment\n    }\n  }\n":
    types.YouTubeColumn_FragmentFragmentDoc,
  "\n  fragment YouTube_Fragment on YouTubeEmbed {\n    embedUrl\n    width\n    height\n  }\n":
    types.YouTube_FragmentFragmentDoc,
  "\n  query PageQuery($tutorial: String!, $step: String, $openFilePath: String) {\n    page(tutorial: $tutorial, step: $step) {\n      ...VisibleColumn_Fragment\n      step\n      focusColumn\n      autoNextSeconds\n    }\n  }\n":
    types.PageQueryDocument,
  "\n  query appTestSourcecodeFilecontentPage {\n    _test {\n      appTestSourcecodeFilecontentPage {\n        ...GqlFileContentPane\n      }\n    }\n  }\n":
    types.AppTestSourcecodeFilecontentPageDocument,
  "\n  query appTestTerminalPage($step: Int) {\n    _test {\n      appTestTerminalPage(step: $step) {\n        ...GqlTerminalColumn\n      }\n    }\n  }\n":
    types.AppTestTerminalPageDocument,
  "\n  query appTestTutorialColumnsPage {\n    _test {\n      appTestTutorialColumnsPage {\n        ...GqlColumnWrappers\n      }\n    }\n  }\n":
    types.AppTestTutorialColumnsPageDocument,
  "\n  query appTestTutorialTutorialPage {\n    _test {\n      appTestTutorialTutorialPage {\n        ...GqlTutorialComponent\n      }\n    }\n  }\n":
    types.AppTestTutorialTutorialPageDocument,
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
  source: "\n  fragment BrowserColumn_Fragment on BrowserColumn {\n    width\n    height\n    path\n  }\n",
): (typeof documents)["\n  fragment BrowserColumn_Fragment on BrowserColumn {\n    width\n    height\n    path\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment Carousel_Fragment on Page {\n    columns {\n      ...ColumnWrapperComponent_Fragment\n      name\n    }\n    focusColumn\n    step\n  }\n",
): (typeof documents)["\n  fragment Carousel_Fragment on Page {\n    columns {\n      ...ColumnWrapperComponent_Fragment\n      name\n    }\n    focusColumn\n    step\n  }\n"];
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
  source: "\n  fragment ColumnWrapperComponent_Fragment on ColumnWrapper {\n    name\n    column {\n      __typename\n      ... on SourceCodeColumn {\n        ...SourceCodeColumn_Fragment\n      }\n\n      ... on TerminalColumn {\n        ...TerminalColumn_Fragment\n      }\n\n      ... on BrowserColumn {\n        ...BrowserColumn_Fragment\n      }\n\n      ... on DevToolsColumn {\n        ...DevToolsColumn_Fragment\n      }\n\n      ... on MarkdownColumn {\n        ...MarkdownColumn_Fragment\n      }\n\n      ... on YouTubeColumn {\n        ...YouTubeColumn_Fragment\n      }\n    }\n  }\n",
): (typeof documents)["\n  fragment ColumnWrapperComponent_Fragment on ColumnWrapper {\n    name\n    column {\n      __typename\n      ... on SourceCodeColumn {\n        ...SourceCodeColumn_Fragment\n      }\n\n      ... on TerminalColumn {\n        ...TerminalColumn_Fragment\n      }\n\n      ... on BrowserColumn {\n        ...BrowserColumn_Fragment\n      }\n\n      ... on DevToolsColumn {\n        ...DevToolsColumn_Fragment\n      }\n\n      ... on MarkdownColumn {\n        ...MarkdownColumn_Fragment\n      }\n\n      ... on YouTubeColumn {\n        ...YouTubeColumn_Fragment\n      }\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment VisibleColumn_Fragment on Page {\n    ...ColumnHeader_Fragment\n    ...Carousel_Fragment\n    columns {\n      ...ColumnWrapperComponent_Fragment\n      name\n    }\n    modal {\n      ...ModalComponentFragment\n    }\n    ...Navigation_Fragment\n  }\n",
): (typeof documents)["\n  fragment VisibleColumn_Fragment on Page {\n    ...ColumnHeader_Fragment\n    ...Carousel_Fragment\n    columns {\n      ...ColumnWrapperComponent_Fragment\n      name\n    }\n    modal {\n      ...ModalComponentFragment\n    }\n    ...Navigation_Fragment\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment ModalComponentFragment on Modal {\n    text\n    position\n  }\n",
): (typeof documents)["\n  fragment ModalComponentFragment on Modal {\n    text\n    position\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment DevToolsColumn_Fragment on DevToolsColumn {\n    width\n    height\n    path\n  }\n",
): (typeof documents)["\n  fragment DevToolsColumn_Fragment on DevToolsColumn {\n    width\n    height\n    path\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment MarkdownColumn_Fragment on MarkdownColumn {\n    description {\n      ...MarkdownFragment\n    }\n    contentsPosition\n  }\n",
): (typeof documents)["\n  fragment MarkdownColumn_Fragment on MarkdownColumn {\n    description {\n      ...MarkdownFragment\n    }\n    contentsPosition\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment MarkdownFragment on Markdown {\n    contents\n    alignment\n  }\n",
): (typeof documents)["\n  fragment MarkdownFragment on Markdown {\n    contents\n    alignment\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment Navigation_Fragment on Page {\n    step\n    nextStep\n    prevStep\n    durationSeconds\n    isTrivialStep\n  }\n",
): (typeof documents)["\n  fragment Navigation_Fragment on Page {\n    step\n    nextStep\n    prevStep\n    durationSeconds\n    isTrivialStep\n  }\n"];
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
  source: "\n  fragment FileTreePane_Fragment on SourceCode {\n    ...FileTreeHeader_Fragment\n    ...FileTreeComponent_Fragment\n    isFoldFileTree\n  }\n",
): (typeof documents)["\n  fragment FileTreePane_Fragment on SourceCode {\n    ...FileTreeHeader_Fragment\n    ...FileTreeComponent_Fragment\n    isFoldFileTree\n  }\n"];
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
  source: "\n  fragment GqlFileContentPane on OpenFile {\n    ...GqlFileNameTabBar\n    ...GqlSourceCodeEditor\n  }\n",
): (typeof documents)["\n  fragment GqlFileContentPane on OpenFile {\n    ...GqlFileNameTabBar\n    ...GqlSourceCodeEditor\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlSourceCodeEditor on OpenFile {\n    content\n    language\n  }\n",
): (typeof documents)["\n  fragment GqlSourceCodeEditor on OpenFile {\n    content\n    language\n  }\n"];
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
  source: "\n  fragment GqlFileNameTabBar on OpenFile {\n    fileName\n  }\n",
): (typeof documents)["\n  fragment GqlFileNameTabBar on OpenFile {\n    fileName\n  }\n"];
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
  source: "\n  fragment TerminalCommand_Fragment on TerminalCommand {\n    command\n    beforeExecution\n    ...TerminalCommandTooltip\n  }\n",
): (typeof documents)["\n  fragment TerminalCommand_Fragment on TerminalCommand {\n    command\n    beforeExecution\n    ...TerminalCommandTooltip\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment TerminalCommandTooltip on TerminalCommand {\n    tooltip\n  }\n",
): (typeof documents)["\n  fragment TerminalCommandTooltip on TerminalCommand {\n    tooltip\n  }\n"];
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
  source: "\n  fragment GqlTerminalColumn on TerminalColumn2 {\n    ...GqlTerminalHeader\n    terminals {\n      ...GqlTerminalContents\n    }\n  }\n",
): (typeof documents)["\n  fragment GqlTerminalColumn on TerminalColumn2 {\n    ...GqlTerminalHeader\n    terminals {\n      ...GqlTerminalContents\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlTerminalContents on Terminal2 {\n    entries {\n      id\n      ...GqlTerminalEntryComponent\n    }\n    tooltip {\n      ...GqlTerminalTooltip\n    }\n  }\n",
): (typeof documents)["\n  fragment GqlTerminalContents on Terminal2 {\n    entries {\n      id\n      ...GqlTerminalEntryComponent\n    }\n    tooltip {\n      ...GqlTerminalTooltip\n    }\n  }\n"];
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
  source: "\n  fragment GqlTerminalHeader on TerminalColumn2 {\n    terminals {\n      name\n      currentDirectory\n    }\n  }\n",
): (typeof documents)["\n  fragment GqlTerminalHeader on TerminalColumn2 {\n    terminals {\n      name\n      currentDirectory\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlTerminalTooltip on TerminalTooltip2 {\n    markdownBody\n    timing\n  }\n",
): (typeof documents)["\n  fragment GqlTerminalTooltip on TerminalTooltip2 {\n    markdownBody\n    timing\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlTutorialComponent on Page2 {\n    ...GqlTutorialHeader\n    ...GqlColumnWrappers\n  }\n",
): (typeof documents)["\n  fragment GqlTutorialComponent on Page2 {\n    ...GqlTutorialHeader\n    ...GqlColumnWrappers\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlColumnWrapper on ColumnWrapper2 {\n    columnName\n    column {\n      # if you forget this, the resulting fragment will have __typename = undefined\n      __typename\n      #\n      # for each column type\n      #\n      ... on TerminalColumn2 {\n        ...GqlTerminalColumn\n      }\n    }\n  }\n",
): (typeof documents)["\n  fragment GqlColumnWrapper on ColumnWrapper2 {\n    columnName\n    column {\n      # if you forget this, the resulting fragment will have __typename = undefined\n      __typename\n      #\n      # for each column type\n      #\n      ... on TerminalColumn2 {\n        ...GqlTerminalColumn\n      }\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlColumnWrappers on Page2 {\n    columns {\n      columnName\n      ...GqlColumnWrapper\n    }\n  }\n",
): (typeof documents)["\n  fragment GqlColumnWrappers on Page2 {\n    columns {\n      columnName\n      ...GqlColumnWrapper\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlTutorialHeader on Page2 {\n    ...GqlColumnTabs\n  }\n",
): (typeof documents)["\n  fragment GqlTutorialHeader on Page2 {\n    ...GqlColumnTabs\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlColumnTab on ColumnWrapper2 {\n    columnName\n    ...GqlColumnTabIcon\n  }\n",
): (typeof documents)["\n  fragment GqlColumnTab on ColumnWrapper2 {\n    columnName\n    ...GqlColumnTabIcon\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlColumnTabIcon on ColumnWrapper2 {\n    column {\n      __typename\n    }\n  }\n",
): (typeof documents)["\n  fragment GqlColumnTabIcon on ColumnWrapper2 {\n    column {\n      __typename\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment GqlColumnTabs on Page2 {\n    columns {\n      columnName\n      ...GqlColumnTab\n    }\n  }\n",
): (typeof documents)["\n  fragment GqlColumnTabs on Page2 {\n    columns {\n      columnName\n      ...GqlColumnTab\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment YouTubeColumn_Fragment on YouTubeColumn {\n    youtube {\n      ...YouTube_Fragment\n    }\n  }\n",
): (typeof documents)["\n  fragment YouTubeColumn_Fragment on YouTubeColumn {\n    youtube {\n      ...YouTube_Fragment\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  fragment YouTube_Fragment on YouTubeEmbed {\n    embedUrl\n    width\n    height\n  }\n",
): (typeof documents)["\n  fragment YouTube_Fragment on YouTubeEmbed {\n    embedUrl\n    width\n    height\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  query PageQuery($tutorial: String!, $step: String, $openFilePath: String) {\n    page(tutorial: $tutorial, step: $step) {\n      ...VisibleColumn_Fragment\n      step\n      focusColumn\n      autoNextSeconds\n    }\n  }\n",
): (typeof documents)["\n  query PageQuery($tutorial: String!, $step: String, $openFilePath: String) {\n    page(tutorial: $tutorial, step: $step) {\n      ...VisibleColumn_Fragment\n      step\n      focusColumn\n      autoNextSeconds\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  query appTestSourcecodeFilecontentPage {\n    _test {\n      appTestSourcecodeFilecontentPage {\n        ...GqlFileContentPane\n      }\n    }\n  }\n",
): (typeof documents)["\n  query appTestSourcecodeFilecontentPage {\n    _test {\n      appTestSourcecodeFilecontentPage {\n        ...GqlFileContentPane\n      }\n    }\n  }\n"];
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
  source: "\n  query appTestTutorialColumnsPage {\n    _test {\n      appTestTutorialColumnsPage {\n        ...GqlColumnWrappers\n      }\n    }\n  }\n",
): (typeof documents)["\n  query appTestTutorialColumnsPage {\n    _test {\n      appTestTutorialColumnsPage {\n        ...GqlColumnWrappers\n      }\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: "\n  query appTestTutorialTutorialPage {\n    _test {\n      appTestTutorialTutorialPage {\n        ...GqlTutorialComponent\n      }\n    }\n  }\n",
): (typeof documents)["\n  query appTestTutorialTutorialPage {\n    _test {\n      appTestTutorialTutorialPage {\n        ...GqlTutorialComponent\n      }\n    }\n  }\n"];

export function graphql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> =
  TDocumentNode extends DocumentNode<infer TType, any> ? TType : never;
