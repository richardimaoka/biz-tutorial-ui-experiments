"use client";
import { useSearchParams } from "next/navigation";
import { EditOperation, SourceCodeEditor } from "./SourceCodeEditor";
import { ReactNode } from "react";

interface Props {
  currentContents: string;
  oldContents: string;
  language: string;

  /**
   * optional props below
   */
  editSequence?: {
    id: string;
    edits: EditOperation[];
  };
  tooltip?: {
    lineNumber: number;
    children: ReactNode;
    timing: "START" | "END";
  };
  defaultFocusColumn?: string;
}

export function SourceCodeAnimationHandler(props: Props) {
  const searchParams = useSearchParams();

  // doAnimate - see if we need to animate editSequence
  const columnParam = searchParams.get("column");
  const currentColumn = columnParam ? columnParam : props.defaultFocusColumn;
  const isSelected = currentColumn === "SourceCode";
  const doAnimate = props.editSequence && isSelected;

  console.log("doAnimate source", doAnimate, currentColumn);

  // To animate, start from `editorText = oldContents`
  // then make animation with editSequence
  const editorText = doAnimate ? props.oldContents : props.currentContents;
  const editSequence = doAnimate ? props.editSequence : undefined;

  console.log("editorText", editorText);

  return (
    <SourceCodeEditor
      editorText={editorText}
      language={props.language}
      editSequence={editSequence}
      tooltip={props.tooltip}
    />
  );
}
