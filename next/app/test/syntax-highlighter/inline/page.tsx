import React from "react";
import SyntaxHighlighter from "react-syntax-highlighter/dist/esm/default-highlight";

import { docco } from "react-syntax-highlighter/dist/esm/styles/hljs";

export default function Page() {
  const codeString = `markdown-it`;
  return (
    <p style={{ backgroundColor: "white" }}>
      The killer feature of
      <SyntaxHighlighter
        language="javascript"
        style={docco}
        PreTag={React.Fragment}
      >
        {codeString}
      </SyntaxHighlighter>
      is very effective support of
    </p>
  );
}
