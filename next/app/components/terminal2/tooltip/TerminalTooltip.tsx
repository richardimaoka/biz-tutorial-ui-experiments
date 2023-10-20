import { MarkdownDefaultStyle } from "../../markdown2/server-component/MarkdownDefaultStyle";
import styles from "./TerminalTooltip.module.css";

interface Props {
  markdownBody: string;
  hidden?: boolean;
}

export type TerminalTooltipProps = Props;

export function TerminalTooltip(props: Props) {
  return (
    <div className={styles.component}>
      <div className={styles.tooltip}>
        <MarkdownDefaultStyle markdownBody={props.markdownBody} />
      </div>
    </div>
  );
}
