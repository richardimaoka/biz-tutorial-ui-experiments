"use client";

// To avoid an error `ReferenceError: navigator is not defined`, dynamic import with ssr false is needed.
// This is because "monaco-editor" module uses browser-side `navigator` inside.
import dynamic from "next/dynamic";
const EditorEditableInner = dynamic(
  () => import("./EditorEditableOnlyDynamicallyImportable"),
  {
    ssr: false,
  }
);

// this is ok with static import
import { EditorEditableInnerProps } from "./EditorEditableOnlyDynamicallyImportable";

type Props = EditorEditableInnerProps;

export function EditorEditable(props: Props) {
  return (
    <EditorEditableInner
      editorText={props.editorText}
      language={props.language}
      edits={props.edits}
    />
  );
}
