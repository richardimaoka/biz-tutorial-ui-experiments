"use client";

import { EditorWithTooltip } from "@/app/components/sourcecode2/EditorWithTooltip";
import { useState } from "react";

interface Props {
  editorText: string;
}

export function Toggler(props: Props) {
  const [showToolTip, setShowTooltip] = useState(false);

  return (
    <div style={{ height: "600px" }}>
      <button onClick={() => setShowTooltip(!showToolTip)}>toggle</button>
      <EditorWithTooltip
        editorText={props.editorText}
        language="go"
        tooltip={showToolTip ? { startLineNumber: 3, numLines: 2 } : undefined}
      />
    </div>
  );
}
