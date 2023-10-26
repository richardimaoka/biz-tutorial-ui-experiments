"use client";

import { DiffEditorBare } from "@/app/components/sourcecode2/diff-editor/DiffEditorBare";
import { EditorEditable } from "@/app/components/sourcecode2/editor/EditorEditable";
import { Carousel } from "@/app/components/tutorial/carousel/Carousel";
import { useState } from "react";

interface Props {
  original: string;
  modified: string;
  language: string;
}

export function Interactive(props: Props) {
  const [index, setIndex] = useState(1);

  return (
    <div style={{ height: "100%" }} onClick={() => setIndex((index + 1) % 2)}>
      <Carousel columnWidth={700} currentIndex={index}>
        <div style={{ display: "flex", height: "100%", width: "1400px" }}>
          <div style={{ width: "700px" }}>
            <EditorEditable
              editorText={props.modified}
              language={props.language}
            />
          </div>
          <div style={{ width: "700px" }}>
            <DiffEditorBare
              original={props.original}
              modified={props.modified}
              language={props.language}
            />
          </div>
        </div>
      </Carousel>
    </div>
  );
}
