import { expect, test } from "vitest";

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
    expect(result).toBe(c.toText);
  });
});
