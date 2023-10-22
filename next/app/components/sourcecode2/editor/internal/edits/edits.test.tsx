import { expect, test } from "vitest";

function insertAt(original: string, at: number, toInsert: string): string {
  return original.slice(0, at) + toInsert + original.slice(at);
}

test("edits add", () => {
  const oldText = `import Editor from "@monaco-editor/react";`;
  const newText = `import Editor, { OnChange } from "@monaco-editor/react";`;

  const edit = {
    at: 13, // zero-start index in string
    diff: `, { OnChange }`,
  };

  const result = insertAt(oldText, edit.at, edit.diff);
  expect(result).toBe(newText);
});
