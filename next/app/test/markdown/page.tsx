import { MarkdownDefaultStyle } from "../../components/markdown/server-component/MarkdownDefaultStyle";
import { promises as fs } from "fs";

export default async function Page() {
  const md = await fs.readFile(
    process.cwd() + "/app/test/markdown/nextjs.md",
    "utf-8"
  );
  return <MarkdownDefaultStyle markdownBody={md} />;
}
