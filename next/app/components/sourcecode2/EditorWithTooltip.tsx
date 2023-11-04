import { EditorSimple } from "./editor/EditorSimple";
import styles from "./EditorWithTooltip.module.css";
import { EditorTooltip } from "./tooltip/EditorTooltip";

interface Props {
  editorText: string;
  language: string;
  tooltip?: {
    lineNumber: number;
    markdownBody: string;
    hidden?: boolean;
    offsetContent?: boolean;
  };
}

export function EditorWithTooltip(props: Props) {
  /**
   * If tooltip is passed and not-hidden, then render tooltip
   * <EditorTooltip>, a server-side component needs to be called
   * outiside <EditorSimple>, a
   */
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
      <EditorSimple
        editorText={props.editorText}
        language={props.language}
        tooltip={tooltip}
      />
    </div>
  );
}
