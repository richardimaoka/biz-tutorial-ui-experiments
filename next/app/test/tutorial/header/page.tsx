import { TutorialHeader } from "@/app/components/column2/TutorialHeader";
import { ColumnName } from "@/app/components/column2/definitions";

export default async function Page() {
  const tabs = [
    { href: "/", name: "Terminal" as ColumnName },
    { href: "/", name: "SourceCode" as ColumnName },
    { href: "/", name: "Terminal" as ColumnName },
    { href: "/", name: "SourceCode" as ColumnName },
    { href: "/", name: "Terminal" as ColumnName },
    { href: "/", name: "Browser" as ColumnName },
    { href: "/", name: "SourceCode" as ColumnName },
    { href: "/", name: "Terminal" as ColumnName },
    { href: "/", name: "SourceCode" as ColumnName },
    { href: "/", name: "SourceCode" as ColumnName },
    { href: "/", name: "Terminal" as ColumnName },
    { href: "/", name: "SourceCode" as ColumnName },
  ];

  return (
    <div>
      <TutorialHeader tabs={tabs} selectTab="Browser" />
    </div>
  );
}
