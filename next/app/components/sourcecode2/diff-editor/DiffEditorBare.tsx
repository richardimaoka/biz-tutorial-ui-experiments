"use client";

import { DiffEditor } from "@monaco-editor/react";
import { editor } from "monaco-editor";

interface Props {
  original: string;
  modified: string;
  language: string;
  onDidMount?: (editorInstance: editor.IDiffEditor) => void;
  // pass-in a callback like below to manipulate editor instance
  //
  //   function handleEditorDidMount(editorInstance: editor.IStandaloneCodeEditor) {
  //     editorRef.current = editorInstance;
  //   }
}

export function DiffEditorBare(props: Props) {
  console.log("DiffEditorBare component is re-rendered");

  return (
    <DiffEditor
      original={props.original}
      modified={props.modified}
      language={props.language}
      options={{
        readOnly: true,
        renderSideBySide: false,
        lineNumbers: "off",
        renderOverviewRuler: false,
      }}
      theme={"vs-dark"}
      onMount={props.onDidMount}
    />
  );
}
