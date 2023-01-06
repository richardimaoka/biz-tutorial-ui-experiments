import { css } from "@emotion/react";
import { useEffect, useRef } from "react";
import { IDEEditorTab } from "./IDEEditorTab";
import { IDESideBar } from "./IDESideBar";

import Prism from "prismjs";
import "prismjs/components/prism-protobuf"; //ts 7016 error suppressed by prism-fix.d.ts
// prism stylesheet /styles/prism-xxx.css is imported from /pages/_app.tsx, as global stylesheet import is only allowed there.
// https://nextjs.org/docs/messages/css-global

const sourceCode = `syntax = "proto3";

// The greeting service definition.
service Greeter {
  // Sends a greeting
  rpc SayHello (HelloRequest) returns (HelloReply) {} rpc SayHello (HelloRequest) returns (HelloReply) {} 
}

// The request message containing the user's name.
message HelloRequest {
  string name = 1;
}

// The response message containing the greetings
message HelloReply {
  string message = 1;
}

`;

export const IDE = (): JSX.Element => {
  const ref = useRef<HTMLElement>(null);
  useEffect(() => {
    if (ref.current) {
      Prism.highlightElement(ref.current);
    }
  }, []);
  return (
    <div
      css={css`
        display: flex;
      `}
    >
      <div
        css={css`
          flex-grow: 0;
        `}
      >
        <IDESideBar />
      </div>
      <div
        css={css`
          flex-grow: 1; //necessary for narrower-than-width source code
          max-width: 520px; //necessary for wider-than-width source code
        `}
      >
        <IDEEditorTab filename="package.json" />
        <div
          css={css`
            height: 312px;
            overflow: scroll; //scroll here, not to include file name tabe in the vertical scroll
            ::-webkit-scrollbar {
              width: 5px;
              height: 5px;
              background-color: #252526; /* or add it to the track */
            }
            ::-webkit-scrollbar-thumb {
              background: #a0a0a0;
              border-radius: 5px;
            }
          `}
        >
          <pre
            css={css`
              width: fit-content;
            `}
          >
            <code className="language-protobuf" ref={ref}>
              {sourceCode}
            </code>
          </pre>
        </div>
      </div>
    </div>
  );
};
