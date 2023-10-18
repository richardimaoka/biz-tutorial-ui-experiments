import { MarkdownDefaultStyle } from "../../components/markdown2/MarkdownDefaultStyle";
import { promises as fs } from "fs";

export default async function Page() {
  const md = await fs.readFile(
    process.cwd() + "/app/test/markdown/sample-js.md",
    "utf-8"
  );
  return <MarkdownDefaultStyle markdownBody={md} />;
}
