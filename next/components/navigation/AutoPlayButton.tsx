import { css } from "@emotion/react";

export const AutoPlayButton = () => (
  <button
    css={css`
      //positioning
      position: fixed;
      bottom: 0px;
      left: 50%;
      transform: translate(-50%, 0%);

      //sizing
      width: 120px;
      height: 40px;

      // color and styles
      background-color: rgba(255, 255, 255, 0.8);
      color: black;
      border-style: none;
    `}
  >
    <div
      css={css`
        font-size: 16px;
        height: 18px;
      `}
    >
      Auto Play
    </div>
    <div
      css={css`
        font-size: 8px;
        line-height: 8px;
      `}
    >
      to next milestone
    </div>
  </button>
);
