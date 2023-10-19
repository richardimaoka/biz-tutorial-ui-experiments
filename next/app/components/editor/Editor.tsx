"use client";
import MonacoEditor from "@monaco-editor/react";
import { editor } from "monaco-editor";
import * as monaco from "monaco-editor/esm/vs/editor/editor.api";
import { useEffect, useRef } from "react";

interface Props {
  editorText: string;
}

export function Editor(props: Props) {
  const editorRef = useRef<editor.IStandaloneCodeEditor | null>(null);
  const current = editorRef.current;
  console.log("Editor component is re-rendered");

  const contentWidget = {
    getId: function () {
      return "my.content.widget";
    },
    getDomNode: function (): HTMLElement {
      const domNode = document.createElement("div");
      domNode.innerHTML = "My content widget";
      domNode.style.background = "grey";
      domNode.style.opacity = "0.6";
      domNode.style.width = "500px";
      domNode.style.height = `${50}px`;
      domNode.hidden = true;
      return domNode;
    },
    getPosition: function (): monaco.editor.IContentWidgetPosition {
      return {
        position: {
          lineNumber: 1,
          column: 20,
        },
        preference: [monaco.editor.ContentWidgetPositionPreference.EXACT],
      };
    },
    // allowEditorOverflow: true,
  };

  useEffect(() => {
    console.log("editorText useEffect", editorRef.current);
    current?.setValue(props.editorText);
  }, [props.editorText, current]);

  useEffect(() => {
    editorRef.current?.addContentWidget(contentWidget);
  });

  function handleEditorDidMount(editorInstance: editor.IStandaloneCodeEditor) {
    // here is the editor instance
    // you can store it in `useRef` for further usage
    editorRef.current = editorInstance;
    console.log("editorRef.current", editorRef.current);
  }

  return (
    <MonacoEditor
      defaultValue={props.editorText}
      options={{
        readOnly: true, // also tried with readOnly: true
      }}
      onMount={handleEditorDidMount}
    />
  );
}
