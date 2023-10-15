import { MarkdownNoStyle } from "./MarkdownNoStyle";
import styles from "./MarkdownDefaultStyle.module.css";

interface Props {
  markdownBody: string;
}

export async function MarkdownDefaultStyle(props: Props) {
  return (
    <div className={styles.defaultStyle}>
      <MarkdownNoStyle markdownBody={props.markdownBody} />
    </div>
  );
}
