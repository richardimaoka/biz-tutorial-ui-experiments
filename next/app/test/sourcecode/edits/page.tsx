import React from "react";
import { promises as fs } from "fs";
import { EditorEditable } from "@/app/components/sourcecode2/editor/EditorEditable";

// https://github.com/golang/go/commit/693def151adff1af707d82d28f55dba81ceb08e1

export default async function Page() {
  // Necessary to hardcode this, as the only other way to get `pathname` is usePathname(),
  // but that requires client component
  const pathname = "app/test/sourcecode/edits";

  const cwd = process.cwd();
  const newContents = await fs.readFile(
    `${cwd}/${pathname}/os_windows_new.go.txt`,
    "utf-8"
  );
  const oldContents = await fs.readFile(
    `${cwd}/${pathname}/os_windows_old.go.txt`,
    "utf-8"
  );

  return <div style={{ height: "700px" }}></div>;
}
