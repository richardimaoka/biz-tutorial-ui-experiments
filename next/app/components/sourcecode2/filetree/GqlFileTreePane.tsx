"use client";

import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { useEffect, useState } from "react";
import { GqlFileTreeHeader } from "./header/GqlFileTreeHeader";
import styles from "./GqlFileTreePane.module.css";
import { GqlFileTreeComponent } from "./tree/GqlFileTreeComponent";

const fragmentDefinition = graphql(`
  fragment GqlFileTreePane on SourceCode {
    ...GqlFileTreeHeader
    ...GqlFileTreeComponent
    isFoldFileTree
  }
`);

export interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
  currentDirectory?: string;
  step: string;
}

export const GqlFileTreePane = (props: Props): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  // TODO: CAUTION: this keeps refreshed upon parent component state change.
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
      <GqlFileTreeHeader
        fragment={fragment}
        isFolded={isFolded}
        onButtonClick={() => {
          setIsFolded(!isFolded);
        }}
      />
      <GqlFileTreeComponent
        isFolded={isFolded}
        fragment={fragment}
        step={props.step}
      />
    </div>
  );
};
