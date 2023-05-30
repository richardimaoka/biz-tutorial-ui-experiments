import { css } from "@emotion/react";
import { dark1MainBg, dark5, gray, themeBlue } from "../../libs/colorTheme";

import rehypeReact from "rehype-react";
import remarkParse from "remark-parse";
import remarkRehype from "remark-rehype";
import { unified } from "unified";
import { Fragment, createElement, useEffect, useState } from "react";
import { FragmentType, graphql, useFragment } from "../../libs/gql";

const MarkdownPane_Fragment = graphql(`
  fragment MarkdownPane_Fragment on Markdown {
    contents
  }
`);

export interface MarkdownPaneProps {
  fragment: FragmentType<typeof MarkdownPane_Fragment>;
}

export const MarkdownPane = (props: MarkdownPaneProps): JSX.Element => {
  const fragment = useFragment(MarkdownPane_Fragment, props.fragment);
  const [mdElem, setMdElem] = useState<JSX.Element | null>(null);

  useEffect(() => {
    if (fragment.contents || fragment.contents == "") {
      unified()
        .use(remarkParse)
        .use(remarkRehype)
        .use(rehypeReact, { createElement, Fragment })
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
    }

    h2 {
      font-size: 24px;
      font-weight: 700;
      margin: 19px 0px;
    }

    h3 {
      font-size: 18px;
      font-weight: 700;
      margin: 18px 0px;
    }

    p {
      font-size: 14px;
      margin: 16px 0px;
    }

    ul {
      margin: 16px 0px;
    }

    ul > li {
      font-size: 14px;
      margin: 4px 0px;
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
