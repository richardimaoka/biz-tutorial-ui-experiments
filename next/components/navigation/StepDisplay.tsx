import { css } from "@emotion/react";

export const StepDisplay = ({ step }: { step: string }) => (
  <div
    css={css`
      position: fixed;
      bottom: 40px;
      right: 0px;
      font-size: 8px;
      padding: 0px 5px;
      background-color: rgba(255, 255, 255, 0.5);
      color: black;
      border-style: none;
      z-index: 100;
    `}
  >
    {step}
  </div>
);
