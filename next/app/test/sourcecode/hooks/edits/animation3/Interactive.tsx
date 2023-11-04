"use client";

import { DiffEditorBare } from "@/app/components/sourcecode2/diff-editor/DiffEditorBare";
import { EditorWithTooltip } from "@/app/components/sourcecode2/editor/EditorWithTooltip";
import { editor } from "monaco-editor";
import { useState } from "react";

interface Props {
  oldEditorText: string;
  newEditorText: string;
  edits: editor.IIdentifiedSingleEditOperation[];
}

type State = "old" | "new" | "animate" | "diff" | "tooltip";

export function Interactive(props: Props) {
  const [state, setState] = useState<State>("old");

  const editorText =
    state === "animate" || state === "old"
      ? props.oldEditorText
      : props.newEditorText;

  const editsSeq =
    state === "animate" ? { edits: props.edits, animate: true } : undefined;

  const showDiff = state === "diff";

  return (
    <>
      <button style={{ marginRight: "10px" }} onClick={() => setState("old")}>
        old
      </button>
      <button
        style={{ marginRight: "10px" }}
        onClick={() => setState("animate")}
      >
        animate
      </button>
      <button style={{ marginRight: "10px" }} onClick={() => setState("diff")}>
        diff
      </button>
      <button
        style={{ marginRight: "10px" }}
        onClick={() => setState("tooltip")}
      >
        tooltip
      </button>
      <div style={{ height: "100%", display: showDiff ? "block" : "none" }}>
        <DiffEditorBare
          original={props.oldEditorText}
          modified={props.newEditorText}
          language={"typescript"}
        />
      </div>
      <div style={{ height: "100%", display: showDiff ? "none" : "block" }}>
        <EditorWithTooltip
          editorText={editorText}
          language={"typescript"}
          editSequence={editsSeq}
        />
      </div>
    </>
  );
}
