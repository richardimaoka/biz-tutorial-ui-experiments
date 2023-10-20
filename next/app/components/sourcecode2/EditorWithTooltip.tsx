import { EditorEditable } from "./editor/EditorEditable";
import styles from "./EditorWithTooltip.module.css";
import { EditorTooltip } from "./tooltip/EditorTooltip";

interface Props {
  editorText: string;
  language: string;
  tooltip?: {
    lineNumber: number;
    markdownBody: string;
    hidden?: boolean;
  };
}

export function EditorWithTooltip(props: Props) {
  const lineHeight = 19;

  //tooltip is passed and not-hidden, then render tooltip
  const tooltip =
    props.tooltip && !props.tooltip.hidden
      ? {
          lineNumber: props.tooltip.lineNumber,
          children: (
            <EditorTooltip markdownBody={props.tooltip?.markdownBody} />
          ),
        }
      : undefined;

  return (
    <div className={styles.component}>
      <EditorEditable
        editorText={props.editorText}
        language={props.language}
        lineHeight={19}
        tooltip={tooltip}
      />
    </div>
  );
}
