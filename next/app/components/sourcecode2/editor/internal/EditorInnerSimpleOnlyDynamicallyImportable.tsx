"use client";

// !!!!
// Can only be used via Next.js dynamic import with ssr false option,
// due to "monaco-editor" module using browser-side `navigator` inside.
// !!!!

import { EditorBare } from "./EditorBare";
import {
  useEditorInstance,
  useEditorTextUpdate,
  useLanguageUpdate,
} from "./hooks";

interface Props {
  editorText: string;
  language: string;
}

// `default` export, for easier use with Next.js dynamic import
export default function EditorInnerOnlyDynamicallyImportable(props: Props) {
  const [editorInstance, onDidMount] = useEditorInstance();
  useEditorTextUpdate(editorInstance, props.editorText);
  useLanguageUpdate(editorInstance, props.language);

  return <EditorBare onDidMount={onDidMount} />;
}
