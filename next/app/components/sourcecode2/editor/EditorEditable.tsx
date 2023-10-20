"use client";

// To avoid an error `ReferenceError: navigator is not defined`, dynamic import with ssr false is needed.
// This is because "monaco-editor" module uses browser-side `navigator` inside.
import dynamic from "next/dynamic";
const EditorInner = dynamic(
  () => import("./internal/EditorInnerOnlyDynamicallyImportable"),
  {
    ssr: false,
  }
);

// this is ok with static import
import { editor } from "monaco-editor";
import { ReactNode } from "react";

interface Props {
  editorText: string;
  language: string;

  // `edits` are immediately executed by useEffect,
  // so the resulting component = editorText + edits
  edits?: editor.IIdentifiedSingleEditOperation[];
  tooltip?: {
    lineNumber: number;
    children: ReactNode;
  };
}

export function Editor(props: Props) {
  return (
    <EditorInner
      editorText={props.editorText}
      language={props.language}
      edits={props.edits}
      tooltip={props.tooltip}
    />
  );
}
