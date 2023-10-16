import { Column } from "@/app/components/column2/Column";
import { ColumnName } from "@/app/components/column2/tabTypes";

export default async function Page() {
  const tabs = [
    { href: "/", name: "Terminal" as ColumnName },
    { href: "/", name: "SourceCode" as ColumnName },
    { href: "/", name: "Browser" as ColumnName },
  ];

  return (
    <Column tabs={tabs} selectColumn="SourceCode">
      <div style={{ backgroundColor: "white", height: "100%" }} />
      <div style={{ backgroundColor: "white", height: "100%" }} />
      <div style={{ backgroundColor: "white", height: "100%" }} />
    </Column>
  );
}
