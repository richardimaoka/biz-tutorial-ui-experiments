import Link from "next/link";
import { promises as fs } from "fs";

export default async function Page() {
  const routingPath = "/test/tutorial";
  const cwd = process.cwd();
  const pwd = `${cwd}/app` + routingPath;

  const subDirEnts = await fs.readdir(pwd, { withFileTypes: true });
  const subDirs = subDirEnts
    .filter((dirent) => dirent.isDirectory())
    .map((dirent) => dirent.name);

  return (
    <ul style={{ margin: "40px" }}>
      {subDirs.map((l) => (
        <li key={l} style={{ marginBottom: "10px", fontSize: "20px" }}>
          <Link
            href={routingPath + "/" + l}
            style={{ color: "blue", textDecoration: "underline" }}
          >
            {l}
          </Link>
        </li>
      ))}
    </ul>
  );
}
