"use client";

// !!!!
// Can only be used via Next.js dynamic import with ssr false option,
// due to "monaco-editor" module using browser-side `navigator` inside.
// !!!!
import { editor } from "monaco-editor";
import { EditorBare } from "./EditorBare";
import { useEditSequence } from "./hooks/useEditSequence";
import { useEditorInstance } from "./hooks/useEditorInstance";
import { useEditorTextUpdate } from "./hooks/useEditorTextUpdate";
import { useLanguageUpdate } from "./hooks/useLanguageUpdate";
import { useTooltip } from "./hooks/useTooltip";
import { ReactNode } from "react";
import styles from "./EditorInnerOnlyDynamicallyImportable.module.css";

interface Props {
  editorText: string;
  language: string;
  editSequence?: {
    edits: editor.IIdentifiedSingleEditOperation[];
    skipAnimation?: boolean;
  };
  tooltip?: {
    lineNumber: number;
    children: ReactNode;
    timing: "START" | "END";
  };
}

// `default` export, for easier use with Next.js dynamic import
export default function EditorInnerOnlyDynamicallyImportable(props: Props) {
  const [editorInstance, onDidMount] = useEditorInstance();

  /**
   * Basic editor text and its language
   */
  useEditorTextUpdate(editorInstance, props.editorText);
  useLanguageUpdate(editorInstance, props.language);

  /**
   * Edits
   */
  const { isEditCompleted } = useEditSequence(
    editorInstance,
    props.editSequence
  );

  /**
   * Tooltip
   */

  const canRender =
    // if edits are already completed, then ok to render tooltip
    isEditCompleted ||
    // if no edits, ok to render tooltip
    !props.editSequence ||
    props.editSequence.edits.length === 0 ||
    // if tooltip timing is set "START", then immediately render the tooltip
    props.tooltip?.timing === "START";

  const tooltip = props.tooltip
    ? { ...props.tooltip, canRender: canRender }
    : undefined;
  const { boundingBoxRef, resizeWindowCallback } = useTooltip(
    editorInstance,
    tooltip
  );

  return (
    // Needs the outer <div> for bounding box size retrieval
    <div className={styles.component} ref={boundingBoxRef}>
      <EditorBare onMount={onDidMount} onChange={resizeWindowCallback} />
    </div>
  );
}
