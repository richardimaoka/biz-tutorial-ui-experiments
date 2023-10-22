"use client";

// !!!!
// Can only be used via Next.js dynamic import with ssr false option,
// due to "monaco-editor" module using browser-side `navigator` inside.
// !!!!
import { editor } from "monaco-editor";
import { EditorBare } from "./EditorBare";
import { useEditSequence } from "./edits/useEditSequence";
import { useEditorInstance } from "./useEditorInstance";
import { useEditorTextUpdate } from "./useEditorTextUpdate";
import { useLanguageUpdate } from "./useLanguageUpdate";

interface Props {
  editorText: string;
  language: string;

  editSequence?: {
    edits: editor.IIdentifiedSingleEditOperation[];
    animate?: boolean;
  };
}

// `default` export, for easier use with Next.js dynamic import
export default function EditorInnerOnlyDynamicallyImportable(props: Props) {
  const [editorInstance, onDidMount] = useEditorInstance();

  useEditorTextUpdate(editorInstance, props.editorText);

  useLanguageUpdate(editorInstance, props.language);

  useEditSequence(editorInstance, props.editSequence);

  return <EditorBare onDidMount={onDidMount} />;
}
