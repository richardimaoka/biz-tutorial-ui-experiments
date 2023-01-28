import { css } from "@emotion/react";

import { FragmentType, graphql, useFragment } from "../../libs/gql";
import { FileTreePane } from "./file-tree/FileTreePane";
import { FileContentPane } from "./open-file/FileContentPane";

const SourceCodeViewer_Fragment = graphql(`
  fragment SourceCodeViewer_Fragment on SourceCode {
    ...FileTreePane_Fragment
    openFile {
      ...FileContentPane_Fragment
    }
  }
`);

export interface SourceCodeViewerProps {
  fragment: FragmentType<typeof SourceCodeViewer_Fragment>;
}

export const SourceCodeViewer = (props: SourceCodeViewerProps): JSX.Element => {
  const fragment = useFragment(SourceCodeViewer_Fragment, props.fragment);
  const sourceCodeHeight = 400;
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
        <FileTreePane fragment={fragment} sourceCodeHeight={sourceCodeHeight} />
      </div>
      <div
        css={css`
          flex-grow: 1; //necessary for narrower-than-width source code
          max-width: 520px; //necessary for wider-than-width source code
        `}
      >
        {fragment.openFile && (
          <FileContentPane
            fragment={fragment.openFile}
            sourceCodeHeight={sourceCodeHeight}
          />
        )}
      </div>
    </div>
  );
};
