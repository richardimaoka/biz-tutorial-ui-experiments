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

  editSequence?: {
    edits: editor.IIdentifiedSingleEditOperation[];
    animate?: boolean;
  };
}

function executeEditCallback(
  editorInstance: editor.IStandaloneCodeEditor,
  editCallback: () => void
) {
  try {
    editorInstance.updateOptions({ readOnly: false });
    editCallback();
  } finally {
    editorInstance.updateOptions({ readOnly: true });
  }
}

function executeEditsOneshot(
  editorInstance: editor.IStandaloneCodeEditor,
  edits: editor.IIdentifiedSingleEditOperation[]
) {
  executeEditCallback(editorInstance, () => {
    const result = editorInstance.executeEdits("", edits);
    if (!result) {
      // TODO: throw error to trigger error.tsx
      console.log("executeEdits for monaco editor failed!");
    }
  });
}

function executeEditsAnimation(
  editorInstance: editor.IStandaloneCodeEditor,
  edits: editor.IIdentifiedSingleEditOperation[]
) {
  const setTimeoutInterval = 20; // milliseconds

  function executeAtomicEdit(at: number) {
    executeEditCallback(editorInstance, () => {
      const result = editorInstance.executeEdits("", [edits[at]]);
      if (!result) {
        // TODO: throw error to trigger error.tsx
        console.log("executeEdits for monaco editor failed!");
      }
    });

    if (at < edits.length - 1) {
      window.setTimeout(() => executeAtomicEdit(at + 1), setTimeoutInterval);
    }
  }

  window.setTimeout(() => {
    executeAtomicEdit(0);
  }, setTimeoutInterval);
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
      const edits = props.editSequence?.edits;
      if (edits && edits.length > 0) {
        // clear previous edits upon props change
        if (isEditsMade) {
          executeEditCallback(editorInstance, () => {
            editorInstance.trigger("", "undo", null);
          });
        }

        if (props.editSequence?.animate) {
          executeEditsAnimation(editorInstance, edits);
        } else {
          executeEditsOneshot(editorInstance, edits);
        }
        isEditsMade.current = true;
      } else {
        // clear edits if edits is undefined or []
        editorInstance.updateOptions({ readOnly: false });

        editorInstance.trigger("", "undo", null);
        isEditsMade.current = false;

        editorInstance.updateOptions({ readOnly: true });
      }
    }
  }, [editorInstance, props.editSequence]);

  return <EditorBare onDidMount={onDidMount} />;
}
