type Props = React.DetailedHTMLProps<
  React.HTMLAttributes<HTMLElement>,
  HTMLElement
>;

export function CustomElementCode(props: Props) {
  const children = props.children;
  console.log(
    `CustomElementCode called for typeof children = ${typeof children}`,
    children
  );
  // if (!children) return <></>;
  // if (!isReactNodeArray(children)) return <></>;
  // console.log(children[0]);
  // console.log(children[0].type);
  return <code>{children}</code>;
}
