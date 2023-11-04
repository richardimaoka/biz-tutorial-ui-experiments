import { editor } from "monaco-editor";
import { useEffect } from "react";

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
