"use client";

// To avoid an error `ReferenceError: navigator is not defined`, dynamic import with ssr false is needed.
// This is because "monaco-editor" module uses browser-side `navigator` inside.
import dynamic from "next/dynamic";
const EditorInnerSimple = dynamic(
  () => import("./internal/EditorInnerSimpleOnlyDynamicallyImportable"),
  {
    ssr: false,
  }
);
import { editor } from "monaco-editor";
import { EditorTooltip } from "../tooltip/EditorTooltip";

interface Props {
  editorText: string;
  language: string;
  editSequence?: {
    edits: editor.IIdentifiedSingleEditOperation[];
    animate?: boolean;
  };
  tooltip?: {
    lineNumber: number;
    markdownBody: string;
    hidden?: boolean;
    offsetContent?: boolean;
  };
}

export function EditorSimple(props: Props) {
  // If tooltip is passed and not-hidden, then render it
  const tooltip =
    props.tooltip && !props.tooltip.hidden
      ? {
          lineNumber: props.tooltip.lineNumber,
          children: (
            <EditorTooltip markdownBody={props.tooltip?.markdownBody} />
          ),
        }
      : undefined;

  return (
    <EditorInnerSimple
      editorText={props.editorText}
      language={props.language}
      editSequence={props.editSequence}
      tooltip={tooltip}
    />
  );
}
