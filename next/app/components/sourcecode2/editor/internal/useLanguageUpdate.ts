import { editor } from "monaco-editor";
import { useEffect } from "react";

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
    if (editorInstance) {
      const model = editorInstance.getModel();
      if (model) {
        editor.setModelLanguage(model, language);
      }
    }
  }, [editorInstance, language]);
}
