import { css } from "@emotion/react";
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
            <code>{sourceCode}</code>
          </pre>
        </div>
      </div>
    </div>
  );
};
