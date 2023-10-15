import { MarkdownNoStyle } from "./MarkdownNoStyle";
import styles from "./MarkdownDefaultStyle.module.css";

interface Props {
  markdownBody: string;
}

export function MarkdownDefaultStyle(props: Props) {
  return (
    <div className={styles.component}>
      <MarkdownNoStyle markdownBody={props.markdownBody} />
    </div>
  );
}
