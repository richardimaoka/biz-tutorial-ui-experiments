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
import { ReactNode, useState } from "react";
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
  /**
   * Monaco editor instance and readiness
   */
  const [editorInstance, onMount] = useEditorInstance();
  // isReady = true, if the initial rendering is finished
  const [isReady, setIsReady] = useState(false);

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
    // if no edits, ok to render tooltip
    !props.editSequence ||
    // if no edits, ok to render tooltip
    props.editSequence.edits.length === 0 ||
    // if there is a tooltip and timing is set "START", then immediately render the tooltip
    props.tooltip?.timing === "START" ||
    // if edits are already completed, then ok to render tooltip
    isEditCompleted;

  const tooltip = props.tooltip
    ? { ...props.tooltip, canRender: canRender }
    : undefined;

  const { boundingBoxRef, resizeWindowCallback } = useTooltip(
    editorInstance,
    tooltip
  );

  return (
    // Needs the outer <div> for bounding box size retrieval.
    <div className={styles.component} ref={boundingBoxRef}>
      {/* The outer <div> has to be separate from the inner <div> because
       ** the inner div has `display: "none"` at the beginning which makes
       ** the bounding box zero-sized.
       */}
      <div
        // Until initial rendering is done (i.e.) onChange is at least called once,
        // delay the display of this display-control component. This is necessary
        // because otherwise the monaco editor moves the carousel unexpectedly by
        // (seemingly) calling scrollIntoView().
        //
        // By setting `display: "none"`, scrollIntoView() will not take effect and
        // the carousel does not move.
        style={{ display: isReady ? "block" : "none" }}
        className={styles.displayControl}
      >
        <EditorBare
          onMount={onMount}
          onChange={() => {
            resizeWindowCallback();
            setIsReady(true);
          }}
        />
      </div>
    </div>
  );
}
