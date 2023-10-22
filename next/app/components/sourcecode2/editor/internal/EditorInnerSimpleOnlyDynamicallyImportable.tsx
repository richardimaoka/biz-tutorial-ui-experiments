"use client";

// !!!!
// Can only be used via Next.js dynamic import with ssr false option,
// due to "monaco-editor" module using browser-side `navigator` inside.
// !!!!
import { editor } from "monaco-editor";
import { EditorBare } from "./EditorBare";
import {
  useEditorInstance,
  useEditorTextUpdate,
  useLanguageUpdate,
} from "./hooks";
import { useEffect, useRef } from "react";

interface Props {
  editorText: string;
  language: string;

  edits?: editor.IIdentifiedSingleEditOperation[];
}

// `default` export, for easier use with Next.js dynamic import
export default function EditorInnerOnlyDynamicallyImportable(props: Props) {
  const [editorInstance, onDidMount] = useEditorInstance();
  useEditorTextUpdate(editorInstance, props.editorText);
  useLanguageUpdate(editorInstance, props.language);

  const isEditsMade = useRef(false);

  // Execute edits
  useEffect(() => {
    if (editorInstance) {
      if (props.edits && props.edits.length > 0) {
        editorInstance.updateOptions({ readOnly: false });

        // clear previous edits upon props change
        if (isEditsMade) {
          editorInstance.trigger("", "undo", null);
        }

        const result = editorInstance.executeEdits("", props.edits);
        if (!result) {
          // TODO: throw error to trigger error.tsx
          console.log("executeEdits for monaco editor failed!");
        }
        isEditsMade.current = true;

        editorInstance.updateOptions({ readOnly: true });
      } else {
        // clear edits if edits is undefined or []
        editorInstance.updateOptions({ readOnly: false });

        editorInstance.trigger("", "undo", null);
        isEditsMade.current = false;

        editorInstance.updateOptions({ readOnly: true });
      }
    }
  }, [editorInstance, props.edits]);

  return <EditorBare onDidMount={onDidMount} />;
}
