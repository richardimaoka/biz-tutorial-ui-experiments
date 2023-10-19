"use client";
import MonacoEditor from "@monaco-editor/react";
import { editor } from "monaco-editor";
import * as monaco from "monaco-editor/esm/vs/editor/editor.api";
import { useEffect, useRef } from "react";
import styles from "./Editor.module.css";
type Monaco = typeof monaco;

interface Props {
  name: string;
}

export function Editor(props: Props) {
  const editorRef = useRef<editor.IStandaloneCodeEditor | null>(null);
  console.log("Editor component is re-rendered, name=", props.name);

  useEffect(() => {
    console.log("editorRef.current?.getModel()", editorRef.current?.getModel());
    editorRef.current?.setValue(props.name);
  }, [props.name]);

  function handleEditorDidMount(
    editorInstance: editor.IStandaloneCodeEditor,
    monaco: Monaco
  ) {
    // here is the editor instance
    // you can store it in `useRef` for further usage
    editorRef.current = editorInstance;
    console.log("editorRef.current", editorRef.current);
  }

  return (
    <div className={styles.component}>
      <MonacoEditor
        height="90vh"
        defaultLanguage="javascript"
        defaultValue="console.log('A')"
        options={{
          readOnly: true, // also tried with readOnly: true
        }}
        onMount={handleEditorDidMount}
      />
    </div>
  );
}
