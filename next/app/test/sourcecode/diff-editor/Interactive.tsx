"use client";

import { DiffEditorBare } from "@/app/components/sourcecode2/diff-editor/DiffEditorBare";
import { EditorEditable } from "@/app/components/sourcecode2/editor/EditorEditable";
import { useState } from "react";

interface Props {
  original: string;
  modified: string;
  language: string;
}

export function Interactive(props: Props) {
  const [showDiff, setShowDiff] = useState(false);

  return (
    <div style={{ height: "100%" }}>
      <button
        onClick={() => {
          setShowDiff(!showDiff);
        }}
      >
        switch
      </button>
      {showDiff ? (
        <DiffEditorBare
          original={props.original}
          modified={props.modified}
          language={props.language}
        />
      ) : (
        <EditorEditable editorText={props.modified} language={props.language} />
      )}
    </div>
  );
}
