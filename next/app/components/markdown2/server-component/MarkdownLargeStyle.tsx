import { MarkdownConfigurable } from "./MarkdownConfigurable";
import styles from "./MarkdownLargeStyle.module.css";
import { CustomElementCode } from "../custom/CustomElementCode";
import { CustomElementPre } from "../custom/CustomElementPre";
import { Components } from "rehype-react";

interface Props {
  markdownBody: string;
}

export function MarkdownLargeStyle(props: Props) {
  // Custom React component mappings
  const components: Partial<Components> = {
    pre: CustomElementPre,
    code: CustomElementCode,
    // a: CustomLink,
  };

  return (
    <MarkdownConfigurable
      customComponents={components}
      className={styles.defaultStyle}
      markdownBody={props.markdownBody}
    />
  );
}
