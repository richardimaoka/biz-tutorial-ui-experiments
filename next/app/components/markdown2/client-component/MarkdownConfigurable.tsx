"use client";

import React, { useEffect, useState } from "react";

import rehypeReact from "rehype-react";
import { ComponentsWithoutNodeOptions } from "rehype-react/lib/complex-types";
import remarkParse from "remark-parse";
import remarkRehype from "remark-rehype";
import { unified } from "unified";

interface Props {
  markdownBody: string;
  className?: string;
  customComponents?: ComponentsWithoutNodeOptions["components"];
  onRenderComplete?: () => void;
}

async function MarkdownInner({
  markdownBody,
}: {
  markdownBody: string;
}): Promise<JSX.Element> {
  const file = await unified()
    .use(remarkParse)
    .use(remarkRehype)
    .use(rehypeReact, {
      createElement: React.createElement,
      Fragment: React.Fragment,
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
