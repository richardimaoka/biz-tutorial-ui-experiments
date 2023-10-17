import { TerminalTabs } from "@/app/components/terminal2/TerminalTabs";

export default async function Page() {
  return (
    <div>
      <TerminalTabs
        tabs={[
          { name: "default", href: "test/terminal/tabs" },
          { name: "another", href: "test/terminal/tabs" },
        ]}
        selectTab="default"
      />
    </div>
  );
}
