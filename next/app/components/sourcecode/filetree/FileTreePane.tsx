"use client";

import { useEffect, useState } from "react";
import { FileTreeHeader } from "./FileTreeHeader";
import styles from "./style.module.css";

import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { FileTreeComponent } from "./FileTreeComponent";

const fragmentDefinition = graphql(`
  fragment FileTreePane_Fragment on SourceCode {
    ...FileTreeHeader_Fragment
    ...FileTreeComponent_Fragment
    isFoldFileTree
  }
`);

export interface FileTreePaneProps {
  fragment: FragmentType<typeof fragmentDefinition>;
  currentDirectory?: string;
  step: string;
}

export const FileTreePane = (props: FileTreePaneProps): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  // CAUTION: this keeps refreshed upon parent component state change.
  // So some workaround or redesign is needed.
  const [isFolded, setIsFolded] = useState(true);

  useEffect(() => {
    if (typeof fragment.isFoldFileTree === "boolean") {
      setIsFolded(fragment.isFoldFileTree);
    }
  }, [props.step, fragment.isFoldFileTree]);

  return (
    <div
      className={`${styles.pane} ${isFolded ? styles.folded : styles.expanded}`}
    >
      <FileTreeHeader
        fragment={fragment}
        isFolded={isFolded}
        onButtonClick={() => {
          setIsFolded(!isFolded);
        }}
      />
      <FileTreeComponent
        isFolded={isFolded}
        fragment={fragment}
        step={props.step}
      />
    </div>
  );
};
