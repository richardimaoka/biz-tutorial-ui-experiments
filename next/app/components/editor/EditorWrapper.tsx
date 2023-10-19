"use client";

import { useEffect } from "react";
import { EditorBare } from "./EditorBare";
import { useEditorInstance } from "./useEditorInstance";

interface Props {
  dummy: string;
}

export function EditorWrapper(props: Props) {
  const [editorInstance, onDidMount] = useEditorInstance();

  useEffect(() => {
    console.log("EditorWrapper useEffect", editorInstance);
    editorInstance.current &&
      console.log(
        "line Height=",
        editorInstance.current?.getOption(
          // somehow, NOT `enum lineHeight` but `enum lineHeight - 1`*/
          65
        )
      );
  });

  return <EditorBare onDidMount={onDidMount} />;
}
