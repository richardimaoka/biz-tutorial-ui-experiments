import { editor } from "monaco-editor";
import { EditorEditable } from "./EditorEditable";

function calculateEdit(
  startLineNumber: number,
  numLines: number
): editor.ISingleEditOperation {
  const a = new Array(numLines);
  const newLinesForOffset = a.fill("\n", 0, a.length); // fill() is necessary, map doesn't work for new Array()
  const insertText = newLinesForOffset.join("");
  return {
    range: {
      startLineNumber: startLineNumber,
      startColumn: 1,
      endLineNumber: startLineNumber,
      endColumn: 1,
    },
    text: insertText,
  };
}

interface Props {
  editorText: string;
  language: string;
  tooltip?: {
    startLineNumber: number;
    numLines: number;
  };
}

export function EditorWithTooltip(props: Props) {
  const edits = props.tooltip
    ? [calculateEdit(props.tooltip.startLineNumber, props.tooltip.numLines)] // single operation in array
    : undefined;

  return (
    <EditorEditable
      editorText={props.editorText}
      language={props.language}
      edits={edits}
    />
  );
}
