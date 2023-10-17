import { TerminalHeader } from "@/app/components/terminal2/TerminalHeader";

export default async function Page() {
  return (
    <div>
      <TerminalHeader
        currentDirectory=""
        tabs={[
          { name: "default", href: "test/terminal/tabs" },
          { name: "another", href: "test/terminal/tabs" },
        ]}
        selectTab="another"
      />
    </div>
  );
}
