import { EditorEditable } from "./editor/EditorEditable";
import styles from "./EditorWithTooltip.module.css";
import { EditorTooltip } from "./tooltip/EditorTooltip";

interface Props {
  editorText: string;
  language: string;
  tooltip?: {
    markdownBody: string;
    hidden?: boolean;
  };
}

export function EditorWithTooltip(props: Props) {
  const lineHeight = 19;

  return (
    <div className={styles.component}>
      <EditorEditable
        editorText={props.editorText}
        language={props.language}
        lineHeight={19}
      />
      {
        //tooltip is passed and not-hidden, then render tooltip
        props.tooltip && !props.tooltip.hidden && (
          <div className={styles.tooltip} style={{ top: "100px" }}>
            <EditorTooltip markdownBody={props.tooltip?.markdownBody} />
          </div>
        )
      }
    </div>
  );
}
