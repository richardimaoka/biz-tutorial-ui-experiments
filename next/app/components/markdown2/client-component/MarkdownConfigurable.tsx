"use client";

import React, { useEffect, useState } from "react";

import rehypeReact, { Components } from "rehype-react";
import remarkParse from "remark-parse";
import remarkRehype from "remark-rehype";
import { unified } from "unified";

interface Props {
  markdownBody: string;
  className?: string;
  customComponents?: Partial<Components>;
  onRenderComplete?: () => void;
}

// @ts-expect-error: the react types are missing.
const production = { Fragment: prod.Fragment, jsx: prod.jsx, jsxs: prod.jsxs };

async function MarkdownInner({
  markdownBody,
  customComponents,
}: {
  markdownBody: string;
  customComponents?: Partial<Components>;
}): Promise<JSX.Element> {
  const file = await unified()
    .use(remarkParse)
    .use(remarkRehype)
    .use(rehypeReact, {
      ...production,
      components: customComponents,
    })
    .process(markdownBody);

  return file.result;
}

export function MarkdownConfigurable(props: Props) {
  const [children, setChildren] = useState<JSX.Element>(<></>);

  useEffect(() => {
    if (props.markdownBody) {
      MarkdownInner({ markdownBody: props.markdownBody }).then((result) => {
        setChildren(result);
      });
    }
  }, [props.markdownBody]);

  useEffect(() => {
    if (props.onRenderComplete) {
      props.onRenderComplete();
    }
  }, [children, props, props.onRenderComplete]);

  return (
    <div className={props.className} data-testid="markdown">
      {children}
    </div>
  );
}
