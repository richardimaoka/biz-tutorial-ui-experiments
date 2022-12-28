import { css } from "@emotion/react";
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
          width: 1080px;

          margin: 0 auto;
          background-color: white;
        `}
      >
        <Terminal />
      </div>
    </main>
  );
}
