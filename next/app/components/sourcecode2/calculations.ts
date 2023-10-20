import { editor } from "monaco-editor";

export function calculateEdit(
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

export function calculateRect(elem: HTMLElement) {
  return elem.offsetHeight;
}
