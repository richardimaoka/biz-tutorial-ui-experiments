import { css } from "@emotion/react";
import { Header } from "../components/Header";
import { SourceCodeViewer } from "../components/sourcecode/SourceCodeViewer";

export default function Home() {
  return (
    <>
      <Header />
      <main
        css={css`
          background-color: #333333;
        `}
      >
        <div
          css={css`
            width: 680px;
            margin: 0 auto;
            background-color: white;
          `}
        >
          <SourceCodeViewer />
        </div>
      </main>
    </>
  );
}
