"use client";

import { EditorWithTooltip } from "@/app/components/sourcecode2/editor/EditorWithTooltip";
import { editor } from "monaco-editor";
import { useState } from "react";

interface Props {
  editorText: string;
  edits: editor.IIdentifiedSingleEditOperation[];
}

export function Interactive(props: Props) {
  const [state, setState] = useState(true);

  const editsSeq = state ? { edits: props.edits, animate: true } : undefined;

  function onClick() {
    setState(!state);
  }

  return (
    <>
      <button onClick={onClick}> animate</button>
      <EditorWithTooltip
        editorText={props.editorText}
        language={"typescript"}
        editSequence={editsSeq}
      />
    </>
  );
}
