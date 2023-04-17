import { css } from "@emotion/react";

import { useQuery } from "@apollo/client";
import { useEffect, useState } from "react";
import { FragmentType, graphql, useFragment } from "../../libs/gql";
import { FileTreePane } from "./file-tree/FileTreePane";
import { FileContentPane } from "./open-file/FileContentPane";

const SourceCodeViewer_Fragment = graphql(`
  fragment SourceCodeViewer_Fragment on SourceCode {
    ...FileTreePane_Fragment
    defaultOpenFile {
      ...FileContentPane_Fragment
    }
  }
`);

const OpenFileQuery = graphql(/* GraphQL */ `
  query OpenFileQuery($step: String, $openFilePath: String!) {
    pageState(step: $step) {
      sourceCode {
        openFile(filePath: $openFilePath) {
          ...FileContentPane_Fragment
        }
      }
    }
  }
`);

export interface SourceCodeViewerProps {
  fragment: FragmentType<typeof SourceCodeViewer_Fragment>;
  step: string | undefined;
  currentDirectory?: string;
}

const EmptyFileContentPane = () => (
  <div
    css={css`
      background-color: #1e1e1e;
      width: auto;
      height: 100%; //expand up to the outer element
    `}
  />
);

export const SourceCodeViewer = (props: SourceCodeViewerProps): JSX.Element => {
  const fragment = useFragment(SourceCodeViewer_Fragment, props.fragment);
  const sourceCodeHeight = 400;
  const [openFilePath, setOpenFilePath] = useState<string>("");
  const { data } = useQuery(OpenFileQuery, {
    variables: { step: props.step, openFilePath: openFilePath },
  });
  const openFile = data?.pageState?.sourceCode?.openFile;

  return (
    <div
      css={css`
        display: flex;
      `}
    >
      <div
        css={css`
          flex-grow: 0; //flex-grow distributes the "remaining" space, and FileTreePane should give any remaining space to FileContentPane
        `}
      >
        <FileTreePane
          fragment={fragment}
          sourceCodeHeight={sourceCodeHeight}
          currentDirectory={props.currentDirectory}
          updateOpenFilePath={setOpenFilePath}
        />
      </div>
      <div
        css={css`
          flex-grow: 1; //necessary for narrower-than-width source code
          overflow: hidden; //necessary for wider-than-width source code
        `}
      >
        {openFile ? (
          <FileContentPane
            fragment={openFile}
            sourceCodeHeight={sourceCodeHeight}
          />
        ) : (
          <EmptyFileContentPane />
        )}
      </div>
    </div>
  );
};
