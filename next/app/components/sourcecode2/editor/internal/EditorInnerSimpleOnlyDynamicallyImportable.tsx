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
import { ReactNode, useRef } from "react";
import styles from "./EditorInnerOnlyDynamicallyImportable.module.css";
import { useEditorBoundingBox } from "./hooks/useBoundingBox";

interface Props {
  editorText: string;
  language: string;
  editSequence?: {
    edits: editor.IIdentifiedSingleEditOperation[];
    animate?: boolean;
  };
  tooltip?: {
    lineNumber: number;
    children: ReactNode;
  };
}

// `default` export, for easier use with Next.js dynamic import
export default function EditorInnerOnlyDynamicallyImportable(props: Props) {
  const [editorInstance, onDidMount] = useEditorInstance();

  // Basic editor text and its language
  useEditorTextUpdate(editorInstance, props.editorText);
  useLanguageUpdate(editorInstance, props.language);

  // Edits
  useEditSequence(editorInstance, props.editSequence);

  // Tooltip
  const { boundingBoxRef, rect, resizeWindow } = useEditorBoundingBox();
  const { resizeContentWidget } = useTooltip(
    editorInstance,
    rect,
    props.tooltip
  );
  return (
    // Needs the outer <div> for bounding box size retrieval
    <div className={styles.component} ref={boundingBoxRef}>
      <EditorBare onMount={onDidMount} onChange={resizeWindow} />
    </div>
  );
}
