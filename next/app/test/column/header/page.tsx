import { ColumnHeader } from "@/app/components/column2/ColumnHeader";
import { TabName } from "@/app/components/column2/tabTypes";

export default async function Page() {
  const tabs = [
    { href: "/", name: "Terminal" as TabName },
    { href: "/", name: "SourceCode" as TabName },
    { href: "/", name: "Terminal" as TabName },
    { href: "/", name: "SourceCode" as TabName },
    { href: "/", name: "Terminal" as TabName },
    { href: "/", name: "Browser" as TabName },
    { href: "/", name: "SourceCode" as TabName },
    { href: "/", name: "Terminal" as TabName },
    { href: "/", name: "SourceCode" as TabName },
    { href: "/", name: "SourceCode" as TabName },
    { href: "/", name: "Terminal" as TabName },
    { href: "/", name: "SourceCode" as TabName },
  ];

  return (
    <div>
      <ColumnHeader tabs={tabs} selectTab="Browser" />
    </div>
  );
}
