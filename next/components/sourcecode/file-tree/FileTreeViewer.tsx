import { css } from "@emotion/react";
import { FileNameComponent } from "../../IDE/filetree/FileNameComponent";

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
          width: 5px;
          height: 5px;
          background-color: #252526; /* or add it to the track */
        }
        ::-webkit-scrollbar-thumb {
          background: #a0a0a0;
          border-radius: 5px;
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
          <FileNameComponent
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
