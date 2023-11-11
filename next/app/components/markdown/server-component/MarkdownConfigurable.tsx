import React from "react";

import rehypeReact, { Components } from "rehype-react";
import remarkParse from "remark-parse";
import remarkRehype from "remark-rehype";
import { unified } from "unified";
import * as prod from "react/jsx-runtime";

interface Props {
  markdownBody: string;
  className?: string;
  customComponents?: Partial<Components>;
}

// @ts-expect-error: the react types are missing.
const production = { Fragment: prod.Fragment, jsx: prod.jsx, jsxs: prod.jsxs };

export async function MarkdownConfigurable(props: Props) {
  const processed = await unified()
    .use(remarkParse)
    .use(remarkRehype)
    .use(rehypeReact, {
      ...production,
      components: props.customComponents,
    })
    // ({
    //   // To use custom compenents, instead of intrinsic elements like <p>, <h1>, etc.
    //   // https://github.com/rehypejs/rehype-react#components
    //   // Each key is a tag name typed in JSX.IntrinsicElements. Each value is either a different tag name, or a component accepting the corresponding props (and an optional node prop if passNode is on).
    //   components: props.customComponents,
    // })
    .process(props.markdownBody);

  return (
    <div className={props.className} data-testid="markdown">
      {processed.result}
    </div>
  );
}
