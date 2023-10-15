import { Tooltip } from "@/app/components/tooltip/Tooltip";
import { promises as fs } from "fs";

export default async function Page() {
  const md = await fs.readFile(
    process.cwd() + "/app/test/tooltip/tooltip.md",
    "utf-8"
  );
  return (
    <div
      style={{ height: "100%", backgroundColor: "grey", paddingTop: "100px" }}
    >
      <Tooltip body={md} />
    </div>
  );
}
