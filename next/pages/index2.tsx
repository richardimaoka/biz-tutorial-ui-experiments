import { css } from "@emotion/react";
import { Header } from "../components/Header";
import { TerminalCommandComponent } from "../components/terminal/TerminalCommandComponent";

export default function Home2() {
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
          <pre
            css={css`
              margin: 0px 0px;
              padding: 4px;
              background-color: #1e1e1e;
              color: #f1f1f1;
              border-bottom: 1px solid #333333;
            `}
          >
            <code>
              {
                "protoc \\\n  --go_out=outdir --go_opt=paths=source_relative \\\n  helloworld.proto"
              }
            </code>
          </pre>
          <pre
            css={css`
              margin: 0px 0px;
              padding: 4px;
              background-color: #1e1e1e;
              color: #f1f1f1;
              border-bottom: 1px solid #333333;
            `}
          >
            <code>outdir/: No such file or directory</code>
          </pre>
        </div>
      </main>
    </>
  );
}
