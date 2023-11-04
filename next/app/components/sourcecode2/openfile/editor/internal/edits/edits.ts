import { editor } from "monaco-editor";

type Edit = {
  at: number; // zero-start index in string
  lineNumber: number;
  diff: string;
};

export function toOperation(e: Edit): editor.IIdentifiedSingleEditOperation {
  return {
    range: {
      startLineNumber: e.lineNumber,
      startColumn: e.at + 1,
      endLineNumber: e.lineNumber,
      endColumn: e.at + 1,
    },
    text: e.diff,
  };
}
