import { css } from "@emotion/react";
import Link from "next/link";

interface PrevButtonProps {
  href: string;
}

export const PrevButton = ({ href }: PrevButtonProps) => (
  <Link href={href}>
    <button
      css={css`
        position: fixed;
        bottom: 0px;
        font-size: 20px;
        width: 100px;
        height: 40px;
        background-color: rgba(255, 255, 255, 0.8);
        color: black;
        border-style: none;
      `}
    >
      prev
    </button>
  </Link>
);
