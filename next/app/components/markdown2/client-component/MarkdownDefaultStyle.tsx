import { MarkdownConfigurable } from "./MarkdownConfigurable";
import styles from "./MarkdownDefaultStyle.module.css";
import { CustomElementCode } from "../custom/CustomElementCode";
import { CustomElementPre } from "../custom/CustomElementPre";
import { ComponentsWithoutNodeOptions } from "rehype-react/lib/complex-types";

interface Props {
  markdownBody: string;
  onRenderComplete?: () => void;
}

export function MarkdownDefaultStyle(props: Props) {
  // Custom React component mappings
  const components: ComponentsWithoutNodeOptions["components"] = {
    //              ComponentsWithoutNodeOptions["components"] is a trick to get friendly type error message for `components`.
    // Directly placing this `components` into `use(rehypeReact, {...})` will cause an unfriendly type error,
    // because TypeScript unexpectedly thinks the second argumetn to `use(rehypeReact, {...})` became boolean due to function overload
    pre: CustomElementPre,
    code: CustomElementCode,
    // a: CustomLink,
  };

  return (
    <MarkdownConfigurable
      customComponents={components}
      className={styles.defaultStyle}
      markdownBody={props.markdownBody}
      onRenderComplete={props.onRenderComplete}
    />
  );
}
