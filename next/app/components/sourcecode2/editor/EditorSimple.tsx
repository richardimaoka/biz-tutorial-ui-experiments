"use client";

// To avoid an error `ReferenceError: navigator is not defined`, dynamic import with ssr false is needed.
// This is because "monaco-editor" module uses browser-side `navigator` inside.
import dynamic from "next/dynamic";
const EditorInnerSimple = dynamic(
  () => import("./internal/EditorInnerOnlyDynamicallyImportable"),
  {
    ssr: false,
  }
);
import { editor } from "monaco-editor";
import { ReactNode } from "react";

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

export function EditorSimple(props: Props) {
  return (
    <EditorInnerSimple
      editorText={props.editorText}
      language={props.language}
      editSequence={props.editSequence}
      tooltip={props.tooltip}
    />
  );
}
