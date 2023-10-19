"use client";

import { useEffect } from "react";
import { EditorBare } from "./EditorBare";
import { useEditorInstance } from "./useEditorInstance";

interface Props {
  editorText: string;
  launguage: string;
}

export function EditorEditable(props: Props) {
  const [editorInstance, onDidMount] = useEditorInstance();

  useEffect(() => {
    console.log("EditorEditable useEffect");
    if (editorInstance) {
      editorInstance.setValue(props.editorText);
    }
  }, [editorInstance, props.editorText]);

  return <EditorBare onDidMount={onDidMount} />;
}
