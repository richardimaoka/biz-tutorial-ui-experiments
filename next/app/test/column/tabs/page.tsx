import { ColumnTabs } from "@/app/components/column2/ColumnTabs";
import { TabName } from "@/app/components/column2/tabTypes";

export default async function Page() {
  const tabs = [
    { href: "/", name: "Terminal" as TabName },
    { href: "/", name: "SourceCode" as TabName },
    { href: "/", name: "Browser" as TabName },
  ];

  return (
    <div>
      <ColumnTabs tabs={tabs} selectTab="SourceCode" />
    </div>
  );
}
