import { css } from "@emotion/react";
import { useEffect, useRef } from "react";

import { File } from "./sourcecode/file-tree/FileTreeViewer";
import { FileContentPane } from "./sourcecode/file-content/FileContentPane";
import { FileTreePane } from "./sourcecode/file-tree/FileTreePane";

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

const files: File[] = [
  { offset: 0, filepath: [], __typename: "directory", filename: "next" },
  { offset: 1, filepath: [], __typename: "directory", filename: "cache" },
  { offset: 1, filepath: [], __typename: "directory", filename: "server" },
  { offset: 1, filepath: [], __typename: "directory", filename: "static" },
  {
    offset: 1,
    filepath: [],
    __typename: "file",
    filename: "build-manifest.json",
  },
  { offset: 1, filepath: [], __typename: "file", filename: "package.json" },
  {
    offset: 1,
    filepath: [],
    __typename: "file",
    filename: "react-loadable-manifest.json",
  },
  { offset: 1, filepath: [], __typename: "directory", filename: "trace" },
  { offset: 0, filepath: [], __typename: "directory", filename: "components" },
  { offset: 1, filepath: [], __typename: "file", filename: "Terminal.tsx" },
  { offset: 0, filepath: [], __typename: "directory", filename: "node_module" },
  { offset: 0, filepath: [], __typename: "directory", filename: "pages" },
  { offset: 1, filepath: [], __typename: "directory", filename: "api" },
  { offset: 2, filepath: [], __typename: "file", filename: "hello.ts" },
  { offset: 1, filepath: [], __typename: "file", filename: "_app.tsx" },
  { offset: 1, filepath: [], __typename: "file", filename: "_document.tsx" },
  { offset: 1, filepath: [], __typename: "file", filename: "index.tsx" },
  { offset: 0, filepath: [], __typename: "directory", filename: "public" },
  { offset: 0, filepath: [], __typename: "directory", filename: "styles" },
  { offset: 1, filepath: [], __typename: "file", filename: "global.css" },
  { offset: 1, filepath: [], __typename: "file", filename: "Home.module.css" },
  { offset: 1, filepath: [], __typename: "directory", filename: ".babelrc" },
  { offset: 0, filepath: [], __typename: "file", filename: ".eslintrc.json" },
  { offset: 0, filepath: [], __typename: "directory", filename: ".gitignore" },
  { offset: 0, filepath: [], __typename: "file", filename: "next-env.d.ts" },
  { offset: 0, filepath: [], __typename: "file", filename: "next.config.js" },
  {
    offset: 0,
    filepath: [],
    __typename: "file",
    filename: "package-lock.json",
  },
  { offset: 0, filepath: [], __typename: "file", filename: "package.json" },
  { offset: 0, filepath: [], __typename: "file", filename: "README.md" },
  { offset: 0, filepath: [], __typename: "file", filename: "tsconfig.json" },
];

const sourceCodeHeight = 400;

export const IDE = (): JSX.Element => {
  return (
    <div
      css={css`
        display: flex;
      `}
    >
      <div
        css={css`
          flex-grow: 1;
        `}
      >
        <FileTreePane files={files} sourceCodeHeight={sourceCodeHeight} />
      </div>
      <div
        css={css`
          flex-grow: 1; //necessary for narrower-than-width source code
          max-width: 520px; //necessary for wider-than-width source code
        `}
      >
        <FileContentPane
          fileContent={sourceCode}
          prismLanguage="protobuf"
          sourceCodeHeight={sourceCodeHeight}
        />
      </div>
    </div>
  );
};
