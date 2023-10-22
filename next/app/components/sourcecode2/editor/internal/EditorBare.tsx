"use client";

import Editor, { OnChange, OnMount } from "@monaco-editor/react";

interface Props {
  // onDidMount: pass-in a callback like below to manipulate editor instance
  //
  //   function handleEditorDidMount(editorInstance: editor.IStandaloneCodeEditor) {
  //     editorRef.current = editorInstance;
  //   }
  onDidMount?: OnMount;

  // onChange: this is also called when the first rendering is finisehd
  onChange?: OnChange;

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
      onChange={props.onChange}
    />
  );
}
