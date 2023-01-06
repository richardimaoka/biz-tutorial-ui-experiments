import { css } from "@emotion/react";
import Prism from "prismjs";
import "prismjs/components/prism-protobuf"; //ts 7016 error suppressed by prism-fix.d.ts
import "prismjs/themes/prism-tomorrow.css";
import { useEffect, useRef } from "react";
import { IDEEditorTab } from "./IDEEditorTab";
import { IDESideBar } from "./IDESideBar";

const sourceCode = `syntax = "proto3";

// The greeting service definition.
service Greeter {
  // Sends a greeting
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

// The request message containing the user's name.
message HelloRequest {
  string name = 1;
}

// The response message containing the greetings
message HelloReply {
  string message = 1;
}`;

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
      <IDESideBar />
      <div>
        <IDEEditorTab filename="package.json" />
        <div>
          <pre>
            <code
              style={{ fontSize: "12px" }}
              // css={css`
              //   font-size: 12px;
              // `}
              className="language-protobuf"
              ref={ref}
            >
              {sourceCode}
            </code>
          </pre>
        </div>
      </div>
    </div>
  );
};
