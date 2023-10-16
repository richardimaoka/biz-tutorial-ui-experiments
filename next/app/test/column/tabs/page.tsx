import { ColumnTab } from "@/app/components/column2/ColumnTab";

export default async function Page() {
  return (
    <div>
      <ColumnTab href="/test/column/tabs" isSelected name="Terminal" />
      <ColumnTab href="/test/column/tabs" isSelected={false} name="Terminal" />
    </div>
  );
}
