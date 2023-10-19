import { promises as fs } from "fs";
import { Toggler } from "./Toggler";

export default async function Page() {
  // Necessary to hardcode this, as the only other way to get `pathname` is usePathname(),
  // but that requires client component, which can't import "fs"
  const pathname = "app/test/editor/tooltip";

  const cwd = process.cwd();
  const srcStr = await fs.readFile(
    `${cwd}/${pathname}/common_test.go`,
    "utf-8"
  );

  return <Toggler editorText={srcStr} />;
}
