import { css } from "@emotion/react";
import { FileNameNode } from "./FileNameNode";

export interface File {
  filepath: string[];
  filename: string;
  __typename: string;
  offset: number;
}

interface FileTreeViewerProps {
  sourceCodeHeight: number;
  files: File[];
}

export const FileTreeViewer = ({
  files,
  sourceCodeHeight,
}: FileTreeViewerProps): JSX.Element => {
  return (
    <div
      css={css`
        height: ${sourceCodeHeight}px;
        max-width: 160px;
        overflow: scroll;
        ::-webkit-scrollbar {
          width: 8px;
          height: 8px;
          background-color: #252526; /* or add it to the track */
        }
        ::-webkit-scrollbar-thumb {
          background: #37373d;
          border-radius: 8px;
        }
        ::-webkit-scrollbar-corner {
          background-color: #252526;
        }
      `}
    >
      <div
        css={css`
          width: fit-content;
          min-width: 100%;
          min-height: 100%; //expand up to the outer element
          background-color: #252526;
        `}
      >
        {files.map((elem) => (
          <FileNameNode
            key={elem.filepath.join("/")}
            type={elem.__typename}
            offset={elem.offset}
            name={elem.filename}
          />
        ))}
      </div>
    </div>
  );
};
