import { promises as fs } from "fs";
import { Child } from "./Child";
import { EditorWithTooltip } from "@/app/components/editor/EditorWithTooltip";

export default async function Page() {
  // Necessary to hardcode this, as the only other way to get `pathname` is usePathname(),
  // but that requires client component, which can't import "fs"
  const pathname = "app/test/editor/tooltip";

  const cwd = process.cwd();
  const goSource = await fs.readFile(
    `${cwd}/${pathname}/common_test.go`,
    "utf-8"
  );
  const mdContents = await fs.readFile(
    `${cwd}/${pathname}/tooltip.md`,
    "utf-8"
  );

  return (
    <div>
      <div style={{ height: "500px" }}>
        <EditorWithTooltip editorText={goSource} language="go" />
      </div>
      <Child markdownBody={mdContents} />
    </div>
  );
}
