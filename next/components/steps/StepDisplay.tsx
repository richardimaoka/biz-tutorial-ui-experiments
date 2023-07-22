import { css } from "@emotion/react";

export const StepDisplay = ({ step }: { step: string }) => (
  <div
    css={css`
      position: fixed;
      top: 0px;
      font-size: 8px;
      width: 100%;
      padding: 5px;
      background-color: rgba(145, 145, 145, 0.8);
      color: white;
      border-style: none;
    `}
  >
    {step}
  </div>
);
