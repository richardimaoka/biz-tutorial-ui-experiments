import React from "react";
import { promises as fs } from "fs";
import { EditorSimple } from "@/app/components/sourcecode2/editor/EditorSimple";
import { editor } from "monaco-editor";

const predefinedAnimationEdits: editor.IIdentifiedSingleEditOperation[] = [
  {
    range: {
      startLineNumber: 3,
      startColumn: 14,
      endLineNumber: 3,
      endColumn: 14,
    },
    text: ",",
  },
  {
    range: {
      startLineNumber: 3,
      startColumn: 15,
      endLineNumber: 3,
      endColumn: 15,
    },
    text: " ",
  },
  {
    range: {
      startLineNumber: 3,
      startColumn: 16,
      endLineNumber: 3,
      endColumn: 16,
    },
    text: "{ O",
  },
  {
    range: {
      startLineNumber: 3,
      startColumn: 19,
      endLineNumber: 3,
      endColumn: 19,
    },
    text: "n",
  },
  {
    range: {
      startLineNumber: 3,
      startColumn: 20,
      endLineNumber: 3,
      endColumn: 20,
    },
    text: "Change }",
  },
];

function range(line: number, column: number) {
  return {
    startLineNumber: line,
    startColumn: column,
    endLineNumber: line,
    endColumn: column,
  };
}

const predefinedAnimationEditsByWord: editor.IIdentifiedSingleEditOperation[] =
  [
    {
      range: range(7, 2),
      text: "\n",
    },
    {
      range: range(7, 2),
      text: "// ",
    },
    {
      range: range(7, 5),
      text: "onDidMount: ",
    },
    {
      range: range(7, 2),
      text: "pass-in ",
    },
    {
      range: range(7, 2),
      text: "a callback",
    },
    {
      range: range(7, 2),
      text: "like ",
    },
    {
      range: range(7, 2),
      text: "below ",
    },
    {
      range: range(7, 2),
      text: "to ",
    },
    {
      range: range(7, 2),
      text: "manipulate ",
    },
    {
      range: range(7, 3),
      text: "editor instance",
    },
  ];

export default async function Page() {
  // Necessary to hardcode this, as the only other way to get `pathname` is usePathname(),
  // but that requires client component
  const pathname = "app/test/sourcecode/hooks/edits/animation";

  const cwd = process.cwd();
  const oldSrcStr = await fs.readFile(
    `${cwd}/${pathname}/EditorBare.tsx.old.txt`,
    "utf-8"
  );

  return (
    <div style={{ height: "700px" }}>
      <EditorSimple
        editorText={oldSrcStr}
        language={"typescript"}
        editSequence={{
          edits: predefinedAnimationEditsByWord,
          animate: true,
        }}
      />
    </div>
  );
}
