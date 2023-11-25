import Link from "next/link";

export default async function Page() {
  return (
    <div style={{ height: "100svh" }}>
      <Link href={"/docker-cmd-entrypoint"}>docker-cmd-entrypoint</Link>
    </div>
  );
}
