type Props = JSX.IntrinsicElements["pre"];

export function CustomElementPre(props: Props) {
  const children = props.children;

  // since <pre> is added by react-syntax-highlight in CustomCodeElement,
  // do not add outer <pre></pre> here
  return <>{children}</>;
}
