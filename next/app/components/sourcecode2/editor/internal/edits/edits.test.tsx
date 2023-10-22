import { expect, test } from "vitest";
import { toOperation } from "./edits";
import { useEditorInstance } from "../useEditorInstance";
import { EditorBare } from "../EditorBare";
import { render } from "@testing-library/react";

function insertAt(original: string, at: number, toInsert: string): string {
  return original.slice(0, at) + toInsert + original.slice(at);
}

test("edits add", () => {
  const cases = [
    {
      fromText: `import Editor from "@monaco-editor/react";`,
      toText: `import Editor, { OnChange } from "@monaco-editor/react";`,
      edit: {
        at: 13, // zero-start index in string
        diff: `, { OnChange }`,
      },
    },
    {
      fromText: `bbbccc`,
      toText: `aaabbbccc`,
      edit: {
        at: 0, // zero-start index in string
        diff: `aaa`,
      },
    },
  ];

  cases.forEach((c) => {
    const result = insertAt(c.fromText, c.edit.at, c.edit.diff);
    expect(result).toStrictEqual(c.toText);
  });
});

test("edits convert", () => {
  const cases = [
    {
      edit: {
        at: 13, // zero-start index in string
        diff: `, { OnChange }`,
        lineNumber: 1,
      },
      op: {
        range: {
          startLineNumber: 1,
          startColumn: 14,
          endLineNumber: 1,
          endColumn: 14,
        },
        text: `, { OnChange }`,
      },
    },
  ];

  cases.forEach((c) => {
    const result = toOperation(c.edit);
    expect(result).toStrictEqual(c.op);
  });
});

// test("useEditorInstance", () => {
//   const [editorInstance, onDidMount] = useEditorInstance();

//   render(<EditorBare onDidMount={onDidMount} />);
//   expect(editorInstance).toBeInstanceOf("object");
// });
