import { editor } from "monaco-editor";
import { useCallback, useEffect, useRef, useState } from "react";

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
    id: string;
    edits: editor.IIdentifiedSingleEditOperation[];
    skipAnimation?: boolean;
  }
) {
  // Save the ID of the last edit to clear previous edits, upon editSequence change
  // useRef, since monaco editor is separate from React state.
  const lastEditID = useRef("");

  // Similar to `isEditMade` flag above, but this is for the caller, to act on edit completion
  const [isEditCompleted, setEditCompleted] = useState(false);

  function markCompletion() {
    setEditCompleted(true);
  }

  // handle editSequence update, even when it becomes undefined
  useEffect(() => {
    if (editorInstance) {
      const edits = editSequence?.edits;

      if (edits && edits.length > 0) {
        /**
         * If edits are non-empty
         */
        if (editSequence.id !== lastEditID.current) {
          // clear previous edits upon props change
          if (lastEditID.current !== "") {
            executeEditCallback(editorInstance, () => {
              editorInstance.trigger("", "undo", null);
            });
          }

          // execute edits
          if (editSequence.skipAnimation) {
            executeEditsStatic(editorInstance, edits, markCompletion);
          } else {
            executeEditsAnimation(editorInstance, edits, markCompletion);
          }

          // save the edit-made flag
          lastEditID.current = editSequence.id;
        }
      } else {
        /**
         * Else if edits are empty
         */

        // clear previous edits
        executeEditCallback(editorInstance, () => {
          editorInstance.trigger("", "undo", null);
        });

        // clear the edit-made flag
        lastEditID.current = "";
      }
    }
  }, [editorInstance, editSequence]);

  useEffect(() => {
    // Whenever edit sequence is updated, set the completed flag as false
    if (lastEditID.current !== editSequence?.id) {
      setEditCompleted(false);
    }
  }, [editSequence]);

  return { isEditCompleted };
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

function executeEditsStatic(
  editorInstance: editor.IStandaloneCodeEditor,
  edits: editor.IIdentifiedSingleEditOperation[],
  markCompletion: () => void
) {
  for (const e of edits) {
    // for-loop is necessary - cannnot pass-in the whole `edits` to executeEdits()
    // in one-shot, because that could cause the below error:
    //   Error: Overlapping ranges are not allowed!
    executeEditCallback(editorInstance, () => {
      const result = editorInstance.executeEdits("", [e]);
      if (!result) {
        // TODO: throw error to trigger error.tsx
        console.log("executeEdits for monaco editor failed!");
      }
    });
  }
  markCompletion();
}

function executeEditsAnimation(
  editorInstance: editor.IStandaloneCodeEditor,
  edits: editor.IIdentifiedSingleEditOperation[],
  markCompletion: () => void
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
    } else {
      markCompletion();
    }
  }

  window.setTimeout(() => {
    executeAtomicEdit(0);
  }, setTimeoutInterval);
}
