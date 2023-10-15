import React from "react";

import rehypeReact from "rehype-react";
import remarkParse from "remark-parse";
import remarkRehype from "remark-rehype";
import { unified } from "unified";

interface Props {
  markdownBody: string;
}

export async function MarkdownNoStyle(props: Props) {
  const processed = await unified()
    .use(remarkParse)
    .use(remarkRehype)
    .use(rehypeReact, {
      createElement: React.createElement,
      Fragment: React.Fragment,
      // To use custom compenents, instead of intrinsic elements like <p>, <h1>, etc.
      // https://github.com/rehypejs/rehype-react#components
      // Each key is a tag name typed in JSX.IntrinsicElements. Each value is either a different tag name, or a component accepting the corresponding props (and an optional node prop if passNode is on).
      // components: {}
    })
    .process(props.markdownBody);

  return <div data-testid="markdown">{processed.result}</div>;
}
