import { css } from "@emotion/react";
import Link from "next/link";

interface NextButtonProps {
  href: string;
}

export const NextButton = ({ href }: NextButtonProps) => (
  <Link href={href}>
    <button
      css={css`
        position: fixed;
        bottom: 0px;
        right: 0px;
        font-size: 20px;
        width: 100px;
        height: 40px;
        background-color: rgba(255, 255, 255, 0.8);
        color: black;
        border-style: none;
      `}
    >
      next
    </button>
  </Link>
);
