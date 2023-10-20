"use client";
import Editor from "@monaco-editor/react";
import { editor } from "monaco-editor";
import { useState } from "react";
import { createRoot } from "react-dom/client";

interface Props {
  onDidMount?: (editorInstance: editor.IStandaloneCodeEditor) => void;
  // pass-in a callback like below to manipulate editor instance
  //
  //   function handleEditorDidMount(editorInstance: editor.IStandaloneCodeEditor) {
  //     editorRef.current = editorInstance;
  //   }
  lineHeight?: number;
}

export function EditorBare(props: Props) {
  console.log("EditorBare component is rendered");
  const lineHeight = props.lineHeight ? props.lineHeight : 19;

  return (
    <Editor
      options={{
        readOnly: true,
        theme: "vs-dark",

        // save width for mobile - folding displays small vertical bar
        folding: false,
        lineNumbers: "off",
        minimap: {
          enabled: false,
        },
        // scroll bar visible, as "hidden" doesn't completely hide it, and meks it just awkward
        // scrollbar: {
        //   vertical: "visible",
        //   horizontal: "visible",
        // },

        lineHeight: lineHeight,
      }}
      onMount={props.onDidMount}
    />
  );
}
