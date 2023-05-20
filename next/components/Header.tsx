import { css } from "@emotion/react";
import { themeBlue } from "../libs/colorTheme";

export const Header = () => {
  return (
    <header
      css={css`
        background-color: ${themeBlue};
      `}
    >
      <div
        css={css`
          height: 40px;
          padding: 4px 20px;
          display: flex;
          align-items: center;
        `}
      >
        <div
          css={css`
            color: white;
            font-size: 16px;
          `}
        >
          The application name
        </div>
      </div>
    </header>
  );
};
