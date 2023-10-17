import { TerminalComponent } from "@/app/components/terminal2/TerminalComponent";

export default async function Page() {
  return (
    <div style={{ height: "100svh" }}>
      {/* <div
        style={{ width: "500px", height: "300px", backgroundColor: "white" }}
      >
        <div style={{ height: "100px", backgroundColor: "aqua" }}></div>
      </div> */}
      <TerminalComponent
        tabs={[
          { name: "default", href: "/test/terminal/" },
          { name: "another", href: "/test/terminal/" },
        ]}
        currentDirectory="/test/terminal"
        selectTab="default"
        entries={[]}
      />
    </div>
  );
}
