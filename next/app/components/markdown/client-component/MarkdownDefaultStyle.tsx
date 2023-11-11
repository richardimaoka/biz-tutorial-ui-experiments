import { MarkdownConfigurable } from "./MarkdownConfigurable";
import styles from "./MarkdownDefaultStyle.module.css";
import { CustomElementCode } from "../custom/CustomElementCode";
import { CustomElementPre } from "../custom/CustomElementPre";
import { Components } from "rehype-react";

interface Props {
  markdownBody: string;
  onRenderComplete?: () => void;
}

export function MarkdownDefaultStyle(props: Props) {
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
      onRenderComplete={props.onRenderComplete}
    />
  );
}
