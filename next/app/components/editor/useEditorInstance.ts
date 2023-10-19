import { editor } from "monaco-editor";
import { MutableRefObject, useRef } from "react";

export function useEditorInstance(): [
  MutableRefObject<editor.IStandaloneCodeEditor | null>,
  (editorInstance: editor.IStandaloneCodeEditor) => void
] {
  const editorRef = useRef<editor.IStandaloneCodeEditor | null>(null);

  function handleEditorDidMount(editorInstance: editor.IStandaloneCodeEditor) {
    editorRef.current = editorInstance;
  }

  return [editorRef, handleEditorDidMount];
}
