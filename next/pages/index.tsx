import { css } from "@emotion/react";
import { Command } from "../components/Command";
import { Terminal } from "../components/Terminal";

export default function Home() {
  return (
    <main
      css={css`
        background-color: #f8f8f8;
      `}
    >
      <div
        css={css`
          width: 680px;
          margin: 0 auto;
          background-color: white;
        `}
      >
        <Terminal />
      </div>
    </main>
  );
}
