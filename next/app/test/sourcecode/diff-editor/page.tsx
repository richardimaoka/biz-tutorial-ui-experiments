import React from "react";
import { promises as fs } from "fs";
import { DiffEditorBare } from "@/app/components/sourcecode2/DiffEditorBare";

export default async function Page() {
  // Necessary to hardcode this, as the only other way to get `pathname` is usePathname(),
  // but that requires client component
  const pathname = "app/test/sourcecode/diff-editor";

  const cwd = process.cwd();
  const oldSrc = await fs.readFile(
    `${cwd}/${pathname}/EditorBare.tsx.old.txt`,
    "utf-8"
  );
  const newSrc = await fs.readFile(
    `${cwd}/${pathname}/EditorBare.tsx.new.txt`,
    "utf-8"
  );
  return (
    <div style={{ height: "700px" }}>
      <DiffEditorBare original={oldSrc} modified={newSrc} />
    </div>
  );
}
