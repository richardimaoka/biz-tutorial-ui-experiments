import { editor } from "monaco-editor";
import { useState } from "react";

export function useEditorInstance(): [
  editor.IStandaloneCodeEditor | null,
  (editorInstance: editor.IStandaloneCodeEditor) => void
] {
  const [editorInstance, setEditorInstance] =
    useState<editor.IStandaloneCodeEditor | null>(null);

  function handleEditorDidMount(instance: editor.IStandaloneCodeEditor) {
    setEditorInstance(instance);
  }

  return [editorInstance, handleEditorDidMount];
}
