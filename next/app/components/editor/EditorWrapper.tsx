"use client";

import { useEffect } from "react";
import { EditorBare } from "./EditorBare";
import { useEditorInstance } from "./useEditorInstance";

interface Props {
  dummy: string;
}

export function EditorWrapper(props: Props) {
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
