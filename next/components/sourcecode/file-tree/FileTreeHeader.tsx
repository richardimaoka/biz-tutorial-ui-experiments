import { css } from "@emotion/react";

export const FileTreeHeader = (): JSX.Element => {
  const maxWidth = 160;
  return (
    <div
      css={css`
        display: flex;
        padding: 6px 10px 6px 6px;
        max-width: ${maxWidth - 16}px;
        justify-content: end;
        background-color: #222121;
      `}
    >
      <img
        width="16"
        height="16"
        css={css`
          display: block;
          background-color: #f7f7f7;
          border-radius: 2px;
        `}
        src="/images/ide-sidebar-shrink.svg"
      />
    </div>
  );
};
