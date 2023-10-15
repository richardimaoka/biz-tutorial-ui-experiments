import { MarkdownDefaultStyle } from "../markdown/MarkdownDefaultStyle";
import styles from "./Tooltip.module.css";
interface Props {
  body: string; // can be markdown
}

export function Tooltip(props: Props) {
  return (
    <div className={styles.component}>
      <div className={styles.tooltip}>
        <MarkdownDefaultStyle markdownBody={props.body} />
      </div>
    </div>
  );
}
