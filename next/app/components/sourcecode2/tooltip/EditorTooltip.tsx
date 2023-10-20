import styles from "./EditorTooltipCC.module.css";
import { MarkdownDefaultStyle } from "../../markdown2/server-component/MarkdownDefaultStyle";

interface Props {
  markdownBody: string;
}

export type EditorTooltipProps = Props;

export function EditorTooltip(props: Props) {
  return (
    <div className={styles.component}>
      <div className={styles.tooltip}>
        <MarkdownDefaultStyle markdownBody={props.markdownBody} />
      </div>
    </div>
  );
}
