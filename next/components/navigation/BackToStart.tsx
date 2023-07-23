import { css } from "@emotion/react";
import Link from "next/link";

export const BackToStart = () => (
  <Link href="/">
    <div
      css={css`
        position: fixed;
        top: 0px;
        right: 0px;
        font-size: 8px;
        padding: 0px 5px;
        background-color: rgba(255, 255, 255, 0.5);
        color: black;
        border-style: none;
        z-index: 100;
      `}
    >
      back to start
    </div>
  </Link>
);
