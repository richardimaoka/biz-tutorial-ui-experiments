"use client";

import { TerminalComponent } from "@/app/components/terminal2/__TerminalComponent";
import { entries } from "./entries";
import React from "react";

export default function Page() {
  const [sliceEnd, setSliceEnd] = React.useState(3);

  function onClick() {
    setSliceEnd(sliceEnd + 1);
  }

  return (
    <div style={{ height: "100svh", width: "100%" }} onClick={onClick}>
      <TerminalComponent
        tabs={[
          { name: "default", href: "/test/terminal/" },
          { name: "another", href: "/test/terminal/" },
        ]}
        currentDirectory="/test/terminal"
        selectTab="default"
        entries={entries.slice(0, sliceEnd)}
        isAnimate
      />
    </div>
  );
}
