"use client";

// !!!!
// Can only be used via Next.js dynamic import with ssr false option,
// due to "monaco-editor" module using browser-side `navigator` inside.
// !!!!

import { useEffect } from "react";
import { EditorBare } from "./EditorBare";
import { useEditorInstance } from "./useEditorInstance";
import { editor } from "monaco-editor";

interface Props {
  editorText: string;
  language: string;

  // `edits` are immediately executed by useEffect,
  // so the resulting component = editorText + edits
  edits?: editor.IIdentifiedSingleEditOperation[];
}

export type EditorEditableInnerProps = Props;

// `default` export, for easier use with Next.js dynamic import
export default function EditorEditableOnlyDynamicallyImportable(props: Props) {
  const [editorInstance, onDidMount] = useEditorInstance();

  // update editorText
  useEffect(() => {
    if (editorInstance) {
      editorInstance.setValue(props.editorText);
    }
  }, [editorInstance, props.editorText]);

  // update language
  useEffect(() => {
    const model = editorInstance?.getModel();
    if (model) {
      editor.setModelLanguage(model, props.language);
    }
  }, [editorInstance, props.language]);

  // execute edits
  useEffect(() => {
    if (editorInstance && props.edits) {
      editorInstance.updateOptions({ readOnly: false });
      const result = editorInstance.executeEdits("", props.edits);
      editorInstance.updateOptions({ readOnly: true });
    }
  }, [editorInstance, props.edits]);

  return <EditorBare onDidMount={onDidMount} />;
}
