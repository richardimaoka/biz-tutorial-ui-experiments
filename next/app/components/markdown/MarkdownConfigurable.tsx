import React from "react";

import rehypeReact from "rehype-react";
import { ComponentsWithoutNodeOptions } from "rehype-react/lib/complex-types";
import remarkParse from "remark-parse";
import remarkRehype from "remark-rehype";
import { unified } from "unified";

interface Props {
  markdownBody: string;
  className?: string;
  components?: ComponentsWithoutNodeOptions["components"];
}

// type CCProps = any;

export async function MarkdownConfigurable(props: Props) {
  const processed = await unified()
    .use(remarkParse)
    .use(remarkRehype)
    .use(rehypeReact, {
      createElement: React.createElement,
      Fragment: React.Fragment,
      // To use custom compenents, instead of intrinsic elements like <p>, <h1>, etc.
      // https://github.com/rehypejs/rehype-react#components
      // Each key is a tag name typed in JSX.IntrinsicElements. Each value is either a different tag name, or a component accepting the corresponding props (and an optional node prop if passNode is on).
      components: props.components,
    })
    .process(props.markdownBody);

  return (
    <div className={props.className} data-testid="markdown">
      {processed.result}
    </div>
  );
}
