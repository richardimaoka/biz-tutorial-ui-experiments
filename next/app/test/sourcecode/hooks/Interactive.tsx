"use client";

import { EditorSimple } from "@/app/components/sourcecode2/editor/EditorSimple";
import { useEffect, useState } from "react";

interface Props {
  newEditorText: string;
  newLanguage: string;
  oldEditorText: string;
  oldLanguage: string;
}

export function Interactive(props: Props) {
  const [editorText, setEditorText] = useState("");
  const [language, setLanguage] = useState("");
  const [newOrOld, setNewOrOld] = useState<"new" | "old">("new");

  useEffect(() => {
    if (newOrOld === "new") {
      setEditorText(props.newEditorText);
      setLanguage(props.newLanguage);
    } else {
      // supposedly newOrOld === "old"
      setEditorText(props.oldEditorText);
      setLanguage(props.oldEditorText);
    }
  }, [newOrOld, props]);

  function toggleNewOrOld() {
    const updatedValue = newOrOld === "new" ? "old" : "new";
    setNewOrOld(updatedValue);
  }

  return (
    <>
      <button
        style={{ backgroundColor: "lightblue", color: "white" }}
        onClick={toggleNewOrOld}
      >
        toggle
      </button>
      <EditorSimple editorText={editorText} language={language} />
    </>
  );
}
