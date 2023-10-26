import React from "react";
import { promises as fs } from "fs";
import { Interactive } from "./Interactive";

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
      <Interactive original={oldSrc} modified={newSrc} language="typescript" />
    </div>
  );
}
