import { css } from "@emotion/react";
import { DirectoryIcon } from "./DirectoryIcon";
import { FileIcon } from "./FileIcon";

interface FileNameComponentProps {
  type: string; //"directory" | "file";
  offset: number;
  name: string;
}

export const FileNameComponent = ({
  type,
  offset,
  name,
}: FileNameComponentProps): JSX.Element => {
  return (
    <div
      css={css`
        display: flex;
        gap: 4px;
        background-color: #252526;
        color: white;
        font-size: 12px;
        padding-top: 4px;
        padding-bottom: 4px;
        padding-right: 8px;
        padding-left: ${8 * offset + 8}px;
      `}
    >
      {type === "directory" ? <DirectoryIcon /> : <FileIcon />}
      <div
        css={css`
          width: fit-content;
          white-space: nowrap;
        `}
      >
        {name}
      </div>
    </div>
  );
};
