import React, { DetailedHTMLProps, HTMLAttributes } from "react";

import rehypeReact from "rehype-react";
import { ComponentsWithoutNodeOptions } from "rehype-react/lib/complex-types";
import remarkParse from "remark-parse";
import remarkRehype from "remark-rehype";
import { unified } from "unified";

interface Props {
  markdownBody: string;
}

// type CCProps = any;
type CCProps = DetailedHTMLProps<
  HTMLAttributes<HTMLPreElement>,
  HTMLPreElement
>;

const CustomFC: React.FC<CCProps> = (props) => {
  const value = props.children?.valueOf();
  console.log(value);
  return <pre>{props.children}</pre>;
};

export async function MarkdownNoStyle(props: Props) {
  // Custom React component mappings
  const components: ComponentsWithoutNodeOptions["components"] = {
    //              ComponentsWithoutNodeOptions["components"] is a trick to get friendly type error message for `components`.
    // Directly placing this `components` into `use(rehypeReact, {...})` will cause an unfriendly type error,
    // because TypeScript unexpectedly thinks the second argumetn to `use(rehypeReact, {...})` became boolean due to function overload
    pre: CustomFC,
  };

  const processed = await unified()
    .use(remarkParse)
    .use(remarkRehype)
    .use(rehypeReact, {
      createElement: React.createElement,
      Fragment: React.Fragment,
      // To use custom compenents, instead of intrinsic elements like <p>, <h1>, etc.
      // https://github.com/rehypejs/rehype-react#components
      // Each key is a tag name typed in JSX.IntrinsicElements. Each value is either a different tag name, or a component accepting the corresponding props (and an optional node prop if passNode is on).
      components: components,
    })
    .process(props.markdownBody);

  return <div data-testid="markdown">{processed.result}</div>;
}