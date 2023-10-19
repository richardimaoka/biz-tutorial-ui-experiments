"use client";
import Editor from "@monaco-editor/react";
import { editor } from "monaco-editor";

interface Props {
  onDidMount?: (editorInstance: editor.IStandaloneCodeEditor) => void;
  // pass-in a callback like below to manipulate editor instance
  //
  //   function handleEditorDidMount(editorInstance: editor.IStandaloneCodeEditor) {
  //     editorRef.current = editorInstance;
  //   }
}

export function EditorBare(props: Props) {
  console.log("EditorBare component is rendered");

  return (
    <Editor
      options={{
        readOnly: true,
      }}
      onMount={props.onDidMount}
    />
  );
}
