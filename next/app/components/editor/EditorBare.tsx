"use client";
import MonacoEditor from "@monaco-editor/react";
import { editor } from "monaco-editor";
import { useRef } from "react";

export function EditorBare() {
  console.log("EditorBare component is re-rendered");
  const editorRef = useRef<editor.IStandaloneCodeEditor | null>(null);

  function handleEditorDidMount(editorInstance: editor.IStandaloneCodeEditor) {
    // here is the editor instance you can store it in `useRef` for further usage
    editorRef.current = editorInstance;
  }

  return (
    <MonacoEditor
      options={{
        readOnly: true, // also tried with readOnly: true
      }}
      onMount={handleEditorDidMount}
    />
  );
}
