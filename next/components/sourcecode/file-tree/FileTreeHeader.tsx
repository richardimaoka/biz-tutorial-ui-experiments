import { css } from "@emotion/react";
import Image from "next/image";

interface FileTreeHeaderProps {
  isFolded: boolean;
  onButtonClick: () => void;
}

export const FileTreeHeader = ({
  isFolded,
  onButtonClick,
}: FileTreeHeaderProps): JSX.Element => {
  const padding = 4;
  const buttonSize = 16;
  const width = isFolded ? buttonSize + 2 * padding : 160 - 2 * padding;

  return (
    <div
      css={css`
        display: flex;
        padding: ${padding}px;
        width: ${width}px;
        justify-content: end;
        background-color: #222121;
      `}
    >
      <button
        css={css`
          margin: 0px;
          padding: 0px;
        `}
        onClick={onButtonClick}
      >
        <Image
          width={buttonSize}
          height={buttonSize}
          css={css`
            display: block;
            background-color: #f7f7f7;
            border-radius: 2px;
          `}
          src={
            isFolded
              ? "/images/ide-sidebar-expand.svg"
              : "/images/ide-sidebar-shrink.svg"
          }
          alt={isFolded ? "file tree expand icon" : "file tree shrink icon"}
        />
      </button>
    </div>
  );
};
