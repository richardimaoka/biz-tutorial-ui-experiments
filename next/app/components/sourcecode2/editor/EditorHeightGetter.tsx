"use client";

import { useEffect } from "react";
import { EditorBare } from "./internal/EditorBare";
import { useEditorInstance } from "./internal/useEditorInstance";

interface Props {
  dummy: string;
}

export function EditorHeightGetter(props: Props) {
  const [editorInsntance, onDidMount] = useEditorInstance();

  useEffect(() => {
    console.log("EditorWrapper useEffect", editorInsntance);
    editorInsntance &&
      console.log(
        "line Height=",
        editorInsntance.getOption(
          // somehow, NOT `enum lineHeight` but `enum lineHeight - 1`
          65
        )
      );
  });

  return <EditorBare onDidMount={onDidMount} />;
}
