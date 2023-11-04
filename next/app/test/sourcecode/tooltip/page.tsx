import { promises as fs } from "fs";
import { EditorWithTooltip } from "@/app/components/sourcecode2/editor/EditorWithTooltip";

export default async function Page() {
  // Necessary to hardcode this, as the only other way to get `pathname` is usePathname(),
  // but that requires client component, which can't import "fs"
  const pathname = "app/test/sourcecode/tooltip";

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
    <div style={{ height: "80svh", maxWidth: "768px" }}>
      <EditorWithTooltip
        editorText={goSource}
        language="go"
        tooltip={{ markdownBody: mdContents, lineNumber: 5 }}
      />
    </div>
  );
}
