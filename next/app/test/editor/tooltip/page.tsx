"use client";

import React from "react";
import * as fs from "fs";
import { EditorEditable } from "@/app/components/editor/EditorEditable";
import { usePathname } from "next/navigation";

export default function Page() {
  const pathname = usePathname();
  const cwd = process.cwd();
  const srcStr = fs.readFileSync(`${cwd}/${pathname}/common_test.go`, "utf-8");

  return (
    <div style={{ height: "700px" }}>
      <EditorEditable editorText={srcStr} language="go" />
    </div>
  );
}
