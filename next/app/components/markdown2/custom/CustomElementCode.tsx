import React from "react";
import SyntaxHighlighter from "react-syntax-highlighter/dist/esm/default-highlight";
import { a11yDark } from "react-syntax-highlighter/dist/esm/styles/hljs";

function isReactNodeArray(children: any): children is string[] {
  return typeof children == "object" && typeof children[0] === "string";
}

type Props = JSX.IntrinsicElements["code"];

export function CustomElementCode(props: Props) {
  let codeString = "";
  if (isReactNodeArray(props.children)) {
    codeString = props.children[0];
  } else if (typeof props.children === "string") {
    codeString = props.children;
  } else {
    // empty code
    return <code></code>;
  }

  const multiLine = codeString.includes("\n");

  return multiLine ? (
    // language is auto-detected?
    <SyntaxHighlighter style={a11yDark}>{codeString}</SyntaxHighlighter>
  ) : (
    // inline code snippet
    <SyntaxHighlighter
      customStyle={{ display: "inline", padding: "0" }}
      style={a11yDark}
      PreTag={"span"}
    >
      {codeString}
    </SyntaxHighlighter>
  );
}
