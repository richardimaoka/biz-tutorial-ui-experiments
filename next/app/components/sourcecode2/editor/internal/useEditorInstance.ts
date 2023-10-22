import { OnMount } from "@monaco-editor/react";
import { editor } from "monaco-editor";
import { useState } from "react";

/**
 * Custom hook to hold monaco-editor's editor instance
 * @return tuple (multi-value) return
 */

export function useEditorInstance(): [
  // @return Underlying monaco-editor's instance
  editor.IStandaloneCodeEditor | null,

  // You should pass this callback to @monaco-editor's onMount props, to initialize the editor instance.
  // To see how this OnMount type is definied, see the import at the top
  OnMount
] {
  const [editorInstance, setEditorInstance] =
    useState<editor.IStandaloneCodeEditor | null>(null);

  function handleEditorDidMount(instance: editor.IStandaloneCodeEditor) {
    setEditorInstance(instance);
  }

  return [editorInstance, handleEditorDidMount];
}
