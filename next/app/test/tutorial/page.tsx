import { Tutorial } from "@/app/components/tutorial/Tutorial";
import { ColumnName } from "@/app/components/tutorial/definitions";

export default async function Page() {
  const tabs = [
    { href: "/", name: "Terminal" as ColumnName },
    { href: "/", name: "SourceCode" as ColumnName },
    { href: "/", name: "Browser" as ColumnName },
  ];

  return (
    <Tutorial tabs={tabs} selectColumn="SourceCode">
      <div
        style={{ backgroundColor: "white", height: "90%", margin: "100px" }}
      />
    </Tutorial>
  );
}
