"use client";

import { DiffEditorBare } from "@/app/components/sourcecode2/diff-editor/DiffEditorBare";
import { useState } from "react";
import { EditorWithTooltip } from "@/app/components/sourcecode2/editor/EditorWithTooltip";

interface Props {
  original: string;
  modified: string;
  language: string;
}

export function Interactive(props: Props) {
  const [showDiff, setShowDiff] = useState(false);
  return (
    <div style={{ height: "100%" }}>
      <button
        onClick={() => {
          setShowDiff(!showDiff);
        }}
      >
        switch
      </button>
      <div style={{ height: "100%", display: showDiff ? "none" : "block" }}>
        <DiffEditorBare
          original={props.original}
          modified={props.modified}
          language={props.language}
        />
      </div>
      <div style={{ height: "100%", display: showDiff ? "block" : "none" }}>
        <EditorWithTooltip
          editorText={props.modified}
          language={props.language}
        />
      </div>
    </div>
  );
}

// export function Interactive(props: Props) {
//   const [index, setIndex] = useState(1);

//   return (
//     <div style={{ height: "100%" }} onClick={() => setIndex((index + 1) % 2)}>
//       <Carousel columnWidth={700} currentIndex={index}>
//         <div style={{ display: "flex", height: "100%", width: "1400px" }}>
//           <div style={{ width: "700px" }}>
//             <EditorEditable
//               editorText={props.modified}
//               language={props.language}
//             />
//           </div>
//           <div style={{ width: "700px" }}>
//             <DiffEditorBare
//               original={props.original}
//               modified={props.modified}
//               language={props.language}
//             />
//           </div>
//         </div>
//       </Carousel>
//     </div>
//   );
// }
