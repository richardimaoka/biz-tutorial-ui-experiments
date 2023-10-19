"use client";

import { DiffEditor } from "@monaco-editor/react";
import { editor } from "monaco-editor";

interface Props {
  onDidMount?: (editorInstance: editor.IDiffEditor) => void;
  // pass-in a callback like below to manipulate editor instance
  //
  //   function handleEditorDidMount(editorInstance: editor.IStandaloneCodeEditor) {
  //     editorRef.current = editorInstance;
  //   }
}

export function EditorBare(props: Props) {
  console.log("EditorBare component is re-rendered");

  return (
    <DiffEditor
      options={{
        readOnly: true,
      }}
      onMount={props.onDidMount}
    />
  );
}
