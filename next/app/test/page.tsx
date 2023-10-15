import { MarkdownNoStyle } from "../components/markdown/MarkdownNoStyle";
import { promises as fs } from "fs";

export default async function Page() {
  const md = await fs.readFile(process.cwd() + "/app/test/sample.md", "utf-8");
  return <MarkdownNoStyle markdownBody={md} />;
}
