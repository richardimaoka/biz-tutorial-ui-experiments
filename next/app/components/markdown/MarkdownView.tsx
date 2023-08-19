"use client";
import React, { ReactNode, useEffect, useState } from "react";

import rehypeReact from "rehype-react";
import remarkParse from "remark-parse";
import remarkRehype from "remark-rehype";
import { unified } from "unified";

import styles from "./style.module.css";

import { FragmentType, graphql, useFragment } from "@/libs/gql";

const fragmentDefinition = graphql(`
  fragment MarkdownFragment on Markdown {
    contents
    alignment
  }
`);

const Markdown = async ({ input }: { input: string }): Promise<JSX.Element> => {
  const file = await unified()
    .use(remarkParse)
    .use(remarkRehype)
    .use(rehypeReact, {
      createElement: React.createElement,
      Fragment: React.Fragment,
    })
    .process(input);

  return file.result;
};

export interface MarkdownViewProps {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export const MarkdownView = /*async*/ (
  props: MarkdownViewProps
): JSX.Element => {
  // eslint-disable-next-line react-hooks/rules-of-hooks
  const fragment = useFragment(fragmentDefinition, props.fragment);

  const [children, setChildren] = useState<JSX.Element>(<></>);

  useEffect(() => {
    if (fragment.contents) {
      Markdown({ input: fragment.contents }).then((result) => {
        setChildren(result);
      });
    }
  }, [fragment.contents]);

  const alignmentStyle =
    fragment.alignment === "CENTER" ? styles.center : styles.left; //default alignment = styles.left

  if (!fragment.contents) {
    return <div></div>;
  }

  return (
    <div className={`${styles.markdown} ${alignmentStyle}`}>
      {children /* <Markdown input={fragment.contents} /> */}
    </div>
  );
};
