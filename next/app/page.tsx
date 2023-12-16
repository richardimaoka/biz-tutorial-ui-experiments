import Link from "next/link";

export default async function Page() {
  return (
    <div style={{ height: "100svh" }}>
      <div>
        <Link href={"/docker-cmd-entrypoint"}>docker-cmd-entrypoint</Link>
      </div>
      <div>
        <Link href={"/nextjs-deploy-cloudrun"}>nextjs-deploy-cloudrun</Link>
      </div>
    </div>
  );
}
