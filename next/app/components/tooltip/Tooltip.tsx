import { MarkdownDefaultStyle } from "../markdown2/server-component/MarkdownDefaultStyle";
import styles from "./Tooltip.module.css";
interface Props {
  markdownBody: string; // can be markdown
}

export function Tooltip(props: Props) {
  return (
    <div className={styles.component}>
      <div className={styles.tooltip}>
        <MarkdownDefaultStyle markdownBody={props.markdownBody} />
      </div>
    </div>
  );
}
