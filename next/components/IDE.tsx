import { css } from "@emotion/react";
import { useEffect, useRef } from "react";
import { IDEEditorTab } from "./IDEEditorTab";
import { FileNameComponent } from "./IDE/filetree/FileNameComponent";

import Prism from "prismjs";
import "prismjs/components/prism-protobuf"; //ts 7016 error suppressed by prism-fix.d.ts
import { FileNameTabBar } from "./sourcecode/file-content/FileNameTabBar";
import { FileContentViewer } from "./sourcecode/file-content/FileContentViewer";
import { FileContentPane } from "./sourcecode/file-content/FileContentPane";
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

const files = [
  { offset: 0, __typename: "directory", filename: "next" },
  { offset: 1, __typename: "directory", filename: "cache" },
  { offset: 1, __typename: "directory", filename: "server" },
  { offset: 1, __typename: "directory", filename: "static" },
  { offset: 1, __typename: "file", filename: "build-manifest.json" },
  { offset: 1, __typename: "file", filename: "package.json" },
  {
    offset: 1,
    __typename: "file",
    filename: "react-loadable-manifest.json",
  },
  { offset: 1, __typename: "directory", filename: "trace" },
  { offset: 0, __typename: "directory", filename: "components" },
  { offset: 1, __typename: "file", filename: "Terminal.tsx" },
  { offset: 0, __typename: "directory", filename: "node_module" },
  { offset: 0, __typename: "directory", filename: "pages" },
  { offset: 1, __typename: "directory", filename: "api" },
  { offset: 2, __typename: "file", filename: "hello.ts" },
  { offset: 1, __typename: "file", filename: "_app.tsx" },
  { offset: 1, __typename: "file", filename: "_document.tsx" },
  { offset: 1, __typename: "file", filename: "index.tsx" },
  { offset: 0, __typename: "directory", filename: "public" },
  { offset: 0, __typename: "directory", filename: "styles" },
  { offset: 1, __typename: "file", filename: "global.css" },
  { offset: 1, __typename: "file", filename: "Home.module.css" },
  { offset: 1, __typename: "directory", filename: ".babelrc" },
  { offset: 0, __typename: "file", filename: ".eslintrc.json" },
  { offset: 0, __typename: "directory", filename: ".gitignore" },
  { offset: 0, __typename: "file", filename: "next-env.d.ts" },
  { offset: 0, __typename: "file", filename: "next.config.js" },
  { offset: 0, __typename: "file", filename: "package-lock.json" },
  { offset: 0, __typename: "file", filename: "package.json" },
  { offset: 0, __typename: "file", filename: "README.md" },
  { offset: 0, __typename: "file", filename: "tsconfig.json" },
];

const sourceCodeHeight = 400;

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
          flex-grow: 1;
        `}
      >
        <div
          css={css`
            display: flex;
            padding: 6px 10px 6px 6px;
            justify-content: end;
            background-color: #222121;
          `}
        >
          <img
            width="16"
            height="16"
            css={css`
              display: block;
              background-color: #f7f7f7;
              border-radius: 2px;
            `}
            src="/images/ide-sidebar-shrink.svg"
          />
        </div>
        <div
          css={css`
            height: ${sourceCodeHeight}px;
            max-width: 160px;
            overflow: scroll;
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
          <div
            css={css`
              width: fit-content;
              min-width: 100%;
              min-height: 100%; //expand up to the outer element
              background-color: #252526;
            `}
          >
            {files.map((elem) => (
              <FileNameComponent
                key={elem.filename}
                type={elem.__typename}
                offset={elem.offset}
                name={elem.filename}
              />
            ))}
          </div>
        </div>
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
