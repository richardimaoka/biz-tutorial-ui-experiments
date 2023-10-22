import { OnMount } from "@monaco-editor/react";
import { editor } from "monaco-editor";
import { useEffect, useState } from "react";

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

/**
 * Custm hook to handle @param `editorText` update
 *
 * @return nothing, as it is an effectful hook
 */
export function useEditorTextUpdate(
  // Shared monaco-editor's editor instance, possibly null
  editorInstance: editor.IStandaloneCodeEditor | null,

  // Prop to be passed from the parent component
  editorText: string
) {
  useEffect(() => {
    if (editorInstance) {
      editorInstance.setValue(editorText);
    }
  }, [editorInstance, editorText]);
}

/**
 * Custm hook to handle @param `language` update
 *
 * @return nothing, as it is an effectful hook
 */
export function useLanguageUpdate(
  // Shared monaco-editor's editor instance, possibly null
  editorInstance: editor.IStandaloneCodeEditor | null,

  // Prop to be passed from the parent component
  language: string
) {
  useEffect(() => {
    const model = editorInstance?.getModel();
    if (model) {
      editor.setModelLanguage(model, language);
    }
  }, [editorInstance, language]);
}
