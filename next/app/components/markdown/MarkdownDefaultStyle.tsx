import { MarkdownConfigurable } from "./MarkdownConfigurable";
import styles from "./MarkdownDefaultStyle.module.css";
import { CustomElementCode } from "./CustomElementCode";
import { CustomElementPre } from "./CustomElementPre";
import { ComponentsWithoutNodeOptions } from "rehype-react/lib/complex-types";

interface Props {
  markdownBody: string;
}

export async function MarkdownDefaultStyle(props: Props) {
  // // Custom React component mappings
  const components: ComponentsWithoutNodeOptions["components"] = {
    //              ComponentsWithoutNodeOptions["components"] is a trick to get friendly type error message for `components`.
    // Directly placing this `components` into `use(rehypeReact, {...})` will cause an unfriendly type error,
    // because TypeScript unexpectedly thinks the second argumetn to `use(rehypeReact, {...})` became boolean due to function overload
    pre: CustomElementPre,
    code: CustomElementCode,
    // a: CustomLink,
  };

  return (
    <div className={styles.defaultStyle}>
      <MarkdownConfigurable
        components={components}
        markdownBody={props.markdownBody}
      />
    </div>
  );
}
