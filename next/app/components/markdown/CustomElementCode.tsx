import React from "react";
import SyntaxHighlighter from "react-syntax-highlighter/dist/esm/default-highlight";
import { a11yDark } from "react-syntax-highlighter/dist/esm/styles/hljs";

function isReactNodeArray(children: any): children is string[] {
  return typeof children == "object" && typeof children[0] === "string";
}

type Props = JSX.IntrinsicElements["code"];

export async function CustomElementCode(props: Props) {
  if (!isReactNodeArray(props.children)) return <code></code>; //empty code
  //const className = props.className;
  const codeString = props.children[0];
  const multiLine = codeString.includes("\n");

  return multiLine ? (
    // language is auto-detected?
    <SyntaxHighlighter style={a11yDark}>{codeString}</SyntaxHighlighter>
  ) : (
    // inline code snippet
    <SyntaxHighlighter
      customStyle={{ display: "inline" }}
      style={a11yDark}
      PreTag={"span"}
    >
      {codeString}
    </SyntaxHighlighter>
  );
}
