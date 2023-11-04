import styles from "./EditorWithTooltip.module.css";
import { EditorTooltip } from "./tooltip/EditorTooltip";

// To avoid an error `ReferenceError: navigator is not defined`, dynamic import with ssr false is needed.
// This is because "monaco-editor" module uses browser-side `navigator` inside.
import dynamic from "next/dynamic";
const EditorInnerSimple = dynamic(
  () => import("./editor/internal/EditorInnerOnlyDynamicallyImportable"),
  {
    ssr: false,
  }
);

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
      <EditorInnerSimple
        editorText={props.editorText}
        language={props.language}
        tooltip={tooltip}
      />
    </div>
  );
}
