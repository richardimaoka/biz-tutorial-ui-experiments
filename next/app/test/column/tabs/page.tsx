import { ColumnTabs } from "@/app/components/column2/ColumnTabs";
import { ColumnTabProps } from "@/app/components/column2/ColumnTab";

export default async function Page() {
  const tabs: ColumnTabProps[] = [
    { href: "/", isSelected: true, name: "Terminal" },
    { href: "/", isSelected: false, name: "SourceCode" },
    { href: "/", name: "Browser" },
  ];

  return (
    <div>
      <ColumnTabs tabs={tabs} />
    </div>
  );
}
