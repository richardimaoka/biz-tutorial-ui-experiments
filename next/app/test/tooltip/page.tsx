import { Tooltip } from "@/app/components/tooltip/Tooltip";
import { promises as fs } from "fs";

export default async function Page() {
  const md = await fs.readFile(
    process.cwd() + "/app/test/tooltip/tooltip.md",
    "utf-8"
  );
  return <Tooltip body={md} />;
}
