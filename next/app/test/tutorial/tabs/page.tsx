import { ColumnTabs } from "@/app/components/tutorial/ColumnTabs";
import { ColumnName } from "@/app/components/tutorial/definitions";

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
