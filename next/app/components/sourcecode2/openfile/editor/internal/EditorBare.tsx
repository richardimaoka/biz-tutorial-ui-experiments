"use client";

import Editor, { OnChange, OnMount } from "@monaco-editor/react";

interface Props {
  // onMount: pass-in a callback like below to manipulate editor instance
  //
  //  function handleEditorDidMount(editorInstance: editor.IStandaloneCodeEditor) {
  //    editorRef.current = editorInstance;
  //  }
  //
  //  https://github.com/suren-atoyan/monaco-react#props
  //  An event is emitted when the editor is mounted. It gets the editor instance as a first argument and the monaco instance as a second
  onMount?: OnMount;

  // onChange: this is called on model change, as well as when first rendering is finisehd
  //   https://github.com/suren-atoyan/monaco-react#props
  //   An event is emitted when the content of the current model is changed
  onChange?: OnChange;

  lineHeight?: number;
}
/**
 * This component simply calls <Editor> from @monaco-editor/react
 * with preferred settings passed to <Editor> as props.
 *
 * It doesn't hold logic nor any React hook for simplicity and separation
 * of concern - those should be implemented in the caller of this component
 */
export function EditorBare(props: Props) {
  // console.log("EditorBare component is rendered");
  const lineHeight = props.lineHeight ? props.lineHeight : 19;

  return (
    <Editor
      options={{
        readOnly: true,
        domReadOnly: true,
        theme: "vs-dark",

        // save width for mobile - folding displays small vertical bar
        folding: false,
        // lineNumbers: "off",
        // lineNumbers: (n: number) => `${n}`.slice(-1),
        lineNumbersMinChars: 0,

        // lineDecorationsWidth: "10px",
        minimap: {
          enabled: false,
        },
        wordWrap: "on",
        // scroll bar visible, as "hidden" doesn't completely hide it, and meks it just awkward
        // scrollbar: {
        //   vertical: "visible",
        //   horizontal: "visible",
        // },

        lineHeight: lineHeight,
      }}
      onMount={props.onMount}
      onChange={props.onChange}
    />
  );
}
