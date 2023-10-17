import { TutorialComponent } from "@/app/components/tutorial/TutorialComponent";
import { ColumnName } from "@/app/components/tutorial/definitions";

export default async function Page() {
  const tabs = [
    { href: "/", name: "Terminal" as ColumnName },
    { href: "/", name: "SourceCode" as ColumnName },
    { href: "/", name: "Browser" as ColumnName },
  ];

  return (
    <TutorialComponent tabs={tabs} selectColumn="SourceCode" columns={[]}>
      <div
        style={{ backgroundColor: "white", height: "90%", margin: "100px" }}
      />
    </TutorialComponent>
  );
}
