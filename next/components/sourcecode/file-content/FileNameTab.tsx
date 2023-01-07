import { css } from "@emotion/react";

interface FileNameTabProps {
  filename: string;
}

export const FileNameTab = ({ filename }: FileNameTabProps): JSX.Element => {
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
