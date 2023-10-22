import React from "react";
import { promises as fs } from "fs";
import { EditorSimple } from "@/app/components/sourcecode2/editor/EditorSimple";
import { Interactive } from "./Interactive";

export default async function Page() {
  // Necessary to hardcode this, as the only other way to get `pathname` is usePathname(),
  // but that requires client component
  const pathname = "app/test/sourcecode/hooks";

  const cwd = process.cwd();
  const newSrcStr = await fs.readFile(
    `${cwd}/${pathname}/EditorBare.tsx.new.txt`,
    "utf-8"
  );
  const oldSrcStr = await fs.readFile(
    `${cwd}/${pathname}/EditorBare.tsx.old.txt`,
    "utf-8"
  );

  return (
    <div style={{ height: "700px" }}>
      <Interactive
        newEditorText={newSrcStr}
        oldEditorText={oldSrcStr}
        newLanguage={"typescript"}
        oldLanguage="typescript"
      />
    </div>
  );
}
