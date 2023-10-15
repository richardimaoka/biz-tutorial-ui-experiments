import React from "react";
import SyntaxHighlighter from "react-syntax-highlighter/dist/esm/default-highlight";

import { docco } from "react-syntax-highlighter/dist/esm/styles/hljs";

export default function Page() {
  const codeString = `(num) => num + 1
(num) => num + 1
(num) => num + 1
(num) => num + 1
`;
  return (
    <div>
      <SyntaxHighlighter
        language="javascript"
        style={docco}
        PreTag={React.Fragment}
      >
        {codeString}
      </SyntaxHighlighter>
      <div>add add add</div>
    </div>
  );
}
