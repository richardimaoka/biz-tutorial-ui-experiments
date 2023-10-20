"use client";

// To avoid an error `ReferenceError: navigator is not defined`, dynamic import with ssr false is needed.
// This is because "monaco-editor" module uses browser-side `navigator` inside.
import dynamic from "next/dynamic";
const EditorEditableInner = dynamic(
  () => import("./EditorEditableOnlyDynamicallyImportable"),
  {
    ssr: false,
  }
);

// this is ok with static import
import { EditorEditableInnerProps } from "./EditorEditableOnlyDynamicallyImportable";
import { useEffect, useRef, useState } from "react";
import { createRoot } from "react-dom/client";

type Props = EditorEditableInnerProps;

export function EditorEditable(props: Props) {
  const [container] = useState<HTMLDivElement>(document.createElement("div"));

  useEffect(() => {
    container.id = "containercontainecontainer";
    container.style.height = "100px";
    container.style.width = "100px";
    container.style.backgroundColor = "white";

    const root = createRoot(container);
    root.render(<div>aaa</div>);
    return container.remove();
  }, [container]);

  return (
    <EditorEditableInner
      editorText={props.editorText}
      language={props.language}
      edits={props.edits}
      contentWidgetElement={container}
    />
  );
}
