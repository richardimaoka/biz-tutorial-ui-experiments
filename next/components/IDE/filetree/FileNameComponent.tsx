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
        padding: 4px 8px;
        background-color: #252526;
        color: white;
      `}
    >
      {type === "directory" ? <DirectoryIcon /> : <FileIcon />}
      <div>{name}</div>
    </div>
  );
};
