import { editor } from "monaco-editor";
import { useEffect, useRef } from "react";

/**
 * Custm hook to handle @param `editSequence` props update
 *
 * @return nothing, as it is an effectful hook
 */
export function useEditSequence(
  // Shared monaco-editor's editor instance, possibly null
  editorInstance: editor.IStandaloneCodeEditor | null,

  // Prop to be passed from the parent component
  editSequence?: {
    edits: editor.IIdentifiedSingleEditOperation[];
    animate?: boolean;
  }
) {
  // useRef, since monaco editor is separate from React state.
  const isEditsMade = useRef(false);

  // handle editSequence update, even when it becomes undefined
  useEffect(() => {
    if (editorInstance) {
      const edits = editSequence?.edits;

      /**
       * If edits are non-empty
       */
      if (edits && edits.length > 0) {
        // clear previous edits upon props change
        if (isEditsMade) {
          executeEditCallback(editorInstance, () => {
            editorInstance.trigger("", "undo", null);
          });
        }

        // execute edits
        if (editSequence?.animate) {
          executeEditsAnimation(editorInstance, edits);
        } else {
          executeEditsOneshot(editorInstance, edits);
        }

        // save the edit-made flag
        isEditsMade.current = true;
      } else {
        /**
         * Else if edits are empty
         */
        // clear previous edits
        executeEditCallback(editorInstance, () => {
          editorInstance.trigger("", "undo", null);
        });

        // clear the edit-made flag
        isEditsMade.current = false;
      }
    }
  }, [editorInstance, editSequence]);
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
