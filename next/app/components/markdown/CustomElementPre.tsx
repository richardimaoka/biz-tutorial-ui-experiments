import { DetailedHTMLProps, HTMLAttributes, ReactElement } from "react";

function isReactNodeArray(children: any): children is ReactElement[] {
  return (
    typeof children == "object" &&
    children.isArray && // isArray() func exists
    children.isArray()
  );
}

function isSingleCodeElement(
  children: any
): children is JSX.IntrinsicElements["code"] {
  return (
    typeof children == "object" &&
    children.isArray && // isArray() func exists
    children.isArray()
  );
}

type Props = DetailedHTMLProps<HTMLAttributes<HTMLPreElement>, HTMLPreElement>;

export function CustomElementPre(props: Props) {
  const children = props.children;
  console.log("CustomPre, called");
  // if (!children) return <></>;
  // if (!isReactNodeArray(children)) return <></>;
  // console.log(children[0]);
  // console.log(children[0].type);
  return <pre>{children}</pre>;
}
