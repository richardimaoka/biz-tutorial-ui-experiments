import React from "react";

import rehypeReact from "rehype-react";
import remarkParse from "remark-parse";
import remarkRehype from "remark-rehype";
import { unified } from "unified";

interface Props {
  markdownBody: string;
}

export async function MarkdownComponent(props: Props) {
  const file = await unified()
    .use(remarkParse)
    .use(remarkRehype)
    .use(rehypeReact, {
      createElement: React.createElement,
      Fragment: React.Fragment,
    })
    .process(props.markdownBody);

  return <div data-testid="markdown">{props.markdownBody}</div>;
}
