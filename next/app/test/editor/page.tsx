"use client";

import { Editor } from "@/app/components/editor/Editor";
import React from "react";

export default function Page() {
  const [name, setName] = React.useState("");
  return (
    <div>
      <input
        value={name}
        onChange={(event) => {
          setName(event.target.value);
        }}
      />
      <Editor name={name} />
    </div>
  );
}
