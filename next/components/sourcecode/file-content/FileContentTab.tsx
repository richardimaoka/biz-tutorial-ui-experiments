import { css } from "@emotion/react";

interface FileContentTabProps {
  filename: string;
}

export const FileContentTab = ({
  filename,
}: FileContentTabProps): JSX.Element => {
  return (
    <div
      css={css`
        width: fit-content;
        font-size: 13px;
        padding: 4px;
        background-color: #232a36;
        color: white;
      `}
    >
      {filename}
    </div>
  );
};
