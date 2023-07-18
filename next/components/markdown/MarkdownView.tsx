import rehypeReact from "rehype-react";
import remarkParse from "remark-parse";
import remarkRehype from "remark-rehype";
import { unified } from "unified";

import { css } from "@emotion/react";
import {
  dark1MainBg,
  dark5,
  gray,
  themeBlue,
  white,
} from "../../libs/colorTheme";
import { FragmentType, graphql, useFragment } from "../../libs/gql";
import React, { ComponentType, ReactNode, useEffect, useState } from "react";

interface PProps {
  children?: ReactNode; //children needs to be optional, otherwise, type error in rehype-react's components argument
}

const MarkdownFragment = graphql(`
  fragment MarkdownFragment on Markdown {
    contents
    alignment
  }
`);

export interface MarkdownViewProps {
  fragment: FragmentType<typeof MarkdownFragment>;
}

export const MarkdownView = (props: MarkdownViewProps): JSX.Element => {
  const fragment = useFragment(MarkdownFragment, props.fragment);
  const [mdElem, setMdElem] = useState<JSX.Element | null>(null);
  const textAlign = fragment.alignment
    ? fragment.alignment.toLowerCase()
    : "left";

  useEffect(() => {
    if (fragment.contents || fragment.contents == "") {
      unified()
        .use(remarkParse)
        .use(remarkRehype)
        .use(rehypeReact, {
          createElement: React.createElement,
          Fragment: React.Fragment,
        })
        .process(fragment.contents)
        .then((file) => {
          setMdElem(file.result as JSX.Element);
        });
    }
  }, [fragment.contents]);

  const markdownCss = css`
    background-color: ${dark1MainBg};

    h1 {
      font-size: 32px;
      font-weight: 700;
      margin: 21px 0px;
      color: ${white};
    }

    h2 {
      font-size: 24px;
      font-weight: 700;
      margin: 19px 0px;
      color: ${white};
    }

    h3 {
      font-size: 18px;
      font-weight: 700;
      margin: 18px 0px;
      color: ${white};
    }

    p {
      font-size: 14px;
      margin: 16px 0px;
      color: ${white};
      text-align: ${textAlign};
    }

    ul {
      margin: 16px 0px;
      color: ${white};
    }

    ol,
    ul {
      margin: 16px 0px;
      padding-left: 24px;
      color: ${white};
    }

    li {
      font-size: 14px;
      color: ${white};
    }

    code {
      font-family: "Roboto Mono";
      font-weight: 500;
      font-size: 14px;
      line-height: 18px;
      padding: 1px 4px;
      background-color: ${dark5};
    }

    pre {
      margin: 16px 0px;
      padding: 8px;
      background-color: ${dark5};
    }

    pre > code {
      font-size: 14px;
      line-height: 20px;
    }

    blockquote {
      margin: 16px 0px;
      margin: 0px;
      padding-left: 18px;
      border-left: 2px solid ${themeBlue};
      color: ${gray};
    }

    img {
      margin: 16px 0px;
      display: block;
      margin: 0 auto;
      max-width: 100%;
    }

    iframe {
      display: block;
      margin: 0 auto;
      max-width: 100%;
    }
  `;

  return (
    <div
      css={css`
        background-color: ${dark1MainBg};
        padding: 2px 10px;
      `}
    >
      <div
        // options.components (https://github.com/rehypejs/rehype-react#optionscomponents) to avoid the following Next Lint error:
        //   > Using `<img>` could result in slower LCP and higher bandwidth. Use `<Image />` from `next/image` instead to utilize Image Optimization.
        //   > See: https://nextjs.org/docs/messages/no-img-elementeslint@next/next/no-img-element
        css={markdownCss}
      >
        {mdElem}
      </div>
    </div>
  );
};
