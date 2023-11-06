import styles from "./SourceCodeEditor.module.css";
import { EditorTooltip } from "../tooltip/EditorTooltip";

// To avoid an error `ReferenceError: navigator is not defined`, dynamic import with ssr false is needed.
// This is because "monaco-editor" module uses browser-side `navigator` inside.
import dynamic from "next/dynamic";
const EditorInner = dynamic(
  () => import("./internal/EditorInnerOnlyDynamicallyImportable"),
  {
    ssr: false,
  }
);
import { editor } from "monaco-editor";

export type EditOperation = editor.IIdentifiedSingleEditOperation;

interface Props {
  editorText: string;
  language: string;
  editSequence?: {
    edits: EditOperation[];
    skipAnimation?: boolean;
  };
  tooltip?: {
    lineNumber: number;
    markdownBody: string;
    timing: "START" | "END";
    hidden?: boolean;
    offsetContent?: boolean;
  };
}

/**
 * Source Code Editor component, which is purely based on React (i.e.) non-GraphQL component.
 * This component serves two purposes:
 *   - set clean props to control the editor behavior, exposed to GraphQL components
 *   - call <EditorTooltip> **server** component and pass it to <EditorInner> **client** component
 *       if you call <EditorTooltip> from inside <EditorInner>, that will cause a runtime error saying
 *       async component (tooltip uses async/await for remark) cannot be called from client component
 */
export function SourceCodeEditor(props: Props) {
  /**
   * If tooltip is passed and not-hidden, then render tooltip
   * <EditorTooltip>, a server-side component needs to be called
   * outiside <EditorInner>
   */
  const tooltip =
    props.tooltip && !props.tooltip.hidden
      ? {
          lineNumber: props.tooltip.lineNumber,
          children: (
            // Convert the markdownBody string into a React component
            <EditorTooltip markdownBody={props.tooltip?.markdownBody} />
          ),
          timing: props.tooltip.timing,
        }
      : undefined;

  return (
    <div className={styles.component}>
      <EditorInner
        editorText={props.editorText}
        language={props.language}
        editSequence={props.editSequence}
        tooltip={tooltip}
      />
    </div>
  );
}
