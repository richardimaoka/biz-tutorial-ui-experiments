import React from "react";
import { promises as fs } from "fs";
import { EditorSimple } from "@/app/components/sourcecode2/editor/EditorSimple";
import { editor } from "monaco-editor";
import { Interactive } from "./Interactive";

export default async function Page() {
  // Necessary to hardcode this, as the only other way to get `pathname` is usePathname(),
  // but that requires client component
  const pathname = "app/test/sourcecode/hooks/edits/animation2";

  const cwd = process.cwd();
  const oldSrcStr = await fs.readFile(
    `${cwd}/${pathname}/EditorBare.tsx.old.txt`,
    "utf-8"
  );
  const editsJson = await fs.readFile(`${cwd}/${pathname}/edits.json`, "utf-8");
  const edits: editor.IIdentifiedSingleEditOperation[] = JSON.parse(editsJson);

  return (
    <div style={{ height: "700px" }}>
      <Interactive editorText={oldSrcStr} edits={edits} />
    </div>
  );
}
