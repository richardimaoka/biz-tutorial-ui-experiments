"use client";

// To avoid an error `ReferenceError: navigator is not defined`, dynamic import with ssr false is needed.
// This is because "monaco-editor" module uses browser-side `navigator` inside.
import dynamic from "next/dynamic";
const EditorInnerSimple = dynamic(
  () => import("./internal/EditorInnerSimpleOnlyDynamicallyImportable"),
  {
    ssr: false,
  }
);

interface Props {
  editorText: string;
  language: string;
}

export function Editor(props: Props) {
  return (
    <EditorInnerSimple
      editorText={props.editorText}
      language={props.language}
    />
  );
}
