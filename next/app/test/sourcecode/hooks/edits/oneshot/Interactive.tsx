"use client";

import { EditorSimple } from "@/app/components/sourcecode2/editor/EditorSimple";
import { useEffect, useState } from "react";
import { editor } from "monaco-editor";
interface Props {
  newEditorText: string;
  newLanguage: string;
  oldEditorText: string;
  oldLanguage: string;
}

const predefinedOneshotEdits: editor.IIdentifiedSingleEditOperation[] = [
  {
    range: {
      startLineNumber: 3,
      startColumn: 14,
      endLineNumber: 3,
      endColumn: 14,
    },
    text: ",",
  },
  {
    range: {
      startLineNumber: 3,
      startColumn: 14,
      endLineNumber: 3,
      endColumn: 14,
    },
    text: "*",
  },
  {
    range: {
      startLineNumber: 3,
      startColumn: 14,
      endLineNumber: 3,
      endColumn: 14,
    },
    text: "{",
  },
  {
    range: {
      startLineNumber: 3,
      startColumn: 14,
      endLineNumber: 3,
      endColumn: 14,
    },
    text: "*",
  },
  {
    range: {
      startLineNumber: 3,
      startColumn: 14,
      endLineNumber: 3,
      endColumn: 14,
    },
    text: "O",
  },
];

export function Interactive(props: Props) {
  const [editorText, setEditorText] = useState("");
  const [language, setLanguage] = useState("");
  const [newOrOld, setNewOrOld] = useState<"new" | "old">("old");
  const [edits, setEdits] = useState<editor.IIdentifiedSingleEditOperation[]>(
    []
  );

  useEffect(() => {
    if (newOrOld === "new") {
      setEditorText(props.newEditorText);
      setLanguage(props.newLanguage);
    } else {
      // supposedly newOrOld === "old"
      setEditorText(props.oldEditorText);
      setLanguage(props.oldLanguage);
    }
  }, [newOrOld, props]);

  function toggleNewOrOld() {
    const updatedValue = newOrOld === "new" ? "old" : "new";
    setNewOrOld(updatedValue);
  }

  function proceedToNextEdit() {
    const nextEnd =
      edits.length < predefinedOneshotEdits.length ? edits.length + 1 : 0;

    setEdits(predefinedOneshotEdits.slice(0, nextEnd));
  }

  return (
    <>
      <button
        style={{ backgroundColor: "lightblue", color: "black" }}
        onClick={toggleNewOrOld}
      >
        toggle
      </button>
      <button
        style={{ backgroundColor: "black", color: "white" }}
        onClick={proceedToNextEdit}
      >
        next edit
      </button>
      <EditorSimple
        editorText={editorText}
        language={language}
        editSequence={{
          edits: edits,
        }}
      />
    </>
  );
}
