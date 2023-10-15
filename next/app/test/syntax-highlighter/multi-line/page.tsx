import React from "react";
import SyntaxHighlighter from "react-syntax-highlighter/dist/esm/default-highlight";

import { docco } from "react-syntax-highlighter/dist/esm/styles/hljs";
import * as fs from "fs/promises";

export default async function Page() {
  const codeString = await fs.readFile(
    process.cwd() + "/app/components/column/VisibleColumn.tsx",
    "utf-8"
  );

  return (
    <SyntaxHighlighter language="javascript" style={docco}>
      {codeString}
    </SyntaxHighlighter>
  );
}
