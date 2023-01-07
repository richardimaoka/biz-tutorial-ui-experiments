import { css } from "@emotion/react";
import { FileNameTab } from "./FileNameTab";

interface FileNameTabBarProps {
  filename: string;
}

export const FileNameTabBar = ({
  filename,
}: FileNameTabBarProps): JSX.Element => {
  return (
    <div
      css={css`
        background-color: #121212;
      `}
    >
      <FileNameTab filename={filename} />
    </div>
  );
};
