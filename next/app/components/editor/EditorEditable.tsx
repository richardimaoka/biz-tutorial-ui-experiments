"use client";

import { useEffect } from "react";
import { EditorBare } from "./EditorBare";
import { useEditorInstance } from "./useEditorInstance";
// import { editor } from "monaco-editor";

interface Props {
  editorText: string;
  language: string;
}

export function EditorEditable(props: Props) {
  const [editorInstance, onDidMount] = useEditorInstance();

  // update editorText
  useEffect(() => {
    if (editorInstance) {
      editorInstance.setValue(props.editorText);
    }
  }, [editorInstance, props.editorText]);

  // // update language
  // useEffect(() => {
  //   const model = editorInstance?.getModel();
  //   if (model) {
  //     editor.setModelLanguage(model, props.language);
  //   }
  // }, [editorInstance, props.language]);

  return <EditorBare onDidMount={onDidMount} />;
}
