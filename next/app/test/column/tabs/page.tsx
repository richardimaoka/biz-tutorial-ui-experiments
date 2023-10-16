import { ColumnTabs } from "@/app/components/column2/ColumnTabs";
import { ColumnName } from "@/app/components/column2/definitions";

export default async function Page() {
  const tabs = [
    { href: "/", name: "Terminal" as ColumnName },
    { href: "/", name: "SourceCode" as ColumnName },
    { href: "/", name: "Browser" as ColumnName },
  ];

  return (
    <div>
      <ColumnTabs tabs={tabs} selectTab="SourceCode" />
    </div>
  );
}
