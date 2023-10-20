import { EditorHeightGetter } from "@/app/components/editor/EditorHeightGetter";
import React from "react";
import { promises as fs } from "fs";

export default async function Page() {
  // Necessary to hardcode this, as the only other way to get `pathname` is usePathname(),
  // but that requires client component
  const pathname = "app/test/editor";

  const cwd = process.cwd();
  const srcStr = await fs.readFile(
    `${cwd}/${pathname}/common_test.go`,
    "utf-8"
  );

  return (
    <div style={{ height: "700px" }}>
      <EditorHeightGetter dummy={srcStr} />
    </div>
  );
}
